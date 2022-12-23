package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func AddToCart() gin.HandlerFunc {
	return nil
}

func RemoteItem() gin.HandlerFunc {
	return nil
}

func GetItemFromCart() gin.HandlerFunc {
	return nil
}

func BuyFromCart() gin.HandlerFunc {
	return nil
}

func InstantBuy() gin.HandlerFunc {
	return nil
}
