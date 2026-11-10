package easyturno

type semana struct {
	anio   int
	numero int
}

type CuadranteSemanal struct {
	semana       semana
	jornadas     []Jornada
	trabajadores []Trabajador
	asignaciones []AsignacionTurno
}
