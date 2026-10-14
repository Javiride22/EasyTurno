package domain

// Trabajador es una Entidad que tiene identidad propia (ID)
type Trabajador struct {
	ID                     string
	Nombre                 string
	HorasContratoSemanales int
	Habilidades            []Puesto
	Indisponibilidad       []FranjaHoraria
}

// Puesto representa la habilidad requerida o que posee un trabajador (Objeto Valor)
type Puesto string

const (
	PuestoCocinero   Puesto = "Cocinero"
	PuestoRepartidor Puesto = "Repartidor"
	PuestoCajero     Puesto = "Horno"
	PuestoTelefono   Puesto = "Telefono"
)
