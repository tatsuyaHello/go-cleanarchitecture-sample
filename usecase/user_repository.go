package usecase

import "github.com/tatsuyaHello/go-cleanarchitecture-sample/tree/master/app/domain"

type UserRepository interface {
	Store(domain.User) (int, error)
	FindById(int) (domain.User, error)
	FindAll() (domain.Users, error)
}
