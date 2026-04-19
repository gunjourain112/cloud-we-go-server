package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gunjourain112/cloud-we-go-server/gin/internal/infra/config"
	"github.com/gunjourain112/cloud-we-go-server/gin/internal/domain/user"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) error
	Login(ctx context.Context, req *LoginRequest) (*TokenResponse, error)
}

type service struct {
	userRepo user.Repository
	cfg      *config.Config
}

func NewService(userRepo user.Repository, cfg *config.Config) Service {
	return &service{userRepo: userRepo, cfg: cfg}
}

func (s *service) Register(ctx context.Context, req *RegisterRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = s.userRepo.Create(ctx, req.Email, string(hash), req.Name)
	return err
}

func (s *service) Login(ctx context.Context, req *LoginRequest) (*TokenResponse, error) {
	u, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.ID.String(),
		"exp": time.Now().Add(time.Duration(s.cfg.JWT.ExpireHours) * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}

	return &TokenResponse{AccessToken: tokenString}, nil
}
