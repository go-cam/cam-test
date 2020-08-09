package migrations

import (
	"github.com/go-cam/cam"
	"github.com/go-cam/cam/base/camStatics"
	"github.com/go-cam/cam/component/camConsole"
)

func init() {
	m := new(m200805_075136_create_new_table)
	cam.App.AddMigration(m)
}

type m200805_075136_create_new_table struct {
	camConsole.BaseMigration
}

// up
func (m *m200805_075136_create_new_table) Up() {
	m.CreateTable("new_table", []camStatics.MysqlColumnBuilderInterface{
		m.Column("id", m.IntPrimaryKey()),
		m.Column("dec", m.Decimal(10, 2).NotNull().Default(0.00).Comment("数字类型")),
		m.Column("text", m.Text().Comment("字符类型")),
		m.Column("create_at", m.Datetime().Comment("创建时间").Index()),
		m.Column("name", m.Varchar(50).Comment("name")),
	}, "CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ENGINE=InnoDB COMMENT = 'table comment'")

	m.RenameTable("new_table", "right_table")

	m.CreateIndex("dec", "right_table", "dec")
	m.CreateIndex("name", "right_table", "name")

	m.AddColumn("right_table", "remark", m.Mediumtext().Comment("备注").After("text"))
	m.RenameColumn("right_table", "remark", "note")
	m.AlterColumn("right_table", "note", "note", m.Tinytext().Comment("123123").After("dec"))
	m.DropColumn("right_table", "name")

	m.DropIndex("right_table", "dec")
	m.DropTable("right_table")
}

// down
func (m *m200805_075136_create_new_table) Down() {

}
