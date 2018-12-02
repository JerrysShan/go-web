package models

type Role struct {
	ID      int `xorm:"pk autoincr 'id'"`
	Name    string
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
	Status  int   `xorm:"default(0)"`
}
