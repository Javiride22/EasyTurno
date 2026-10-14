package domain

import "time"

// FranjaHoraria representa un bloque de tiempo inmutable (Objeto Valor)
type FranjaHoraria struct {
	Inicio time.Time
	Fin    time.Time
}
