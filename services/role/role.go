package role

import (
	"go-web/dao/mysql"
	"go-web/models"
	"go-web/schemas"
)

func Query(s *schemas.Role) ([]models.Role, error) {
	where := make(map[string]interface{})
	if s.ID > 0 {
		where["id"] = s.ID
	}
	if s.Name != "" {
		where["name"] = s.Name
	}
	return mysql.RoleDAO.Select(where, 1, 10)
}

func FindByName(name string) ([]models.Role, error) {
	where := make(map[string]interface{})
	where["name"] = name
	return mysql.RoleDAO.Select(where, 1, 10)
}

func Insert(s *schemas.Role) (int64, error) {
	r := &models.Role{}
	r.Name = s.Name
	return mysql.RoleDAO.Insert(r)
}

func Update(s *schemas.Role) (int64, error) {
	r := &models.Role{}
	r.Name = s.Name
	return mysql.RoleDAO.Update(s.ID, r)
}

func Delete(id int) (int64, error) {
	return mysql.RoleDAO.Delete(id)
}
