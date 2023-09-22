package repository

import (
	"AnaOrderService/model"
	"AnaOrderService/request"
	"database/sql"
	"errors"
	"log"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (u *OrderRepository) CheckoutItems(param request.OrderRequest) ([]model.Product, error) {
	var products []model.Product
	var productStockDeducted []model.ProductStockDeducted
	tx, err := u.db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	queryValidateItem := "WITH _param AS (SELECT\n                    $1::BIGINT AS product_id,\n                    $2::BIGINT AS deduct_stock)\n\n   , _stock AS (SELECT\n                    rank() OVER (ORDER BY stock_quantity) AS rank,\n                    (SELECT deduct_stock FROM _param) AS deducted_remaining,\n                    stock.*\n                FROM product_stock stock\n                JOIN warehouse\n                     ON warehouse.id = stock.warehouse_id\n                         AND warehouse.is_active IS TRUE\n                WHERE\n                    product_id = (SELECT product_id FROM _param) AND\n                    stock_quantity > 0)\n\n   , _data_stock AS (SELECT\n                         id AS product_stock_id,\n                         rank,\n                         max(rank) OVER (),\n                         product_id,\n                         warehouse_id,\n                         CASE\n                             WHEN rank > 1 THEN\n                                 stock_quantity\n                             ELSE\n                                 CASE\n                                     WHEN stock_quantity -\n                                          coalesce(lag(deducted_remaining) OVER (ORDER BY rank), deducted_remaining) >=\n                                          0\n                                         THEN\n                                             stock_quantity -\n                                             coalesce(lag(deducted_remaining) OVER (ORDER BY rank), deducted_remaining)\n                                     ELSE 0 END END AS stock_quantity,\n                         CASE\n                             WHEN stock_quantity -\n                                  coalesce(lag(deducted_remaining) OVER (ORDER BY rank), deducted_remaining) >= 0\n                                 THEN\n                                 CASE\n                                     WHEN max(rank) OVER () = rank THEN\n                                         0\n                                     ELSE coalesce(lag(deducted_remaining) OVER (ORDER BY rank), deducted_remaining) END\n                             ELSE\n                                 abs(stock_quantity -\n                                     coalesce(lag(deducted_remaining) OVER (ORDER BY rank),\n                                              deducted_remaining))\n                             END AS deducted_remaining\n                     FROM _stock)\n\nSELECT\n    product_stock_id,\n    product_id,\n    warehouse_id,\n    CASE\n        WHEN stock_quantity -\n             coalesce(lag(deducted_remaining) OVER (ORDER BY rank), deducted_remaining) >=\n             0\n            THEN\n                stock_quantity -\n                coalesce(lag(deducted_remaining) OVER (ORDER BY rank), deducted_remaining)\n        ELSE 0 END AS stock_quantity,\n    deducted_remaining,\n    CASE\n        WHEN min(deducted_remaining) OVER () = 0 AND (SELECT deduct_stock FROM _param) > 0 THEN\n            TRUE\n        ELSE FALSE END is_valid\nFROM _data_stock"
	rows, err := u.db.Query(queryValidateItem, param.ProductID, param.Quantity)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	for rows.Next() {
		var deducted model.ProductStockDeducted
		if err := rows.Scan(&deducted.ProductStockId, &deducted.ProductId, &deducted.WarehouseId, &deducted.StockQuantity, &deducted.DeductedRemaining, &deducted.IsValid); err != nil {
			tx.Rollback()
			log.Fatal(err)
			return nil, err
		}
		productStockDeducted = append(productStockDeducted, deducted)
	}

	if len(productStockDeducted) == 0 {
		tx.Rollback()
		log.Fatal(err)
		return nil, errors.New("Stock not found")
	}

	if !productStockDeducted[0].IsValid {
		tx.Rollback()
		log.Fatal(err)
		return nil, errors.New("Stock not enough")
	}

	for _, stock := range productStockDeducted {
		queryUpdateStock := "UPDATE product_stock SET stock_quantity = $2 WHERE id = $1"
		_, err = tx.Exec(queryUpdateStock, stock.ProductStockId, stock.StockQuantity)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return nil, err
		}
	}

	queryInsertCart := "INSERT INTO cart (users_id, product_id, quantity) VALUES ($1, $2, $3)"
	_, err = tx.Exec(queryInsertCart, param.UserId, param.ProductID, param.Quantity)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return nil, err
	}

	tx.Commit()

	return products, nil
}

func (u *OrderRepository) ReleaseUnusedStock() error {
	var productStockCarts []model.ProductStockCart
	tx, err := u.db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	queryUnusedCart := "SELECT\n    id,\n    users_id,\n    product_id,\n    quantity\nFROM cart\nWHERE\n    is_paid IS FALSE AND\n    released_stock IS FALSE AND\n    updated_at + hold_duration < now()"
	rows, err := u.db.Query(queryUnusedCart)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for rows.Next() {
		var deducted model.ProductStockCart
		if err := rows.Scan(&deducted.ID, &deducted.UsersId, &deducted.ProductId, &deducted.Quantity); err != nil {
			tx.Rollback()
			log.Fatal(err)
			return err
		}
		productStockCarts = append(productStockCarts, deducted)
	}

	for _, cart := range productStockCarts {
		queryUpdateStock := "UPDATE product_stock SET stock_quantity = $2 WHERE id = $1"
		_, err = tx.Exec(queryUpdateStock, cart.ProductId, cart.Quantity)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return err
		}

		queryUpdateCart := "UPDATE cart SET released_stock = TRUE WHERE id = $1"
		_, err = tx.Exec(queryUpdateCart, cart.ID)
		if err != nil {
			tx.Rollback()
			log.Fatal(err)
			return err
		}
	}

	tx.Commit()

	return nil
}
