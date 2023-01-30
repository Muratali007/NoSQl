package Handlers

import (
	"MongoDb/internal/data"
	"MongoDb/pkg/logging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"html/template"
	"net/http"
)

func Shop(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/shop.html"))
	tmpl.Execute(w, data.GetUser(w))
}

func List(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/list.html"))
	tmpl.Execute(w, data.GetUser(w))
	logger := logging.GetLogger()

	if r.Method == http.MethodPost {
		value := r.FormValue("element")

		switch value {
		case "cpu":
			data.Init("shop", "cpu")
			var items []data.Cpu

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.Cpu

				err := cur.Decode(&item)
				if err != nil {
					logger.Infof("error :", err)
				}
				items = append(items, item)
			}
			if err := cur.Err(); err != nil {
				logger.Infof("error :", err)
			}
			fmt.Println("Found multiple items:", items)

			err = tmpl.Execute(w, items)
			if err != nil {
				logger.Infof("error :", err)
			}
			data.Init("test", "users")
		case "cooling":
			tmpl := template.Must(template.ParseFiles("html/list2.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "cooling")
			var items []data.Cooling

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.Cooling

				err := cur.Decode(&item)
				if err != nil {
					logger.Infof("error :", err)
				}
				items = append(items, item)
			}
			if err := cur.Err(); err != nil {
				logger.Infof("error :", err)
			}
			fmt.Println("Found multiple items:", items)

			err = tmpl.Execute(w, items)
			if err != nil {
				logger.Infof("error :", err)
			}
			data.Init("test", "users")
		case "hdd":
			tmpl := template.Must(template.ParseFiles("html/list3.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "hdd")
			var items []data.Hdd

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.Hdd

				err := cur.Decode(&item)
				if err != nil {
					logger.Infof("error :", err)
				}
				items = append(items, item)
			}
			if err := cur.Err(); err != nil {
				logger.Infof("error :", err)
			}
			fmt.Println("Found multiple items:", items)

			err = tmpl.Execute(w, items)
			if err != nil {
				logger.Infof("error :", err)
			}
			data.Init("test", "users")

		case "housing":
			tmpl := template.Must(template.ParseFiles("html/list4.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "housing")
			var items []data.Housing

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.Housing

				err := cur.Decode(&item)
				if err != nil {
					logger.Infof("error :", err)
				}
				items = append(items, item)
			}
			if err := cur.Err(); err != nil {
				logger.Infof("error :", err)
			}
			fmt.Println("Found multiple items:", items)

			err = tmpl.Execute(w, items)
			if err != nil {
				logger.Infof("error :", err)
			}
			data.Init("test", "users")
		case "ram":
			tmpl := template.Must(template.ParseFiles("html/list5.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "ram")
			var items []data.Ram

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.Ram

				err := cur.Decode(&item)
				if err != nil {
					logger.Infof("error :", err)
				}
				items = append(items, item)
			}
			if err := cur.Err(); err != nil {
				logger.Infof("error :", err)
			}
			fmt.Println("Found multiple items:", items)

			err = tmpl.Execute(w, items)
			if err != nil {
				logger.Infof("error :", err)
			}
			data.Init("test", "users")

		case "motherboard":
			tmpl := template.Must(template.ParseFiles("html/list6.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "motherboard")
			var items []data.Motherboard

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.Motherboard

				err := cur.Decode(&item)
				if err != nil {
					logger.Infof("error :", err)
				}
				items = append(items, item)
			}
			if err := cur.Err(); err != nil {
				logger.Infof("error :", err)
			}
			fmt.Println("Found multiple items:", items)

			err = tmpl.Execute(w, items)
			if err != nil {
				logger.Infof("error :", err)
			}
			data.Init("test", "users")
		}
	}
}
