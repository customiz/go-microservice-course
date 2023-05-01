package service

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

// custom return จาก Database เพราะ Respository จะ return เป็น db
//service จะ return เป็น json สามารถ custom ส่วนที่จะให้ดึงได้

//go:generate mockgen -destination=../mock/mock_service/mock_customer_service.go bank/service CustomerService
type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(int) (*CustomerResponse, error)
}
