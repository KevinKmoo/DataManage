package request

type CreateModuleDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   int    `json:"projectId"`
	VersionId   int    `json:"versionId"`
}

type UpdateModuleDto struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ProjectId   int    `json:"projectId"`
	VersionId   int    `json:"versionId"`
}

type DeleteModuleDto struct {
}

type GetModuleListDto struct {
	ProjectId int `json:"projectId"`
}
