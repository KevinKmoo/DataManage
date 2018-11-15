package model

import (
	"database/sql"
	"time"
)

type StringsElement struct {
	Id         int    `json:"id"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	Note       string `json:"note"`
	VersionId  int    `json:"versionId"`
	PageId     int    `json:"pageId"`
	Status     int    `json:"status"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type StringsModel struct{}

func NewStringsModel() StringsModel {
	return StringsModel{}
}

/**
 * 创建页面上需要翻译的文本
 */
func (model *StringsModel) Create(db *sql.DB, key string, value string, note string, versionId int, pageId int) (StringsElement, error) {
	insertSql := "insert into mb_strings (key , value , note , version_id , page_id) values (?,?,?,?,?)"
	lastRow, err := db.Exec(insertSql, key, value, note, versionId, pageId)
	if err != nil {
		return StringsElement{}, err
	}
	insertedId, err := lastRow.LastInsertId()
	if err != nil {
		return StringsElement{}, err
	}
	return model.findById(db, int(insertedId))
}

/**
 * 更新需要翻译的文本信息
 */
func (model *StringsModel) Update(db *sql.DB, id int, value string, note string, versionId int) (StringsElement, error) {
	stringsElement, err := model.findById(db, id)
	if err != nil {
		return StringsElement{}, err
	}
	if value != "" {
		stringsElement.Value = value
	}
	if note != "" {
		stringsElement.Note = note
	}
	if versionId != 0 {
		stringsElement.VersionId = versionId
	}

	stringsElement.UpdateTime = time.Now().String()

	updateSql := "update mb_strings set value = ?, note = ?, version_id = ?, update_time = ? where id = ?"
	_, err = db.Exec(updateSql, stringsElement.Value, stringsElement.Note, stringsElement.VersionId, id)
	if err != nil {
		return StringsElement{}, err
	}
	return stringsElement, nil
}

/**
 * 删除需要翻译的文本信息
 */
func (model *StringsModel) Delete(db *sql.DB, id int) error {
	updateSql := "update mb_strings set status = 0 where id = ?"
	_, err := db.Exec(updateSql, id)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 通过id查找需要翻译的文本
 */
func (model *StringsModel) findById(db *sql.DB, id int) (StringsElement, error) {
	selectSql := "select * from mb_strings where id = ?"
	resultRow := db.QueryRow(selectSql, id)
	var result StringsElement
	err := resultRow.Scan(&result.Id, &result.Key, &result.Value, &result.VersionId, &result.PageId, &result.CreateTime, &result.UpdateTime)
	if err != nil {
		return StringsElement{}, err
	}
	return result, nil
}
