package resource

import (
	"go-web/dao/mysql"
	"go-web/models"
	"go-web/schemas"
)

func Query(s *schemas.Resource) ([]models.Resource, error) {
	where := make(map[string]interface{})
	if s.ID > 0 {
		where["id"] = s.ID
	}
	if s.Code != "" {
		where["code"] = s.Code
	}
	if s.Name != "" {
		where["name"] = s.Name
	}
	data, err := mysql.ResourceDAO.Select(where, 1, 10)
	return data, err
}

func Insert(s *schemas.Resource) (int64, error) {
	r := models.Resource{}
	r.Name = s.Name
	r.Code = s.Code
	r.URL = s.URL
	affected, err := mysql.ResourceDAO.Insert(r)
	return affected, err
}

func Update(s *schemas.Resource) (int64, error) {
	r := &models.Resource{}
	r.Name = s.Name
	r.Code = s.Code
	r.URL = s.URL
	return mysql.ResourceDAO.Update(s.ID, r)
}

func Delete(id int) (int64, error) {
	return mysql.ResourceDAO.Delete(id)
}
