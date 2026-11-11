package easyturno

import "time"

type Trabajador struct {
	horasContratoSemanales int
	habilidades            map[Puesto]bool
	indisponibilidad       map[time.Weekday]bool
}
