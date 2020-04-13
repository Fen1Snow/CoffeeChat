//
//  IMChatContentViewController.swift
//  Coffchat
//
//  Created by xuyingchun on 2020/3/28.
//  Copyright © 2020 Xuyingchun Inc. All rights reserved.
//

import SnapKit
import UIKit

let kSendMsgBarHeight = 50
let kStartAnimateOffset = CGFloat(30)

/// 聊天页面
class IMChatContentViewController: UIViewController, UITableViewDataSource, UIScrollViewDelegate, UITableViewDelegate,
    IMChatSendMsgBarDelegate, IMChatManagerDelegate {
    @IBOutlet var msgTabView: UITableView!
    @IBOutlet var refreshView: UIView! // tabview顶部视图，背景色透明，里面放了一个UIActivityIndicatorView
    @IBOutlet var indicatorView: UIActivityIndicatorView! // 一个系统自带的菊花，可以显示动画

    // 消息发送框
    var sendMsgBar: IMChatSendMsgBar?
    // 消息发送框的下边距
    var sendMsgBarBottomConstranit: Constraint?

    // 会话信息
    var session: SessionModel
    var msgList: [IMMessage] = []

    // 是否正在刷新历史记录中
    var isRefreshMsgLst: Bool = false

    init(session: SessionModel) {
        self.session = session

        super.init(nibName: nil, bundle: nil)
    }

    required init?(coder: NSCoder) {
        fatalError()
    }

    override func viewDidLoad() {
        super.viewDidLoad()
        msgTabView.delegate = self
        msgTabView.dataSource = self

        // 注册自定义Cell
        msgTabView.register(UINib(nibName: "IMMessageTextCell", bundle: nil), forCellReuseIdentifier: "IMMessageTextCell")
        // 不显示分割线
        msgTabView.separatorStyle = .none
        msgTabView.backgroundColor = IMUIResource.chatBackground
        // msgTabView.tableHeaderView = refreshView // 顶部显示
        msgTabView.estimatedRowHeight = 0
        msgTabView.estimatedSectionHeaderHeight = 0
        msgTabView.estimatedSectionFooterHeight = 0

        // 添加子控件
        setupSubviews()

        // 注册键盘出现事件 #selector 绑定object-c
        NotificationCenter.default.addObserver(self, selector: #selector(keyboardShow(note:)), name: UIResponder.keyboardWillShowNotification, object: nil)
        NotificationCenter.default.addObserver(self, selector: #selector(keyboardHidden(note:)), name: UIResponder.keyboardWillHideNotification, object: nil)
        NotificationCenter.default.addObserver(self, selector: #selector(keyboardDidShow(note:)), name: UIResponder.keyboardDidShowNotification, object: nil)

        // 查询历史聊天记录
        queryMsgList(endMsgId: 0)

        // 注册消息委托
        IMManager.singleton.chatManager.register(key: "IMChatContentViewController", delegate: self)
    }

    deinit {
        // 取消注册
        NotificationCenter.default.removeObserver(self)
        IMManager.singleton.chatManager.unregister(key: "IMChatContentViewController")
    }

    // 添加子控件
    func setupSubviews() {
        // 聊天框
        sendMsgBar = getNibInstance(IMChatSendMsgBar.self)
        sendMsgBar!.delegate = self
        // 先添加到父视图，在设置约束，否则会崩溃
        view.addSubview(sendMsgBar!)
        sendMsgBar!.snp.makeConstraints { [weak self] (make) -> Void in
            guard let strongSelf = self else { return }
            // 左 右 底部 高度
            make.left.equalTo(strongSelf.view.snp.left)
            make.right.equalTo(strongSelf.view.snp.right)
            // 获取下边距的引用，便于处理键盘显示时，把输入框挤上去
            strongSelf.sendMsgBarBottomConstranit = make.bottom.equalTo(strongSelf.view.snp.bottom).constraint
            make.height.equalTo(kSendMsgBarHeight)
        }

        // 键盘出现时，改变tabview的下边距，以显示全部内容
        msgTabView.snp.makeConstraints { (make) -> Void in
            // 上 左 右 全屏绑定
            make.top.left.right.equalTo(self.view)
            // 下边距是消息输入框的上边距，这里很关键。
            make.bottom.equalTo(self.sendMsgBar!.snp.top)
        }
    }

    // 初始化控件的实例，根据nib文件
    func getNibInstance<T>(_ aClass: T.Type) -> T {
        let name = String(describing: aClass)
        if Bundle.main.path(forResource: name, ofType: "nib") != nil {
            return UINib(nibName: name, bundle: nil).instantiate(withOwner: nil, options: nil)[0] as! T
        } else {
            fatalError("\(String(describing: aClass)) nib is not exist")
        }
    }

    /// 刷新历史记录
    @objc
    func refreshMsgList() {
        if !isRefreshMsgLst {
            indicatorView.startAnimating()
            
            if msgList.count > 0{
                queryMsgList(endMsgId: msgList[msgList.count - 1].serverMsgId!)
            }else{
                queryMsgList(endMsgId: 0)
            }
        }
    }

    func queryMsgList(endMsgId: UInt64) {
        isRefreshMsgLst = true
        
        let sId = session.rectSession.session.sessionId
        let sType = session.rectSession.session.sessionType
        // var endMsgId = session.rectSession.latestMsg.serverMsgId
        let limitCount = 20

        // 查询历史消息
        IMManager.singleton.conversationManager.queryMsgList(sessionId: sId, sessionType: sType, endMsgId: endMsgId, limitCount: limitCount, callback: { rsp in
            // print("success query msg list\(rsp)")
            for item in rsp.msgList {
                let msg = IMMessage(clientId: item.clientMsgID, sessionType: item.sessionType, fromId: item.fromUserID, toId: item.toSessionID, time: item.createTime, msgType: item.msgType, data: String(data: item.msgData, encoding: .utf8)!)
                msg.serverMsgId = item.serverMsgID
                msg.msgResCode = item.msgResCode
                msg.msgFeature = item.msgFeature
                msg.senderClientType = item.senderClientType
                msg.attach = item.attach
                msg.msgStatus = item.msgStatus
                self.msgList.append(msg)
            }

            DispatchQueue.main.async {
                // 如果超过20个，可能有历史聊天记录，显示刷新按钮
                self.isRefreshMsgLst = false
                self.indicatorView.stopAnimating()
                if self.msgList.count >= 20 {
                    if self.msgTabView.tableHeaderView == nil {
                        self.msgTabView.tableHeaderView = self.refreshView
                    }
                } else {
                    self.msgTabView.tableHeaderView = nil // 隐藏刷新按钮
                }
                self.msgTabView.reloadData()

                // 滚动到底部，不要动画
                self.msgTabView.scrollToBottom(animated: false)
            }
        }, timeout: {
            print("timeout")
        })
    }
}

// MARK: UITableViewDelegate

extension IMChatContentViewController {
    // 禁止选中效果
    func tableView(_ tableView: UITableView, didSelectRowAt indexPath: IndexPath) {
        tableView.deselectRow(at: indexPath, animated: true)
    }
}

// MARK: UITableViewDataSource

extension IMChatContentViewController {
    func numberOfSections(in tableView: UITableView) -> Int {
        return 1
    }

    func tableView(_ tableView: UITableView, heightForRowAt indexPath: IndexPath) -> CGFloat {
        let msg = msgList[indexPath.row]
        // 动态计算每一行文本的高度
//        if msg.msgType == .kCimMsgTypeText {
//            return IMMessageTextCell.getTextHeight(text: msg.msgData)
//        }
//        return 100
        let height = IMMessageTextCell.getCellHeight(text: msg.msgData)
        return height
    }

    func tableView(_ tableView: UITableView, cellForRowAt indexPath: IndexPath) -> UITableViewCell {
        let msg = msgList[indexPath.row]
        let cellId = String(describing: IMMessageTextCell.self)

//        switch msg.msgType {
//        case .kCimMsgTypeText:
//
//        default:
//            print("unknown msg type")
//        }

        var cell = tableView.dequeueReusableCell(withIdentifier: cellId) as? IMMessageTextCell
        if cell == nil {
            cell = IMMessageTextCell(style: .default, reuseIdentifier: cellId)
            // 选中效果，不需要
            cell?.selectionStyle = .none
        }
        cell!.setContent(message: msg)
        return cell!
    }

    func tableView(_ tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
        return msgList.count
    }
}

// MARK: UIScrollViewDelegate

extension IMChatContentViewController {
    // 滚动过程中
    func scrollViewDidScroll(_ scrollView: UIScrollView) {
        if scrollView.contentOffset.y < kStartAnimateOffset {
            IMLog.info(item: "pull to refresh msg list")
            refreshMsgList()
        }
    }

    // 停止拖拽滚动条
    func scrollViewDidEndDragging(_ scrollView: UIScrollView, willDecelerate decelerate: Bool) {
        if scrollView.contentOffset.y - scrollView.contentInset.top < kStartAnimateOffset {
            IMLog.info(item: "pull to refresh msg list")
            refreshMsgList()
        }
    }
}

// MARK: IMChatSendMsgBarDelegate

extension IMChatContentViewController {
    func onSendMsg(_ text: String) {
        if text.isEmpty {
            let alert = UIAlertController(title: "提示", message: "请输入内容", preferredStyle: .alert)
            let action = UIAlertAction(title: "ok", style: .cancel, handler: nil)
            alert.addAction(action)
            present(alert, animated: true, completion: nil)
            return
        }

        sendMsgBar?.textInput.text = ""

        let sessionId = session.rectSession.session.sessionId
        let sessionType = session.rectSession.session.sessionType

        // 追加
        let fromId = IMManager.singleton.loginManager.userId
        let msg = IMMessage(clientId: "", sessionType: sessionType, fromId: fromId!, toId: sessionId, time: 0, msgType: .kCimMsgTypeText, data: text)
        appendMsg(msg: msg)

        // 发消息
        IMManager.singleton.chatManager.sendTextMessage(toSessionId: sessionId, sessionType: sessionType, text: text)
    }

    func appendMsg(msg: IMMessage) {
        // 追加
        msgList.append(msg)

        // 刷新tabview
        let index = IndexPath(row: msgList.count - 1, section: 0)
        msgTabView.insertRowsAtBottom([index])
    }
}

// MARK: IMChatManagerDelegate

extension IMChatContentViewController {
    // 发送消息结果
    func sendMessageResult(msg: IMMessage, result: IMSendResult, _ code: CIM_Def_CIMResCode) {
        // 自己的消息
        if msg.toSessionId == session.rectSession.session.sessionId {
            IMLog.info(item: "消息发送结果，msgId:\(msg.clientMsgId),result:\(result),sessionId:\(msg.toSessionId)")
        }
    }

    // 收到一条消息
    func onRecvMessage(msg: IMMessage) {
        if msg.toSessionId == session.rectSession.session.sessionId {
            DispatchQueue.main.async {
                self.appendMsg(msg: msg)
            }
        }
    }
}
