package service

import (
	"context"
	"errors"
	"munayfund-api2/internal/core/domain"
	"munayfund-api2/internal/core/port"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var secretKey string = os.Getenv("SECRETKEY")

type UserService struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Login(ctx context.Context, user *domain.User) (string, error) {
	existingUser := &domain.User{}
	err := s.repo.FindOne(ctx, existingUser, bson.M{"email": user.Email})
	if err != nil {
		return "", errors.New("User not found: " + user.Email)
	}

	// Comparar la contraseña proporcionada con la almacenada en la base de datos
	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("Wrong password")
	}

	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 5).Unix(),
		"iat":      time.Now().Unix(),
		"user_exp": existingUser.ID,
	}

	// Crear el token usando el conjunto de claims y la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", errors.New("Error al generar el token JWT")
	}

	return signedToken, nil
}

// Signup registra un nuevo usuario y devuelve un token JWT si el registro es exitoso
func (s *UserService) Signup(ctx context.Context, newUser *domain.User) (string, error) {
	// Verificar si el usuario ya existe
	existingUser := &domain.User{}
	err := s.repo.FindOne(ctx, existingUser, bson.M{"email": newUser.Email})
	if err == nil {
		return "", errors.New("Email already registered")
	} else if err != nil && err != mongo.ErrNoDocuments {
		return "", errors.New("Error checking for existing user")
	}

	// Hash de la contraseña antes de almacenarla
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("Error hashing password")
	}

	// Almacenar el nuevo usuario en la base de datos
	newUser.Password = string(hashedPassword)
	err = s.repo.InsertOne(ctx, newUser)
	if err != nil {
		return "", errors.New("Error creating user")
	}

	// Generar un token JWT para el nuevo usuario
	claims := jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour * 5).Unix(),
		"iat":      time.Now().Unix(),
		"user_exp": newUser.ID,
	}

	// Crear el token usando el conjunto de claims y la clave secreta
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		defer s.repo.DeleteOne(ctx, bson.M{"_id": newUser.ID})

		return "", errors.New("Error generating JWT")
	}

	return signedToken, nil
}

func (s *UserService) Update(ctx context.Context, user *domain.User) (*domain.User, error) {
	// Verificar si el usuario existe
	existingUser := &domain.User{}
	err := s.repo.FindOne(ctx, existingUser, bson.M{"_id": user.ID})
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("Usuario no encontrado")
		}
		return nil, errors.New("Error al buscar el usuario")
	}

	// Verificar si el nuevo username ya está en uso
	if user.Username != "" && user.Username != existingUser.Username {
		existingUserByUsername := &domain.User{}
		err := s.repo.FindOne(ctx, existingUserByUsername, bson.M{"username": user.Username})
		if err == nil {
			return nil, errors.New("Nombre de usuario ya está en uso")
		} else if err != mongo.ErrNoDocuments {
			return nil, errors.New("Error al verificar el nombre de usuario")
		}
	}

	// Verificar si el nuevo email ya está en uso
	if user.Email != "" && user.Email != existingUser.Email {
		existingUserByEmail := &domain.User{}
		err := s.repo.FindOne(ctx, existingUserByEmail, bson.M{"email": user.Email})
		if err == nil {
			return nil, errors.New("Correo electrónico ya está en uso")
		} else if err != mongo.ErrNoDocuments {
			return nil, errors.New("Error al verificar el correo electrónico")
		}
	}

	// Actualizar campos necesarios
	if user.Username != "" {
		existingUser.Username = user.Username
	}

	if user.Email != "" {
		existingUser.Email = user.Email
	}

	if user.WalletID != "" {
		existingUser.WalletID = user.WalletID
	}

	// Verificar y hashear la nueva contraseña si se proporciona
	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("Error al hashear la contraseña")
		}
		existingUser.Password = string(hashedPassword)
	}

	// Actualizar el usuario en la base de datos
	err = s.repo.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": existingUser})
	if err != nil {
		return nil, errors.New("Error al actualizar el usuario en la base de datos")
	}

	return existingUser, nil
}

func (s *UserService) Delete(ctx context.Context, userID string) error {
	// Verificar si el usuario existe
	existingUser := &domain.User{}
	err := s.repo.FindOne(ctx, existingUser, bson.M{"_id": userID})
	if err != nil {
		return errors.New("Usuario no encontrado")
	}

	// Eliminar el usuario en la base de datos
	err = s.repo.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		return errors.New("Error al eliminar el usuario en la base de datos")
	}

	return nil
}
