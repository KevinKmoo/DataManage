package request

type CreatePageDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Key         string `json:"key" binding:"required"`
	VersionId   int    `json:"versionId" binding:"required"`
	ModuleId    int    `json:"moduleId" binding:"required"`
}

type UpdatePageDto struct {
	Id          int    `json:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
