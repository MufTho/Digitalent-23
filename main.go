package main

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"strings"
)
var client *db.Client
var ctx context.Context

func init()  {
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://digitelent-be-23.firebaseio.com/",
	}

	opt := option.WithCredentialsFile("firebase-admin-sdk.json")
	app, err := firebase.NewApp(ctx, conf,opt)
	if err != nil{
		log.Fatalf("Error initializing app", err)
	}
	client, err = app.Database(ctx)
	if err != nil{
		log.Fatalf("Error Initializing Database", err)
	}

}

//var data []Antrian
//Antrian @dec queue
type Antrian struct {
	ID string `json:"id"`
	Status bool `json:"status"`

}

func main()  {
	router := gin.Default()
	router.POST("/api/v1/antrian", AddAntrianHandler)
	router.GET("/api/v1/antrian/status", GetAntrianHandler)
	router.PUT("/api/v1/antrian/id/:idAntrian", UpdateAntrianHandler)
	router.DELETE("/api/v1/antrian/id/:idAntrain/delete", DeleteAntrianHandler)
	router.Run(":8080")
}

func getAntrian() (bool, []map[string]interface{}, error) {
	var data []map[string]interface{}
	ref := client.NewRef("Antrian")
	if err := ref.Get(ctx, &data); err != nil{
		log.Fatalln("Error Reading Form Database", err)
		return false, nil, err
	}
	return true, data, nil
}

func addAntrian() (bool, error) {
	_,dataAntrian,_ := getAntrian()
	var ID string
	var antrianRef *db.Ref
	ref := client.NewRef("Antrian")

	if dataAntrian == nil{
		ID = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	}else {
		ID = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}
	antrian := Antrian{
		ID: ID,
		Status: false,
	}

	if err := antrianRef.Set(ctx, antrian); err != nil{
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func updateAntrian(idAntrian string) (bool, error)  {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	antrian := Antrian{
		ID: idAntrian,
		Status: true,
	}
	if err := childRef.Set(ctx, antrian); err != nil{
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func deleteAntrian(idAntrian string) (bool, error)  {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	if err := childRef.Delete(ctx); err != nil{
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

func AddAntrianHandler(c *gin.Context)  {
	flag, err := addAntrian()
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	}else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}

func GetAntrianHandler(c *gin.Context)  {
	flag, data, err := getAntrian()
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
			"data": data,
		})
	}else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}

func UpdateAntrianHandler(c *gin.Context)  {
	idAntrian := c.Param("idAntrian")
	flag, err := updateAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	}else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}

func DeleteAntrianHandler(c *gin.Context)  {
	idAntrian := c.Param("idAntrian")
	flag, err := deleteAntrian(idAntrian)
	if flag {
		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "success",
		})
	}else{
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status": "failed",
			"error": err,
		})
	}
}