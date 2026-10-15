package domain

type Cuadrante struct {
	semana       string
	demandas     []DemandaTurno
	asignaciones map[string][]FranjaHoraria
}
