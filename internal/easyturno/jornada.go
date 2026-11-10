package easyturno

import "time"

type DemandaPuestos map[Puesto]int

type Jornada struct {
	dia     time.Weekday
	demanda DemandaPuestos
}
