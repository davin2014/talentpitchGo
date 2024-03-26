// El paquete handlers contiene los manejadores de las rutas HTTP.
package handlers

import (
	"encoding/json"
	"net/http"
	"talentpitchGo/server"
)

// Home es una estructura que representa la respuesta que se enviará
// cuando se acceda a la ruta de inicio ("/"). Contiene un mensaje
// y un estado que indica si la solicitud fue exitosa.
type Home struct {
    Message string `json:"message"` // El mensaje a mostrar al usuario
    Status  bool   `json:"status"`  // El estado de la solicitud
}

// HomeHandler es una función que toma un servidor y devuelve un 
// manejador de funciones HTTP. Este manejador de funciones se encarga
// de manejar las solicitudes a la ruta de inicio ("/"). Cuando se accede
// a esta ruta, se envía una respuesta con un mensaje de bienvenida y un 
// estado de éxito.
func HomeHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Establecer el tipo de contenido de la respuesta a JSON
		w.Header().Set("Content-Type", "application/json")
		// Establecer el código de estado de la respuesta a 200 (OK)
		w.WriteHeader(http.StatusOK)
		// Codificar la respuesta como JSON y enviarla
		json.NewEncoder(w).Encode(Home{Message: "Welcome to TalentPitch", Status: true})
		
	}
}