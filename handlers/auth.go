package handlers

import (
	"context"
	"log"
	"time"

	"auth-service/models"
	"auth-service/proto"
	"auth-service/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
	db *mongo.Collection
}

func NewAuthService() *AuthService {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	db := client.Database("authdb").Collection("users")
	return &AuthService{db: db}
}

func (s *AuthService) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.AuthResponse, error) {
	var existing models.User
	err := s.db.FindOne(ctx, bson.M{"email": req.Email}).Decode(&existing)
	if err == nil {
		return &auth.AuthResponse{Message: "User already exists"}, nil
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email:    req.Email,
		Password: hashedPassword,
		CreatedAt: time.Now(),
	}
	_, err = s.db.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &auth.AuthResponse{Message: "User registered successfully"}, nil
}

func (s *AuthService) Login(ctx context.Context, req *auth.LoginRequest) (*auth.AuthResponse, error) {
	var user models.User
	err := s.db.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return &auth.AuthResponse{Message: "Invalid credentials"}, nil
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return nil, err
	}

	return &auth.AuthResponse{Message: "Login successful", Token: token}, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *auth.ValidateTokenRequest) (*auth.ValidateTokenResponse, error) {
	isValid := utils.ValidateJWT(req.Token)
	return &auth.ValidateTokenResponse{IsValid: isValid}, nil
}
