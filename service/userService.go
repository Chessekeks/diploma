package service

import (
	"diploma/domain"
	"diploma/repo"
)

//UserService is service of User's scenarios
type UserService struct {
	rep repo.UserRepo
}

func NewUserService(conn repo.Mysql) UserService {
	return UserService{
		rep:conn,
	}
}

func (uServ UserService) CreateUser(u domain.User) error {
	return uServ.rep.CreateUser(u)
}

//SingIn is realization of UserService interface
func (uServ UserService) FindByEmail(email string) (domain.User,error) {
	return uServ.rep.FindByEmail(email)
}

func (uServ UserService) AllUsers() ([]domain.User,error) {
	return uServ.rep.AllUsers()
}

func (uServ UserService) UpdateUser(id int, u domain.User) error {
	return uServ.rep.UpdateUser(id,u)
}
func (uServ UserService) DeleteUser(id int) error {
	return uServ.rep.DeleteUser(id)
}

func (uServ UserService) FindByToken(token string) (domain.User,error) {
	return uServ.rep.FindByToken(token)
}

func (uServ UserService) FindUser(id int) (domain.User, error) {
	 return uServ.rep.FindUser(id)
}

func(uServ UserService) SpecialistProfile(token string) (domain.Specialist, error) {
	return uServ.rep.SpecialistProfile(token)
}