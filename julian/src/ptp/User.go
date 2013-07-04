package ptp

// ~~~~~~~~~~~~~~~~~
// U S E R
// ~~~~~~~~~~~~~~~~~
type User struct {
	Name string
	Other string
}

// ctor
func NewUser(name, other string ) *User {
	return &User{
		name, 
		other,
	}
}