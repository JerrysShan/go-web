package mysql

import (
	"go-web/models"
)

type userRoleDao struct{}

var UserRoleDAO userRoleDao

func (ur *userRoleDao) Select(where map[string]interface{}, page, pageSize int) (userRole []models.UserRole, err error) {
	err = db.Where(where).Limit(pageSize, (page-1)*pageSize).Find(&userRole)
	return
}

func (ur *userRoleDao) Get(where map[string]interface{}) (*models.UserRole, error) {
	var userRole models.UserRole
	_, err := db.Where(where).Get(&userRole)
	return &userRole, err
}

func (ur *userRoleDao) Insert(u *models.UserRole) (int64, error) {
	affected, err := db.Insert(u)
	return affected, err
}

func (ur *userRoleDao) InsertMany(u []*models.UserRole) (int64, error) {
	affected, err := db.Insert(u)
	return affected, err
}

func (ur *userRoleDao) Update(id int, userRole *models.UserRole) (int64, error) {
	affected, err := db.ID(id).Update(userRole)
	return affected, err
}

func (ur *userRoleDao) Delete(id int) (int64, error) {
	var userRole models.UserRole
	affected, err := db.ID(id).Delete(&userRole)
	return affected, err
}
