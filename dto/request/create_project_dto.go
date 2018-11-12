package request

type CreateProjectDto struct {
	Name string `json:"name" binding:"required"`
}
