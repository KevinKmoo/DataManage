package model

import (
	"database/sql"
)

type Role struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"UpdateTime"`
}

type RolesModel struct {
}

func NewRolesModel() RolesModel {
	return RolesModel{}
}

func (rolesModel *RolesModel) FindById(db *sql.DB, id int) (Role, error) {
	selectSql := "select * from mb_roles where id = ?"
	resultRow := db.QueryRow(selectSql, id)
	var role Role
	err := resultRow.Scan(&role.Id, &role.Name, &role.Description, &role.CreateTime, &role.UpdateTime)

	if err != nil {
		return Role{}, err
	}
	return role, err
}
