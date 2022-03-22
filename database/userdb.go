package database

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres",
		"user=philip dbname=gwp password=marcus01 sslmode=disable")
	if err != nil {
		fmt.Print("WOWOOW \n")
		panic(err)
	}
}

func (user User) Create() (err error) {
	statement := "insert into users (email, name, password) values ($1, $2, $3) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Email, user.Name, user.Password).Scan(&user.Id)
	user.CreatedAt = time.Now()
	return
}

func GetUser(email string) (user User, err error){
	user = User{}
	err = Db.QueryRow("select id, name, email, password from users where email = $1",
		email).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	return
}

func GetAllUsers() (users []User, err error) {
	rows, err := Db.Query("select id, name, email, password from users limit $1", 10)
	if err != nil {
		return
	}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return
		}
		users = append(users, user)
	}
	rows.Close()
	return
}

type User struct {
	Id int
	Email string
	Name string
	Password string
	CreatedAt time.Time
}
