package payload

type UserRegisterPayload struct {
	Name  string `json:"name" validate:"required"`
	Age   int    `json:"age"`
	Email string `json:"email" validate:"required,email"`
}
