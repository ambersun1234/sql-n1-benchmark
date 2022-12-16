package database

import (
	"database/sql"
	"fmt"

	"sql-n1-benchmark/database/post"
	"sql-n1-benchmark/database/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	tableList = []interface{}{
		&user.User{}, &post.Post{},
	}
)

type DatabaseClient struct {
	Conn *sql.DB
	DB   *gorm.DB
}

func NewDatabaseClient(dbName string) (*DatabaseClient, error) {
	db := &DatabaseClient{}
	if err := db.newConnection(dbName); err != nil {
		return nil, err
	}
	if err := db.newDatabase(); err != nil {
		return nil, err
	}

	return db, nil
}

func (c *DatabaseClient) newConnection(dbName string) error {
	url := fmt.Sprintf("root:@tcp(127.0.0.1:3306)/%v?charset=utf8mb4&parseTime=True&loc=Local", dbName)
	conn, err := sql.Open("mysql", url)
	if err != nil {
		return err
	}

	c.Conn = conn

	return nil
}

func (c *DatabaseClient) newDatabase() error {
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: c.Conn}), &gorm.Config{})
	if err != nil {
		return err
	}

	c.DB = db

	return nil
}

func (c *DatabaseClient) AutoMigration() error {
	return c.DB.AutoMigrate(tableList...)
}

func (c *DatabaseClient) Init(size int) error {
	userList := make([]interface{}, 0)
	postList := make([]interface{}, 0)

	for i := 0; i < size; i++ {
		if i%10 == 0 {
			userList = append(userList, &user.User{
				FirstName: fmt.Sprintf("ambersun%v", i),
				LastName:  "",
			})
		}

		postList = append(postList, &post.Post{
			Title:         fmt.Sprintf("post%v", i),
			Description:   fmt.Sprintf("description%v", i),
			LikesCount:    10,
			CommentsCount: 10,
			UserID:        uint((i / 10) + 1),
		})
	}

	for _, data := range append(userList, postList...) {
		if err := c.DB.Create(data).Error; err != nil {
			return err
		}
	}

	return nil
}
