package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/naoina/toml"
)

type tomlConfig struct {
	Title    string
	Database database
}

type database struct {
	Server   string
	Port     int
	Charset  string
	Dbname   string `toml:"db"`
	Username string
	Password string
}

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (u database) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", u.Username, u.Password, u.Server, u.Port, u.Dbname, u.Charset)
}

// getConfigFilePath: 获取config file 的地址
func getConfigFilePath(configFlag string) string {

	if configFlag == "" {
		if runtime.GOOS == "darwin" {
			return "~/.jindowin/config.toml"
		}
		return "/root/.jindowin/config.toml"
	}

	return configFlag
}

func main() {

	configPtr := flag.String("c", "", "config file dest")
	dbinitPtr := flag.Bool("d", false, "db init flag default false")
	flag.Parse()

	configFile := getConfigFilePath(*configPtr)

	fmt.Println("config file:", configFile)
	fmt.Println("dbinit flag:", *dbinitPtr)

	f, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	var config tomlConfig
	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		panic(err)
	}
	fmt.Println("title", config.Title)
	fmt.Println("Database", config.Database)

	db, err := gorm.Open("mysql", config.Database.String())
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	// db.Delete(&product)
}
