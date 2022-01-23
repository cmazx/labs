package handlers

import (
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
)

type add struct {
	db *gorm.DB
}

func NewAddHandler(db *gorm.DB) *add {
	return &add{db: db}
}

func (impl *add) Handle(params user.CreateUserParams) middleware.Responder {
	err := impl.db.Create(params.Body).Error
	if err != nil {
		return user.NewCreateUserInternalServerError()
	}
	return user.NewCreateUserOK().WithPayload(params.Body)
}
