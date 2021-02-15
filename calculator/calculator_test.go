package calculator

import (
	"testing"
)

type DiscountRepositoryMock struct{}

func (dc *DiscountRepositoryMock) FindCurrentDiscount() int {
	return 20
}

func Test_Calculator(t *testing.T) {
	testCases := []struct {
		name                  string
		mininumPurchaseAmount int
		purchaseAmount        int
		output                int
	}{
		{
			name:                  "Purchase with discount",
			mininumPurchaseAmount: 100,
			purchaseAmount:        150,
			output:                120,
		},
		{
			name:                  "Purchase without discount",
			mininumPurchaseAmount: 50,
			purchaseAmount:        350,
			output:                280,
		},
	}
	discountRepositoryMock := &DiscountRepositoryMock{}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			calculator, err := NewDiscountCalculator(tc.mininumPurchaseAmount, discountRepositoryMock)
			if err != nil {
				t.Fatalf("cloud not instantiate the calculator: %v", err.Error())
			}
			result := calculator.Calculate(tc.purchaseAmount)
			if result != tc.output {
				t.Errorf("expected %d but got %d", tc.output, result)
			}
		})
	}
}

func Test_CalculatorShouldFail(t *testing.T) {
	discountRepositoryMock := &DiscountRepositoryMock{}
	_, err := NewDiscountCalculator(0, discountRepositoryMock)
	if err == nil {
		t.Fatalf("Should not instatiate the calculator, zero purchase amount")
	}
}
