package req

type Register struct {
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
	PasswordRepeat string `json:"passwordRepeat" validate:"required"`
}
