package request

type LoginDto struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
