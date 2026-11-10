package domain

import "time"

type Trabajador struct {
	dni                    string
	horasContratoSemanales int
	habilidades            []Puesto
	indisponibilidad       []time.Weekday
}
