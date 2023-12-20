package daos

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/config/database"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/do"
	"log"
	"strings"
)

func LicenseMesSave(saveLicenseMes do.LicenseMes) (err error) {

	licenseMes, err := findLicenseMesBySourceUrl(saveLicenseMes.SourceUrl)

	if err != nil {
		return err
	}

	if licenseMes != nil &&
		licenseMes.SourceUrl != "" {
		// update
		if strings.Compare(saveLicenseMes.SourceUrl, "") != 0 {
			licenseMes.SourceUrl = saveLicenseMes.SourceUrl
		}

		if strings.Compare(saveLicenseMes.LicenseUrl, "") != 0 {
			licenseMes.LicenseUrl = saveLicenseMes.LicenseUrl
		}

		if strings.Compare(saveLicenseMes.LicenseContent, "") != 0 {
			licenseMes.LicenseContent = saveLicenseMes.LicenseContent
		}

		if strings.Compare(saveLicenseMes.CopyrightFlag, "") != 0 {
			licenseMes.CopyrightFlag = saveLicenseMes.CopyrightFlag
		}

		if strings.Compare(saveLicenseMes.LicenseName, "") != 0 {
			licenseMes.LicenseName = saveLicenseMes.LicenseName
		}

		if saveLicenseMes.LicenseType != 0 {
			licenseMes.LicenseType = saveLicenseMes.LicenseType
		}

		if strings.Compare(saveLicenseMes.Licensor, "") != 0 {
			licenseMes.Licensor = saveLicenseMes.Licensor
		}

		if saveLicenseMes.AibomId != 0 {
			licenseMes.AibomId = saveLicenseMes.AibomId

		}

		err = database.DB.Save(&licenseMes).Error

		if err != nil {
			return err
		}

		return
	}

	// add new data line
	err = database.DB.
		Create(&saveLicenseMes).
		Error

	if err != nil {
		return err
	}

	return
}

func findLicenseMesBySourceUrl(sourceUrl string) (licenseMes *do.LicenseMes, err error) {

	//var _licenseMes do.LicenseMes

	dbErr := database.
		DB.Model(&do.LicenseMes{}).
		Where("source_url = ?", sourceUrl).
		First(&licenseMes).Error

	if dbErr != nil &&
		strings.Compare(dbErr.Error(), "record not found") != 0 {
		log.Printf("cur error:%s", dbErr.Error())

		return nil, dbErr
	}

	return
}
