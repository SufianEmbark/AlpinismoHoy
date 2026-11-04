package alpinismo

import "time"

type Cumbre struct {
	datosTiempo            map[time.Time]map[string]DatosTiempo
	coberturaNieveSatelite map[time.Time]float64
}
