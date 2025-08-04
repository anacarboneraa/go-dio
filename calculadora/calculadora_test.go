package main

import "testing"

func TestShouldSumCorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldSumIncorrect(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 7
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldSubtractCorrect(t *testing.T) {
	teste := subtrair(10, 3, 2)
	resultado := 5
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldSubtractIncorrect(t *testing.T) {
	teste := subtrair(10, 3, 2)
	resultado := 4
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldDivideCorrect(t *testing.T) {
	teste := dividir(100, 2, 5)
	resultado := 10
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldDivideIncorrect(t *testing.T) {
	teste := dividir(100, 2, 5)
	resultado := 5
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldMultiplyCorrect(t *testing.T) {
	teste := multiplicar(2, 3, 4)
	resultado := 24
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}

func TestShouldMultiplyIncorrect(t *testing.T) {
	teste := multiplicar(2, 3, 4)
	resultado := 20
	if teste != resultado {
		t.Errorf("Valor esperado: %d, Valor retornado: %d", resultado, teste)
	}
}
