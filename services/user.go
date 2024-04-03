package services

import (
	"context"
	"myapp/models"
	"myapp/tools"

	"gorm.io/gorm"
)

func (s *Service) CreateUser(ctx context.Context, userInput *models.UserCreateInput) (*models.UserReturn, error) {
	hashedPassword, err := tools.HashAndSalt(userInput.Password)
	if err != nil {
		return nil, err
	}

	newUser := models.User{
		Username: userInput.Username,
		Password: hashedPassword,
	}

	userInput.Password = hashedPassword

	if err := s.DB.Model(&newUser).Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &models.UserReturn{
		ID:       newUser.ID,
		Username: newUser.Username,
	}, nil
}

func (s *Service) GetUsers(ctx context.Context) ([]*models.UserReturn, error) {
	var (
		users  []*models.User
		_users []*models.UserReturn
	)
	if err := s.DB.Find(&users).Error; err != nil {
		panic(err)
	}

	for _, user := range users {
		_users = append(_users, &models.UserReturn{
			ID:       user.ID,
			Username: user.Username,
		})
	}

	return _users, nil
}

func (s *Service) GetUserByID(ctx context.Context, user *models.User) (*models.UserReturn, error) {
	if err := s.DB.First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return &models.UserReturn{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func (s *Service) UpdateUser(ctx context.Context, user *models.UserUpdateInput) (*models.UserReturn, error) {
	if err := s.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		panic(err)
	}

	return &models.UserReturn{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func (s *Service) DeleteUser(ctx context.Context, user *models.User) (string, error) {
	if err := s.DB.First(&user).Error; err == gorm.ErrRecordNotFound {
		return "User not found", err
	} else if err != nil {
		return "Error", err
	}

	if err := s.DB.Delete(&user).Error; err != nil {
		return "Error Deleting User", err
	}

	return "Success", nil
}

func (s *Service) Login(ctx context.Context, user *models.UserLoginInput) (string, error) {
	var foundUser *models.User

	if err := s.DB.Model(&foundUser).Where("username = ?", user.Username).First(&foundUser).Error; err == gorm.ErrRecordNotFound {
		return "Username or password doesn't match", err
	} else if err != nil {
		return "Error", err
	}

	if err := tools.ComparePasswords(foundUser.Password, user.Password); err != nil {
		return "Username or password doesn't match", err
	}

	token := tools.JwtGenerate(foundUser.ID)

	return token, nil

}

func (s *Service) GetMe(ctx context.Context) (*models.UserReturn, error) {
	var user *models.User
	userId := tools.AuthCtx(ctx).ID

	if err := s.DB.Model(&user).Where("id = ?", userId).First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	return &models.UserReturn{
		ID:       user.ID,
		Username: user.Username,
	}, nil
}

func (s *Service) GetMyPost(ctx context.Context) (*models.UserWithPostReturn, error) {
	var (
		user   *models.User
		posts  []*models.Post
		_posts []*models.PostReturn
	)
	userId := tools.AuthCtx(ctx).ID

	if err := s.DB.Model(&user).Where("id = ?", userId).First(&user).Error; err == gorm.ErrRecordNotFound {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	if err := s.DB.Model(&posts).Where("user_id = ?", userId).Find(&posts).Error; err != nil {
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

	return &models.UserWithPostReturn{
		UserID:   user.ID,
		Username: user.Username,
		Post:     _posts,
	}, nil
}
