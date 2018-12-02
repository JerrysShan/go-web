package models

type RoleResource struct {
	ID         int   `xorm:"pk autoincr 'id'"`
	RoleID     int   `xorm:"role_id"`
	ResourceID int   `xorm:"resource_id"`
	Created    int64 `xorm:"created"`
	Updated    int64 `xorm:"updated"`
}
