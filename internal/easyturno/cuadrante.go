package easyturno

import "time"

type semana struct {
	anio   int
	numero int
}

type demandaPuestos map[Puesto]int

type CuadranteSemanal struct {
	semana       semana
	jornadas     map[time.Weekday]demandaPuestos
	trabajadores map[string]Trabajador
	asignaciones map[AsignacionTurno]struct{}
}
