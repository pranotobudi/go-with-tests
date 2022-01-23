package repository

import "github.com/pranotobudi/go-with-tests/mock/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
