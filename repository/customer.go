package repository

//เอาไม้ทำ get all เพราะต้องมี struct []Customer
type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zipcode"`
	Status      int    `db:"status"`
}

//ตั้งชื่อตัวใหญ่ เพราะจะให้ main เข้าถึงแค่ pods เท่านั้น ไม่ให้เข้าถึง adepter
//pod เต้ารับ เป็น interface
//go:generate mockgen -destination=../mock/mock_repository/mock_customer_repository.go bank/repository CustomerRepository
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(int) (*Customer, error)
	//ใส่ pointer เพราะต้อง return new ที่เป็น struck ไปด้วย ถ้าไม่้เป็นมัน return ไม่ได้
}
