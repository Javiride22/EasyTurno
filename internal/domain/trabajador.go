package domain

import "time"

type Trabajador struct {
	Nombre                 string
	HorasContratoSemanales int
	Habilidades            []Puesto
	Indisponibilidad       []time.Weekday
}
