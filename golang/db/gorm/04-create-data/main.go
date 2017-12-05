package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Cloud struct {
	gorm.Model
	Account string `gorm:"size:20", json:",omitempty"`
	Cloud   string `gorm:"size:20"`
}

func (Cloud) TableName() string {
	return "cloud"
}

type SecurityGroup struct {
	gorm.Model

	CloudId int
	Cloud   Cloud

	GroupId     string
	VpcId       string
	GroupName   string `gorm:"column:groupname"`
	Description string
}

func (SecurityGroup) TableName() string {
	return "securitygroup"
}

func main() {
	SQLConnectingString := "root:@tcp(127.0.0.1:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", SQLConnectingString)
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	db.LogMode(true)

	sg := SecurityGroup{CloudId: 1, GroupId: "xxx", GroupName: "xxx_name", VpcId: "Vpc_xxx"}
	db.Create(&sg)

	sg2 := SecurityGroup{CloudId: 2, GroupId: "xxx_2", GroupName: "xxx_name", VpcId: "Vpc_xxx"}
	db.Create(&sg2)

	var cloud Cloud
	db.Model(&sg2).Related(&cloud)
	fmt.Println(cloud)
}
