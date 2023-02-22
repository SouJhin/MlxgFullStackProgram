package models

type entity struct {
	specId          string
	tradeMark       string
	id              string
	stock           float64
	specText        string
	num             int
	price           string
	membershipPrice int
}

type goodsList struct {
	sellTimeStatus  int
	id              int
	isSell          bool
	packCost        float64
	sales           int
	goodsType       int
	coverImg        string
	property        string
	goodsMealsInfo  string
	isAdd           int
	useSpec         bool
	stallCode       string
	sort            int
	price           float64
	unit            string
	imageArr        string
	membershipPrice int
	useProperty     int
	unitType        int
	minBuyNum       int
	specs           string
	content         string
	isFollowSuit    int
	stock           float64
	originalType    int `xorm:"type"`
	isLabel         int
	name            string
	image           string
}

type Goods struct {
	name            string
	sort            int
	icon            string
	id              int
	isShowBackStage int
	entity          []entity
	goodsList       []goodsList
}
