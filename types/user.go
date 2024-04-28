package types

type Loginform struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
