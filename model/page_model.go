package model

import (
	"database/sql"
)

type Page struct {
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VersionId   int    `json:"versionId"` //加入版本
	ModuleId    int    `json:"moduleId"`
	Status      int    `json:"status"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type PageModel struct {
}

/**
 * 创建页面
 */
func (p *PageModel) CreatePage(db *sql.DB, name string, description string, key string, versionId int, moduleId int) (page Page, err error) {
	insertSql := "insert into mb_page (name , description , key , version_id , moduleId) values (?,?,?,?,?)"
	result, err := db.Exec(insertSql, name, description, key, versionId, moduleId)
	if err != nil {
		return Page{}, err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return Page{}, err
	}
	return p.findPageById(db, int(insertId))
}

/**
 * 更新页面
 */
func (model *PageModel) UpdatePage(db *sql.DB, id int, name string, description string) (Page, error) {
	page, err := model.findPageById(db, id)
	if err != nil {
		return Page{}, err
	}
	if name != "" {
		page.Name = name
	}
	if description != "" {
		page.Description = description
	}
	updateSql := "update mb_page set name = ?,description = ? where id = ?"
	_, err = db.Exec(updateSql, page.Name, page.Description, id)
	if err != nil {
		return Page{}, err
	}
	return page, nil
}

/**
 * 通过id查找页面信息
 */
func (p *PageModel) findPageById(db *sql.DB, id int) (page Page, err error) {
	selectSql := "select * from mb_page where id = ?"
	resultRow := db.QueryRow(selectSql, page.Id)
	error := resultRow.Scan(&page.Id, &page.Key, &page.Name, &page.Description, &page.VersionId, &page.ModuleId, &page.Status, &page.CreateTime, &page.UpdateTime)
	if err != nil {
		return Page{}, err
	}
	return page, error
}
