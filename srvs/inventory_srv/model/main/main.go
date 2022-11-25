package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"srvs/inventory_srv/model"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/ecs_inventory_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL Threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color printing
		},
	)

	// global mode
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	_ = db.AutoMigrate(&model.Inventory{}, &model.StockSellDetail{})
	orderDetail := model.StockSellDetail{
		OrderSn: "L2ncE",
		Status:  1,
		Detail:  []model.GoodsDetail{{1, 2}, {2, 3}},
	}
	db.Create(&orderDetail)
}
