// Code generated by goctl. DO NOT EDIT.
package types

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Code     string `json:"code"`
}

type RegisterResponse struct {
	Msg string `json:"msg"`
}

type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type UserDetailRequest struct {
	UUID string `path:"uuid"`
}

type UserDetailResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type EmailCodeRequest struct {
	Email string `json:"email"`
}

type EmailCodeResponse struct {
	Msg string `json:"msg"`
}
