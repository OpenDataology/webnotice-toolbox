package do

type LicenseMes struct {
	Id             int    `gorm:"type:int;primary_key;autoIncrement" json:"id"`
	LicenseName    string `gorm:"type:varchar" json:"license_name"`
	LicenseUrl     string `gorm:"type:varchar" json:"license_url"`
	LicenseType    int    `gorm:"type:int" json:"license_type"`
	CopyrightFlag  string `gorm:"type:varchar" json:"copyright_flag"`
	Licensor       string `gorm:"type:varchar" json:"licensor"`
	LicenseContent string `gorm:"type:LONGTEXT" json:"license_content"`
	SourceUrl      string `gorm:"type:varchar" json:"source_url"`
	AibomId        int    `gorm:"type:int" json:"aibom_id"`
}

// TableName 设置表名
func (LicenseMes) TableName() string {
	return "t_license_mes"
}
