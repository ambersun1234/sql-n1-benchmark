package benchmark

import (
	"fmt"
	"os"
	"time"

	"sql-n1-benchmark/database"
	"sql-n1-benchmark/database/post"
	"sql-n1-benchmark/database/user"

	"gorm.io/gorm"
)

type Data struct {
	user.User `gorm:"embedded"`
	post.Post `gorm:"embedded"`
}

func Benchmark(db *database.DatabaseClient, round int) error {
	var (
		start time.Time
		end   time.Time
	)

	n1File, err := os.Create("./n1_benchmark.txt")
	if err != nil {
		return err
	}
	defer n1File.Close()

	optimizeFile, err := os.Create("./optimize_benchmark.txt")
	if err != nil {
		return err
	}
	defer optimizeFile.Close()

	for i := 0; i < round; i++ {
		start = time.Now()
		n1(db.DB)
		end = time.Now()

		n1File.WriteString(fmt.Sprintf("%v %v\n", i, end.UnixNano()-start.UnixNano()))

		start = time.Now()
		optimize(db.DB)
		end = time.Now()
		optimizeFile.WriteString(fmt.Sprintf("%v %v\n", i, end.UnixNano()-start.UnixNano()))
	}

	return nil
}

func n1(db *gorm.DB) error {
	posts := make([]*post.Post, 0)

	if err := db.Model(&post.Post{}).
		Where(&post.Post{LikesCount: 10}).
		Select("*").
		Find(&posts).Error; err != nil {
		return err
	}

	for _, post := range posts {
		var author user.User
		if err := db.Model(&user.User{}).
			Where(&user.User{ID: uint(post.UserID)}).
			Select("*").
			Take(&author).Error; err != nil {
			return err
		}
	}

	return nil
}

func optimize(db *gorm.DB) error {
	posts := make([]*Data, 0)
	if err := db.Model(&post.Post{}).
		Joins("LEFT JOIN user ON post.user_id = user.id").
		Where(&post.Post{LikesCount: 10}).
		Select("user.*, post.*").
		Find(&posts).Error; err != nil {
		return err
	}

	return nil
}
