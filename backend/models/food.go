package models

import "fmt"

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
	engine := Init()
	_, err := engine.Insert(&food)
	findFood := make([]Food, 0)
	err = engine.Find(&findFood)
	for _, value := range findFood {
		fmt.Printf("food ====> %v\n", value)
	}
	fmt.Printf("summary %v\n\n", findFood)
	if err != nil {
		return ""
	}

	return "success"
}
