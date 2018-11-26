package mysql

import "go-web/models"

type roleDao struct{}

var RoleDAO roleDao

// Select return a list of role
func (rd *roleDao) Select(where string, page, pageSize int) (roles []models.Role, err error) {
	err = db.Where(where).Limit(pageSize, (page-1)*pageSize).Find(&roles)
	return
}

func (rd *roleDao) Insert(role ...interface{}) (int64, error) {
	affected, err := db.Insert(role...)
	return affected, err
}

func (rd *roleDao) Update(id int64, role *models.Role) (int64, error) {
	affected, err := db.ID(id).Update(role)
	return affected, err
}

func (rd *roleDao) Delete(id int64) (int64, error) {
	var role models.Role
	affected, err := db.ID(id).Delete(&role)
	return affected, err
}
