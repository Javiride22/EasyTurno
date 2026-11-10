package easyturno

import "time"

type semana struct {
	anio   int
	numero int
}

type CuadranteSemanal struct {
	semana       semana
	jornadas     map[time.Weekday]Jornada
	trabajadores map[nombre]Trabajador
	asignaciones map[AsignacionTurno]struct{}
}
