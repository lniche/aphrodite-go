package service

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/model"
	"aphrodite-go/internal/repository"
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) error
	Login(ctx context.Context, req *v1.LoginRequest) (string, error)
	GetProfile(ctx context.Context, userCode string) (*v1.GetProfileResponseData, error)
	UpdateProfile(ctx context.Context, userCode string, req *v1.UpdateProfileRequest) error
	SendVerifyCode(ctx context.Context, req *v1.SendVerifyCodeRequest) error
}

func NewUserService(
	service *Service,
	userRepo repository.UserRepository,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	userRepo repository.UserRepository
	*Service
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) error {
	// check phone
	user, err := s.userRepo.GetByPhone(ctx, req.Phone)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if err == nil && user != nil {
		return v1.ErrPhoneAlreadyUse
	}
	// check verify code
	storedCode, err := s.userRepo.GetVerifyCode(ctx, req.Phone)
	if err != nil {
		return err
	}
	if storedCode != req.VerifyCode {
		return fmt.Errorf("verify code check fail")
	}

	// Generate user code and no
	userCode, err := s.sid.GenString()
	if err != nil {
		return err
	}
	userNo, err := s.userRepo.GenerateUserNo(ctx)
	if err != nil {
		return err
	}
	// set default if nickname is blank
	if len(req.Nickname) == 0 {
		req.Nickname = "SUGAR" + req.Phone[len(req.Phone)-4:]
	}
	user = &model.User{
		UserCode: userCode,
		UserNo:   userNo,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
	}
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	// Transaction demo
	err = s.tm.Transaction(ctx, func(ctx context.Context) error {
		// Create a user
		if err = s.userRepo.Create(ctx, user); err != nil {
			return err
		}
		// TODO: other repo
		return nil
	})
	return err
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (string, error) {
	user, err := s.userRepo.GetByPhone(ctx, req.Phone)
	if err != nil || user == nil {
		return "", v1.ErrUnauthorized
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}
	token, err := s.jwt.GenToken(user.UserCode, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, userCode string) (*v1.GetProfileResponseData, error) {
	user, err := s.userRepo.GetByCode(ctx, userCode)
	if err != nil {
		return nil, err
	}

	return &v1.GetProfileResponseData{
		UserNo:   strconv.FormatInt(user.UserNo, 10),
		UserCode: user.UserCode,
		Nickname: user.Nickname,
		Email:    desensitizeEmail(user.Email),
		Phone:    desensitizePhone(user.Phone),
	}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, userCode string, req *v1.UpdateProfileRequest) error {
	user, err := s.userRepo.GetByCode(ctx, userCode)
	if err != nil {
		return err
	}
	if req.Nickname != "" {
		user.Nickname = req.Nickname
	}
	if req.Email != "" {
		user.Email = req.Email
	}
	if req.OldPassword != "" {
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.OldPassword))
		if err != nil {
			return err
		}

		if req.NewPassword != "" {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			user.Password = string(hashedPassword)
		}
	}
	if req.VerifyCode != "" && req.OldPhone != "" && req.NewPhone != "" {
		storedCode, err := s.userRepo.GetVerifyCode(ctx, req.OldPhone)
		if err != nil {
			return err
		}
		if storedCode != req.VerifyCode {
			return fmt.Errorf("verify code check fail")
		}
		user.Phone = req.NewPhone
	}

	if err = s.userRepo.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
func (s *userService) SendVerifyCode(ctx context.Context, req *v1.SendVerifyCodeRequest) error {
	code := generateVerificationCode()
	storedCode, err := s.userRepo.GetVerifyCode(ctx, req.Phone)
	if err != nil {
		return err
	}
	if storedCode != "" {
		return fmt.Errorf("verify code already sent, please wait 1 minute")
	}
	// TODO real send msg
	err = s.userRepo.CacheVerifyCode(ctx, req.Phone, code)
	if err != nil {
		return err
	}
	return nil
}

func generateVerificationCode() string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return fmt.Sprintf("%04d", r.Intn(10000))
}

func desensitizePhone(phone string) string {
	re := regexp.MustCompile(`(\d{3})\d{4}(\d{4})`)
	return re.ReplaceAllString(phone, "$1****$2")
}

func desensitizeEmail(email string) string {
	re := regexp.MustCompile(`(\w{3})[\w.-]+@([\w.]+)`)
	return re.ReplaceAllString(email, "$1***@$2")
}
