package service

import (
	"bank/repository"
	"database/sql"
)

// ที่ประกาศเป็นตัวเล็กเพราะว่าไม่อยากให้เข้าถึงโดยตรงเพราะป้องกัน error ที่เวลา user ส่งข้อมูลมาไม่ครบ
type customerService struct {
	custRepo repository.CustomerRepository
}

//ประกาศ struck ไม่ใช่ adapter แต่ต้องการอ้างถึง interface CustomerRepository ไม่ใช่ DB

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

// expose ผ่าน New เท่านั้น ทำหน้าที่แบบ constructor
func (s customerService) GetCustomers() ([]CustomerResponse, error) {

	customers, err := s.custRepo.GetAll()
	//เนื่องจาก service connect db เองไม่ได้ ต้องดึงผ่าน repo เท่านั้น
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	//ที่จริง return customer ออกไปเลยก็ได้ แต่ว่าไม่สวยเท่าใหร่ ปั้้น return ใหม่ดีกว่า
	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)
	if err != nil {

		//ดักว่าหาไม่เจอ
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
