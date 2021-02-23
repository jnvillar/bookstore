package login

type Request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *Request) GetUserName() string {
	if r == nil {
		return ""
	}
	return r.Username
}

func (r *Request) GetPassword() string {
	if r == nil {
		return ""
	}
	return r.Password
}
