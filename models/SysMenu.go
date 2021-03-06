//该代码使用CodeSmith
//author：tsui

package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterModel(new(SysMenu))
}

type SysMenu struct {
	ID         int64     `orm:"column(ID);pk;unique;index;auto;"`
	Direction  bool      `orm:"column(Direction);not null;"`
	Parent     int64     `orm:"column(Parent);"`
	Name       string    `orm:"column(Name);size(100);"`
	Icon       string    `orm:"column(Icon);size(100);"`
	Clazz      string    `orm:"column(Clazz);size(100);"`
	Area       string    `orm:"column(Area);size(100);"`
	Controller string    `orm:"column(Controller);size(100);"`
	Method     string    `orm:"column(Method);size(100);"`
	Parameter  string    `orm:"column(Parameter);size(100);"`
	URL        string    `orm:"column(Url);size(100);"`
	Reserve    string    `orm:"column(Reserve);null;size(50);"`
	Remark     string    `orm:"column(Remark);null;size(250);"`
	Creator    int64     `orm:"column(Creator);null;"`
	CreateTime time.Time `orm:"column(CreateTime);null;type(datetime);"`
	Auditor    int64     `orm:"column(Auditor);null;"`
	AuditTime  time.Time `orm:"column(AuditTime);null;type(datetime);"`
	Canceler   int64     `orm:"column(Canceler);null;"`
	CancelTime time.Time `orm:"column(CancelTime);null;type(datetime);"`
}

func (this *SysMenu) TableName() string {
	return "SysMenu"
}

func (this *SysMenu) Add() (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(this)
	return id, err
}

func (this *SysMenu) Count(condation *orm.Condition) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		return qs.SetCond(condation).Count()
	} else {
		return qs.Count()
	}
}

func (this *SysMenu) Update(cols ...string) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Update(this, cols...)
	return num, err
}

func (this *SysMenu) Delete() (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(this)
	return num, err
}

func (this *SysMenu) Read(cols ...string) (*SysMenu, error) {
	o := orm.NewOrm()
	err := o.Read(this, cols...)
	if err != nil {
		return this, err
	}
	return this, nil
}

func (this *SysMenu) GetAll(condation *orm.Condition, sort string) (models []SysMenu) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation).All(&models)
	} else {
		qs.All(&models)
	}
	return models
}

func (this *SysMenu) Getlist(condation *orm.Condition, page int64, page_size int64, sort string) (models []orm.Params, count int64) {
	o := orm.NewOrm()
	qs := o.QueryTable(this)
	if condation != nil {
		qs.SetCond(condation)
	}
	qs.Limit(page_size, page).OrderBy(sort).Values(&models)
	count, _ = qs.Count()
	return models, count
}

func (this *SysMenu) Validation() (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&this)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func (this *SysMenu) ReadByName() *SysMenu {
	o := orm.NewOrm()
	o.Read(&this, "Name")
	return this
}

type MenuNode struct {
	ID           int64
	Iocn         string
	Name         string
	URL          string
	Direction    bool
	ChildrenNode []MenuNode
}

func (this *SysMenu) GetMenuHorizontal(id int64, roles string) []MenuNode {
	pNode := new(MenuNode)
	return this.GetMenu(id, roles, false, pNode)
}
func (this *SysMenu) GetMenusVertical(id int64, roles string) []MenuNode {
	pNode := new(MenuNode)
	return this.GetMenu(id, roles, true, pNode)
}
func (this *SysMenu) GetMenu(parent int64, roles string, direction bool, pNode *MenuNode) []MenuNode {
	menus := this.GetMenusByRole(roles, parent, direction)
	menuNodes := make([]MenuNode, 0, 0)
	for _, value := range menus {
		pNode.ID = value.ID
		pNode.URL = value.URL
		pNode.Name = value.Name
		pNode.Direction = value.Direction
		pNode.Iocn = value.Icon
		cnt := len(this.GetMenusByRole(roles, value.ID, direction))
		if cnt > 0 {
			pNode.ChildrenNode = this.GetMenu(value.ID, roles, direction, new(MenuNode))
		}
		menuNodes = append(menuNodes, *pNode)
	}
	return menuNodes
}

func (this *SysMenu) GetMenusByRole(roles string, parentMenu int64, direction bool) []SysMenu {
	o := orm.NewOrm()
	var menusResult []SysMenu
	o.Raw(` 
			SELECT c.* from sysroleoperator a
			left JOIN sysrolemenu  b on a.Role=b.Role
			left join sysmenu c on b.Menu=c.ID
			where  a.Role in (?) and c.Parent=? and c.Direction=? and c.ID<>0
	   `, roles, parentMenu, direction).QueryRows(&menusResult)
	return menusResult
}

//GetMenuByUser 通过用户获取菜单
func GetMenusByUser(user SysOperator) []SysMenu {
	o := orm.NewOrm()
	var menusResult []SysMenu

	o.Raw(` 
			SELECT c.* from sysroleoperator a
			left JOIN sysrolemenu  b on a.Role=b.Role
			left join sysmenu c on b.Menu=c.ID
			where  a.Operator=? 
	   `, user.ID).QueryRows(&menusResult)

	return menusResult
}
