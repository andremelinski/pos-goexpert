package tax

import "errors"

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0, nil
	}
	if amount >= 1000.0 && amount < 20000{
		return 10.0, nil
	}
	if amount >= 20000 {
		return 20.0, nil
	}
	return 5.0, errors.New("deu erro")
}

// interface que seria usada para uma comunicacao com o banco
// separa a comunicacao do banco com a aplicacao, assim as camadas ficam independente e pode testar
type Repository interface {
	SaveTax(amount float64) error
}

func CalculateTaxAndSave(amount float64, repo Repository) error{
	tax := CalculateTax2(amount) 
	// com a interface vc "emula"/"forja" essa logica com o db para nao atrapalhar nos testes
	return repo.SaveTax(tax)
}

func CalculateTax2(amount float64) float64 {
	if amount <= 0 {
		return 0
	}
	if amount >= 1000.0 && amount < 20000{
		return 10.0
	}
	if amount >= 20000 {
		return 20.0
	}
	return 5.0
}