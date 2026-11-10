package domain

type Trabajador struct {
	dni                    string
	horasContratoSemanales int
	habilidades            []Puesto
	indisponibilidad       []FranjaHoraria
}
