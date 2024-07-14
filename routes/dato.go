package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lucas5z/arduino1/db"
	"github.com/lucas5z/arduino1/models"
)

// para enviar los datos
func Get_arduinio(w http.ResponseWriter, r *http.Request) {
	var dato models.Datos
	ardu := make([]int, 1)
	ardu[0] = 1

	da := db.DB.First(&dato, ardu[0])
	if da.Error != nil {
		log.Println("error interno get")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato)

}

// para llenar el primer dato
func Post_arduino(w http.ResponseWriter, r *http.Request) {
	var dato models.Datos

	err := json.NewDecoder(r.Body).Decode(&dato)
	if err != nil {
		log.Println("error interno post")
		return
	}
	da := db.DB.Create(&dato)
	if da.Error != nil {
		log.Println("error interno post")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dato)
}

// actualiza los datos

/* func Put_arduino(w http.ResponseWriter, r *http.Request) {
	var dato models.Datos
	ardu := make([]int, 1)
	ardu[0] = 1

	c := &serial.Config{Name: "COM3", Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		http.Error(w, "No se pudo abrir el archivo JSON", http.StatusInternalServerError)
		return
	}
	defer s.Close()

	err = json.NewDecoder(s).Decode(&dato)
	if err != nil {
		log.Println("no encontrado put-arduino")
		return
	}
	ar := db.DB.Save(&dato)
	if ar.Error != nil {
		log.Println("error interno put")
	}
	log.Println(&dato)

} */
