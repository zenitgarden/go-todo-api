package service

import (
	"database/sql"
	"time"
	"todo-api/app/exception"
	"todo-api/helper"
	"todo-api/model/domain"
	"todo-api/model/web"
	"todo-api/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Register(c *fiber.Ctx, r web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, errorDB := service.DB.Begin()
	helper.PanicIfError(errorDB)
	defer helper.CommitOrRollback(tx)

	_, err = service.UserRepository.FindByUsername(c, tx, r.Username)
	if err == nil {
		panic(exception.NewUniqueError("username"))
	}

	password := []byte(r.Password)
	password, err = bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user := domain.User{Name: r.Name, Username: r.Username, Password: string(password)}
	user = service.UserRepository.Save(c, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(c *fiber.Ctx, r web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, errorDB := service.DB.Begin()
	helper.PanicIfError(errorDB)
	defer helper.CommitOrRollback(tx)

	user := domain.User{Name: r.Name, Username: r.Username}
	user = service.UserRepository.Update(c, tx, user)
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(c *fiber.Ctx, r web.LoginRequest) web.LoginResponse {
	err := service.Validate.Struct(r)
	helper.PanicIfError(err)

	tx, errorDB := service.DB.Begin()
	helper.PanicIfError(errorDB)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(c, tx, r.Username)
	isValidPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.Password))
	if err != nil || isValidPassword != nil {
		panic(exception.NewLoginError("username or password wrong"))
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"name":     user.Name,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte("secret"))
	helper.PanicIfError(err)

	return web.LoginResponse{Token: t}
}
