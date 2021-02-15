package database

import "database/sql"

type Repository interface {
	FindCurrentDiscount() int
}

type DiscountRepository struct {
	db *sql.DB
}

func NewDiscountRepository(db *sql.DB) *DiscountRepository {
	return &DiscountRepository{
		db: db,
	}
}

func (discountRepository *DiscountRepository) FindCurrentDiscount() (discount int) {
	sql := "SELECT value FROM discount ORDER BY data DESC LIMIT 1"
	row := discountRepository.db.QueryRow(sql)
	row.Scan(&discount)
	return discount
}
