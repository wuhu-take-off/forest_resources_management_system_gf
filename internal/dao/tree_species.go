// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"forest_resources_management_system_gf/internal/dao/internal"
)

// internalTreeSpeciesDao is internal type for wrapping internal DAO implements.
type internalTreeSpeciesDao = *internal.TreeSpeciesDao

// treeSpeciesDao is the data access object for table tree_species.
// You can define custom methods on it to extend its functionality as you wish.
type treeSpeciesDao struct {
	internalTreeSpeciesDao
}

var (
	// TreeSpecies is globally public accessible object for table tree_species operations.
	TreeSpecies = treeSpeciesDao{
		internal.NewTreeSpeciesDao(),
	}
)

// Fill with you ideas below.
