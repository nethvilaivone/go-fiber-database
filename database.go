package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func createProducts(product *Product) (error) {
	fmt.Println(product)
	_ , err := db.Exec("INSERT INTO products(name, price) VALUES ($1, $2);", product.Name, product.Price)

	if err != nil {
		log.Fatal(err)
	}

	return err
}


func getProduct(id int)(Product, error) {

	var p = Product{}
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id=$1", id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return p, err
	}

	return p, nil
}

func getProducts() ([]Product, error) {
	var products []Product
	var p Product

	rows, err := db.Query("SELECT id , name , price FROM products ORDER BY id;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		
		if err != nil {
			return nil,err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, err
}

func UpdateProducts(id int, product *Product) (Product, error) {

	var p = Product{}

	row := db.QueryRow("UPDATE products SET name=$1, price=$2  WHERE id=$3 RETURNING id , name , price", product.Name, product.Price, id)

	err := row.Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return Product{}, err
	}

	return p, err
}

func deleteProduct(id int) (error) {
	_, err := db.Exec("DELETE FROM  products WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
