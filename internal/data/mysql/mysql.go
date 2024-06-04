package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"kratos_test/internal/data/entity"
	"log"
)

// Mysql 此接口定义了一组方法，交给MysqlClient实现
type Mysql interface {
	InsertOneGrocery(grocery entity.GroceryInfo) QueryResult
}

// MysqlClient 是Mysql接口的实现
type MysqlClient struct {
	DB *gorm.DB
}

// QueryResult 响应结果
type QueryResult struct {
	status    string
	msg       string
	performId any
}

// InsertOneGrocery FIXME 新增一个物品方法
func (mc *MysqlClient) InsertOneGrocery(grocery entity.GroceryInfo) QueryResult {
	newGrocery := entity.GroceryInfo{
		Id:          999,
		Name:        "strawberry",
		Description: "a red fruit",
		Price:       30,
		ImageURL:    "https://baidu.com",
		Category:    "fruit",
	}
	result := mc.DB.Create(&newGrocery)
	if result.Error != nil {
		log.Fatal("新增数据失败: ", result.Error)
	}
	log.Println("新增数据成功: ", newGrocery.Id)
	newResult := QueryResult{
		status:    "200",
		msg:       "新增数据成功",
		performId: newGrocery.Id,
	}
	return newResult
}

// NewMysql !!!生成的客户端实例需要去data.go注册（控制反转）!!!
func NewMysql() (*MysqlClient, error) {
	// 定义你的数据库连接信息
	dsn := "root:123456@tcp(192.168.1.204:3308)/kratos_test?charset=utf8mb4&parseTime=True&loc=Local"
	// 使用GORM打开一个数据库连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// 测试连接是否成功
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to test database connection: %v", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开到数据库的最大连接数。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(0)
	log.Println("connected to mysql！")
	return &MysqlClient{DB: db}, nil
}
