package domain

// DemandaTurno representa las necesidades del local en una franja horaria
type DemandaTurno struct {
	Franja            FranjaHoraria
	PuestosRequeridos map[Puesto]int
}
