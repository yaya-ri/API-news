package db

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // MySQL driver
)

//DB Global bariabel for response initial DB sql
var DB *gorm.DB

//InitDBSQL Initial database sql
func InitDBSQL() (*gorm.DB, error) {

	var (
		USERNAME = os.Getenv("SQL_USER")
		PASSWORD = os.Getenv("SQL_PASSWORD")
		HOST     = os.Getenv("SQL_HOST")
		PORT     = os.Getenv("SQL_PORT")
		DATABASE = os.Getenv("SQL_DATABASE")

		mysqlCon = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
			USERNAME,
			PASSWORD,
			HOST,
			PORT,
			DATABASE,
		)
	)

	DB, err := gorm.Open("mysql", mysqlCon)

	if err == nil {
		greenOutput := color.New(color.FgGreen)
		successOutput := greenOutput.Add(color.Bold)
		successOutput.Println("")
		successOutput.Println("!!! Info")
		successOutput.Println(fmt.Sprintf("Successfully connected to database %s", mysqlCon))
		successOutput.Println("")
	}

	return DB, err
}
