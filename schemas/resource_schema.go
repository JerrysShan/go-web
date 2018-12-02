package schemas

type Resource struct {
	ID       int    `form:"id" json:"id"`
	Name     string `form:"name" json:"name" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
	URL      string `form:"url" json:"url" binding:"required"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"pageSize" json:"pageSize"`
}
