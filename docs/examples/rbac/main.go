package main

import (
	"sync"

	"github.com/casbin/casbin/v3/rbac"
)

// NewRoleManager is the constructor for creating an instance of the
// default RoleManager implementation.
func NewRoleManager(maxHierarchyLevel int) rbac.RoleManager {
	rm := RoleManager{}
	rm.allRoles = &sync.Map{}
	rm.maxHierarchyLevel = maxHierarchyLevel
	rm.hasPattern = false

	return &rm
}
