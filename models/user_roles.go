package models

type UserRole struct {
	ID      int   `xorm:"pk autoincr 'id'"`
	UserID  int   `xorm:"user_id"`
	RoleID  int   `xorm:"role_id"`
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
}
