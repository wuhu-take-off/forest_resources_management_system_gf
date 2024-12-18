// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"forest_resources_management_system_gf/internal/dao/internal"
)

// internalPrivateChatDao is internal type for wrapping internal DAO implements.
type internalPrivateChatDao = *internal.PrivateChatDao

// privateChatDao is the data access object for table private_chat.
// You can define custom methods on it to extend its functionality as you wish.
type privateChatDao struct {
	internalPrivateChatDao
}

var (
	// PrivateChat is globally public accessible object for table private_chat operations.
	PrivateChat = privateChatDao{
		internal.NewPrivateChatDao(),
	}
)

// Fill with you ideas below.
