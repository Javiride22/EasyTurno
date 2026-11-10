package domain

type DemandaPuestos struct {
	franja            FranjaHoraria
	puestosRequeridos map[Puesto]int
}
