package dto

type CopyrightComplianceHandlerRequestDTO struct {
	AibomId   int
	SourceUrl string
}

type WebPageLicenseDTO struct {
	LicenseContent string
	LicenseUrl     string
}

type WebPageCopyrightFlagDTO struct {
	Flag     string
	licensor string
}
