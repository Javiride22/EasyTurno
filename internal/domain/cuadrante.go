package domain

type Cuadrante struct {
	semana       string
	trabajadores map[string]Trabajador
	demandas     []DemandaTurno
	asignaciones map[string][]FranjaHoraria
}
