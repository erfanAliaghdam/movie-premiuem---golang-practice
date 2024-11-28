package license_handlers

import (
	"context"
	"log"
	"movie_premiuem/core"
	"movie_premiuem/core/entity"
	"movie_premiuem/core/repository"
	utils2 "movie_premiuem/core/utils"
	"net/http"
	"time"
)

func LicenseListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db := core.AppInstance.GetDB()
	redisClient := core.AppInstance.GetRedis()
	var licenseList []entity.License
	cacheKey := "license_list"
	cacheService := utils2.NewCacheService(redisClient)

	// find data in cache
	cacheFound, cachedData, _ := cacheService.GetCache(ctx, cacheKey)

	if cacheFound {
		response := utils2.Response{
			Status:  "success",
			Message: "License list fetched successfully",
			Code:    "",
			Data:    cachedData,
		}
		utils2.WriteJSONResponse(w, response, http.StatusOK)
		return
	}

	// get data from db
	licenseRepository := repository.NewLicenseRepository(db)

	licenseList, licenseListErr := licenseRepository.GetAllLicenses()
	if licenseListErr != nil {
		log.Println("error occurred in license list handler :", licenseListErr)
		utils2.InternalServerError500(w)
		return
	}

	// store in cache
	setCacheErr := cacheService.SetCache(ctx, cacheKey, licenseList, time.Hour*2)
	if setCacheErr != nil {
		log.Println("error occurred in license list handler :", setCacheErr)
	}

	response := utils2.Response{
		Status:  "success",
		Message: "License list fetched successfully",
		Code:    "",
		Data:    licenseList,
	}
	utils2.WriteJSONResponse(w, response, http.StatusOK)
}
