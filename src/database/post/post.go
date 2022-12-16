package post

import (
	"sql-n1-benchmark/database/user"
)

type Post struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Title         string `gorm:"column:title" json:"title"`
	Description   string `gorm:"column:description" json:"description"`
	LikesCount    int    `gorm:"column:likes_count" json:"likesCount"`
	CommentsCount int    `gorm:"column:comments_count" json:"commentsCount"`

	UserID uint
	User   user.User `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (p *Post) TableName() string {
	return "post"
}
