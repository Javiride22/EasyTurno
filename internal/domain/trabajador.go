package domain

type Trabajador struct {
	id                     string
	nombre                 string
	horasContratoSemanales int
	habilidades            []Puesto
	indisponibilidad       []FranjaHoraria
}

type Puesto string

const (
	PuestoCocinero   Puesto = "Cocinero"
	PuestoRepartidor Puesto = "Repartidor"
	PuestoCajero     Puesto = "Horno"
	PuestoTelefono   Puesto = "Telefono"
)
