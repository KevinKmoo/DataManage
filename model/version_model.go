package model

import (
	"database/sql"
)

type Version struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	PublishTime string `json:"publishTime"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"UpdateTime"`
}

type VersionModel struct {
}

/**
 * 创建版本
 */
func (model *VersionModel) CreateVersion(db *sql.DB, name string, description string, publishTime string) (Version, error) {
	insertSql := "insert into mb_version (name, description, publish_time) values (?,?,?)"
	resultRow, err := db.Exec(insertSql, name, description, publishTime)
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
		err = rows.Scan(&version.Id, &version.Name, &version.Description, &version.PublishTime, &version.CreateTime, &version.UpdateTime)

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
