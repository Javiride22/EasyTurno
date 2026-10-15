package domain

type Cuadrante struct {
	trabajadores map[string]Trabajador
	demandas     []DemandaTurno
	asignaciones map[string][]FranjaHoraria
}

type Planificacion map[string]Cuadrante
