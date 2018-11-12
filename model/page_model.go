package model

import (
	"database/sql"
	"log"
)

type Page struct {
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	Description string `json:"description"`
	VersionId   int    `json:"versionId"` //加入版本
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type PageModel struct {
}

func (p *PageModel) CreatePage(db *sql.DB, name string, description string, key string, versionId int) (page Page, err error) {
	insertSql := "insert into mb_page (name , description , key , version_id) values (?,?,?,?)"
	result, err := db.Exec(insertSql, name, description, key, versionId)
	if err != nil {
		log.Fatal(err)
		return Page{}, err
	}
	insertId, err := result.LastInsertId()
	if err != nil {
		return Page{}, err
	}
	return p.findPageById(db, int(insertId))
}

func (p *PageModel) findPageById(db *sql.DB, id int) (page Page, err error) {
	selectSql := "select * from mb_page where id = ?"
	resultRow := db.QueryRow(selectSql, page.Id)
	error := resultRow.Scan(&page.Id, &page.Name, &page.Description, &page.Key, &page.CreateTime, &page.UpdateTime)
	if err != nil {
		return Page{}, err
	}
	return page, error
}
