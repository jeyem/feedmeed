package v1

type loginForm struct {
	Identifier string `json:"identifier" form:"identifier" valid:"required"`
	Password   string `json:"password" form:"password" valid:"required"`
}

type registerForm struct {
	Username string `json:"username" form:"username" valid:"required"`
	Nikname  string `json:"nikname" form:"nikname" valid:"required"`
	Password string `json:"password" form:"password" valid:"required"`
}

type postform struct {
	Message string `json:"message" form:"message"`
}
