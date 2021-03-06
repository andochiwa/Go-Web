package model

import (
	"fmt"
	"goweb/02-db/utils"
)

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Email    string `db:"email"`
}

func (user *User) Save() error {
	sql := "insert into users(username, email) values(?, ?)"
	_, err := utils.Db.Exec(sql, user.Username, user.Email)
	if err != nil {
		fmt.Println("error by User.save()")
		return err
	}
	return nil
}

func (user *User) GetById() error {
	sql := "select id, username, email from users where id = ?"

	row := utils.Db.QueryRow(sql, user.Id)

	err := row.Scan(&user.Id, &user.Username, &user.Email)
	if err != nil {
		fmt.Println("error by User.getById()")
		return err
	}
	return nil
}

func (user *User) GetAll() ([]User, error) {
	sql := "select id, username, email from users"

	rows, err := utils.Db.Query(sql)
	if err != nil {
		fmt.Println("error by User.GetAll() 1")
		return nil, err
	}
	var users []User
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email)
		if err != nil {
			fmt.Println("error by User.GetAll() 2")
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
