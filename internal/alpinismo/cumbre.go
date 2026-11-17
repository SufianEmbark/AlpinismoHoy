package alpinismo

import "fmt"

type Cumbre struct {
	altitud                float64
	estaciones             map[string]Estacion
	datosTiempo            map[string]map[Hora]DatosTiempo
	coberturaNieveSatelite map[Dia]float64
}

type CoberturaNieveFueraDeRangoError struct {
	Cobertura float64
}

func (e *CoberturaNieveFueraDeRangoError) Error() string {
	return fmt.Sprintf("la cobertura de nieve es %v %% y tiene que estar entre 0 y 100", e.Cobertura)
}
