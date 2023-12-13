package service

import (
	"errors"
	"sadensmol/go/test_faraway/pkg/domain"
	"sadensmol/go/test_faraway/pkg/utils"
)

const (
	requestStringLength = 10
)

type IHasher interface {
	Mint(string) (string, error)
	Check(string) bool
}

type POWService struct {
	hasher IHasher
}

func NewPOWService() *POWService {
	return &POWService{hasher: domain.NewStdHashcash()}
}

func (h POWService) GenerateRequest() (*string, error) {
	return utils.GenRandomString(requestStringLength)
}

func (h POWService) Mint(req string) (*string, error) {
	res, err := h.hasher.Mint(req)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (h POWService) Check(hash string) error {
	ok := h.hasher.Check(hash)
	if !ok {
		return errors.New("hash is not valid")
	}
	return nil
}
