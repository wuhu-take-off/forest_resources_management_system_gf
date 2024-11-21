// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"forest_resources_management_system_gf/internal/dao/internal"
)

// internalIdentityDao is internal type for wrapping internal DAO implements.
type internalIdentityDao = *internal.IdentityDao

// identityDao is the data access object for table identity.
// You can define custom methods on it to extend its functionality as you wish.
type identityDao struct {
	internalIdentityDao
}

var (
	// Identity is globally public accessible object for table identity operations.
	Identity = identityDao{
		internal.NewIdentityDao(),
	}
)

// Fill with you ideas below.
