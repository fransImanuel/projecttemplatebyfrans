package schemas

type CreateUserRequest struct {
	Name   string `json:"name" binding:"required" gorm:"column:Name"`
	Age    string `json:"age" binding:"required" gorm:"column:Age"`
	Email  string `json:"email" binding:"required" gorm:"column:Email"`
	Phone  string `json:"phone" binding:"required" gorm:"column:Phone"`
	RoleID int64  `json:"role_id" binding:"required" gorm:"column:RoleID"`
}
