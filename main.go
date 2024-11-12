package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Mostrar el formulario HTML
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Verificar Par o Impar</title>
		</head>
		<body>
			<h1>Ingrese un número</h1>
			<form action="/check" method="post">
				<input type="number" name="number" required>
				<input type="submit" value="Verificar">
			</form>
		</body>
		</html>
	`)
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Obtener el número del formulario
		numStr := r.FormValue("number")

		// Convertir el número a un entero
		num, err := strconv.Atoi(numStr)
		if err != nil {
			http.Error(w, "Número inválido", http.StatusBadRequest)
			return
		}

		// Determinar si es par o impar
		result := ""
		if num%2 == 0 {
			result = fmt.Sprintf("%d es un número par.", num)
		} else {
			result = fmt.Sprintf("%d es un número impar.", num)
		}

		// Mostrar el resultado y un botón para volver
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Resultado</title>
			</head>
			<body>
				<h1>%s</h1>
				<a href="/"> <button>Volver</button> </a>
			</body>
			</html>
		`, result)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/check", checkHandler)
	fmt.Println("Servidor escuchando en el puerto 8084...")
	err := http.ListenAndServe(":8084", nil)
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}