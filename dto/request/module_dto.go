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
	Status      int    `json:"status"`
}

type DeleteModuleDto struct {
}
