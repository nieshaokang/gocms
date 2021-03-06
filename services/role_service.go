package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/cuua/gocms/models"
	"sync"
)

var RoleService = &roleService{
	mutex: &sync.Mutex{},
}

type roleService struct {
	mutex *sync.Mutex
}

// RolePageList 获取分页数据
func (s *roleService) RolePageList(params *models.RoleQueryParam) ([]*models.Role, int64) {
	query := orm.NewOrm().QueryTable(models.RoleTBName())
	data := make([]*models.Role, 0)
	//默认排序
	sortorder := "Id"
	switch params.Sort {
	case "Id":
		sortorder = "Id"
	case "Seq":
		sortorder = "Seq"
	}
	if params.Order == "desc" {
		sortorder = "-" + sortorder
	}
	query = query.Filter("name__istartswith", params.NameLike)
	total, _ := query.Count()
	_, _ = query.OrderBy(sortorder).Limit(params.Limit, params.Offset).All(&data)
	return data, total
}

// RoleDataList 获取角色列表
func (s *roleService) RoleDataList(params *models.RoleQueryParam) []*models.Role {
	params.Limit = -1
	params.Sort = "Seq"
	params.Order = "asc"
	data, _ := s.RolePageList(params)
	return data
}

// RoleBatchDelete 批量删除
func (s *roleService) RoleBatchDelete(ids []int) (int64, error) {
	query := orm.NewOrm().QueryTable(models.RoleTBName())
	num, err := query.Filter("id__in", ids).Delete()
	return num, err
}

// RoleOne 获取单条
func (s *roleService) RoleOne(id int) (*models.Role, error) {
	o := orm.NewOrm()
	m := models.Role{Id: id}
	err := o.Read(&m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
