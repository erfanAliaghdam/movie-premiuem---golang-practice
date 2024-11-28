package license_handlers

import (
	"context"
	"log"
	"movie_premiuem/app"
	"movie_premiuem/entity"
	"movie_premiuem/entity/repositories"
	"movie_premiuem/utils"
	"net/http"
	"time"
)

func LicenseListHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	db := app.AppInstance.GetDB()
	redisClient := app.AppInstance.GetRedis()
	var licenseList []entity.License
	cacheKey := "license_list"
	cacheService := utils.NewCacheService(redisClient)

	// find data in cache
	cacheFound, cachedData, _ := cacheService.GetCache(ctx, cacheKey)

	if cacheFound {
		response := utils.Response{
			Status:  "success",
			Message: "License list fetched successfully",
			Code:    "",
			Data:    cachedData,
		}
		utils.WriteJSONResponse(w, response, http.StatusOK)
		return
	}

	// get data from db
	licenseRepository := repositories.NewLicenseRepository(db)

	licenseList, licenseListErr := licenseRepository.GetAllLicenses()
	if licenseListErr != nil {
		log.Println("error occurred in license list handler :", licenseListErr)
		utils.InternalServerError500(w)
		return
	}

	// store in cache
	setCacheErr := cacheService.SetCache(ctx, cacheKey, licenseList, time.Hour*2)
	if setCacheErr != nil {
		log.Println("error occurred in license list handler :", setCacheErr)
	}

	response := utils.Response{
		Status:  "success",
		Message: "License list fetched successfully",
		Code:    "",
		Data:    licenseList,
	}
	utils.WriteJSONResponse(w, response, http.StatusOK)
}
