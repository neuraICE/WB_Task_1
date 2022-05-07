package _1_adapter

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Сервис аналитики который принимает данные в формате xml
type AnalyticalDataService interface {
	SendXmlData()
}

type XmlDocument struct {
}

func (doc XmlDocument) SendXmlData() {
	fmt.Println("Отправка xml документа!")
}
