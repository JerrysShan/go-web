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

func IsExist(name string) (bool, error) {
	return mysql.RoleDAO.Exist(&models.Role{Name: name})
}

func Insert(s *schemas.Role) (int64, int, error) {
	r := &models.Role{}
	r.Name = s.Name
	affectd, err := mysql.RoleDAO.Insert(r)
	return affectd, r.ID, err
}

func CreateResources(roleId int, resourceIds []int) (int64, error) {
	var roleResource []*models.RoleResource
	for _, resourceID := range resourceIds {
		roleResource = append(roleResource, &models.RoleResource{RoleID: roleId, ResourceID: resourceID})
	}
	return mysql.RoleResourceDAO.InsertMany(roleResource)
}

func Update(s *schemas.Role) (int64, error) {
	r := &models.Role{}
	r.Name = s.Name
	return mysql.RoleDAO.Update(s.ID, r)
}

func Delete(id int) (int64, error) {
	return mysql.RoleDAO.Delete(id)
}
