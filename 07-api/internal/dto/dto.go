package dto

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BearerToken struct {
	AccessToken string `json:"access_token"`
}
