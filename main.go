package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/lucas5z/arduino1/db"
	"github.com/lucas5z/arduino1/models"
	"github.com/lucas5z/arduino1/routes"
	"github.com/tarm/serial"
)

var serialPort *serial.Port

func main() {
	// Conexión a la base de datos
	db.Conex()
	db.DB.AutoMigrate(&models.Datos{})

	// Configuración del puerto serial
	var err error
	serialPort, err = serial.OpenPort(&serial.Config{Name: "COM3", Baud: 9600})
	if err != nil {
		log.Fatal(err)
	}
	defer serialPort.Close()

	r := mux.NewRouter()

	// Rutas
	r.HandleFunc("/prueba", routes.Get_arduinio).Methods("GET")
	r.HandleFunc("/prueba", routes.Post_arduino).Methods("POST")
	go Put_arduino_time()

	// Ruta para servir el archivo HTML del frontend
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./frontend/")))

	port := ":9090"
	fmt.Printf("Servidor escuchando en el puerto %s...\n", port)
	err = http.ListenAndServe(port, r)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}

func Put_arduino_time() {
	for {
		Put_arduino(serialPort, nil)
		time.Sleep(1000 * time.Millisecond)
	}
}

func Put_arduino(serialPort *serial.Port, w http.ResponseWriter) {
	var dato models.Datos

	// Leer datos del puerto serial
	err := json.NewDecoder(serialPort).Decode(&dato)
	if err != nil {
		log.Println("Error al decodificar datos del Arduino:", err)
		if w != nil {
			http.Error(w, "No se pudo leer datos del Arduino", http.StatusInternalServerError)
		}
		return
	}

	// Guardar datos en la base de datos
	ar := db.DB.Save(&dato)
	if ar.Error != nil {
		log.Println("Error al guardar datos en la base de datos:", ar.Error)
		if w != nil {
			http.Error(w, "Error al guardar datos en la base de datos", http.StatusInternalServerError)
		}
		return
	}

	log.Println("Datos recibidos:", dato)
}
