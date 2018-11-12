package request

type CreatePageDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	Key         string `json:"key"`
	VersionId   int    `json:"versionId" binding:"required"`
}
