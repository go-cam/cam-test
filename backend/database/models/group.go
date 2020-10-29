package models

type Group struct {
	Id   int    `xorm:"not null pk autoincr INT(10)"`
	Name string `xorm:"VARCHAR(50)"`
}
