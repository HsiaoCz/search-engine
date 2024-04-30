package types

type Loginform struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type User struct {
	ID       string `bson:"_id" json:"id"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
}
