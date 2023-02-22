package models

type benefitsItemSummary struct {
	UnitType     int
	BenefitsType int
	num          int
}

type benefitsSummary struct {
	BenefitsName      string
	BenefitsType      int
	benefitsSummaries []benefitsItemSummary
}

type LevelBenefit struct {
	Picture           string
	CardName          string
	BenefitsSummaries []benefitsSummary
	Level             uint
}
