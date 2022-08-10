package users

import (
	"encoding/json"

	"github.com/chnsz/golangsdk/pagination"
)

// ListResp is the structure that represents the API response of 'List' method.
type ListResp struct {
	// Total number of query results.
	TotalCount int `json:"total_count"`
	// List of users to query.
	Users string `json:"users"`
}

// UserResp is the structure that represents the detail of the database user.
type UserResp struct {
	// Whether role is built-in.
	IsBuiltin bool `json:"isBuiltin"`
	// Role name.
	Name string `json:"user"`
	// Database name.
	DbName string `json:"db"`
	// The list of privileges inherited by the newly created role.
	Privileges []Privilege `json:"privileges"`
	// The list of privileges inherited by the newly created role, includes all privileges inherited by inherited roles.
	InheritedPrivileges []Privilege `json:"inheritedPrivileges"`
	// The list of roles inherited by the newly created role.
	Roles []RoleDetail `json:"roles"`
	// The list of roles inherited by the newly created role, includes all roles inherited by inherited roles.
	InheritedRoles []RoleDetail `json:"inheritedRoles"`
}

// Privilege is the structure that represents the privilege detail for database.
type Privilege struct {
	// The details of the resource to which the privilege belongs.
	Resource Resource `json:"resource"`
	// The operation permission list.
	Actions []string `json:"actions"`
}

// Resource is the structure that represents the database details to which the role and user belongs.
type Resource struct {
	// The database to which the privilege belongs.
	Collection string `json:"collection"`
	// The database name.
	DbName string `json:"db"`
}

// RoleDetail is the structure that represents the inherited role details.
type RoleDetail struct {
	// Role name.
	Name string `json:"role"`
	// The database name to which the role belongs.
	DbName string `json:"db"`
}

// UserPage is a single page maximum result representing a query by offset page.
type UserPage struct {
	pagination.OffsetPageBase
}

// IsEmpty checks whether a RolePage struct is empty.
func (p UserPage) IsEmpty() (bool, error) {
	arr, err := ExtractUsers(p)
	return len(arr) == 0, err
}

// ExtractUsers is a method to extract the list of database role for DDS service.
func ExtractUsers(p pagination.Page) ([]UserResp, error) {
	var r ListResp
	err := (p.(UserPage)).ExtractInto(&r)
	if err != nil {
		return nil, err
	}
	var ur []UserResp
	if r.Users != "" {
		err = json.Unmarshal([]byte(r.Users), &ur)
	}
	return ur, err
}
