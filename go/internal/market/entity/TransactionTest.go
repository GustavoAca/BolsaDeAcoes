package entity

import (
	"testing"
)

func TestNewTransaction(t *testing.T) {
	// Criar instâncias de Order para usar no teste
	sellingOrder := &Order{}
	buyingOrder := &Order{}

	// Criar uma nova transação
	transaction := NewTransaction(sellingOrder, buyingOrder, 100, 10.0)

	// Verificar se os campos foram configurados corretamente
	if transaction.ID == "" {
		t.Error("ID não foi gerado corretamente")
	}

	if transaction.SellingOrder != sellingOrder {
		t.Error("sellingOrder não foi configurado corretamente")
	}

	if transaction.BuyingOrder != buyingOrder {
		t.Error("buyingOrder não foi configurado corretamente")
	}

	if transaction.Shares != 100 {
		t.Error("Shares não foi configurado corretamente")
	}

	if transaction.Price != 10.0 {
		t.Error("Price não foi configurado corretamente")
	}

	// Certifique-se de que o Total foi calculado corretamente
	expectedTotal := 100 * 10.0
	if transaction.Total != expectedTotal {
		t.Errorf("Total foi calculado incorretamente. Esperado: %f, Obtido: %f", expectedTotal, transaction.Total)
	}

	// Certifique-se de que DateTime foi definido
	if transaction.DateTime.IsZero() {
		t.Error("DateTime não foi configurado corretamente")
	}
}

func TestCalculateTotal(t *testing.T) {
	// Criar uma transação para usar no teste
	transaction := &Transaction{}

	// Chamar CalculateTotal
	transaction.CalculateTotal(50, 15.0)

	// Certifique-se de que o Total foi calculado corretamente
	expectedTotal := 50 * 15.0
	if transaction.Total != expectedTotal {
		t.Errorf("Total foi calculado incorretamente. Esperado: %f, Obtido: %f", expectedTotal, transaction.Total)
	}
}


func TestMinPendingShares(t *testing.T) {
	// Criar uma transação com ordens que têm 30 e 40 PendingShares
	transaction := &Transaction{
		SellingOrder: &Order{PendingShares: 30},
		BuyingOrder:  &Order{PendingShares: 40},
	}

	// Chamar minPendingShares
	result := transaction.minPendingShares()

	// Certifique-se de que o resultado é o mínimo entre 30 e 40, ou seja, 30
	if result != 30 {
		t.Errorf("minPendingShares retornou %d, esperado 30", result)
	}
}
