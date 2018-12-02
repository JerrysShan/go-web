package mysql

import "go-web/models"

type resourceDao struct{}

var ResourceDAO resourceDao

func (rd *resourceDao) Select(where map[string]interface{}, page, pageSize int) (resources []models.Resource, err error) {
	err = db.Where(where).Limit(pageSize, (page-1)*pageSize).Find(&resources)
	return
}

func (rd *resourceDao) Get(where map[string]interface{}) (resource models.Resource, err error) {
	_, err = db.Where(where).Get(&resource)
	return
}

func (rd *resourceDao) Insert(resources ...interface{}) (int64, error) {
	affected, err := db.Insert(resources...)
	return affected, err
}

func (rd *resourceDao) Update(id int, resource *models.Resource) (int64, error) {
	affected, err := db.ID(id).Update(resource)
	return affected, err
}

func (rd *resourceDao) Delete(id int) (int64, error) {
	var resource models.Resource
	affected, err := db.ID(id).Delete(&resource)
	return affected, err
}
