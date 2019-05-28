package domain

//User type defines common user of application
type User struct {
	ID       int    `json:"ID"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	Residense string `json:"residense"`
	Sex string `json:"sex"`
	Photo string `json:"photo"`
}

type Specialist struct {
	University string `json:"university"`
	Subject string `json:"subject"`
	Specialization string `json:"specialization"`
}

type UserResponse struct {
	 User User `json:"user"`
	 Specialist Specialist `json:"specialist"`
}
//UserStore interface provides methods to work with databases for User entity
type UserService interface {
	FindByID(id int) (User, error)
	FindByToken(token string) (User,error)
	FindByEmail(email string) (User, error)
	FindUser(id int) (User, error)
	SpecialistProfile(token string) (Specialist, error)
	AllUsers() ([]User,error)
	CreateUser(u User) error
	UpdateUser(id int, u User) error
	DeleteUser(id int) error
}
