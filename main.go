package main

import (
	"net/http"
	"log"
	"html/template"
	"time"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Button struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

type Components struct {
	PageTitle        string
	pageButtons      []Button
	Answer           string
}

type Activity struct {
	ID          bson.ObjectId  `bson:"_id" json:"id"`
	Time		time.Duration  `bson:"time" json:"time"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
}

type ActivitiesDA struct {
	Server   string
	Database string
}
var db *mgo.Database

const (
	COLLECTION = "Activities"

)
func (m *ActivitiesDA) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func main(){

	http.HandleFunc("/", WeekButtons)
	http.HandleFunc("/selected",Selected)
	//http.HandleFunc("/selcted", LogActivity)


	log.Fatal(http.ListenAndServe(":8000",nil))
}

func WeekButtons(w http.ResponseWriter, r *http.Request){
	Tittle := "log in plans"

	MyButtons := []Button{
		Button{"daySelect", "monday", false, false, "MONDAY"},
		Button{"daySelect", "tuesday", false, false, "TUESDAY"},
		Button{"daySelect", "wednesday", false, false, "WEDNESDAY"},
		Button{"daySelect", "thursdau", false, false, "THURSDAY"},
		Button{"daySelect", "friday", false, false, "FRIDAY"},
	}

	pageComponents := Components{
	PageTitle : Tittle,
	pageButtons : MyButtons,
	}

	t, err := template.ParseFiles("select.html")
	if err != nil{
		log.Print("cant read files")
	}

	err = t.Execute(w, pageComponents)
	if err != nil {
		log.Print("cant write response the html file")
	}
}

func Selected(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	daySelected := r.Form.Get("daySelect")

	Title := "Log into day"
	pageComponents := Components{
		PageTitle : Title,
		Answer:daySelected,
	}

	temp, err := template.ParseFiles("select.html")
	if err != nil{
		log.Print("cant read files")
	}

	err = temp.Execute(w, pageComponents)
	if err != nil {
		log.Print("cant write response the html file")
	}

}

func (m *ActivitiesDA) Insert(activity Activity) error {
	err := db.C(COLLECTION).Insert(&activity)
	return err
}
func (m *ActivitiesDA) Delete(activity Activity) error {
	err := db.C(COLLECTION).Remove(&activity)
	return err
}


