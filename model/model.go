package model
import (
	"github.com/jinzhu/gorm"
)

var Model *gorm.DB

func init() {
	var err error
	//log.Println(config.GetEnv().Database.FormatDSN())
	Model, err = gorm.Open("mysql", "root:sct123@tcp(127.0.0.1:3306)/fastadmin?charset=utf8&parseTime=True&loc=Local")
	//Model, err = gorm.Open("mysql", config.GetEnv().Database.FormatDSN())

	if err != nil {
		panic(err)
	}
}

