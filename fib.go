// fib contiene dos funciones para calcular números de la secuencia fibonacci.
package fib

// Fib calcula el n número de la secuencia fibonacci.
func Fib(n int) int {
	return FibCiclo(n)
}

// FibCiclo calcula un número de la secuencia fibonacci para la posición n.
// Esta versión usa un ciclo for para hacer el cálculo.
func FibCiclo(n int) int {
		a, b := 0, 1
		for i:=0; i<n; i++ {
			a, b = b, a+b
		}
		return a
}


// FibRecursivo calcula un número de la secuencia fibonacci para la posición n.
// Esta versión usa recursión para hacer el cálculo.
func FibRecursivo(n int) int {
	if n < 2 {
		return n
	}

	return FibRecursivo(n-1) + FibRecursivo(n-2)
}

