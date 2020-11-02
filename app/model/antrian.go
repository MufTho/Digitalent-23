package model

import (
	"firebase.google.com/go/db"
	"fmt"
	"log"
	"strings"
)

type Antrian struct {
	ID string `json:"id"`
	Status bool `json:"status"`
}
//func for get all queue from firebase
func GetAntrian() (bool, []map[string]interface{}, error) {
	var data []map[string]interface{}
	ref := client.NewRef("Antrian")
	if err := ref.Get(ctx, &data); err != nil{
		log.Fatalln("Error Reading Form Database", err)
		return false, nil, err
	}
	return true, data, nil
}
//func for get add queue from firebase
func AddAntrian() (bool, error) {
	_,dataAntrian,_ := GetAntrian()
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
//func for get update queue from firebase
func UpdateAntrian(idAntrian string) (bool, error)  {
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

//func for get delete queue from firebase
func DeleteAntrian(idAntrian string) (bool, error)  {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	if err := childRef.Delete(ctx); err != nil{
		log.Fatal(err)
		return false, err
	}
	return true, nil
}

