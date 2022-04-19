package application

import (
	"go-DDD/domain/entity"
	"go-DDD/domain/repository"
)

type LastestNewApp struct {
	rl repository.LatestNewsRepo
}

var _ LastestNewAppInterFace = &LastestNewApp{}

type LastestNewAppInterFace interface {
	GetNewDetail(newID string) *entity.LatestNews
	CreateNew(new *entity.LatestNews) error
}

func (r *LastestNewApp) GetNewDetail(newID string) *entity.LatestNews {
	return r.rl.GetNewDetail(newID)
}

func (r *LastestNewApp) CreateNew(new *entity.LatestNews) error {
	err := r.rl.CreateNew(new)
	return err
}
