package schemas

type Role struct {
	ID       int    `form:"id" json:"id"`
	Name     string `form:"name" json:"name"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"pageSize" json:"pageSize"`
}
