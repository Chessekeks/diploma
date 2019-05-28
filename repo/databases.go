package repo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//Mysql defines realization for mysql database
type Mysql struct {
	DB *sql.DB
}

//Postgres defines realization for postgres database
type Postgres struct {
	DB *sql.DB
}

//OpenMysql create connect to database
func OpenMysql() Mysql {
	db, err := sql.Open("mysql", "root:root@/diploma")
	if err != nil {
		log.Println("Can not create a connection pool: Mysql\n", err)
		return Mysql{}
	}
	return Mysql{
		DB: db,
	}
}

/*
//OpenDB create connect to database
func (m *Postgres) OpenDB() Postgres {

}
*/
