package main

import (
	"gorm.io/gorm"
)

type DatamartRepository struct {
	db *gorm.DB
}

func NewDatamartRepository(client *gorm.DB) DatamartRepository {
	return DatamartRepository{client}
}

func (r DatamartRepository) GetDatamart1() ([]Datamart1, error) {

	var res []Datamart1

	r.db.Raw(
		`
		select
			transaction_id,
			buyer_id,
			c.customer_name buyer_name,
			seller_id,
			c2.customer_name seller_name,
			t.product_id,
			p.product_name,
			null as currency,
			price,
			volume,
			value,
			transaction_date,
			entry_date,
			date_part('month', entry_date) entry_month,
			date_part('year', entry_date) entry_year,
			buy_sell,
			confirm_status,
			complete_status_buyer,
			complete_status_seller
		from transactions t
		left join customers c on t.buyer_id = c.customer_id
		left join customers c2 on t.seller_id = c2.customer_id
		left join products p on t.product_id = p.product_id
		`,
	).Scan(&res)

	return res, nil
}

func (r DatamartRepository) InserDatamart1(data []Datamart1) error {

	if err := r.db.Create(&data).Error; err != nil {
		return err
	}

	return nil

}
