package service

import (
	"movie_premiuem/core/custom_errors"
	"movie_premiuem/user/repository"
)

type UserLicenseService interface {
	CreateUserLicense(licenseID int64, userID int64) (bool, error)
}

type userLicenseService struct {
	repository repository.UserLicenseRepository
}

func NewUserLicenseService(userLicenseRepository repository.UserLicenseRepository) UserLicenseService {
	if userLicenseRepository == nil {
		panic("userLicenseRepository cannot be nil")
	}
	return &userLicenseService{repository: userLicenseRepository}
}

func (s *userLicenseService) CreateUserLicense(licenseID int64, userID int64) (bool, error) {

	userHasActiveLicense, userHasActiveLicenseErr := s.repository.CheckIfUserHasActiveLicense(userID)
	if userHasActiveLicenseErr != nil {
		return false, userHasActiveLicenseErr
	}
	if userHasActiveLicense {
		return false, custom_errors.ErrUserHasActiveLicense
	}

	_, createUserLicenseErr := s.repository.CreateUserLicense(licenseID, userID)
	if createUserLicenseErr != nil {
		return false, createUserLicenseErr
	}

	return true, nil
}
