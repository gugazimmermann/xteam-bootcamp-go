package calc_testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DiscountRepositoryMock struct {
	mock.Mock
}

func (drm *DiscountRepositoryMock) FindCurrentDiscount() int {
	args := drm.Called()
	return args.Int(0)
}

func Test_CalculatorTestify(t *testing.T) {
	testCases := []struct {
		name                  string
		mininumPurchaseAmount int
		discountValue         int
		purchaseAmount        int
		output                int
	}{
		{
			name:                  "Purchase with discount",
			mininumPurchaseAmount: 100,
			discountValue:         20,
			purchaseAmount:        150,
			output:                120,
		},
		{
			name:                  "Purchase without discount",
			mininumPurchaseAmount: 50,
			discountValue:         10,
			purchaseAmount:        350,
			output:                315,
		},
		{
			name:                  "Should not apply if discount value equals zero",
			mininumPurchaseAmount: 250,
			discountValue:         0,
			purchaseAmount:        500,
			output:                500,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			discountRepositoryMock := &DiscountRepositoryMock{}
			discountRepositoryMock.On("FindCurrentDiscount").Return(tc.discountValue)
			calculator, err := NewDiscountCalculator(tc.mininumPurchaseAmount, discountRepositoryMock)
			if err != nil {
				t.Fatalf("cloud not instantiate the calculator: %v", err.Error())
			}
			result := calculator.Calculate(tc.purchaseAmount)
			assert.Equal(t, tc.output, result)
		})
	}
}

func Test_CalculatorShouldFail(t *testing.T) {
	discountRepositoryMock := &DiscountRepositoryMock{}
	discountRepositoryMock.On("FindCurrentDiscount").Return(10)
	_, err := NewDiscountCalculator(0, discountRepositoryMock)
	if err == nil {
		t.Fatalf("Should not instatiate the calculator, zero purchase amount")
	}
}
