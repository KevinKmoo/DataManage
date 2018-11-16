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

type StringsEnglishTranslate struct {
	Id           int    `json:"id"`
	StringId     int    `json:"stringId"`
	Key          string `json:"key"`
	ChineseValue string `json:"chineseValue"`
	Translate    string `json:"translate"`
	PageId       int    `json:"pageId"`
	PageName     string `json:"pageName"`
	ModuleId     int    `json:"moduleId"`
	ModuleName   string `json:"moduleName"`
	VersionId    string `json:"versionId"`
	VersionName  string `json:"versionName"`
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
 * 更新英文翻译
 */
func (model *StringsEnglishModel) Update(db *sql.DB, id int, translate string) (StringsEnglishElement, error) {
	updateSql := "update mb_strings_english set translate = ? where id = ?"
	_, err := db.Exec(updateSql, translate, id)
	if err != nil {
		return StringsEnglishElement{}, err
	}
	return model.findById(db, id)
}

/**
 * 查找英文翻译的列表
 */
func (model *StringsEnglishModel) List(db *sql.DB, pageIndex int, pageSize int) ([]StringsEnglishTranslate, error) {
	selectSql := "select english.id,english.string_id,chinese.key,chinese.value,english.translate,chinese.page_id,page.name,page.module_id,module.name,chinese.version_id,version.name from mb_strings_english as english left join mb_strings as chinese on english.string_id = chinese.id left join mb_page as page on page.id = chinese.page_id left join mb_module as module on page.module_id = module.id left join mb_version as version on chinese.version_id = version.id limit ?,?"
	resultRows, err := db.Query(selectSql)
	if err != nil {
		return nil, err
	}
	var result []StringsEnglishTranslate
	for resultRows.Next() {
		var resultElement StringsEnglishTranslate
		resultRows.Scan(resultElement.Id, resultElement.StringId, resultElement.Key, resultElement.ChineseValue, resultElement.Translate, resultElement.PageId, resultElement.PageName, resultElement.ModuleId, resultElement.ModuleName, resultElement.VersionId, resultElement.VersionName)
		result = append(result, resultElement)
	}
	return result, nil
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
