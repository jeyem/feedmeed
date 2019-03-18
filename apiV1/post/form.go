package post

type form struct {
	Message string `json:"message" form:"message"`
}
