package apiType

// UserInfo 的数据结构  ======== 对外展示 =======
type UserInfo struct {
	Email    string `bson:"Email"`    // 用户 Email
	UserID   string `bson:"UserID"`   // 用户 ID
	Avatar   string `bson:"Avatar"`   // 用户头像
	NickName string `bson:"NickName"` // 用户昵称
}
