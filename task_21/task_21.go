package main

import "fmt"

func main() {
	euSocket := EUSocket{}

	euPlug := EUPlugImpl{}

	usPlug := USPlugImpl{}

	// Переходник с американской вилки на евроскую
	usPlugAdapter := USPlugAdapter{&usPlug}

	euSocket.connect(&euPlug)
	euSocket.connect(&usPlugAdapter)
}

type EUSocket struct{}

// Розетка в которую подключается европейская вилка
func (socket EUSocket) connect(plug EUPlug) {
	fmt.Println(plug.useEUSocket())
}

type EUPlug interface {
	useEUSocket() string
}

type USPlug interface {
	useUSSocket() string
}

type EUPlugImpl struct{}

func (plug *EUPlugImpl) useEUSocket() string {
	return "eu plug"
}

type USPlugImpl struct{}

func (plug *USPlugImpl) useUSSocket() string {
	return "us plug"
}

type USPlugAdapter struct {
	USPlug
}

func (adapter *USPlugAdapter) useEUSocket() string {
	return "us plug adapter"
}
