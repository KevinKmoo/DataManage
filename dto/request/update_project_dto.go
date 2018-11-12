package request

type UpdateProjectDto struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
