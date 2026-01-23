package main

import (
	"fmt"
	"strings"
)

// Estructura para el contenido (Clase 3: Organización de contenidos)
type Video struct {
	Titulo    string
	Rating    string // Ej: "G", "PG-13", "R"
	IsPremium bool
}

// Estructura de Usuario para Ciberseguridad
type Usuario struct {
	Nombre   string
	Edad     int
	Suscrito bool
}

// 1. Función Variádica para Auditoría de Seguridad (Clase 2)
// Permite registrar múltiples eventos de seguridad a la vez.
func auditoriaSeguridad(userId string, eventos ...string) {
	fmt.Printf("\n--- REPORTE DE AUDITORÍA [ID: %s] ---\n", userId)
	for _, evento := range eventos { // Clase 2: Uso de guion bajo para ignorar el índice
		fmt.Println("[LOG]:", strings.ToUpper(evento)) // Clase 2: Manipulación de strings
	}
}

func main() {
	// 2. Uso de Maps para almacenamiento eficiente (Clase 3)
	// La clave es el nombre de usuario para búsquedas rápidas.
	baseUsuarios := make(map[string]Usuario)
	baseUsuarios["seguridad_admin"] = Usuario{Nombre: "Admin", Edad: 25, Suscrito: true}
	baseUsuarios["invitado_01"] = Usuario{Nombre: "Juan", Edad: 15, Suscrito: false}

	// 3. Uso de Slices para el Catálogo (Clase 3)
	catalogo := []Video{
		{Titulo: "Cyber Wars", Rating: "R", IsPremium: true},
		{Titulo: "Algoritmos para Niños", Rating: "G", IsPremium: false},
	}

	// Simulacro de inicio de sesión
	idBuscado := "invitado_01"
	user, existe := baseUsuarios[idBuscado]

	// 4. Lógica de Control de Acceso (Clase 1: Operadores Lógicos y de Comparación)
	if existe {
		fmt.Printf("Validando acceso para: %s...\n", user.Nombre)

		for _, video := range catalogo {
			// Regla de Ciberseguridad: Validar edad y suscripción simultáneamente
			// Usamos operadores: && (AND), >= (Mayor o igual), == (Igualdad)
			if user.Edad < 18 && video.Rating == "R" {
				auditoriaSeguridad(idBuscado, "Bloqueo de contenido sensible", "Restricción por edad")
				continue
			}

			if video.IsPremium && !user.Suscrito { // ! operador NOT
				fmt.Printf("Contenido '%s' requiere suscripción Premium.\n", video.Titulo)
			} else {
				fmt.Printf("Reproduciendo: %s [Acceso Autorizado]\n", video.Titulo)
			}
		}
	} else {
		auditoriaSeguridad("DESCONOCIDO", "Intento de acceso fallido", "Usuario no registrado")
	}
}
