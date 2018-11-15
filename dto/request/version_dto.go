package request

type CreateVersionDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	PublishTime string `json:"publishTime"`
}

type DeleteVersionDto struct {
	Id int `json:"id"`
}
