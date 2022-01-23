package handlers

import (
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/models"
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
)

type del struct {
	db *gorm.DB
}

func NewDeleteHandler(db *gorm.DB) *del {
	return &del{db: db}
}

func (impl *del) Handle(params user.DeleteUserParams) middleware.Responder {
	err := impl.db.Where("id=?", params.ID).Delete(&models.User{}).Error
	if err != nil {
		return user.NewCreateUserInternalServerError()
	}
	return user.NewDeleteUserNoContent()
}
