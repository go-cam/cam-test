package migrations

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component/camConsole"
)

func init() {
	m := new(m200809_211813_create_foreign_key)
	cam.App.AddMigration(m)
}

type m200809_211813_create_foreign_key struct {
	camConsole.BaseMigration
}

// up
func (m *m200809_211813_create_foreign_key) Up() {
	m.CreateTable("group", []camStatics.MysqlColumnBuilderInterface{
		m.Column("id", m.IntPrimaryKey()),
		m.Column("name", m.Varchar(50)),
	}, m.DefaultOption("group"))

	m.CreateTable("group_user", []camStatics.MysqlColumnBuilderInterface{
		m.Column("id", m.IntPrimaryKey()),
		m.Column("group_id", m.Int().Unsigned().Comment("group_id")),
		m.Column("name", m.Varchar(50)),
	}, m.DefaultOption("group_user"))

	m.CreateForeignKey("group_id", "group_user", []string{"group_id"}, "group", []string{"id"})
}

// down
func (m *m200809_211813_create_foreign_key) Down() {
}
