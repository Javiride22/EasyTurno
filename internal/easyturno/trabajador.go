package easyturno

import "time"

type Trabajador struct {
	dni                    int
	horasContratoSemanales int
	habilidades            map[Puesto]bool
	indisponibilidad       map[time.Weekday]bool
}
