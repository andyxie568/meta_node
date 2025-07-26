package utils

type SignUpReq struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreatePostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type UpdatePostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateCommentReq struct {
	Content string `json:"content"`
}
