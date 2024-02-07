package repository

// 用于UserQueryHandler
type UserQueryResult struct {
	NickName string `json:"nickname"`
	Sex      int    `json:"sex"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	IsFriend bool   `json:"is_friend"`
}
