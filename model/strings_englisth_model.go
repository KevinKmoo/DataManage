package model

import (
	"database/sql"
)

type StringsEnglishElement struct {
	Id         int    `json:"id"`
	StringId   int    `json:"stringId"`
	Translate  string `json:"translate"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type StringsEnglishModel struct{}

func NewStringsEnglishModel() StringsEnglishModel {
	return StringsEnglishModel{}
}

/**
 * 创建英文翻译
 */
func (model *StringsEnglishModel) Create(db *sql.DB, stringId int, translate string) (StringsEnglishElement, error) {
	insertSql := "insert into mb_strings_english (string_id , translate) values (?,?)"
	insertRow, err := db.Exec(insertSql, stringId, translate)
	if err != nil {
		return StringsEnglishElement{}, err
	}
	insertRowId, err := insertRow.LastInsertId()
	if err != nil {
		return StringsEnglishElement{}, err
	}
	return model.findById(db, int(insertRowId))
}

/**
 * 通过id查找英文翻译
 */
func (model *StringsEnglishModel) findById(db *sql.DB, id int) (StringsEnglishElement, error) {
	selectSql := "select * from mb_strings_english where id = ?"
	resultRow := db.QueryRow(selectSql, id)
	var element StringsEnglishElement
	err := resultRow.Scan(&element.Id, &element.StringId, &element.Translate, &element.CreateTime, &element.UpdateTime)
	if err != nil {
		return StringsEnglishElement{}, err
	}
	return element, nil
}
