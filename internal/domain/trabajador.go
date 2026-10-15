package domain

type Trabajador struct {
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
