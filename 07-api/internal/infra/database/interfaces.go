package database

import "github.com/hugovallada/go-expert/api/internal/entity"

type UserModel interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type ProductModel interface {
	Create (product *entity.Product) error
	FindAll(page, limit int,sort string) ([]entity.Product, error)
	FindById (id string) (*entity.Product, error)
	Update(product *entity.Product) error
	Delete(id string) error 
}
