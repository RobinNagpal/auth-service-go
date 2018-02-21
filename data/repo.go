package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


func main() {
	db, err := gorm.Open("postgres", "host=localhost port=myport user=gorm dbname=gorm password=mypassword")
	if err {

	}
	defer db.Close()
}