package requests

type UserRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username    string `json:"username"`
	Address  string `json:"address"`
	Status  string `json:"status"`
	Token  string `json:"token"`
	Password string `json:"password"`
}
