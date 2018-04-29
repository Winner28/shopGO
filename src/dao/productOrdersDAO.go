package dao

type (
	ProductOrderDAO interface {
		Create(orderID, productID int) error
	}

	ProductOrderDAOImpl struct{}
)

func GetProductOrderDAO() ProductOrderDAO {
	return &ProductOrderDAOImpl{}
}

func (productOder *ProductOrderDAOImpl) Create(orderID, productID int) error {
	return nil
}
