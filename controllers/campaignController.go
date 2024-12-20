package controllers

import (
	"context"
	"net/http"

	"campaign-app.local/database"
	"campaign-app.local/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var campaignCollection *mongo.Collection

func getCampaignCollection() *mongo.Collection {
	campaignCollection = database.GetCollection("campaigns")
	return campaignCollection
}

func GetCampaigns(c *gin.Context) {
	getCampaignCollection()
	
	cursor, err := campaignCollection.Find(context.TODO(), bson.M{"status": bson.M{"$ne": "deleted"}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(context.TODO())

	var campaigns []models.Campaign
	if err := cursor.All(context.TODO(), &campaigns); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, campaigns)
}

func GetCampaignByID(c *gin.Context) {
	getCampaignCollection()
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var campaign models.Campaign
	if err := campaignCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&campaign); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}
	c.JSON(http.StatusOK, campaign)
}

func CreateCampaign(c *gin.Context) {
	getCampaignCollection()
	var campaign models.Campaign
	if err := c.BindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	campaign.Status = "active"
	result, err := campaignCollection.InsertOne(context.TODO(), campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func UpdateCampaign(c *gin.Context) {
	getCampaignCollection()
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updateData map[string]interface{}
	if err := c.BindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, ok := updateData["status"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status is required"})
		return
	}

	if _, err := campaignCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": updateData}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Campaign updated successfully"})
}

func DeleteCampaign(c *gin.Context) {
	getCampaignCollection()
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if _, err := campaignCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, bson.M{"$set": bson.M{"status": "deleted"}}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}