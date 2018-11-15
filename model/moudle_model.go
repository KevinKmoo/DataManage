package model

import (
	"database/sql"
)

type Module struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   int    `json:"projectId"`
	VersionId   int    `json:"versionId"`
	Status      int    `json:"status"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type ModuleModel struct{}

func NewModuleModel() ModuleModel {
	return ModuleModel{}
}

func (model *ModuleModel) CrateModule(db *sql.DB, name string, description string, projectId int, versionId int) (Module, error) {
	insertSql := "insert into mb_module (name , description , project_id , version_id) values (? , ? , ? , ?)"
	resultRow, err := db.Exec(insertSql, name, description, projectId, versionId)
	if err != nil {
		return Module{}, err
	}
	lastInsertId, err := resultRow.LastInsertId()
	if err != nil {
		return Module{}, err
	}
	return model.findById(db, int(lastInsertId))
}

func (model *ModuleModel) findById(db *sql.DB, id int) (Module, error) {
	selectSql := "select * from mb_module where id = ?"
	resultRow := db.QueryRow(selectSql, id)
	var module Module
	err := resultRow.Scan(&module.Id, &module.Name, &module.Description, &module.ProjectId, &module.VersionId, &module.Status, &module.CreateTime, &module.UpdateTime)
	if err != nil {
		return Module{}, err
	}
	return module, nil
}
