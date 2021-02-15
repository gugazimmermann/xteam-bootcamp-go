package calc_testify

import (
	"errors"

	"github.com/gugazimmermann/xteam-bootcamp-go/database"
)

type DiscountCalculator struct {
	mininumPurchaseAmount int
	discountRepository    database.Repository
}

func NewDiscountCalculator(mininumPurchaseAmount int, discountRepository database.Repository) (*DiscountCalculator, error) {
	if mininumPurchaseAmount == 0 {
		return &DiscountCalculator{}, errors.New("minimum purchase amount could not be zero")
	}
	return &DiscountCalculator{
		mininumPurchaseAmount: mininumPurchaseAmount,
		discountRepository:    discountRepository,
	}, nil
}

func (c *DiscountCalculator) Calculate(purchaseAmount int) int {
	if purchaseAmount > c.mininumPurchaseAmount {
		discount := c.discountRepository.FindCurrentDiscount()
		return (purchaseAmount * (100 - discount)) / 100
	}
	return purchaseAmount
}
