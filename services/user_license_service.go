package services

import (
	"movie_premiuem/custom_errors"
	"movie_premiuem/entity/repositories"
)

type UserLicenseService interface {
	CreateUserLicense(licenseID int64, userID int64) (bool, error)
}

type userLicenseService struct {
	repository repositories.UserLicenseRepository
}

func NewUserLicenseService(userLicenseRepository repositories.UserLicenseRepository) UserLicenseService {
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
