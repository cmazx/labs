package handlers

import (
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/models"
	"github.com/cmazx/mazx-labs/master/k8s/restapp/gen/restapi/operations/user"
	"github.com/go-openapi/runtime/middleware"
	"gorm.io/gorm"
)

type put struct {
	db *gorm.DB
}

func NewPutHandler(db *gorm.DB) *put {
	return &put{db: db}
}

func (impl *put) Handle(params user.UpdateUserParams) middleware.Responder {
	usr := &models.User{}
	err := impl.db.Where("id=?", params.ID).Take(usr).Error
	if err != nil {
		return user.NewUpdateUserDefault(404)
	}
	usr.Email = params.Body.Email
	usr.Username = params.Body.Username
	usr.FirstName = params.Body.FirstName
	usr.LastName = params.Body.LastName
	usr.Phone = params.Body.Phone
	err = impl.db.Updates(&usr).Error
	if err != nil {
		return user.NewUpdateUserDefault(500)
	}
	return user.NewUpdateUserOK().WithPayload(usr)
}
