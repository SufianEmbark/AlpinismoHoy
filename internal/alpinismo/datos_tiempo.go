package alpinismo

import "fmt"

type DatosTiempo struct {
	temperaturaMaxima float64
	temperaturaMinima float64
	precipitacion     float64
}

type MinimaMayorQueMaximaError struct {
	Minima float64
	Maxima float64
}

func (e *MinimaMayorQueMaximaError) Error() string {
	return fmt.Sprintf("la temperatura mínima es %v y no puede ser mayor que la máxima, que es %v", e.Minima, e.Maxima)
}
