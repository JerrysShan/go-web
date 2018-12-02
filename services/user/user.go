package user

import (
	"go-web/dao/mysql"
	"go-web/models"
	"go-web/schemas"
)

func Create(s *schemas.User) (int64, int, error) {
	u := &models.User{}
	u.Name = s.Name
	u.UID = s.UID
	affectedRows, err := mysql.UserDAO.Insert(u)
	return affectedRows, u.ID, err
}

func FindByUid(uid int64) (*models.User, error) {
	where := make(map[string]interface{})
	where["uid"] = uid
	data, err := mysql.UserDAO.Get(where)
	return data, err
}

func IsExist(uid int64) (bool, error) {
	return mysql.UserDAO.Exist(&models.User{UID: uid})
}

func CreateRoles(userId int, roleIds []int) (int64, error) {
	var userRoles []*models.UserRole
	for _, roleId := range roleIds {
		userRoles = append(userRoles, &models.UserRole{UserID: userId, RoleID: roleId})
	}
	return mysql.UserRoleDAO.InsertMany(userRoles)
}

func Query(s *schemas.User) ([]models.User, error) {
	where := make(map[string]interface{})
	if s.Name != "" {
		where["name"] = s.Name
	}
	if s.ID > 0 {
		where["id"] = s.ID
	}
	if s.UID > 0 {
		where["uid"] = s.UID
	}
	data, err := mysql.UserDAO.Select(where, 1, 10)
	return data, err
}

func Update(s *schemas.User) (int64, error) {
	user := &models.User{}
	user.Name = s.Name
	user.UID = s.UID
	affected, err := mysql.UserDAO.Update(s.ID, user)
	return affected, err
}

func Delete(id int) (int64, error) {
	affected, err := mysql.UserDAO.Delete(id)
	return affected, err
}
