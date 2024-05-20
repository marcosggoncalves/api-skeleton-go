package controller

import (
	"ApiSup/internal/config/auth"
	"ApiSup/internal/models"
	"ApiSup/internal/services"
	"ApiSup/pkg/mapear/request"
	"ApiSup/pkg/mapear/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

func (controller *UserController) Login(c echo.Context) error {
	var body request.Login
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: "Falha Body!", Description: err.Error()})
	}

	var usuario *models.Usuario
	usuario, err := controller.UserService.Authenticate(body.CPF, body.Senha)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.Error{Message: "Acesso não Autorizado!", Description: err.Error()})
	}

	tokenString, err := auth.GenerateJWT(*usuario)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: "Falha Autenticação", Description: err.Error()})
	}

	response := response.Token{
		Token:   tokenString,
		Usuario: *usuario,
	}

	return c.JSON(http.StatusOK, response)
}

func (controller *UserController) GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: "ID não foi informado!", Description: err.Error()})
	}

	user, err := controller.UserService.Detalhar(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: "Usuário não encontrado!", Description: err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) GetUsers(c echo.Context) error {
	users, err := controller.UserService.Usuarios(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: "Não foi possivel listar usuários!", Description: err.Error()})
	}

	return c.JSON(http.StatusOK, users)
}

func (controller *UserController) CreateUser(c echo.Context) error {
	var user models.Usuario
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: "Falha Body!", Description: err.Error()})
	}

	if err := controller.UserService.Novo(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: "Não foi possivel cadastrar novo usuário!", Description: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Success{Message: "Usuário cadastrado com sucesso!"})
}

func (controller *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: "ID não foi informado!", Description: err.Error()})
	}

	var updatedUser models.Usuario
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: "Falha Body!", Description: err.Error()})
	}

	user, err := controller.UserService.Editar(id, &updatedUser)
	if err != nil {
		return c.JSON(http.StatusNotFound, response.Error{Message: "Usuário não encontrado!", Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.SuccessBody{Message: "Cadastrado alterado sucesso!", Body: user})
}

func (controller *UserController) DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Error{Message: "ID não foi informado!", Description: err.Error()})
	}

	if err := controller.UserService.Deletar(id); err != nil {
		return c.JSON(http.StatusInternalServerError, response.Error{Message: "Não foi possivel deletar usuário!", Description: err.Error()})
	}

	return c.JSON(http.StatusOK, response.Success{Message: "Cadastrado deletado com sucesso!"})
}
