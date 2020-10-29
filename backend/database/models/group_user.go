package models

type GroupUser struct {
	Id      int    `xorm:"not null pk autoincr INT(10)"`
	GroupId int    `xorm:"comment('group_id') index INT(10)"`
	Name    string `xorm:"VARCHAR(50)"`
}
