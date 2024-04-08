package service

import (
	"github.com/mmfshirokan/positionService/internal/model"
	"github.com/mmfshirokan/positionService/internal/repository"
)

type Prices struct {
	repo *repository.Prices
}

type PrcGeter interface {
	GetAllChanForSymb(symb string) (chanels []chan model.Price, isSuccessfull bool)
}

type PrcManipulator interface {
	Add(key model.SymbOperDTO, ch chan model.Price)
	Delete(key model.SymbOperDTO) (wasDeleted bool)
}

func NewSymbOperMap(repo *repository.Prices) *Prices {
	return &Prices{
		repo: repo,
	}
}

func (s *Prices) GetAllChanForSymb(symb string) ([]chan model.Price, bool) {
	return s.repo.GetAllChanForSymb(symb)
}

func (s *Prices) Add(key model.SymbOperDTO, ch chan model.Price) {
	s.repo.Add(key, ch)
}
func (s *Prices) Get(key model.SymbOperDTO) chan model.Price {
	return s.Get(key)
}

func (s *Prices) Delete(key model.SymbOperDTO) (wasDeleted bool) {
	return s.repo.Delete(key)
}
