package models

type Food struct {
	Image    string
	Name     string
	ID       string
	Price    int
	Describe string
}

func InsetFood(name string, price int) string {
	food := Food{
		Image:    "asdasd",
		Name:     name,
		ID:       "sssss",
		Price:    price,
		Describe: "youde",
	}
	_, err := Init().Insert(&food)
	if err != nil {
		return ""
	}

	return "success"
}
