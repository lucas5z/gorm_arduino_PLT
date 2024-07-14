package routes

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/lucas5z/arduino1/models"
	"github.com/tarm/serial"
)

func Open2() {

	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	var datos models.Datos

	decoder := json.NewDecoder(s)

	for {
		err := decoder.Decode(&datos)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(datos)
	}
}
