package l_chat

import (
	"forest_resources_management_system_gf/api/chat"
)

type defaultChat struct {
	chatConn *ChatConn
}

type IChatLogic interface {
	chat.IChatChar_v1
}

func NewChatLogic() IChatLogic {
	return &defaultChat{
		chatConn: NewChatConn(),
	}
}

func (d *defaultChat) GetChatS() *ChatConn {
	return d.chatConn
}
