package domain

import "time"

type Trabajador struct {
	HorasContratoSemanales int
	Habilidades            []Puesto
	Indisponibilidad       []time.Weekday
}
