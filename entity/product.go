package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (p Product) StockStatus() string {
	status := ""
	if p.Stock <= 3 {
		status = "Stock hampir habis"
	} else if p.Stock < 10 {
		status = "Stock terbatas"
	} else {
		status = "Tersedia"
	}

	return status
}
