package users

import (
	"encoding/json"
	"time"

	"github.com/chnsz/golangsdk"
	"github.com/chnsz/golangsdk/pagination"
)

// User represents a User in the OpenStack Identity Service.
type User struct {
	// DefaultProjectID is the ID of the default project of the user.
	DefaultProjectID string `json:"default_project_id"`

	// Description is a description of the user.
	Description string `json:"description"`

	// DomainID is the domain ID the user belongs to.
	DomainID string `json:"domain_id"`

	// Enabled is whether or not the user is enabled.
	Enabled bool `json:"enabled"`

	// ID is the unique ID of the user.
	ID string `json:"id"`

	// Links contains referencing links to the user.
	Links map[string]interface{} `json:"links"`

	// Name is the name of the user.
	Name string `json:"name"`

	// PasswordExpiresAt is the timestamp when the user's password expires.
	PasswordExpiresAt time.Time `json:"-"`
}

func (r *User) UnmarshalJSON(b []byte) error {
	type tmp User
	var s struct {
		tmp
		PasswordExpiresAt golangsdk.JSONRFC3339MilliNoZ `json:"password_expires_at"`
	}
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*r = User(s.tmp)

	r.PasswordExpiresAt = time.Time(s.PasswordExpiresAt)

	return err
}

type userResult struct {
	golangsdk.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a User.
type GetResult struct {
	userResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a User.
type CreateResult struct {
	userResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a User.
type UpdateResult struct {
	userResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	golangsdk.ErrResult
}

// UserPage is a single page of User results.
type UserPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a UserPage contains any results.
func (r UserPage) IsEmpty() (bool, error) {
	users, err := ExtractUsers(r)
	return len(users) == 0, err
}

// NextPageURL extracts the "next" link from the links section of the result.
func (r UserPage) NextPageURL() (string, error) {
	var s struct {
		Links struct {
			Next     string `json:"next"`
			Previous string `json:"previous"`
		} `json:"links"`
	}
	err := r.ExtractInto(&s)
	if err != nil {
		return "", err
	}
	return s.Links.Next, err
}

// ExtractUsers returns a slice of Users contained in a single page of results.
func ExtractUsers(r pagination.Page) ([]User, error) {
	var s struct {
		Users []User `json:"users"`
	}
	err := (r.(UserPage)).ExtractInto(&s)
	return s.Users, err
}

// Extract interprets any user results as a User.
func (r userResult) Extract() (*User, error) {
	var s struct {
		User *User `json:"user"`
	}
	err := r.ExtractInto(&s)
	return s.User, err
}

type AddMembershipResult struct {
	golangsdk.ErrResult
}
