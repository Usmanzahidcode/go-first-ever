package user

import "fmt"

type User struct {
	Name string
	Role string
}

// Common behavior for all users
func (u User) Login() bool {
	fmt.Println(u.Name, "logged in as", u.Role)
	return true
}

func (u User) GetRole() string {
	return u.Role
}

func (u User) CanAccess(resource string) bool {
	return false
}

// --------------------

type Admin struct {
	User
	Permission string
}

// Admin overrides access behavior
func (a Admin) CanAccess(resource string) bool {
	return true
}

func (a Admin) GetPermissions() string {
	return a.Permission
}

func (a Admin) ManageUsers() {
	fmt.Println(a.Name, "is managing users")
}
