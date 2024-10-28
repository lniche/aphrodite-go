package service

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/model"
	"aphrodite-go/internal/repository"
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(ctx context.Context, clientIp string, req *v1.LoginReq) (string, error)
	GetUser(ctx context.Context, userCode string) (*v1.GetUserRespData, error)
	UpdateUser(ctx context.Context, userCode string, req *v1.UpdateUserReq) error
	SendVerifyCode(ctx context.Context, req *v1.SendVerifyCodeReq) error
	Logout(ctx context.Context, userCode string) error
	DeleteUser(ctx context.Context, userCode string) error
}

func NewUserService(
	service *Service,
	userRepository repository.UserRepository,
) UserService {
	return &userService{
		userRepository: userRepository,
		Service:        service,
	}
}

type userService struct {
	userRepository repository.UserRepository
	*Service
}

func (s *userService) Login(ctx context.Context, clientIp string, req *v1.LoginReq) (string, error) {
	// check phone
	user, err := s.userRepository.GetByPhone(ctx, req.Phone)
	if err != nil {
		return "", v1.ErrInternalServerError
	}
	if user == nil && req.OpenId != "" {
		user, err = s.userRepository.GetByOpenId(ctx, req.OpenId)
		if err != nil {
			return "", v1.ErrInternalServerError
		}
	}
	if user != nil {
		user.LoginAt = time.Now()
		if user.OpenId == "" {
			user.OpenId = req.OpenId
		}
		if user.Phone == "" {
			user.Phone = req.Phone
		}
		if err = s.userRepository.UpdateProfile(ctx, user); err != nil {
			return "", err
		}
	} else {
		// check verify code
		storedCode, err := s.userRepository.GetVerifyCode(ctx, req.Phone)
		if err != nil {
			return "", fmt.Errorf("cache code not exist")
		}
		if storedCode != req.VerifyCode {
			return "", fmt.Errorf("verify code check fail")
		}
		// Generate user code and no
		userCode, err := s.sid.GenUint64()
		if err != nil {
			return "", err
		}
		userNo, err := s.userRepository.GenerateUserNo(ctx)
		if err != nil {
			return "", err
		}
		user = &model.User{
			Nickname: "SUGAR" + req.Phone[len(req.Phone)-4:],
			UserCode: strconv.FormatUint(userCode, 10),
			UserNo:   uint64(100000 + userNo),
			Phone:    req.Phone,
			ClientIp: clientIp,
			OpenId:   req.OpenId,
			LoginAt:  time.Now(),
		}
		// Transaction demo
		err = s.tm.Transaction(ctx, func(ctx context.Context) error {
			// CreateProfile a user
			if err = s.userRepository.CreateProfile(ctx, user); err != nil {
				return err
			}
			return nil
		})
	}
	token, err := s.jwt.GenToken(user.UserCode, time.Now().Add(time.Hour*24*30))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetUser(ctx context.Context, userCode string) (*v1.GetUserRespData, error) {
	user, err := s.userRepository.GetByCodeWithCache(ctx, userCode)
	if err != nil {
		return nil, err
	}

	return &v1.GetUserRespData{
		UserNo:   strconv.FormatUint(user.UserNo, 10),
		UserCode: user.UserCode,
		Nickname: user.Nickname,
		Email:    desensitizeEmail(user.Email),
		Phone:    desensitizePhone(user.Phone),
	}, nil
}

func (s *userService) UpdateUser(ctx context.Context, userCode string, req *v1.UpdateUserReq) error {
	user, err := s.userRepository.GetByCode(ctx, userCode)
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
		storedCode, err := s.userRepository.GetVerifyCode(ctx, req.OldPhone)
		if err != nil {
			return err
		}
		if storedCode != req.VerifyCode {
			return fmt.Errorf("verify code check fail")
		}
		user.Phone = req.NewPhone
	}

	if err = s.userRepository.UpdateProfile(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s *userService) SendVerifyCode(ctx context.Context, req *v1.SendVerifyCodeReq) error {
	code := generateVerificationCode()
	s.logger.Info("send verify code", zap.String("code", code), zap.String("phone", req.Phone))
	storedCode, err := s.userRepository.GetVerifyCode(ctx, req.Phone)
	if storedCode != "" {
		return fmt.Errorf("verify code already sent, please wait 1 minute")
	}
	// TODO real send msg
	err = s.userRepository.CacheVerifyCode(ctx, req.Phone, code)
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

func (s *userService) Logout(ctx context.Context, userCode string) error {
	err := s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := s.userRepository.Logout(ctx, userCode); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) DeleteUser(ctx context.Context, userCode string) error {
	err := s.tm.Transaction(ctx, func(ctx context.Context) error {
		if err := s.userRepository.DeleteUser(ctx, userCode); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
