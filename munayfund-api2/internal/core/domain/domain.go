package domain

import "errors"

var (
	ErrProjectNotFound  = errors.New("project not found")
	ErrInvalidMediaType = errors.New("Tipo de multimedia no válido")
	ErrHistoryNotFound  = errors.New("Historial no encontrado")
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	WalletID string `json:"wallet_id" bson:"wallet_id"`
	Email    string `json:"email" bson:"email"`
}

// Multimedia representa la estructura para almacenar información multimedia en un proyecto.
type Multimedia struct {
	URL  string `json:"url" bson:"url"`
	Type string `json:"type" bson:"type"`
}

// ProjectAdvance representa la estructura para almacenar avances en un proyecto.
type ProjectAdvance struct {
	Date    string       `json:"date" bson:"date"`
	Details string       `json:"details" bson:"details"`
	Media   []Multimedia `json:"media" bson:"media"`
}

// Project representa la estructura para almacenar información de proyectos en MongoDB.
type Project struct {
	ID          string           `json:"id" bson:"_id"`
	UserID      string           `json:"userId" bson:"userId"`
	Title       string           `json:"title" bson:"title"`
	Description string           `json:"description" bson:"description"`
	Advances    []ProjectAdvance `json:"advances" bson:"advances"`
	Multimedia  []Multimedia     `json:"multimedia" bson:"multimedia"`
}
