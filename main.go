package main // Se llama el paquete main

import (

	// Se importa el paquete fmt el cual implementa funciones de entrada y salida

	"encoding/json"
	"fmt"

	"math"
	"text/template"

	// Se importa el paquete math el cual implementa funciones matematicos como trigonométricas,
	// exponenciales y logarítmicas. Raices complejas
	"log"
	"net/http"
)

/*
La función separarImpares toma una serie de números complejos
y devuelve una lista con los elementos impares.
*/

func separarImpares(x ...complex64) []complex64 {
	pos := 0
	n := len(x) / 2
	impar := make([]complex64, n)

	for i := 0; i < len(x); i++ {
		if i%2 != 0 {
			impar[pos] = x[i]
			pos++
		}
	}
	return impar
}

/*
La función separarPares toma una lista de números complejos y devuelve una lista que contiene
solo los elementos en posiciones pares.
*/

func separarPares(x ...complex64) []complex64 {
	pos := 0
	n := len(x) / 2
	par := make([]complex64, n)

	for i := 0; i < len(x); i++ {
		if i%2 == 0 {
			par[pos] = x[i]
			pos++
		}
	}
	return par
}

/*
La funcion transformada F, toma una lista o un arreglo de numeros complejos, para hacer la transformada
de un tamaño n, se utiliza un metodo de divide y venceras. que es un metodo que lo resuelve de forma
mas eficiente desarrollandolo por partes.
*/

func transformadaF(x ...complex64) []complex64 {
	w := complex64(0)
	z := complex64(0)
	n := len(x)

	if n <= 1 {
		return x
	}

	par := make([]complex64, n/2)
	impar := make([]complex64, n/2)

	par = separarPares(x...)
	impar = separarImpares(x...)

	par = transformadaF(par...)
	impar = transformadaF(impar...)

	for i := 0; i < n/2; i++ {
		imaginaria := float32(i) / float32(n)
		w = complex(float32(math.Cos(2*math.Pi*float64(imaginaria))), imag(w))
		w = complex(real(w), float32(-1*math.Sin(2*math.Pi*float64(imaginaria))))
		z = complex(real(w)*real(impar[i])-imag(w)*imag(impar[i]), real(w)*imag(impar[i])+imag(w)*real(impar[i]))
		x[i] = complex(real(par[i])+real(z), imag(x[i]))
		x[i] = complex(real(x[i]), imag(par[i])+imag(z))
		x[i+(n>>1)] = complex(real(par[i])-real(z), imag(x[i+n/2]))
		x[i+(n>>1)] = complex(real(x[i+n/2]), imag(par[i])-imag(z))
	}
	return x

}

func obtMagnitud(x complex64) float32 {
	return float32(math.Sqrt(float64(real(x)*real(x) + imag(x)*imag(x))))
}

func obtFase(x complex64) float32 {
	return float32(math.Atan2(float64(imag(x)), float64(real(x))))
}

func main() {
	x := []complex64{complex(0, 0), complex(1, 0), complex(2, 0), complex(3, 0), complex(4, 0), complex(5, 0), complex(6, 0), complex(7, 0)}
	resultado := transformadaF(x...)
	fmt.Println(resultado)

	for _, result := range resultado {
		magnitud := obtMagnitud(result)
		fase := obtFase(result)
		fmt.Printf("Magnitud: %f, Fase: %f\n", magnitud, fase)
	}

	jsonData := make([]map[string]float32, len(resultado))
	for i, result := range resultado {
		magnitud := obtMagnitud(result)
		fase := obtFase(result)
		jsonData[i] = map[string]float32{"mag": magnitud, "fase": fase}
	}
	// Configurar el manejador para los archivos estáticos
	fs := http.FileServer(http.Dir("."))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Configurar el manejador para la página principal
	http.HandleFunc("/", homeHandler)

	// Iniciar el servidor en el puerto 8080
	log.Println("Servidor escuchando en el puerto 8080...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	x := []complex64{complex(0, 0), complex(1, 0), complex(2, 0), complex(3, 0), complex(4, 0), complex(5, 0), complex(6, 0), complex(7, 0)}
	resultado := transformadaF(x...)

	jsonData := make([]map[string]float32, len(resultado))
	for i, result := range resultado {
		magnitud := obtMagnitud(result)
		fase := obtFase(result)
		jsonData[i] = map[string]float32{"mag": magnitud, "fase": fase}
	}

	jsonBytes, err := json.Marshal(jsonData)
	if err != nil {
		log.Println("Error al convertir a JSON:", err)
		http.Error(w, "Error interno del servidor", http.StatusInternalServerError)
		return
	}

	// Renderizar la página HTML con la gráfica
	tpl := `
    <html>
    <head>
        <title>GRAFICA TRANSFORMADA</title>
        <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    </head>
    <body>
		<h1>GRAFICA TRANSFORMADA DE FOURIER</h1><br> <br> 
        <canvas id="myChart" width="400" height="400"></canvas>
        <script>
            var jsonData = {{.}};
            var labels = [];
            var magnitudes = [];

            for (var i = 0; i < jsonData.length; i++) {
                labels.push("Punto " + (i + 1));
                magnitudes.push(jsonData[i].mag);
            }

            var ctx = document.getElementById('myChart').getContext('2d');
            var chart = new Chart(ctx, {
                type: 'line',
                data: {
                    labels: labels,
                    datasets: [{
                        label: 'Magnitudes',
                        data: magnitudes,
                        backgroundColor: 'rgba(75, 0, 0, 1)',
                        borderColor: 'rgba(75, 0, 0, 1)',
                        borderWidth: 1
                    }]
                },
                options: {
                    scales: {
                        y: {
                            beginAtZero: true
                        }
                    }
                }
            });
        </script>
    </body>
    </html>
    `

	// Establecer la cabecera y escribir la página HTML
	w.Header().Set("Content-Type", "text/html")
	tplParsed := template.Must(template.New("").Parse(tpl))
	tplParsed.Execute(w, string(jsonBytes))
}
