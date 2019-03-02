package datasources

import (
	"log"
	"os"
	"start/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func MySqlConnect(isTesting bool) {
	var err error
	config := config.GetConfig()
	if isTesting == false {
		db, err = gorm.Open("mysql", config.GetString("mysql.user")+":"+
			config.GetString("mysql.password")+"@"+
			"tcp("+config.GetString("mysql.host")+":"+config.GetString("mysql.port")+")"+
			"/"+config.GetString("mysql.db_name")+
			"?charset=utf8&parseTime=True")
		// db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
		log.Println("real db")
		// =======
		// db.LogMode(true)
		// =======
	} else {
		pwd, err := os.Getwd()
		if err != nil {
			glog.Fatalf("Error get current path, %s", err)
		}
		db, err = gorm.Open("sqlite3", pwd+"/test/caro.db")
		if err != nil {
			glog.Fatalf("Error connect database test")
		}

		log.Println("test db")
	}
	if err != nil {
		panic("failed to connect database")
	}
}

func GetDatabase() *gorm.DB {
	return db
}
