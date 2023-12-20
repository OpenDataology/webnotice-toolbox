package daos

import (
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/config/database"
	"github.com/OpenDataology/webnotice-toolbox/webnotice-scanner/models/do"
)

func SourceWebPageCopyrightKeywordFindAll() (sourceWebPageCopyrightKeywordList []do.SourceWebPageCopyrightKeyword, err error) {

	err = database.DB.Model(&do.SourceWebPageCopyrightKeyword{}).Find(&sourceWebPageCopyrightKeywordList).Error
	if err != nil {
		return nil, err
	}
	return
}
