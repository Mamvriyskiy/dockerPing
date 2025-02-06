package services

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/Mamvriyskiy/dockerPing/internal/models"
	"github.com/Mamvriyskiy/dockerPing/internal/repository"
	"github.com/Mamvriyskiy/dockerPing/logger"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt       = "hfdjmaxckdk20"
	signingKey = "jaskljfkdfndnznmckmdkaf3124kfdlsf"
)

type ClientService struct {
	repo repository.ClientPostgresReposipory
}

func NewClientService(repo repository.ClientPostgresReposipory) *ClientService {
	return &ClientService{repo: repo}
}

func generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))

	return hex.EncodeToString(hash.Sum([]byte(salt)))
}

func (s *ClientService) AddClient(client models.ClientHandler) (models.ClientData, error) {
	clientServ := models.ClientService{
		Client:   client.Client,
		Password: generatePasswordHash(client.Password),
	}

	return s.repo.AddClient(clientServ)
}

type tokenClaims struct {
	jwt.StandardClaims
	ClientID string `json:"clientID"`
}

type markerClaims struct {
	jwt.StandardClaims
}

func generateMarker() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &markerClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(signingKey))
}

func (s *ClientService) GenerateToken(client models.ClientHandler) (models.ClientData, string, error) {
	// clientService := models.ClientService{
	// 	Client: client.Client,
	// }

	clientData, err := s.repo.GetClient(client.Email)
	if err != nil {
		logger.Log("Error", "GetClient", "Error get client:", err, client.Email)
		return models.ClientData{}, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		clientData.ClientID,
	})

	result, err := token.SignedString([]byte(signingKey))

	return clientData, result, err
}
