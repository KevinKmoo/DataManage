package model

import (
	"database/sql"
)

type Version struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   int    `json:"projectId"`
	PublishTime string `json:"publishTime"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type VersionWithProjectName struct {
	Version
	ProjectName string `json:"projectName"`
}

type VersionModel struct {
}

/**
 * 创建版本
 */
func (model *VersionModel) CreateVersion(db *sql.DB, name string, description string, projectId int, publishTime string) (Version, error) {
	insertSql := "insert into mb_version (name, description, project_id, publish_time) values (?,?,?,?)"
	resultRow, err := db.Exec(insertSql, name, description, projectId, publishTime)
	if err != nil {
		return Version{}, err
	}
	insertId, err := resultRow.LastInsertId()
	if err != nil {
		return Version{}, err
	}
	return model.findById(db, int(insertId))
}

/**
 * 查找所有的版本
 */
func (model *VersionModel) FindAllVersion(db *sql.DB) ([]Version, error) {
	selectSql := "select * from mb_version"
	rows, err := db.Query(selectSql)
	if err != nil {
		return nil, err
	}
	var result []Version
	for rows.Next() {
		var version Version
		err = rows.Scan(&version.Id, &version.Name, &version.Description, &version.ProjectId, &version.PublishTime, &version.CreateTime, &version.UpdateTime)

		if err != nil {
			return nil, err
		}

		result = append(result, version)
	}
	return result, nil
}

/**
 * 通过项目id查找版本列表
 */
func (model *VersionModel) FindVerionListByProjectId(db *sql.DB, projectId int) ([]VersionWithProjectName, error) {
	selectSql := "select version.*,project.project_name from mb_version as version left join mb_project as project on version.project_id = project.id where version.project_id = ?"
	rows, err := db.Query(selectSql, projectId)
	if err != nil {
		return nil, err
	}
	var result []VersionWithProjectName
	for rows.Next() {
		var version VersionWithProjectName
		err = rows.Scan(&version.Id, &version.Name, &version.Description, &version.ProjectId, &version.PublishTime, &version.CreateTime, &version.UpdateTime, &version.ProjectName)

		if err != nil {
			return nil, err
		}

		result = append(result, version)
	}
	return result, nil
}

/**
 * 通过id删除版本
 */
func (model *VersionModel) DeleteVersionById(db *sql.DB, id int) error {
	deleteSql := "delete from mb_version where id = ?"
	_, err := db.Exec(deleteSql)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 更新版本信息
 */
func (model *VersionModel) UpdateVersion(db *sql.DB, id int, name string, description string) (Version, error) {
	needUpdateVersion, err := model.findById(db, id)
	if err != nil {
		return Version{}, err
	}
	if name != "" {
		needUpdateVersion.Name = name
	}
	if description != "" {
		needUpdateVersion.Description = description
	}
	updateSql := "update mb_version set name = ?,description = ? where id = ?"
	_, err = db.Exec(updateSql, name, description, id)
	if err != nil {
		return Version{}, err
	}
	return needUpdateVersion, nil
}

/**
 * 通过id查找版本
 */
func (model *VersionModel) findById(db *sql.DB, id int) (Version, error) {
	selectSql := "select * from mb_version where id = ?"
	resultRow := db.QueryRow(selectSql, id)
	var version Version
	err := resultRow.Scan(&version.Id, &version.Name, &version.Description, &version.PublishTime, &version.CreateTime, &version.UpdateTime)
	if err != nil {
		return Version{}, err
	}
	return version, nil
}
