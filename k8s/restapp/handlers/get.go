package handlers

import (
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/models"
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
)

type get struct {
	db *gorm.DB
}

func NewGetHandler(db *gorm.DB) *get {
	return &get{db: db}
}

func (impl *get) Handle(params user.GetParams) middleware.Responder {
	usr := &models.User{}
	err := impl.db.Where("id=?", params.ID).Take(usr).Error
	if err != nil {
		return user.NewCreateUserInternalServerError()
	}
	return user.NewGetOK().WithPayload(usr)
}
