package bookstore

import (
	"math"
)

type Basket []int

type Group struct {
	subset []int
	basket Basket
}

func Cost(basket []int) float64 {
	return getLowestPrice(Group{[]int{}, basket})
}

func getLowestPrice(group Group) float64 {
	lowestPrice := getSubsetPrice(group.subset)
	if len(group.basket) != 0 {
		subGroups := getAvailableGroups(group.basket)
		lowestSubPrice := math.MaxFloat64
		for _, subGroup := range subGroups {
			price := getLowestPrice(subGroup)
			if price < lowestSubPrice {
				lowestSubPrice = price
			}
		}
		return lowestSubPrice + lowestPrice
	}
	return lowestPrice
}

func getSubsetPrice(subset []int) float64 {
	price := float64(8)
	discount := float64(0)
	switch len(subset) {
	case 0:
		return 0
	case 1:
		discount = 0
	case 2:
		discount = 0.05
	case 3:
		discount = 0.1
	case 4:
		discount = 0.2
	case 5:
		discount = 0.25
	}
	return price * (1 - discount) * float64(len(subset))
}

func getAvailableGroups(basket Basket) []Group {
	result := make([]Group, 0)
	subset := make([]int, 0)
	remainder := basket
	for v := 1; v <= 5; v++ {
		if ind := remainder.find(v); ind != -1 {
			subset = append(subset, v)
			remainder = remainder.remove(ind)
			result = append(result, Group{subset: subset, basket: NewBasket(remainder)})
		}
	}
	return result
}

func (basket Basket) find(val int) (position int) {
	for i, v := range basket {
		if v == val {
			return i
		}
	}
	return -1
}

func (basket Basket) remove(position int) Basket {
	if position >= len(basket) {
		return basket[:len(basket)-1]
	}
	return append(basket[:position], basket[position+1:]...)
}

func NewBasket(basket Basket) Basket {
	newBasket := make([]int, len(basket))
	copy(newBasket, basket)
	return newBasket
}
