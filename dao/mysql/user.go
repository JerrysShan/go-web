package mysql

import "go-web/models"

type userDao struct{}

var UserDAO userDao

func (ud *userDao) Select(where string, page, pageSize int) (users []models.User, err error) {
	err = db.Where(where).Limit(pageSize, (page-1)*pageSize).Find(&users)
	return
}

func (ud *userDao) Get(where string) (user *models.User, err error) {
	_, err = db.Where(where).Get(user)
	return
}

func (ud *userDao) Insert(u ...interface{}) (int64, error) {
	affected, err := db.Insert(u...)
	return affected, err
}

func (ud *userDao) Update(id int64, user *models.Resource) (int64, error) {
	affected, err := db.ID(id).Update(user)
	return affected, err
}

func (ud *userDao) Delete(id int64) (int64, error) {
	var user models.User
	affected, err := db.ID(id).Delete(&user)
	return affected, err
}
