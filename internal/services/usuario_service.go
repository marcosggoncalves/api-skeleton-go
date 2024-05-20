package services

import (
	"ApiSup/internal/models"
	"ApiSup/internal/repository"
	"ApiSup/pkg/pagination"
	"errors"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Authenticate(cpf, senha string) (*models.Usuario, error)
	Detalhar(id int) (*models.Usuario, error)
	Usuarios(c echo.Context) (*pagination.Pagination, error)
	Novo(user *models.Usuario) error
	Editar(id int, updatedUser *models.Usuario) (*models.Usuario, error)
	Deletar(id int) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Authenticate(cpf, senha string) (*models.Usuario, error) {
	user, err := s.userRepo.GetUserByCPF(cpf)
	if err != nil {
		return nil, errors.New("usuário não foi localizado, inválido")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(senha))
	if err != nil {
		return nil, errors.New("senha inválida")
	}

	return user, nil
}
func (s *userService) Detalhar(id int) (*models.Usuario, error) {
	return s.userRepo.Detalhar(id)
}

func (s *userService) Usuarios(c echo.Context) (*pagination.Pagination, error) {
	return s.userRepo.Usuarios(c)
}

func (s *userService) Novo(user *models.Usuario) error {
	return s.userRepo.Novo(user)
}

func (s *userService) Editar(id int, updatedUser *models.Usuario) (*models.Usuario, error) {
	return s.userRepo.Editar(id, updatedUser)
}

func (s *userService) Deletar(id int) error {
	return s.userRepo.Deletar(id)
}
