package mysql

import "go-web/models"

type resourceDao struct{}

var ResourceDAO resourceDao

func (rd *resourceDao) Select(where string, page, pageSize int) (resources []models.Resource, err error) {
	err = db.Where(where).Limit(pageSize, (page-1)*pageSize).Find(&resources)
	return
}

func (rd *resourceDao) Insert(resources ...interface{}) (int64, error) {
	affected, err := db.Insert(resources...)
	return affected, err
}

func (rd *resourceDao) Update(id int64, resource *models.Resource) (int64, error) {
	affected, err := db.ID(id).Update(resource)
	return affected, err
}

func (rd *resourceDao) Delete(id int64) (int64, error) {
	var resource models.Resource
	affected, err := db.ID(id).Delete(&resource)
	return affected, err
}
