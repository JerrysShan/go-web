package user

import (
	"go-web/dao/mysql"
	"go-web/models"
	"go-web/schemas"
)

func Create(s schemas.User) error {
	u := &models.User{}
	u.Name = s.Name
	u.UID = s.UID
	affectedRows, err := mysql.UserDAO.Insert(u)
	if affectedRows == 1 {
		return nil
	}
	return err
}

func FindByUid(uid int64) (*models.User, error) {
	where := make(map[string]interface{})
	where["uid"] = uid
	data, err := mysql.UserDAO.Get(where)
	return data, err
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
