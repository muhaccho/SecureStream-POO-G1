package main

import (
	"errors"
	"fmt"
	"strings"
)

// --- INTERFACES (Unidad 3) ---
// Define el comportamiento que debe tener cualquier contenido multimedia
type ContenidoMultimedia interface {
	Reproducir(edadUsuario int) error
	ObtenerInfo() string
}

// --- ENCAPSULACIÓN ---
// Atributos privados (minúscula) para proteger la integridad de los datos
type pelicula struct {
	titulo     string
	minimaEdad int
	esPremium  bool
}

// Constructor (Simula lo visto en Clase 3)
func NuevaPelicula(t string, edad int, premium bool) *pelicula {
	return &pelicula{titulo: t, minimaEdad: edad, esPremium: premium}
}

// Implementación de métodos de la Interfaz con Manejo de Errores
func (p *pelicula) Reproducir(edadUsuario int) error {
	if edadUsuario < p.minimaEdad {
		return errors.New("ACCESO DENEGADO: El usuario no cumple con el Rating de edad")
	}
	fmt.Printf("🎥 Reproduciendo ahora: %s... Disfrute su función.\n", p.titulo)
	return nil
}

func (p *pelicula) ObtenerInfo() string {
	return fmt.Sprintf("Título: %s | Edad Mínima: %d | Premium: %v", p.titulo, p.minimaEdad, p.esPremium)
}

func main() {
	// 1. Simulación de Servidor y Conexión (Clase 4)
	go func() {
		// El puerto 8000 se usa para verificar si el servicio de DB está vivo
		fmt.Println("\n--- " + strings.ToUpper("panel de control securestream") + " ---")
	}()
	fmt.Println(">>> SISTEMA SECURESTREAM CONECTADO EXITOSAMENTE A LA DB <<<")

	// 2. Datos de prueba (Catálogo dinámico - Clase 3)
	catalogo := []ContenidoMultimedia{
		NuevaPelicula("John Wick (R)", 18, true),
		NuevaPelicula("Toy Story (G)", 0, false),
		NuevaPelicula("Stranger Things (PG-13)", 13, true),
	}

	// 3. Menú Interactivo con Switch (Clase 3)
	var opcion int
	var edad int

	for {
		fmt.Println("\n--- PANEL DE CONTROL SECURESTREAM ---")
		fmt.Println("1. Ver Catálogo Disponible")
		fmt.Println("2. Intentar Reproducir Contenido")
		fmt.Println("3. Verificar Estado del Servidor")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Println("\n--- Catálogo en Memoria ---")
			for i, v := range catalogo {
				// Uso del identificador en blanco si fuera necesario, aquí usamos el índice i
				fmt.Printf("[%d] %s\n", i+1, v.ObtenerInfo())
			}
		case 2:
			fmt.Print("Ingrese su edad para validación de ciberseguridad: ")
			fmt.Scan(&edad)
			fmt.Println("Seleccione el número de película (1-3): ")
			var sel int
			fmt.Scan(&sel)

			if sel > 0 && sel <= len(catalogo) {
				err := catalogo[sel-1].Reproducir(edad)
				if err != nil {
					fmt.Println("⚠️ Alerta de Auditoría:", err)
				}
			} else {
				fmt.Println("Selección inválida.")
			}
		case 3:
			fmt.Println("Enviando PIN al servidor... [OK] Puerto 8000 Activo.")
		case 4:
			fmt.Println("Cerrando sesión segura...")
			return
		default:
			fmt.Println("Opción no reconocida.")
		}
	}
}
