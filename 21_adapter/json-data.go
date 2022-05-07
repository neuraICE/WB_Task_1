package _1_adapter

import "fmt"

/*
Реализовать паттерн «адаптер» на любом примере.
*/

// Сервис работает с данными в формате json
type JsonDocument struct {
	Id    string
	Value string
}

func (doc JsonDocument) ConvertToXml() string {
	return "<xml></xml>"
}

// Адаптер позволяет взаимодейстовать с обоими сервисами не изменяя их код

// Структура JsonDocumentAdapter использует методы и поля структуры JsonDocument
// и наследует интерфейс AnalyticalDataService, реализуя его метод SendXmlData
type JsonDocumentAdapter struct {
	jsonDocument *JsonDocument
}

func (adapter JsonDocumentAdapter) SendXmlData() {
	adapter.jsonDocument.ConvertToXml()
	fmt.Println("Отправка xml документа!")
}
