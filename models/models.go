package models

type Result struct {
	PostID        int32  `json:"postId"`
	PostName      string `json:"postName"`
	CommentsCount int64  `json:"commnetsCount"`
	UserName      string `json:"userName"`
	Body          string `json:"body"`
}

type Comments struct {
	PostID int32  `json:"postId"`
}

type Posts struct {
	ID     int32  `json:"id"`
	UserID int32  `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Users struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}
