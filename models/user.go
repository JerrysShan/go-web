package models

type User struct {
	ID      int `xorm:"pk autoincr 'id'"`
	Name    string
	UID     int64 `xorm:"'uid"`
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
	Status  int   `xorm:"default(0)"`
}
