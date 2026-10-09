package product

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type MySQLRepository struct {
	DB *sql.DB
}

func (r *MySQLRepository) GetAll() ([]Product, error) {
	query := `
				SELECT id, name, price FROM products 
			 `

	rows, err := r.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []Product

	for rows.Next() {
		var product Product
		var productID string

		err := rows.Scan(
			&productID,
			&product.Name,
			&product.Price,
		)

		if err != nil {
			return nil, err
		}

		product.ID = uuid.MustParse(productID)

		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *MySQLRepository) FindByID(id uuid.UUID) (*Product, error) {
	query := `
				SELECT id, name, price 
				FROM products 
				WHERE id = ?
			 `

	row := r.DB.QueryRow(
		query,
		id.String(),
	)

	var product Product
	var productID string

	err := row.Scan(
		&productID,
		&product.Name,
		&product.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf(
				"Produk dengan ID %s tidak ditemukan",
				id,
			)
		}

		return nil, err
	}

	product.ID = uuid.MustParse(productID)

	return &product, nil
}

func (r *MySQLRepository) Create(product Product) (Product, error) {
	query := `
				INSERT INTO products (id, name, price) 
				VALUES (?, ?, ?)
			`

	_, err := r.DB.Exec(
		query,
		product.ID.String(),
		product.Name,
		product.Price,
	)

	if err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *MySQLRepository) Update(product Product) error {
	query := `
			 UPDATE products SET name = ?, price = ? WHERE id = ?
			 `

	result, err := r.DB.Exec(
		query,
		product.Name,
		product.Price,
		product.ID.String(),
	)

	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowAffected == 0 {
		return fmt.Errorf(
			"produk dengan ID %s tidak ditemukan",
			product.ID,
		)
	}

	return nil

}

func (r *MySQLRepository) Delete(id uuid.UUID) error {
	query := `
				DELETE FROM products WHERE id = ?
			 `

	result, err := r.DB.Exec(
		query,
		id.String(),
	)

	if err != nil {
		return err
	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return nil
	}

	if rowAffected == 0 {
		return fmt.Errorf(
			"produk dengan ID %s tidak ditemukan",
			id,
		)
	}

	return nil

}
