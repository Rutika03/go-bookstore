package cofig

import(
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"

)

var {
	bd *gorm.DB
}

func Connect(){
	d, err := gorm.Open("mysql","rutika:Pavni@123/simpleinterest?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db=d
}

func GetDB() *gorm.DB{
	return db
}