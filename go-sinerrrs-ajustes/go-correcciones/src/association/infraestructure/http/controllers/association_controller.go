package controllers

import (
	"api/src/association/application"
	"api/src/association/domain"
	"api/src/association/infraestructure/http/responses"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AssociationController struct {
	createAssociationUseCase *application.CreateAssociationUseCase
	getAssociationUseCase    *application.GetAssociationUseCase
	updateAssociationUseCase *application.UpdateAssociationUseCase
	deleteAssociationUseCase *application.DeleteAssociationUseCase
}

func NewAssociationController(createUC *application.CreateAssociationUseCase, getUC *application.GetAssociationUseCase, updateUC *application.UpdateAssociationUseCase, deleteUseCase *application.DeleteAssociationUseCase) *AssociationController {
	return &AssociationController{
		createAssociationUseCase: createUC,
		getAssociationUseCase:    getUC,
		updateAssociationUseCase: updateUC,
		deleteAssociationUseCase: deleteUseCase,
	}
}

func (ctrl *AssociationController) Create(ctx *gin.Context) {
	var association domain.Association
	if err := ctx.ShouldBindJSON(&association); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	idAssociation, err := ctrl.createAssociationUseCase.Execute(association)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear asociación", err.Error()))
		return
	}

	association.ID = idAssociation

	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Asociación creada exitosamente", association))
}

func (ctrl *AssociationController) GetAll(ctx *gin.Context) {
	associations, err := ctrl.getAssociationUseCase.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener la asociación", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Asociaciones obtenidas exitosamente", associations))
}

func (ctrl *AssociationController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	association, err := ctrl.getAssociationUseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener la asociación", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Asociación obtenida exitosamente", association))
}

func (ctrl *AssociationController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}
	
	var association domain.Association
	if err := ctx.ShouldBindJSON(&association); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	association.ID = id

	if err := ctrl.updateAssociationUseCase.Execute(association); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al actualizar la asociación", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Asociación actualizada exitosamente", association))
}

func (ctrl *AssociationController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}
	 
	if err := ctrl.deleteAssociationUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al eliminar la asociación", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Asociación eliminada exitosamente", nil))
}