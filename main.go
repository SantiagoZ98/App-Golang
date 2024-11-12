package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Definimos el contenido HTML
	htmlContent := `
        <!DOCTYPE html>
        <html lang="es">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>App con Golang</title>
            <style>
                body {
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    margin: 0;
                    font-family: Arial, sans-serif;
                    background-color: #f0f8ff;
                }
                h1 {
                    color: #333;
                    font-size: 3rem;
                    background: linear-gradient(135deg, #ff6b6b, #f0e68c);
                    -webkit-background-clip: text;
                    color: transparent;
                    text-align: center;
                }
            </style>
        </head>
        <body>
            <h1>Hello world, my name is Santiago and this is my app with Golang and Docker</h1>
        </body>
        </html>
    `
	// Enviamos el contenido HTML como respuesta
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, htmlContent)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
