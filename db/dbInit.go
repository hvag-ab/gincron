package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	// "github.com/go-xorm/core"
	"myproject/pkg/setting"
	"log"
)

var DB *xorm.Engine

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Port		string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{
	Type:setting.Type,
	User:setting.User,
	Password:setting.Password,
	Host:setting.Host,
	Port:setting.Port,
	Name:setting.Name,
	TablePrefix:setting.TablePrefix,
}

func Setup() {
	//读取配置文件
	// setting.MapTo("database", DatabaseSetting)
	// fmt.Println(DatabaseSetting)

	var err error //这里一定要写成 = 如果用:=表示重新申明一个变量会覆盖全局变量DB DB就会变成局部变量就不能导入到其他包中 导入后就会出现空指针引用报错
	DB, err = xorm.NewEngine(DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.Port,
		DatabaseSetting.Name))
	

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	//更改默认表名添加前缀
	// tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, DatabaseSetting.TablePrefix )
	// DB.SetTableMapper(tbMapper)

	//日志是一个接口，通过设置日志，可以显示SQL，警告以及错误等，默认的显示级别为INFO。

	DB.ShowSQL(true)//则会在控制台打印出生成的SQL语句；
	// f, err := os.Create("sql.log")//输出sql语句日志
	// if err != nil {
	// 	println(err.Error())
	// 	return
	// }
	// engine.SetLogger(xorm.NewSimpleLogger(f))
	//设置数据库连接池
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxLifetime(600)

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)//设置查询缓存
	DB.SetDefaultCacher(cacher)

	// defer DB.Close()
}