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

func Query(s *schemas.User) ([]models.User, error) {
	where := ""
	if s.Name != "" {
		where += "name=" + s.Name
	}
	if s.ID > 0 {
		where += ""
	}
	data, err := mysql.UserDAO.Select(where, 1, 10)
	return data, err
}

func Update() {

}

func Del() {

}
