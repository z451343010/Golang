package tool

import (
	"gin/FinancialSystem/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// 定义一个全局变量
var DbEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config) (*Orm, error) {

	database := cfg.Database
	// 数据库连接字符串
	conn := database.User + ":" + database.Password + "@tcp(" +
		database.Host + ":" + database.Port + ")/" +
		database.DbName + "?charset=" + database.Charset

	engine, err := xorm.NewEngine("mysql", conn)

	if err != nil {
		LogRecoder(cfg, "OrmEngine", err)
		return nil, err
	}

	// 显示 SQL 语句
	engine.ShowSQL(database.ShowSql)

	// 把结构体映射成数据库中的表
	err = engine.Sync2(
		// 把结构体 User 映射成数据库的一张表
		new(model.User),
		// 把结构体 Goods 映射成数据库的一张表
		new(model.Goods),
		// 把结构体 PurchaseShoppingCart 映射成数据库的一张表
		new(model.PurchaseShoppingCart),
		// 把结构体 PurchaseOrderDetail 映射成数据库的一张表
		new(model.PurchaseOrderDetail),
		// 把结构体 InStock 映射成数据库的一张表
		new(model.InStock),
		// 把结构体 PurchaseOrder 映射成数据库的一张表
		new(model.PurchaseOrder),
	)
	if err != nil {
		LogRecoder(cfg, "OrmEngine", err)
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine

	DbEngine = orm
	return DbEngine, nil

}
