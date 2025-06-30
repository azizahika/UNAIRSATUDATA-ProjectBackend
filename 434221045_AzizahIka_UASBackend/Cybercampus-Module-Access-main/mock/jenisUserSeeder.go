package mock

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionTemplate *mongo.Collection = configs.GetCOllection(configs.Client, "templates")

func JenisUserSeeder() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	newJenisUser := []interface{}{
		models.TemplateRequest{
			ID: primitive.NewObjectID(),
			JenisUser: "mahasiswa",
			Template:  []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			ID: primitive.NewObjectID(),
			JenisUser: "dosen",
			Template:  []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			ID: primitive.NewObjectID(),
			JenisUser: "tendik",
			Template:  []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			ID: primitive.NewObjectID(),
			JenisUser: "kps",
			Template:  []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			ID: primitive.NewObjectID(),
			JenisUser: "dekanat",
			Template:  []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			ID: primitive.NewObjectID(),
			JenisUser: "ketua_unit",
			Template:  []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		models.TemplateRequest{
			ID: primitive.NewObjectID(),
			JenisUser: "pimpinan_univ",
			Template:  []primitive.ObjectID{},
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	_, err := collectionTemplate.InsertMany(ctx, newJenisUser)
	if err != nil {
		// Handle error gracefully
		fmt.Println("Error inserting data:", err)
		return
	}

	fmt.Println("Jenis User Seeder Created")
}
