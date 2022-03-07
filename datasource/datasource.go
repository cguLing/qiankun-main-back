package datasource

import (
	"bus-backend-go/conf"
	"bus-backend-go/model"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}
func init() {
	dsn := conf.Sysconfig.DBUserName + ":" + conf.Sysconfig.DBPassword + "@tcp(" + conf.Sysconfig.DBIp + ":" + conf.Sysconfig.DBPort + ")/" + conf.Sysconfig.DBName + "?charset=utf8&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 彩色打印
		},
	)
	var err error
	db, err = gorm.Open(mysql.New(mysql.Config{

		DSN:                      dsn,  // data source name
		DefaultStringSize:        256,  // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		//DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		//DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{
		Logger:         newLogger,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Println(err)
		panic(err)
	}
	// 自动迁移表结构
	AutoMigrateTables(db)
}

// 更新数据库表结构
func AutoMigrateTables(db *gorm.DB) {
	db.AutoMigrate(
		&model.SuperAdmin{},
		&model.ServiceType{},
		&model.MicroList{},
		&model.ServiceList{},
		&model.EnshrineList{},
		)
}
