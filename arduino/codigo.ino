#include <DHT.h>
#include <ArduinoJson.h>

#define DHTPIN 8          // Pin digital conectado al sensor DHT11
#define DHTTYPE DHT11     // Define el tipo de sensor DHT
#define FANPIN 12         // Pin digital conectado al ventilador
#define LEDPIN 4          // Pin digital conectado al LED
#define TEMP_THRESHOLD 23 // Umbral de temperatura para encender el ventilador

DHT dht(DHTPIN, DHTTYPE);

int sensorLuz = A1;
int sensorMag = 10;

bool puertaAbierta = false;
unsigned long tiempoCierrePuerta = 0;
const unsigned long RETARDO_LED = 8000; // 8 segundos

void setup() {
  Serial.begin(9600);
  dht.begin();
  pinMode(FANPIN, OUTPUT);  // Configura el pin del ventilador como salida
  pinMode(LEDPIN, OUTPUT);  // Configura el pin del LED como salida
  pinMode(sensorLuz, INPUT);
  pinMode(sensorMag, INPUT);
}

void loop() {
  delay(1000);

  float h = dht.readHumidity();
  float t = dht.readTemperature();

  if (isnan(h) || isnan(t)) {
    Serial.println("Failed to read from DHT sensor!");
    return;
  }

  controlarVentilador(t);
  int luz = analogRead(sensorLuz);
  int mag = digitalRead(sensorMag);

  controlarLED(mag);
  enviarDatosJSON(t, mag, luz);
}

void controlarVentilador(float temperatura) {
  if (temperatura > TEMP_THRESHOLD) {
    digitalWrite(FANPIN, HIGH);  // Enciende el ventilador
  } else {
    digitalWrite(FANPIN, LOW);   // Apaga el ventilador
  }
}

void controlarLED(int estadoMagnetico) {
  if (estadoMagnetico == HIGH) { // Puerta cerrada
    if (puertaAbierta) {
      tiempoCierrePuerta = millis();
      puertaAbierta = false;
    }
  } else { // Puerta abierta
    digitalWrite(LEDPIN, HIGH); // Enciende el LED
    puertaAbierta = true;
  }

  if (!puertaAbierta && (millis() - tiempoCierrePuerta >= RETARDO_LED)) {
    digitalWrite(LEDPIN, LOW);
  }
}

void enviarDatosJSON(float temperatura, int estadoMagnetico, int luz) {
  DynamicJsonDocument jsonDoc(200);
  JsonObject root = jsonDoc.to<JsonObject>();

  root["id"] = 1;
  root["puerta"] = (estadoMagnetico == HIGH) ? "cerrado" : "abierto";
  root["luz"] = (luz > 240) ? "prendido" : "apagado";
  root["personas"] = temperatura;

  serializeJson(root, Serial);
  Serial.println(); // Agrega una línea en blanco al final para indicar el final del JSON
}
