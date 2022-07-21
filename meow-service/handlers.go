package main

import (
	"log"
	"meower/db"
	"meower/event"
	"meower/schema"
	"meower/util"
	"net/http"
	"text/template"
	"time"

	"github.com/segmentio/ksuid"
)

func createMeowHandler(w http.ResponseWriter,r *http.Request){
	type response struct{
		ID string `json:"id"`
	}
	ctx := r.Context()

	body := template.HTMLEscapeString(r.FormValue("body"))
	if len(body) < 1 || len(body) > 140 {
		util.ResponseError(w, http.StatusBadRequest, "Invalid body")
		return
	  }
	createdAt := time.Now().UTC()
	id, err := ksuid.NewRandomWithTime(createdAt)
	if err != nil{
		util.ResponseError(w,http.StatusInternalServerError,"Failed to create Meow")
		return
	}
	meow := schema.Meow{
		ID : id.String(),
		Body : body,
		CreatedAt: createdAt,
	}

	if err := db.InsertMeow(ctx,meow); err != nil{
		log.Println(err)
		util.ResponseError(w,http.StatusInternalServerError,"Failed to create Meow")
	}

	if err := event.PublishMeowCreated(meow); err != nil{
		log.Println(err)
	}

	util.ResponseOk(w, response{ID: meow.ID})
}

