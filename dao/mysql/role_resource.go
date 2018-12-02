package mysql

import (
	"go-web/models"
)

type roleResourceDao struct{}

var RoleResourceDAO roleResourceDao

func (r *roleResourceDao) Select(where map[string]interface{}, page, pageSize int) (roleResources []models.RoleResource, err error) {
	err = db.Where(where).Limit(pageSize, (page-1)*pageSize).Find(&roleResources)
	return
}

func (r *roleResourceDao) Get(where map[string]interface{}) (*models.RoleResource, error) {
	var roleResource models.RoleResource
	_, err := db.Where(where).Get(&roleResource)
	return &roleResource, err
}

func (r *roleResourceDao) Exist(roleResource *models.RoleResource) (bool, error) {
	return db.Exist(roleResource)
}

func (r *roleResourceDao) Insert(u *models.RoleResource) (int64, error) {
	affected, err := db.InsertOne(u)
	return affected, err
}

func (r *roleResourceDao) InsertMany(u []*models.RoleResource) (int64, error) {
	affected, err := db.Insert(u)
	return affected, err
}

func (r *roleResourceDao) Update(id int, roleResource *models.RoleResource) (int64, error) {
	affected, err := db.ID(id).Update(roleResource)
	return affected, err
}

func (r *roleResourceDao) Delete(id int) (int64, error) {
	var roleResource models.RoleResource
	affected, err := db.ID(id).Delete(&roleResource)
	return affected, err
}
