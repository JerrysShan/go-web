package mysql

import (
	"go-web/models"
)

type userDao struct{}

var UserDAO userDao

func (ud *userDao) Select(where map[string]interface{}, page, pageSize int) (users []models.User, err error) {
	err = db.Where(where).Limit(pageSize, (page-1)*pageSize).Find(&users)
	return
}

func (ud *userDao) Get(where map[string]interface{}) (*models.User, error) {
	var user models.User
	_, err := db.Where(where).Get(&user)
	return &user, err
}

func (ud *userDao) Insert(u ...interface{}) (int64, error) {
	affected, err := db.Insert(u...)
	return affected, err
}

func (ud *userDao) Update(id int, user *models.User) (int64, error) {
	affected, err := db.ID(id).Update(user)
	return affected, err
}

func (ud *userDao) Delete(id int) (int64, error) {
	var user models.User
	affected, err := db.ID(id).Delete(&user)
	return affected, err
}
