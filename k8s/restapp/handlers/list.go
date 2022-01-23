package handlers

import (
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/models"
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
)

type list struct {
	db *gorm.DB
}

func NewListHandler(db *gorm.DB) *list {
	return &list{db: db}
}

func (impl *list) Handle(params user.ListParams) middleware.Responder {
	var list []*models.User
	err := impl.db.Model(&models.User{}).Find(&list).Error
	if err != nil {
		panic(err)
	}
	return user.NewListOK().WithPayload(list)
}
