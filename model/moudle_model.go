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

type ModuleFindByProjectId struct {
	Module
	ProjectName string `json:"projectName"`
}

type ModuleModel struct{}

func NewModuleModel() ModuleModel {
	return ModuleModel{}
}

/**
 * 创建模块
 */
func (model *ModuleModel) CrateModule(db *sql.DB, name string, description string, projectId int, versionId int) (Module, error) {
	insertSql := "insert into mb_module (module_name , description , project_id , version_id) values (? , ? , ? , ?)"
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

/**
 * 更新模块信息
 */
func (model *ModuleModel) UpdateModule(db *sql.DB, id int, name string, description string, projectId int, versionId int) (Module, error) {
	updateSql := "update mb_module set module_name = ?,description = ?,project_id = ?,version_id = ? where id = ?"
	result, err := model.findById(db, id)
	if err != nil {
		return Module{}, err
	}
	if name != "" {
		result.Name = name
	}
	if description != "" {
		result.Description = description
	}
	if projectId != 0 {
		result.ProjectId = projectId
	}
	if versionId != 0 {
		result.VersionId = versionId
	}
	_, err = db.Exec(updateSql, result.Name, result.Description, result.ProjectId, result.VersionId, id)
	if err != nil {
		return Module{}, err
	}
	return result, nil
}

/**
 * 根据项目id查找模块
 */
func (model *ModuleModel) GetModuleListByProjectId(db *sql.DB, projectId int) ([]ModuleFindByProjectId, error) {
	selectSql := "select module.*,project.project_name from mb_module as module left join mb_project as project on module.project_id = project.id where module.project_id = ? and module.status = 1"
	resultRows, err := db.Query(selectSql, projectId)
	if err != nil {
		return nil, err
	}
	var resultData []ModuleFindByProjectId
	for resultRows.Next() {
		var module ModuleFindByProjectId
		err = resultRows.Scan(&module.Id, &module.Name, &module.Description, &module.ProjectId, &module.VersionId, &module.Status, &module.CreateTime, &module.UpdateTime, &module.ProjectName)
		if err != nil {
			return nil, err
		}
		resultData = append(resultData, module)
	}
	return resultData, nil
}

/**
 * 通过id查找模块信息
 */
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
