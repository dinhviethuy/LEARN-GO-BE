package dto

type CreateUserBody struct {
	Name     string `json:"name" validate:"required,min=1"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Age      int    `json:"age" validate:"required,min=1"`
}

type UpdateUserBody struct {
	Name  *string `json:"name" validate:"omitempty,min=1"`
	Email *string `json:"email" validate:"omitempty,email"`
	Age   *int    `json:"age" validate:"omitempty,min=1"`
}

type GetParamUser struct {
	UserID string `json:"userid" validate:"required"`
}