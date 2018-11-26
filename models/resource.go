package models

type Resource struct {
	ID      int64 `xorm:"pk autoincr 'id'"`
	Name    string
	URL     string
	Code    string
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
	Status  int   `xorm:"default(0)"`
}
