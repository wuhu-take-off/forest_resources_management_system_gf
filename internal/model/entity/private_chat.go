// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// PrivateChat is the golang structure for table private_chat.
type PrivateChat struct {
	PrivateChatId         int    `json:"privateChatId"         orm:"private_chat_id"          ` // 私聊id
	PrivateChatSendId     int    `json:"privateChatSendId"     orm:"private_chat_send_id"     ` // 发送方id
	PrivateChatReceiverId int    `json:"privateChatReceiverId" orm:"private_chat_receiver_id" ` // 接收方id
	PrivateChatCreateTime int64  `json:"privateChatCreateTime" orm:"private_chat_create_time" ` // 发送时间
	PrivateChatMessage    string `json:"privateChatMessage"    orm:"private_chat_message"     ` //
}
