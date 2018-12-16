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

type PageListBean struct {
	Page
	ModuleName  string `json:"moduleName"`
	VersionName string `json:"versionName"`
}

type PageModel struct {
}

/**
 * 创建页面
 */
func (p *PageModel) CreatePage(db *sql.DB, name string, key string, description string, versionId int, moduleId int) (page Page, err error) {
	insertSql := "insert into mb_page (page_name , description , page_key , version_id , module_id) values (?,?,?,?,?)"
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
	updateSql := "update mb_page set page_name = ?,description = ? where id = ?"
	_, err = db.Exec(updateSql, page.Name, page.Description, id)
	if err != nil {
		return Page{}, err
	}
	return page, nil
}

/**
 * 页面列表
 */
func (model *PageModel) PageList(db *sql.DB, projectId int, moduelId int, versionId int, key string) ([]PageListBean, error) {
	var selectSql string
	selectSql = "select page.*,module.module_name as moduleName,version.version_name as versionName from mb_page as page " +
		"left join mb_module as module on page.module_id = module.id " +
		"left join mb_version as version on page.version_id = version.id"
	var resultRows *sql.Rows
	var err error
	if moduelId != 0 && versionId == 0 {
		selectSql += " where page.module_id = ? and page.page_name like ?"
		resultRows, err = db.Query(selectSql, moduelId, "%"+key+"%")
	} else if moduelId == 0 && versionId != 0 {
		selectSql += " and page.version_id = ? and page.page_name like ?"
		resultRows, err = db.Query(selectSql, versionId, "%"+key+"%")
	} else if moduelId != 0 && versionId != 0 {
		selectSql += " where page.module_id = ? and page.version_id = ? and page.page_name like ?"
		resultRows, err = db.Query(selectSql, moduelId, versionId, "%"+key+"%")
	}
	if err != nil {
		return nil, err
	}
	result := []PageListBean{}
	for resultRows.Next() {
		temp := PageListBean{}
		err = resultRows.Scan(&temp.Id, &temp.Key, &temp.Name, &temp.Description, &temp.VersionId, &temp.ModuleId, &temp.Status, &temp.CreateTime, &temp.UpdateTime, &temp.ModuleName, &temp.VersionName)
		if err != nil {
			return nil, err
		}
		result = append(result, temp)
	}
	return result, nil
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
