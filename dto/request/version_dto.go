package request

type CreateVersionDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   int    `json:"projectId"`
	PublishTime string `json:"publishTime"`
}

type GetVersionListByProjectIdDto struct {
	ProjectId int `json:"projectId"`
}

type UpdateVersionDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type DeleteVersionDto struct {
	Id int `json:"id"`
}
