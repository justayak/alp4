package ptp

// ~~~~~~~~~~~~~~~~~
// U S E R
// ~~~~~~~~~~~~~~~~~
type User struct {
	name string 
	others []User
}

// ctor
func NewUser(name string ) *User {
	return &User{
		name, 
		make( []User, 0),
	}
}