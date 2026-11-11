package easyturno

import "time"

type AsignacionTurno struct {
	trabajador string
	puesto     Puesto
	dia        time.Weekday
}
