package repository

import (
	v1 "aphrodite-go/api/v1"
	"aphrodite-go/internal/constant"
	"aphrodite-go/internal/model"
	"context"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateProfile(ctx context.Context, user *model.User) error
	UpdateProfile(ctx context.Context, user *model.User) error
	GetByCode(ctx context.Context, userCode string) (*model.User, error)
	GetByCodeWithCache(ctx context.Context, userCode string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByPhone(ctx context.Context, phone string) (*model.User, error)
	GetByOpenId(ctx context.Context, openId string) (*model.User, error)
	GenerateUserNo(ctx context.Context) (int64, error)
	CacheVerifyCode(ctx context.Context, phone string, code string) error
	GetVerifyCode(ctx context.Context, phone string) (string, error)
	Logout(ctx context.Context, userCode string) error
	DeleteUser(ctx context.Context, userCode string) error
}

func NewUserRepository(
	repository *Repository,
) UserRepository {
	return &userRepository{
		Repository: repository,
	}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) CreateProfile(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateProfile(ctx context.Context, user *model.User) error {
	if err := r.DB(ctx).Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetByCode(ctx context.Context, userCode string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("user_code = ?", userCode).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByCodeWithCache(ctx context.Context, userCode string) (*model.User, error) {
	cacheData, err := r.rdb.Get(ctx, constant.USER+userCode).Result()
	if err == nil {
		var user model.User
		if err := json.Unmarshal([]byte(cacheData), &user); err != nil {
			return nil, err
		}
		return &user, nil
	}

	var user model.User
	if err := r.DB(ctx).Where("user_code = ?", userCode).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, v1.ErrNotFound
		}
		return nil, err
	}

	userJSON, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	if err := r.rdb.Set(ctx, constant.USER+userCode, userJSON, 0).Err(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("phone = ?", phone).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByOpenId(ctx context.Context, openId string) (*model.User, error) {
	var user model.User
	if err := r.DB(ctx).Where("openId = ?", openId).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GenerateUserNo(ctx context.Context) (int64, error) {
	script := `
    local key = KEYS[1]
    local exists = redis.call('EXISTS', key)
    if exists == 0 then
        redis.call('SET', key, 100000)
    end
    return redis.call('INCR', key)
    `
	result, err := r.rdb.Eval(ctx, script, []string{constant.NEXTIDUNO}).Result()
	if err != nil {
		return 0, err
	}
	return result.(int64), nil
}

func (r *userRepository) CacheVerifyCode(ctx context.Context, phone string, code string) error {
	return r.rdb.Set(ctx, phone, code, 90*time.Second).Err()
}

func (r *userRepository) GetVerifyCode(ctx context.Context, phone string) (string, error) {
	storedCode, err := r.rdb.Get(ctx, phone).Result()
	if err != nil {
		return "", err
	}
	return storedCode, nil
}

func (r *userRepository) Logout(ctx context.Context, userCode string) error {
	user, err := r.GetByCode(ctx, userCode)
	if err != nil {
		return err
	}
	if err = r.DB(ctx).Model(&model.User{}).Where("user_code = ?", userCode).Update("login_token", "").Error; err != nil {
		return err
	}
	expiration := 30 * 24 * time.Hour
	if err = r.rdb.Set(ctx, user.LoginToken, constant.JWT_BLACKLIST, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userCode string) error {
	now := time.Now()
	if err := r.DB(ctx).Model(&model.User{}).Where("user_code =?", userCode).Updates(map[string]interface{}{
		"deleted_at": now,
		"is_deleted": true,
	}).Error; err != nil {
		return err
	}
	return nil
}
