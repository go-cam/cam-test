package models

import (
	"time"
)

type User struct {
	Id       int       `xorm:"not null pk autoincr comment('user id') INT(11)"`
	Account  string    `xorm:"comment('user name') VARCHAR(50)"`
	Password string    `xorm:"comment('password') VARCHAR(64)"`
	DeleteAt time.Time `xorm:"comment('delete timestamp') TIMESTAMP"`
	UpdateAt time.Time `xorm:"default CURRENT_TIMESTAMP comment('update timestamp') TIMESTAMP"`
	CreateAt time.Time `xorm:"default CURRENT_TIMESTAMP comment('create timestamp') TIMESTAMP"`
}
