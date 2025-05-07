package service

import (
	"errors"
	"github.com/ParthSharma272/GoStock/internal/domain/common"                 // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/domain/user"                   // Replace your_project_path
	infraAuth "github.com/ParthSharma272/GoStock/internal/infrastructure/auth" // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/config"         // Replace your_project_path

	"gorm.io/gorm"
)

type AuthService struct {
	db        *gorm.DB
	userRepo  user.Repository
	appConfig *config.Config
}

func NewAuthService(db *gorm.DB, userRepo user.Repository, cfg *config.Config) *AuthService {
	return &AuthService{
		db:        db,
		userRepo:  userRepo,
		appConfig: cfg,
	}
}

func (s *AuthService) Register(firstName, lastName, email, password string, role common.UserRole) (*user.User, string, error) {
	existingUser, _ := s.userRepo.FindByEmail(s.db, email)
	if existingUser != nil {
		return nil, "", errors.New("user with this email already exists")
	}

	newUser := &user.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password, // Will be hashed by BeforeSave hook
		Role:      role,
		IsActive:  true,
	}

	if err := s.userRepo.Create(s.db, newUser); err != nil {
		return nil, "", errors.New("failed to register user: " + err.Error())
	}

	token, err := infraAuth.GenerateToken(newUser, s.appConfig)
	if err != nil {
		return nil, "", errors.New("failed to generate token: " + err.Error())
	}

	return newUser, token, nil
}

func (s *AuthService) Login(email, password string) (*user.User, string, error) {
	u, err := s.userRepo.FindByEmail(s.db, email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, "", errors.New("invalid email or password")
		}
		return nil, "", err
	}

	if !u.IsActive {
		return nil, "", errors.New("account is not active")
	}

	if !u.CheckPassword(password) {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := infraAuth.GenerateToken(u, s.appConfig)
	if err != nil {
		return nil, "", errors.New("failed to generate token: " + err.Error())
	}
	return u, token, nil
}

func (s *AuthService) GetUserByID(id uint) (*user.User, error) {
	return s.userRepo.FindByID(s.db, id)
}
