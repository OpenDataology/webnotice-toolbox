package daos

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/config/database"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/do"
)

func LicenseUrlSuffixFindAll() (licenseUrlSuffixList []do.LicenseUrlSuffix, err error) {

	err = database.DB.Model(&do.LicenseUrlSuffix{}).Find(&licenseUrlSuffixList).Error
	if err != nil {
		return nil, err
	}
	return
}
