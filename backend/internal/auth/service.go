package auth

import (
	"alumni_hub/internal/alumni"
	"alumni_hub/internal/staff"
	"alumni_hub/internal/utils"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"
)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type Service struct {
	DB     *gorm.DB
	secret []byte
}

func NewService(db *gorm.DB) *Service {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	return &Service{DB: db, secret: []byte(secret)}
}

func (s *Service) AuthenticateUser(username, password, role string) (uint, string, error) {
	if role == "alumni" {
		loginName, err := utils.ExtractLoginName(username)
		if err != nil {
			return 0, "", errors.New("invalid login name")
		}

		var alumnus alumni.Alumni
		if err := s.DB.Where("LOWER(matric_no) LIKE ?", loginName).First(&alumnus).Error; err != nil {
			return 0, "", errors.New("invalid credentials: no username found")
		}

		if !utils.CheckPasswordHash(password, alumnus.Password) {
			return 0, "", errors.New("invalid credentials: password does not match")
		}

		return alumnus.ID, "alumni", nil
	}

	if role == "staff" {
		var st staff.Staff
		if err := s.DB.Where("LOWER(username) = ?", strings.ToLower(username)).First(&st).Error; err != nil {
			return 0, "", errors.New("invalid credentials: no username found")
		}

		if !utils.CheckPasswordHash(password, st.Password) {
			return 0, "", errors.New("invalid credentials: password does not match")
		}

		return st.ID, "staff", nil
	}

	return 0, "", errors.New("invalid role")
}

func (s *Service) GenerateToken(userID uint, role string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * 24 * time.Hour)

	claims := Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  nowTime.Unix(),
			Issuer:    "alumni_hub",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(s.secret)
	return token, err
}

func (s *Service) StoreToken(userID uint, role, token string) error {
	var result *gorm.DB
	if role == "alumni" {
		result = s.DB.Model(&alumni.Alumni{}).Where("id = ?", userID).Update("token", token)
	} else if role == "staff" {
		result = s.DB.Model(&staff.Staff{}).Where("id = ?", userID).Update("token", token)
	} else {
		return errors.New("invalid role")
	}

	if result.Error != nil {
		log.Printf("Failed to store token for user %d: %v", userID, result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Printf("No rows affected when storing token for user %d", userID)
		return errors.New("no rows affected")
	}

	return nil
}

func (s *Service) ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return s.secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
