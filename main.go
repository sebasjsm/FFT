package main // Se llama el paquete main

import (
	"fmt"
	// Se importa el paquete fmt el cual implementa funciones de entrada y salida
	"math"
	// Se importa el paquete math el cual implementa funciones matematicos como trigonométricas,
	// exponenciales y logarítmicas. Raices complejas
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

	json := make([]map[string]float32, len(resultado))
	for i, result := range resultado {
		magnitud := obtMagnitud(result)
		fase := obtFase(result)
		json[i] = map[string]float32{"mag": magnitud, "fase": fase}
	}
	fmt.Println(json)

}
