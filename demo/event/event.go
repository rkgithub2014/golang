package event

import (
	"encoding/json"
	"fmt"
	"log"
)

//InitConfig initializes
func InitConfig() *Palos {
	fmt.Println("Event")
	paloConfig := `{"PASW_ENABLED":"true", "PASW_PORT_START":"50001"}`
	var gogPalos Palos

	if err := json.Unmarshal([]byte(paloConfig), &gogPalos); err != nil {
		log.Fatal(err)
	}
	return &gogPalos
}

//Palos ssmConfig
type Palos struct {
	Enabled   bool `json:"PASW_ENABLED"`
	PortStart int  `json:"PASW_PORT_START"`
}

//UnMarshallText to Palos
func (pa *Palos) UnMarshallText(text []byte) error {
	json := string(text)
	fmt.Println(json)
	return nil
}

//MarshallText Palos to Json
func (pa Palos) MarshallText() ([]byte, error) {
	var value string
	return []byte(value), nil
}
