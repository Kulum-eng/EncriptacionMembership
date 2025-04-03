package controllers

import (
	"api/src/membership/application"
	"api/src/membership/domain"
	"api/src/user/infraestructure/http/responses"
	"strconv"

	"net/http"

	"github.com/gin-gonic/gin"
)

type MembershipController struct {
	createMembershipUseCase *application.CreateMembershipUseCase
	getMembershipUseCase   *application.GetMembershipUseCase
	updateaMembershipUseCase *application.UpdateMembershipUseCase
	deleteMembershipUseCase *application.DeleteMembershipUseCase
}

func NewMembershipController(createUC *application.CreateMembershipUseCase, getUC *application.GetMembershipUseCase, updateUC *application.UpdateMembershipUseCase, deleteMembershipUseCase *application.DeleteMembershipUseCase) *MembershipController {
	return &MembershipController{
		createMembershipUseCase: createUC,
		getMembershipUseCase:   getUC,
		updateaMembershipUseCase: updateUC,
		deleteMembershipUseCase: deleteMembershipUseCase,
	}
}

func (ctrl *MembershipController) CreateMembership(ctx *gin.Context) {
	var membership domain.Membership
	if err := ctx.ShouldBindJSON(&membership); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	idMembership, err := ctrl.createMembershipUseCase.Execute(membership)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear el miembro", err.Error()))
		return
	}

	membership.ID = idMembership

	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Miembro creado exitosamente", membership))
}

func (ctrl *MembershipController) GetAllMemberships(ctx *gin.Context) {
	memberships, err := ctrl.getMembershipUseCase.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener el miembro", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Miembros obtenidas exitosamente", memberships))
}

func (ctrl *MembershipController) GetMembershipByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	membership, err := ctrl.getMembershipUseCase.ExecuteByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener el miembro", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Miembro obtenido exitosamente", membership))
}

func (ctrl *MembershipController) UpdateMembership(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	var membership domain.Membership
	if err := ctx.ShouldBindJSON(&membership); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	membership.ID = id

	if err := ctrl.updateaMembershipUseCase.Execute(membership); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al actualizar el miembro", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Miembro actualizada exitosamente", membership))

}

func (ctrl *MembershipController) DeleteMembership(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	if err := ctrl.deleteMembershipUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al eliminar el miembro", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Miembro eliminado exitosamente", nil))
}