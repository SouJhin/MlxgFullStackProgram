package models

type themesList struct {
	ActivityId            int
	ActivityName          string
	ActivityCode          string
	CardShelvesCategoryId int
	ImageUrl              string
}

type categoryList struct {
	Id          int
	ThemesLists []themesList
	Name        string
	ActivityIds string
}

type GiftCard struct {
	BannerActivityId   int
	Content            string
	BannerActivityName string
	ShowType           uint
	Title              string
	CategoryLists      []categoryList
}
