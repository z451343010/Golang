package tool

import (
	"fmt"
	"gin/CloudRestaurant/model"

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
	fmt.Println("conn string = ", conn)
	engine, err := xorm.NewEngine("mysql", conn)

	if err != nil {
		return nil, err
	}

	// 显示 SQL 语句
	engine.ShowSQL(database.ShowSql)

	// 把结构体 SmsCode 映射成数据库的一张表
	err = engine.Sync2(new(model.SmsCode))
	if err != nil {
		fmt.Println("把 SmsCode 映射成表时出错")
		return nil, err
	}

	// 把结构体 Member 映射成数据库的一张表
	err = engine.Sync2(new(model.Member))
	if err != nil {
		fmt.Println("把 Member 映射成表时出错")
		return nil, err
	}

	// 把结构体 FoodCategory 映射成数据库的一张表
	err = engine.Sync2(new(model.FoodCategory))
	if err != nil {
		fmt.Println("把 FoodCategory 映射成表时出错")
		return nil, err
	}

	// 把结构体 Shop 映射成数据库的一张表
	err = engine.Sync2(new(model.Shop))
	if err != nil {
		fmt.Println("把 Shop 映射成表时出错")
		return nil, err
	}

	// 把结构体 Service 映射成数据库的一张表
	err = engine.Sync2(new(model.Service))
	if err != nil {
		fmt.Println("把 Service 映射成表时出错")
		return nil, err
	}

	// 把结构体 ShopService 映射成数据库的一张表
	err = engine.Sync2(new(model.ShopService))
	if err != nil {
		fmt.Println("把 ShopService 映射成表时出错")
		return nil, err
	}

	// 把结构体 Goods 映射成数据库的一张表
	err = engine.Sync2(new(model.Goods))
	if err != nil {
		fmt.Println("把 Goods 映射成表时出错")
		return nil, err
	}

	orm := new(Orm)
	orm.Engine = engine

	DbEngine = orm
	return DbEngine, nil

}

func GetOrmConnect() (*Orm, error) {

	// 解析配置文件
	cfg, err := ParseConfig("../config/app.json")
	if err != nil {
		panic(err.Error())
	}

	conn, err := OrmEngine(cfg)
	return conn, err

}
