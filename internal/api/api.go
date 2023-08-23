package api

import (
	"github.com/Shemistan/Lesson_5/internal/converters"
	"github.com/Shemistan/Lesson_5/internal/models"
	"github.com/Shemistan/Lesson_5/internal/service"
)

type IApi interface {
	Add(req *models.AddRequest) (int, error)
}

func New(serv service.IService) IApi {
	return &api{serv: serv}
}

type api struct {
	serv service.IService
}

func (a *api) Add(req *models.AddRequest) (int, error) {
	res, err := a.serv.Add(converters.ApiModelToServiceModel(*req))
	if err != nil {
		return 0, err
	}

	return res, nil
}
