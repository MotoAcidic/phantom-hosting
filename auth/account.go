package auth

type Account struct {
	ID       int
	Username string `json:"username",storm:"username"`
	Password string `json:"password",storm:"password"`
}
