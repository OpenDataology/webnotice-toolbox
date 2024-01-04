package do

type SourceWebPageCopyrightKeyword struct {
	Id      int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	Keyword string `gorm:"type:varchar" json:"keyword"`
}

// TableName 设置表名
func (SourceWebPageCopyrightKeyword) TableName() string {
	return "t_source_web_page_copyright_keyword"
}
