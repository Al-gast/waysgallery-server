package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindPosts() ([]models.Post, error)
	FindPostsByUserId(ID int) ([]models.Post, error)
	GetPost(ID int) (models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPosts() ([]models.Post, error) {
	var Posts []models.Post
	err := r.db.Preload("User").Find(&Posts).Error

	return Posts, err
}

func (r *repository) FindPostsByUserId(ID int) ([]models.Post, error) {
	var Posts []models.Post
	err := r.db.Preload("User").Where("user_id = ?", ID).Find(&Posts).Error

	return Posts, err
}

func (r *repository) GetPost(ID int) (models.Post, error) {
	var Post models.Post
	err := r.db.Preload("User").First(&Post, ID).Error

	return Post, err
}

func (r *repository) CreatePost(Post models.Post) (models.Post, error) {
	err := r.db.Create(&Post).Error

	return Post, err
}
