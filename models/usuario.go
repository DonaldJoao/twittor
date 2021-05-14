package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* Usuario es el modelo de usuario de la base de datos */
type Usuario struct {
	ID               primitive.ObjectID `bson:"_id,omitempy" json:"id"`
	Nombre           string             `bson:"nombre" json:"nombre,omitempy"`
	Apellidos        string             `bson:"apellidos" json:"apellidos,omitempy"`
	FechaNacimientos time.Time          `bson:"fechaNacimientos" json:"fechaNacimientos,omitempy"`
	Email            string             `bson:"email" json:"email"`
	Password         string             `bson:"password" json:"password,omitempy"`
	Avatar           string             `bson:"avatar" json:"avatar,omitempy"`
	Banner           string             `bson:"banner" json:"banner,omitempy"`
	Biografia        string             `bson:"biografia" json:"biografia,omitempy"`
	SitioWeb         string             `bson:"sitioWeb" json:"sitioWeb,omitempy"`
}
