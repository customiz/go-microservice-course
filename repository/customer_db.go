package repository

import "github.com/jmoiron/sqlx"

//เป็น struck เพราะเป็นเต้าเสียบ
type customerRepositoryDB struct {
	db *sqlx.DB
}

//ให้เข้าถึงส่วน new Struck ของ instant
//เปรียบเทีนยกับ OOP คือส่วน constructor นั่นเอง
func NewCustomerRepositoryDB(db *sqlx.DB) CustomerRepository {
	return customerRepositoryDB{db: db}

} //ให้ main ทำหน้าที่ instant ให้

func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id=?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}
	//ใส่ & เพราะ customer เป็น pointer
	return &customer, nil
}
