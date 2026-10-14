package domain

// Cuadrante agrupa turnos y asignaciones y tiene identidad propia (Semana)
type Cuadrante struct {
	Semana       string // identificador único del cuadrante
	Demandas     []DemandaTurno
	Asignaciones map[string][]FranjaHoraria // ID del trabajador -> Franjas asignadas
}
