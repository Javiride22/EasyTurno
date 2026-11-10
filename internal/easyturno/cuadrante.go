package easyturno

type CuadranteSemanal struct {
	ID           string
	Jornadas     []Jornada
	Trabajadores []Trabajador
	Asignaciones []AsignacionTurno
}
