package main

import (
	"log"
	"os"
	"time"

	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model                //一些默认字段
	Code       sql.NullString //通过nullstring来设置零值问题
	Price      uint
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/sbi?charset=utf8mb4"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		})

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 迁移 schema
	//db.AutoMigrate(&Product{}) //定义一个表结构将表结构直接生成对应的表 -migrations实例化一个空表此处一个有sql语句
	// Create
	db.Create(&Product{Code: sql.NullString{"D42", true}, Price: 100}) //创建一个例子
	product := Product{
		Code:  sql.NullString{String: "D42", Valid: true},
		Price: 100,
	}
	//db.Create(&product)
	// Read
	//var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	db.Model(&product).Where("id = 1", product.ID).Update("Price", 300)
	// Update - 更新多个字段
	//db.Model(&product).Where("id = ?", product.ID).Updates(Product{Price: 200, Code: sql.NullString{String: "D42", Valid: true}}) // 仅更新非零值字段 将code“”写入只是code没有改变将price设为0则price还是没有改变
	//如果我们去更新一个produce 只设置了price200 会导致其他值设为默认值
	//利用nullstring
	/*product = Product{
		Code:  sql.NullString{"", true},
		Price: 100,
	}*/
	//db.Model(&product).Updates(&product)
	db.Model(&product).Where("id = ?", product.ID).Updates(map[string]interface{}{"Price": 300, "Code": sql.NullString{String: "F42", Valid: true}})
	// Delete - 删除 product
	//db.Delete(&product, 1)
}

// Globally mode

//设置全局的logger，这个logger在我们只想每个sql语句时候会打印一行sql
//sql是最重要的
