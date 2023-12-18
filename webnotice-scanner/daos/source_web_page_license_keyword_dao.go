package daos

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/config/database"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/do"
)

func LicenseKeywordFindAll() (sourceWebPageLicenseKeywordList []do.SourceWebPageLicenseKeyword, err error) {

	err = database.DB.Model(&do.SourceWebPageLicenseKeyword{}).Find(&sourceWebPageLicenseKeywordList).Error
	if err != nil {
		return nil, err
	}
	return
}
