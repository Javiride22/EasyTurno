package domain

import "time"

type DemandaPuestos map[Puesto]int

type Jornada struct {
	Dia     time.Weekday
	Demanda DemandaPuestos
}
