package post

type form struct {
	Message   string `json:"message" form:"message"`
	IsPrivate bool   `json:"isPrivate" form:"isPrivate"`
}
