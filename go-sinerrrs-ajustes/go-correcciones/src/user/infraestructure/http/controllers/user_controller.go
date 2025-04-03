package controllers

import (
	"api/src/user/application"
	"api/src/user/domain"
	"api/src/user/infraestructure/http/responses"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUserUseCase *application.CreateUserUseCase
	getUserUseCase *application.GetUserUseCase
	updateUserUserCase *application.UpdateUserUseCase
	deleteUserUseCase *application.DeleteUserUseCase
}

func NewUserController(createUC *application.CreateUserUseCase, getUC *application.GetUserUseCase, updateUC *application.UpdateUserUseCase, deleteUC *application.DeleteUserUseCase) *UserController {
	return &UserController{
		createUserUseCase: createUC,
		getUserUseCase: getUC,
		updateUserUserCase: updateUC,
		deleteUserUseCase: deleteUC,
	}
}

func (ctrl *UserController) Create(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	idUser, err := ctrl.createUserUseCase.Execute(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear usuario", err.Error()))
		return
	}

	user.ID = idUser

	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Usuario creado exitosamente", user))
}

func (ctrl *UserController) GetAll(ctx *gin.Context) {
	users, err := ctrl.getUserUseCase.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener usuarios", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Usuarios obtenidos exitosamente", users))
}

func (ctrl *UserController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	user, err := ctrl.getUserUseCase.ExecuteByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener usuario", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Usuario obtenido exitosamente", user))
}

func (ctrl *UserController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	_, err = ctrl.getUserUseCase.ExecuteByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener usuario", err.Error()))
		return
	}

	user.ID = id
	if err := ctrl.updateUserUserCase.Execute(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al actualizar usuario", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Usuario actualizado exitosamente", user))
}

func (ctrl *UserController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	if err := ctrl.deleteUserUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al eliminar usuario", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Usuario eliminado exitosamente", nil))
}