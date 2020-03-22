//
//  IMClient.swift
//  Coffchat
//
//  Created by xuyingchun on 2020/3/10.
//  Copyright © 2020 Xuyingchun Inc. All rights reserved.
//

import CocoaAsyncSocket
import Foundation

// IM结果回调
typealias IMResultCallback<T> = (_ res: T) -> Void

// IMClient协议回调 - 连接状态
protocol IMClientDelegateStatus {
    /// connected
    /// - Parameters:
    ///   - host: IP
    ///   - port: 端口
    func onConnected(_ host: String, port: UInt16)
    /// disconnected
    /// - Parameter err: 错误
    func onDisconnect(_ err: Error?)
}

// IMClient协议回调 - 数据
protocol IMClientDelegateData {
    /// 收到数据
    /// - Parameters:
    ///   - header: 协议头
    ///   - data: 数据体（裸数据）
    func onHandleData(_ header: IMHeader, _ data: Data)
}

let kClientVersion = "0.0.1"
let DefaultIMClient = IMClient()

// IM连接
// 负责与服务端通信
class IMClient: NSObject, GCDAsyncSocketDelegate {
    fileprivate var tcpClient: GCDAsyncSocket?
    fileprivate var seq: UInt16 = 1 // 序号，发送一次后即递增，没加锁FIXME
    fileprivate var checkTimer: Timer? // 超时检测定时器
    // fileprivate var recvBuffer:Data // TCP缓冲区，粘包处理
    
    // callback
    fileprivate var delegateDicStatus: [String: IMClientDelegateStatus] = [:] // tcp连接状态的回调
    fileprivate var delegateDicData: [String: IMClientDelegateData] = [:] // 业务数据的回调
    fileprivate var requestDic: [UInt16: IMRequest] = [:] // 需要响应的请求列表
    
    // 是否已连接
    public var isConnected: Bool? { return tcpClient?.isConnected }
    public var ip: String = "10.0.106.117"
    public var port: UInt16 = 8000
    
    override init() {
        // recvBuffer = Data()
        
        super.init()
        IMLog.debug(item: "CIMClient init")
        tcpClient = GCDAsyncSocket(delegate: self, delegateQueue: DispatchQueue.main)
        
        checkTimer = Timer.scheduledTimer(withTimeInterval: 1, repeats: true, block: _checkTimer)
        checkTimer?.fire()
    }
    
    deinit {
        tcpClient?.delegate = nil
        tcpClient?.delegateQueue = nil
    }
    
    /// 连接服务器
    /// - Parameters:
    ///   - ip: 服务器地址
    ///   - port: 服务器端口
    ///   - callback: 连接结果回调
    internal func connect(ip: String, port: UInt16) -> Bool {
        self.ip = ip
        self.port = port
        
        if tcpClient!.isConnected {
            tcpClient?.disconnect()
        }
        
        IMLog.debug(item: "CIMClient connect to \(ip):\(port)")
        do {
            try tcpClient?.connect(toHost: ip, onPort: port)
            return true
        } catch {
            IMLog.error(item: "connect error:\(error)")
        }
        return false
    }
    
    /// 断开连接
    internal func disconnect() {
        tcpClient?.disconnect()
    }
    
    /// send raw data to server,with out header
    /// - Parameter data: raw data
    func sendRaw(data: Data) {
        tcpClient?.write(data, withTimeout: -1, tag: 0)
    }
    
    /// 发送请求，包含协议头
    /// - Parameters:
    ///   - cmdId: 命令ID，见CIM_Def_CIMCmdID
    ///   - body: 数据部
    func send(cmdId: CIM_Def_CIMCmdID, body: Data) -> IMHeader {
        IMLog.debug(item: "send cmdId=\(cmdId.rawValue),dataLen=\(body.count)")
        
        // 自增序列号，服务器会返回设置的序号ID，此时可查找响应是属于哪一个请求，当然也可以通过CMDID来判断，可以实现广播的效果
        let tempSeq = seq
        if tempSeq > (UINT16_MAX - 1) {
            seq = 1
        } else {
            seq += 1
        }
        
        // 头部
        let header = IMHeader()
        header.setCommandId(cmdId: UInt16(cmdId.rawValue))
        header.setSeq(seq: tempSeq)
        header.setMsg(msg: body)
        
        // 发送
        // header.getBuffer()包含协议头和数据部
        sendRaw(data: header.getBuffer()!)
        
        // debug big-endian
        // let bytes = [UInt8](header.getBuffer()!)
        // print(bytes)
        return header
    }
    
    /// 发送请求，包含协议头，如果一定时间没有收到响应，会触发超时回调
    /// - Parameters:
    ///   - cmdId: 命令ID，见CIM_Def_CIMCmdID
    ///   - body: 数据
    ///   - timeout: 超时回调，10s
    func sendRequest(cmdId: CIM_Def_CIMCmdID, body: Data, timeout: IMResponseTimeoutCallback?) {
        let header = send(cmdId: cmdId, body: body)
        
        // 加入到待响应字典中
        let req = IMRequest(header: header, timeout: timeout)
        requestDic[header.seqNumber] = req
    }
    
    /// 注册委托 - 连接状态
    /// - Parameters:
    ///   - key: 唯一标识
    ///   - delegate: 委托对象
    func register(key: String, delegateStatus: IMClientDelegateStatus) {
        delegateDicStatus[key] = delegateStatus
    }
    
    /// 注册委托 - 业务数据
    /// - Parameters:
    ///   - key: 唯一标识
    ///   - delegate: 委托对象
    func register(key: String, delegateData: IMClientDelegateData) {
        delegateDicData[key] = delegateData
    }
    
    /// 取消委托  - 连接状态
    /// - Parameter key: 唯一标识
    func unregister(keyStatus: String) {
        delegateDicStatus.removeValue(forKey: keyStatus)
    }
    
    /// 取消委托 - 业务数据
    /// - Parameter key: 唯一标识
    func unregister(keyData: String) {
        delegateDicData.removeValue(forKey: keyData)
    }
    
    func _checkTimer(t: Timer) {
        let nowTimestamp = Int32(NSDate.now.timeIntervalSince1970)
        let maxWaitTimeLen = 10 // 最大等待时间
        
        var it = requestDic.makeIterator()
        var item = it.next()
        while item != nil {
            let req = (item?.value)!
            item = it.next()
            
            if abs(nowTimestamp - req.createTimestamp) > maxWaitTimeLen {
                IMLog.warn(item: "request timeout,seq=\(req.seq),cmdId=\(req.header.commandId) ")
                // next已经改变，放心移除
                requestDic.removeValue(forKey: req.seq)
                req.timeoutCallback?()
            }
        }
    }
}

// MARK: GCDAsyncSocketDelegate

extension IMClient {
    // connect
    func socket(_ sock: GCDAsyncSocket, didConnectToHost host: String, port: UInt16) {
        IMLog.debug(item: "IMClient successful connected to \(host):\(port)")
        
        // 回调 FIXME 非线程安全
        for item in delegateDicStatus {
            item.value.onConnected(host, port: port)
        }
        
        // 监听数据
        tcpClient?.readData(withTimeout: -1, tag: 0)
    }
    
    // receive data
    func socket(_ sock: GCDAsyncSocket, didRead data: Data, withTag tag: Int) {
        IMLog.debug(item: "IMClient socket receive data,len=\(data.count)")
        
        // 是否足够长，数据包完整
        if !IMHeader.isAvailable(data: data) {
            IMLog.error(item: "bad data!")
        } else {
            // 解析协议头
            let header = IMHeader()
            if !header.readHeader(data: data) {
                IMLog.error(item: "readHeader error!")
            } else {
                IMLog.debug(item: "parse IMHeader success,cmd=\(header.commandId),seq=\(header.seqNumber)")
                
                // 处理消息
                let bodyData = data[Int(kHeaderLen)..<data.count] // 去掉头部，只放裸数据
                
                // 回调 FIXME 非线程安全
                for item in delegateDicData {
                    item.value.onHandleData(header, bodyData)
                }
            }
        }
        
        // 监听数据
        tcpClient?.readData(withTimeout: -1, tag: 0)
    }
    
    // disconnect
    func socketDidDisconnect(_ sock: GCDAsyncSocket, withError err: Error?) {
        IMLog.debug(item: "socket disconnected,error:\(String(describing: err))")
        // 回调 FIXME 非线程安全
        for item in delegateDicStatus {
            item.value.onDisconnect(err)
        }
    }
}
