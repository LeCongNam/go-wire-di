package services

import "test/repo"

type UserService struct {
	userRepo *repo.UserRepo
}

func NewUserService(userRepo *repo.UserRepo) *UserService { // 🛠 Dependency Injection
	return &UserService{userRepo: userRepo}
}

func (us *UserService) GetUser() string {
	return us.userRepo.GetList()
}
