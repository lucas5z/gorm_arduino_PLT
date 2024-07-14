function actualizarDatos() {
    fetch('/prueba')
        .then(response => response.json())
        .then(data => {
            console.log(data); // Verificar los datos recibidos

            const door = document.getElementById('door');
            const light = document.getElementById('light');
            const temp = document.getElementById('temp');

            const puerta = data.puerta;
            const luz = data.luz;
            const temperatura = data.temperatura;

            // Mostrar condición de la puerta
            door.innerHTML = `
                <div class="door-frame">
                    <div class="door ${puerta === 'abierto' ? 'open' : ''}">Puerta ${puerta.charAt(0).toUpperCase() + puerta.slice(1)}</div>
                </div>
            `;

            // Mostrar condición de la luz
            light.innerHTML = `
                <div class="light-frame">
                    <div class="light ${luz === 'prendido' ? 'open' : ''}">Luz ${luz.charAt(0).toUpperCase() + luz.slice(1)}</div>
                </div>
            `;

            // Mostrar temperatura y estado del ventilador
            temp.innerHTML = `
                <div class="temp-frame">
                    <div class="temp">${temperatura} °C</div>
                    <div class="fan">${temperatura > 23 ? 'Ventiladores Encendidos' : 'Ventiladores Apagados'}</div>
                </div>
            `;
        })
        .catch(error => {
            console.error('Error:', error);
        });
}

setInterval(actualizarDatos, 1000);
