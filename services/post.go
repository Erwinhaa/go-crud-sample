package services

import (
	"context"
	"myapp/models"
)

func (s *Service) GetPosts(ctx context.Context) ([]*models.PostReturn, error) {
	var posts []*models.Post
	var _posts []*models.PostReturn
	if err := s.DB.Model(&posts).Find(&posts).Error; err != nil {
		return nil, err
	}

	for _, post := range posts {
		_posts = append(_posts, &models.PostReturn{
			ID:    post.ID,
			Title: post.Title,
			Body:  post.Body,
		})
	}

	return _posts, nil
}

func (s *Service) CreatePost(ctx context.Context, input models.Post) (*models.PostReturn, error) {
	post := models.Post{
		Title: input.Title,
		Body:  input.Body,
	}

	if err := s.DB.Model(&post).Create(&post).Error; err != nil {
		return nil, err
	}

	return &models.PostReturn{
		ID:    post.ID,
		Title: post.Title,
		Body:  post.Body,
	}, nil
}
