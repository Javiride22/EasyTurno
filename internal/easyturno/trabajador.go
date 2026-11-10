package easyturno

import "time"

type Trabajador struct {
	nombre                 string
	horasContratoSemanales int
	habilidades            []Puesto
	indisponibilidad       []time.Weekday
}
