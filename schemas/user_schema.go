package schemas

type User struct {
	ID     int    `form:"id" json:"id"`
	Name   string `form:"name" json:"name" binding:"required"`
	UID    int64  `form:"uid" json:"uid" binding:"required"`
	RoleID int    `form:"roleId" json:"roleId"`
}
