package schemas

type User struct {
	ID     int    `form:"id"`
	Name   string `form:"name" binding:"required"`
	UID    int64  `form:"uid" binding:"required"`
	RoleID int    `form:"roleId" binding:"required"`
}
