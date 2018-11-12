package model

import (
	"database/sql"
	"log"
)

type User struct {
	Id         int    `json:"id"`
	UserName   string `json:"username"`
	Password   string `json:"password"`
	Token      string `json:"token"`
	Roles      string `json:"roles"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

/**
 * 用户model层
 */
type UserModel struct {
}

/**
 * 创建用户
 */
func (userModel *UserModel) Create(db *sql.DB, username string, password string, token string) (u User, e error) {
	insertSql := "insert into mb_user (username , password , token) values (? , ? , ?)"
	insertRow, err := db.Exec(insertSql, username, password, token)
	if err != nil {
		log.Fatal(err)
		return User{}, err
	}

	insertRowId, err := insertRow.LastInsertId()
	if err != nil {
		return User{}, err
	}

	user, err := userModel.findById(db, int(insertRowId))
	if err != nil {
		return User{}, err
	}

	return user, nil
}

/**
 * 更新用户名
 */
func (userModel *UserModel) UpdateName(db *sql.DB, id int, username string) (u User, e error) {
	updateSql := "update mb_user set username = ? where id = ?"
	_, err := db.Exec(updateSql, username, id)
	if err != nil {
		log.Fatal(err)
		return User{}, err
	}
	return userModel.findById(db, id)
}

/**
 * 更新token
 */
func (userModel *UserModel) UpdateToken(db *sql.DB, username string, password string, token string) (User, error) {
	updateSql := "update mb_user set token = ? where username = ? and password = ?"
	_, err := db.Exec(updateSql, token, username, password)
	if err != nil {
		return User{}, err
	}
	return userModel.FindByUsernameAndPassword(db, username, password)
}

/**
 * 通过用户名密码查找用户
 */
func (UserModel *UserModel) FindByUsernameAndPassword(db *sql.DB, username string, password string) (User, error) {
	selectsql := "select * from mb_user where username = ? and password = ?"
	log.Println(selectsql)
	resultRows := db.QueryRow(selectsql, username, password)
	var user User
	err := resultRows.Scan(&user.Id, &user.UserName, &user.Password, &user.Token, &user.Roles, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (userModel *UserModel) findById(db *sql.DB, id int) (u User, e error) {
	querySql := "select * from mb_user where id = ?"
	resultRow := db.QueryRow(querySql, id)
	var user User
	err := resultRow.Scan(&user.Id, &user.UserName, &user.Password, &user.Token, &user.Roles, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
