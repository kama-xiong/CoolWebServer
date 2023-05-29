package datasource

import (
	"WebServer/models"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
)

var Db *gorm.DB

func init() {
	var err error
	var dsn string
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn, err = getDsn("e:/dbconfig.json")
	if err != nil {
		panic("Failed to Open Database ConfigFile")
	}
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	Db.AutoMigrate(&models.Category{}, &models.Product{}) //自动迁移 跟py一样
}

func getDsn(filename string) (string, error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Printf("error opening dsn json file")
		return "", err
	}
	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading json file")
		return "", err
	}
	var dsn DSN
	json.Unmarshal(jsonData, &dsn)
	//fmt.Println(dsn)
	//"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsnStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", dsn.User, dsn.Pass, dsn.Host, dsn.Port, dsn.Dbname)
	//fmt.Printf(dsnStr)
	return dsnStr, nil

}

type DSN struct {
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Dbname string `json:"dbname"`
}
