package helpers

import (
	"context"
	"cybercampus_module/configs"
	"cybercampus_module/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collectionUserModule = configs.GetCOllection(configs.Client, "users")
var collectionTemplate = configs.GetCOllection(configs.Client, "templates")


func SyncModuleTemplate(jenis_user primitive.ObjectID, idUser primitive.ObjectID) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Syncing module for user with id:", idUser)
	fmt.Println("Jenis user:", jenis_user)

	
	var userModule models.UserModule
	err := collectionUserModule.FindOne(ctx, bson.M{"_id": idUser}).Decode(&userModule)

	
	
	var template models.TemplateUserModuleRequest
	err = collectionTemplate.FindOne(ctx, bson.M{"_id": jenis_user}).Decode(&template)

	if err != nil {
		return false, err
	}


	updatedModules := []primitive.ObjectID{}
	updatedModules = append(updatedModules, template.Template...)
	// for _, module := range template.Template {
	// 	updatedModules = append(updatedModules, module.ID)
	// }

	
	update := bson.M{
		"$set": bson.M{
			"modules":    updatedModules,
			"updated_at": time.Now(),
		},
	}

	_, err = collectionUserModule.UpdateOne(ctx, bson.M{"_id": idUser}, update)
	if err != nil {
		return false, err
	}

	return true, nil
}



func SyncUpdateTemplate(jenis_user primitive.ObjectID, template []primitive.ObjectID) (bool, error){
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	hexId, err := primitive.ObjectIDFromHex(jenis_user.Hex())

	if err != nil {
		return false , err
	}

	var checkJenisUserModule models.UserModule
	err = collectionTemplate.FindOne(ctx, bson.M{"_id": hexId}).Decode(&checkJenisUserModule)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}
		return false, err
	}

	update := bson.M{
		"$set" : bson.M{
			"modules" : template,
		},
	}

	filter := bson.M{
		"jenis_user" : hexId,
	}

	_, err = collectionUserModule.UpdateMany(ctx, filter, update)

	if err != nil {
		return false, err
	}

	return true, nil
	
}