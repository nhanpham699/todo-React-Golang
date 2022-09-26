package user

import (
	"context"
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/nhanpham699/demo/dto"
	"golang.org/x/crypto/bcrypt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var dbName = "todoDB"
var docCollection = "users"
var SECRET_KEY = []byte("gosecretkey")

type UserRepositoryDb struct {
	client *mongo.Client
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString(SECRET_KEY)
	if err != nil {
		log.Fatal("Error in JWT token generation")
	}
	return tokenString, nil
}

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}

func (d UserRepositoryDb) Create(data dto.RegisterRequest) (*dto.HandleResponse, *error) {
	data.Password = getHash([]byte(data.Password))
	db := d.client.Database(dbName).Collection(docCollection)
	_, err := db.InsertOne(context.TODO(), data)
	if err != nil {
		log.Fatal(err)
	}
	response := &dto.HandleResponse{
		ResultCode: "00",
		Message:    "created successfully",
	}
	return response, nil
}

func (d UserRepositoryDb) Login(data dto.AuthRequest) (*dto.AuthResponse, *error) {
	db := d.client.Database(dbName).Collection(docCollection)
	cur := db.FindOne(context.TODO(), bson.M{"email": data.Email})
	var user User
	cur.Decode(&user)
	userPass := []byte(data.Password)
	dbPass := []byte(user.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		response := &dto.AuthResponse{
			ResultCode: "01",
			Message:    "Login failed",
			Token:      "",
		}
		return response, nil
	}
	jwtToken, err := GenerateJWT()
	fmt.Print(err)

	if err == nil {
		response := &dto.AuthResponse{
			ResultCode: "00",
			Message:    "Login successfully",
			Token:      jwtToken,
		}
		return response, nil
	}

	response := &dto.AuthResponse{
		ResultCode: "99",
		Message:    "System Error",
		Token:      "",
	}
	return response, nil
}

func NewUserRepositoryDb(dbClient *mongo.Client) UserRepositoryDb {
	return UserRepositoryDb{dbClient}
}
