package easyturno

type CuadranteSemanal struct {
	id           string
	jornadas     []Jornada
	trabajadores []Trabajador
	asignaciones []AsignacionTurno
}
