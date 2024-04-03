package services

import (
	"context"
	"myapp/models"
	"myapp/tools"

	"gorm.io/gorm"
)

func (s *Service) GetPosts(ctx context.Context) ([]*models.PostReturn, error) {
	var posts []*models.Post
	var _posts []*models.PostReturn
	if err := s.DB.Model(&posts).Find(&posts).Error; err != nil {
		return nil, err
	}

	for _, post := range posts {
		_posts = append(_posts, &models.PostReturn{
			ID:     post.ID,
			Title:  post.Title,
			Body:   post.Body,
			UserId: post.UserId,
		})
	}

	return _posts, nil
}

func (s *Service) CreatePost(ctx context.Context, input models.Post) (*models.PostReturn, error) {
	userId := tools.AuthCtx(ctx).ID

	post := models.Post{
		Title:  input.Title,
		Body:   input.Body,
		UserId: userId,
	}

	if err := s.DB.Model(&post).Create(&post).Error; err != nil {
		return nil, err
	}

	return &models.PostReturn{
		ID:     post.ID,
		Title:  post.Title,
		Body:   post.Body,
		UserId: userId,
	}, nil
}

func (s *Service) UpdatePost(ctx context.Context, input models.PostInput) (*models.PostReturn, error) {
	var post models.Post

	if err := s.DB.Model(&post).Where("id = ?", input.ID).Updates(&input).Error; err != nil {
		return nil, err
	}

	return &models.PostReturn{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}, nil
}

func (s *Service) DeletePost(ctx context.Context, post *models.Post) (string, error) {
	if err := s.DB.Delete(&post).Error; err == gorm.ErrRecordNotFound {
		return "", err
	} else if err != nil {
		return "", err
	}

	return "Success", nil
}
