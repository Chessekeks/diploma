package repo

import (
	"diploma/domain"
	"log"
)

const findUserByIDQuery = "SELECT * FROM users WHERE id = ?"
const findUserByEmailQuery = "SELECT * FROM users WHERE email = ?"
const readAllUsersQuery = "SELECT * FROM users"
const createUserQuery = "INSERT INTO users (name,surname,pass,email) VALUES (?,?,?,?) "
const updateUserQuery = "UPDATE users SET name = ?, surname = ?, pass = ?, email = ?, token = ?, sex = ?, residense = ?, photo = ? WHERE id = ?"
const deleteUserQuery = "DELETE FROM users WHERE id = ?"
const profileQuery = `SELECT specialists.univercity, ind.name, spec.name  
FROM specialists
LEFT JOIN industries ind ON specialists.indusry_id = ind.ID
LEFT JOIN specializations spec ON specialists.spec_id = spec.ID
WHERE user_id = (SELECT id FROM users WHERE token  = ?)`
const findByTokenQuery = "SELECT ID,name,email,surname,token, sex,residense,photo FROM users WHERE token = ?"

type UserRepo interface {
	CreateUser(u domain.User) error
	UpdateUser(id int, u domain.User) error
	FindByID(id int) (domain.User, error)
	FindByEmail(email string) (domain.User,error)
	FindByToken(token string) (domain.User,error)
	SpecialistProfile(token string) (domain.Specialist, error)
	FindUser(id int) (domain.User, error)
	AllUsers() ([]domain.User, error)
	DeleteUser(id int) error
}

//FindUserByID realization of UserStore interface
func (m Mysql) FindByID(id int) (domain.User, error) {
	var user domain.User
	strPre, err := m.DB.Prepare(findUserByIDQuery)
	if err != nil {
		log.Println("Prepare statement error: FindUserByID")
		return user, err
	}
	defer strPre.Close()

	err = strPre.QueryRow(id).Scan(&user.ID, &user.Name, &user.Email, &user.Surname, &user.Password, &user.Token,&user.Sex,&user.Residense, &user.Photo)
	if err != nil {
		log.Println("QueryRow Scan error: FindUserByID")
		return user,err
	}
	return user,nil
}

//FindUserByEmail realization of UserStore interface
func (m Mysql) FindByEmail(email string) (domain.User,error) {
	var user domain.User
	strPre, err := m.DB.Prepare(findUserByEmailQuery)
	if err != nil {
		log.Println("Prepare statement error: FindUserByName")
		return user,err
	}
	defer strPre.Close()

	err = strPre.QueryRow(email).Scan(&user.ID, &user.Name, &user.Email, &user.Surname, &user.Password, &user.Token,&user.Sex,&user.Residense,&user.Photo)
	if err != nil {
		log.Println("QueryRow Scan error: FindUserByName")
		return user,err
	}
	return user,nil
}

//ReadAllUsers realization of UserStore interface
func (m Mysql) AllUsers() ([]domain.User,error) {
	var result []domain.User
	strPre, err := m.DB.Prepare(readAllUsersQuery)
	if err != nil {
		log.Println("Prepare statement error: ReadAllUsers")
		return nil,err
	}
	defer strPre.Close()

	rows, err := strPre.Query()
	if err != nil {
		log.Println("Query error: ReadAllUsers")
		return nil,err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Surname, &user.Password, &user.Token,&user.Sex,&user.Residense,&user.Photo)
		if err != nil {
			log.Println("Scan error: ReadAllUsers")
			return nil,err
		}
		result = append(result, user)
	}
	return result,nil
}

//CreateUser realization of UserStore interface
func (m Mysql) CreateUser(u domain.User) error {
	strPre, err := m.DB.Prepare(createUserQuery)
	if err != nil {
		log.Println("Prepare statement error: CreateUser")
		return err
	}
	defer strPre.Close()

	_, err = strPre.Exec(u.Name, u.Surname, u.Password, u.Email)
	if err != nil {
		log.Println("Exec error: CreateUser")
		return err
	}
	return nil
}

//UpdateUser realization of UserStore interface
func (m Mysql) UpdateUser(id int, u domain.User) error {
	strPre, err := m.DB.Prepare(updateUserQuery)
	if err != nil {
		log.Println("Prepare statement error: UpdateUser")
		return err
	}
	defer strPre.Close()

	_, err = strPre.Exec(u.Name, u.Surname, u.Password, u.Email, u.Token, u.Sex, u.Residense, u.Photo, id)
	if err != nil {
		log.Println("Exec error: UpdateUser")
		return err
	}
	return nil
}

//DeleteUser realization of UserStore interface
func (m Mysql) DeleteUser(id int) error {
	strPre, err := m.DB.Prepare(deleteUserQuery)
	if err != nil {
		log.Println("Prepare statement error: DeleteUser")
		return err
	}
	defer strPre.Close()

	_, err = strPre.Exec(id)
	if err != nil {
		log.Println("Exec error: DeleteUser")
		return err
	}
	return nil
}

func (m Mysql) SpecialistProfile(token string) (domain.Specialist, error) {
	var spec domain.Specialist
	strPre, err := m.DB.Prepare(profileQuery)
	if err != nil {
		log.Println("Prepare error: SpecialistProfile")
		return spec,err
	}
	defer strPre.Close()

	err = strPre.QueryRow(token).Scan(&spec.University,&spec.Subject,&spec.Specialization)
	if err != nil {
		log.Println("QueryRow Scan error: SpecialistProfile")
		return spec,err
	}

	return spec,nil
}

func(m Mysql) FindUser(id int) (domain.User, error) {
	var user domain.User

	return user,nil
}

func (m Mysql) FindByToken(token string) (domain.User,error) {
	var user domain.User
	strPre,err := m.DB.Prepare(findByTokenQuery)
	if err != nil {
		log.Println("Prepare error: FindByToken")
		return user,err
	}
	defer strPre.Close()

	err = strPre.QueryRow(token).Scan(&user.ID,&user.Name,&user.Email,&user.Surname,&user.Token, &user.Sex, &user.Residense, &user.Photo)
	if err != nil {
		log.Println("Exec error: FindByToken")
		return user,err
	}
	return user,nil
}