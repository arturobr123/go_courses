package repository

import (
	"context"
	"platzi/go/rest_websockets/models"
)

type Repository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserById(ctx context.Context, id string) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetPostById(ctx context.Context, id string) (*models.Post, error)
	UpdatePost(ctx context.Context, post *models.Post) error
	InsertPost(ctx context.Context, post *models.Post) error
	DeletePost(ctx context.Context, id string, userId string) error
	ListPosts(ctx context.Context, page uint64) ([]*models.Post, error)
	Close()
	// GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

// Abstracciones
// Handler - GetUserByIdPostgres...
// Concretas

var implmentation Repository

func SetRepository(r Repository) {
	implmentation = r
}

func InsertUser(ctx context.Context, user *models.User) error {
	return implmentation.InsertUser(ctx, user)
}

func InsertPost(ctx context.Context, post *models.Post) error {
	return implmentation.InsertPost(ctx, post)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	return implmentation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implmentation.GetUserByEmail(ctx, email)
}

func GetPostById(ctx context.Context, id string) (*models.Post, error) {
	return implmentation.GetPostById(ctx, id)
}

func UpdatePost(ctx context.Context, post *models.Post) error {
	return implmentation.UpdatePost(ctx, post)
}

func DeletePost(ctx context.Context, id string, userId string) error {
	return implmentation.DeletePost(ctx, id, userId)
}

func ListPosts(ctx context.Context, page uint64) ([]*models.Post, error) {
	return implmentation.ListPosts(ctx, page)
}

func Close() {
	implmentation.Close()
}
