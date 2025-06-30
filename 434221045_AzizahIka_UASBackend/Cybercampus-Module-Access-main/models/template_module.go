package models

import (
	// "cybercampus_module/models"
	
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TemplateRequest struct {	
	ID primitive.ObjectID `json:"id" bson:"_id"`
	JenisUser string  `json:"jenis_user" bson:"jenis_user"`
	Template  []primitive.ObjectID `json:"template" bson:"template"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type JenisUserResponse struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	JenisUser string `json:"jenis_user" bson:"jenis_user"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

// cuma ngeget jenis usernya saja
// cuma mengembalikan nilai idnya dan jenis usernya apa
// jadi satu collection
// jadi fungsi model misal kalian 1 collection banyak atribut, bagaimana cara kalian mengeluarkannya ke API bagaimana resonnya itu dibatasi lewat models

type TemplateResponse struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
	JenisUser string `json:"JENIS_USER" bson:"jenis_user"`
	Template  []ModuleResponse `json:"TEMPLATE_MODULES" bson:"template"`
	CreatedAt time.Time `json:"CREATED_AT" bson:"created_at"`
	UpdatedAt time.Time `json:"UPDATED_AT" bson:"updated_at"`
}

type TemplateUserModuleRequest struct {
	JenisUser string `json:"jenis_user" bson:"jenis_user"`
	Template []primitive.ObjectID `json:"template" bson:"template"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}
