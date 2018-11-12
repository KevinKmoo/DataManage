package model

import (
	"database/sql"
	"log"
	"time"
)

type MyBittProject struct {
	Id          int64  `json:"id" binding:"required"`
	ProjectName string `json:"projectName" binding:"required"`
	CreateTime  string `json:"createTime" binding:"required"`
	UpdateTime  string `json:"updateTime" binding:"required"`
}

/**
 * 创建一个项目
 */
func CreateProject(db *sql.DB, projectName string) (createdProject MyBittProject, e error) {
	insertSql := "insert into mb_project(project_name) values (?)"
	result, err := db.Exec(insertSql, projectName)
	if err != nil {
		log.Fatal(err)
		return MyBittProject{}, err
	}
	lastInsertId, _ := result.LastInsertId()
	return findById(db, int(lastInsertId))
}

/**
 * 项目列表
 */
func ProjectList(db *sql.DB, pageIndex int, pageSize int) (projectList []MyBittProject, e error) {
	querySql := "select * from mb_project limit?,?"
	result, err := db.Query(querySql, pageIndex-1, pageSize)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	index := 0
	projectListResult := []MyBittProject{}
	for result.Next() {
		temp := MyBittProject{}
		err = result.Scan(&temp.Id, &temp.ProjectName, &temp.CreateTime, &temp.UpdateTime)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		projectListResult = append(projectListResult, temp)
		index++
	}

	return projectListResult, nil
}

/**
 * 更新项目
 */
func UpdateProject(db *sql.DB, id int, name string) (project MyBittProject, e error) {
	updateSql := "update mb_project set project_name = ? , update_time = ? where id = ?"
	_, err := db.Exec(updateSql, name, time.Now(), id)
	if err != nil {
		log.Fatal(err)
		return MyBittProject{}, err
	}
	return findById(db, id)
}

/**
 * 删除项目
 */
func DeleteProject(db *sql.DB, id int) error {
	deleteSql := "delete from mb_project where id = ?"
	_, err := db.Exec(deleteSql, id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

/**
 * 通过id查找项目
 */
func findById(db *sql.DB, id int) (project MyBittProject, e error) {
	insertRow := db.QueryRow("select * from mb_project where id = ?", id)
	err := insertRow.Scan(&project.Id, &project.ProjectName, &project.CreateTime, &project.UpdateTime)
	if err != nil {
		log.Fatal(err)
		return MyBittProject{}, err
	}
	return project, nil
}
