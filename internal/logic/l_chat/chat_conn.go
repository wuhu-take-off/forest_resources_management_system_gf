package l_chat

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"sync"
)

type ChatConn struct {
	sync.RWMutex
	ChatMap map[int]*websocket.Conn
}

func NewChatConn() *ChatConn {
	if defaultChatStruct == nil {
		defaultChatStruct = &ChatConn{
			RWMutex: sync.RWMutex{},
			ChatMap: make(map[int]*websocket.Conn),
		}
	}
	return defaultChatStruct
}

var defaultChatStruct *ChatConn

// AddConn adds a new connection to the chat map
func (c *ChatConn) AddConn(userId int, conn *websocket.Conn) {
	c.Lock()
	defer c.Unlock()
	g.Log().Printf(nil, "add user %d to chat", userId)
	//删除之前的连接
	c.Remove(userId)
	c.ChatMap[userId] = conn
	go c.monitorChat(userId, conn)
}

// MonitorChat 监控指定用户的连接
func (c *ChatConn) monitorChat(userId int, conn *websocket.Conn) {
	for {
		// 不处理任何消息，只等待连接断开
		_, _, err := conn.ReadMessage()
		if err != nil {
			c.Remove(userId)
			break
		}
	}
}

// Remove  删除指定用户的连接
func (c *ChatConn) Remove(userId int, isSafe ...bool) {
	if len(isSafe) > 0 && isSafe[0] {
		c.Lock()
		defer c.Unlock()
	}
	// close connection
	if conn, ok := c.ChatMap[userId]; ok {
		g.Log().Printf(nil, "remove user %d from chat", userId)
		conn.Close()
	}
	// delete from chat map
	delete(c.ChatMap, userId)
}

// SendMessage 发送消息给指定用户
func (c *ChatConn) SendMessage(msg any, userId ...int) {
	for _, u := range userId {
		conn, ok := c.ChatMap[u]
		if ok {
			c.SendMessageToSingle(msg, u, conn, true)
		}
	}
}

// SendMessageToSingle 发送消息给单个用户
func (c *ChatConn) SendMessageToSingle(msg any, userId int, conn *websocket.Conn, isSafe ...bool) {
	if len(isSafe) > 0 && isSafe[0] {
		c.Lock()
		defer c.Unlock()
	}
	if err := conn.WriteJSON(msg); err != nil {
		c.Remove(userId, true)
	}
}
