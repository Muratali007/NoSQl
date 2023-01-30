package Handlers

import (
	"MongoDb/internal/data"
	"MongoDb/pkg/logging"
	"context"
	"encoding/base64"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"log"
	"net/http"
)

var images string
var number int

/*Shop function*/
func Shop(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/shop.html"))
	tmpl.Execute(w, data.GetUser(w))
}

func send(name string, num int) (string, int) {
	images = name
	number = num
	return images, number
}

/* function for images for cpu*/
func Images(w http.ResponseWriter, r *http.Request) {

	if number == 1 {
		data.Init("shop", "powersupplies")
		var doc bson.M

		err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		// Decode the base64 string and serve it as a JPEG image
		imgBase64 := images
		img, err := base64.StdEncoding.DecodeString(imgBase64)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(img)
		data.Init("test", "users")
	}
	if number == 2 {
		data.Init("shop", "cpu")
		var doc bson.M

		err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		// Decode the base64 string and serve it as a JPEG image
		imgBase64 := images
		img, err := base64.StdEncoding.DecodeString(imgBase64)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(img)
		data.Init("test", "users")
	}
	if number == 3 {
		data.Init("shop", "cooling")
		var doc bson.M

		err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		// Decode the base64 string and serve it as a JPEG image
		imgBase64 := images
		img, err := base64.StdEncoding.DecodeString(imgBase64)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(img)
		data.Init("test", "users")
	}

	if number == 4 {
		data.Init("shop", "gpu")
		var doc bson.M

		err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		// Decode the base64 string and serve it as a JPEG image
		imgBase64 := images
		img, err := base64.StdEncoding.DecodeString(imgBase64)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(img)
		data.Init("test", "users")
	}

	if number == 5 {
		data.Init("shop", "housing")
		var doc bson.M

		err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		// Decode the base64 string and serve it as a JPEG image
		imgBase64 := images
		img, err := base64.StdEncoding.DecodeString(imgBase64)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(img)
		data.Init("test", "users")
	}

	if number == 6 {
		data.Init("shop", "motherboards")
		var doc bson.M

		err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
		if err != nil {
			log.Fatal(err)
		}

		// Decode the base64 string and serve it as a JPEG image
		imgBase64 := images
		img, err := base64.StdEncoding.DecodeString(imgBase64)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "image/jpeg")
		w.Write(img)
		data.Init("test", "users")
	}
}

func ImageMotherboard(w http.ResponseWriter, r *http.Request) {

	data.Init("test", "imageMotherboard")
	var doc bson.M

	err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the base64 string and serve it as a JPEG image
	imgBase64 := doc["image"].(string)
	img, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(img)
	data.Init("test", "users")
}

func ImageHousing(w http.ResponseWriter, r *http.Request) {

	data.Init("test", "imagehousing")
	var doc bson.M

	err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the base64 string and serve it as a JPEG image
	imgBase64 := doc["image"].(string)
	img, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(img)
	data.Init("test", "users")
}

func ImageGpu(w http.ResponseWriter, r *http.Request) {

	data.Init("test", "imagegpu")
	var doc bson.M

	err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the base64 string and serve it as a JPEG image
	imgBase64 := doc["image"].(string)
	img, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(img)
	data.Init("test", "users")
}

func ImageCooling(w http.ResponseWriter, r *http.Request) {

	data.Init("test", "imagecooling")
	var doc bson.M

	err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the base64 string and serve it as a JPEG image
	imgBase64 := doc["image"].(string)
	img, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(img)
	data.Init("test", "users")
}

func ImageCpu(w http.ResponseWriter, r *http.Request) {

	data.Init("test", "images")
	var doc bson.M

	err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the base64 string and serve it as a JPEG image
	imgBase64 := doc["image"].(string)
	img, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(img)
	data.Init("test", "users")
}

func ImagePowersupplies(w http.ResponseWriter, r *http.Request) {

	data.Init("test", "imagepower")
	var doc bson.M

	err := data.Collection.FindOne(nil, bson.M{}).Decode(&doc)
	if err != nil {
		log.Fatal(err)
	}

	// Decode the base64 string and serve it as a JPEG image
	imgBase64 := doc["image"].(string)
	img, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "image/jpeg")
	w.Write(img)
	data.Init("test", "users")

}

/*Full List of Gpu*/
func ListGpu(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListGpu.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "gpu")
	var items []data.Gpu

	if r.Method == http.MethodPost {

		name := r.FormValue("ssdbtn")
		filter := bson.M{"image": name}
		number = 4

		send(name, number)

		cur, err := data.Collection.Find(context.TODO(), filter)
		if err != nil {
			logger.Infof("error:", err)
		}
		defer cur.Close(context.TODO())

		for cur.Next(context.TODO()) {
			var item data.Gpu

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

/*Full List of PowerSupply*/
func ListPowerSupply(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListPowerSupply.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "powersupplies")
	var items []data.PowerSupply

	if r.Method == http.MethodPost {

		name := r.FormValue("ssdbtn")
		filter := bson.M{"image": name}
		number = 1

		send(name, number)
		cur, err := data.Collection.Find(context.TODO(), filter)
		if err != nil {
			logger.Infof("error:", err)
		}
		defer cur.Close(context.TODO())

		for cur.Next(context.TODO()) {
			var item data.PowerSupply

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

/*Full List of Motherboard*/
func ListMotherboard(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListMotherboard.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "motherboards")
	var items []data.Motherboard

	if r.Method == http.MethodPost {

		name := r.FormValue("ssdbtn")
		filter := bson.M{"image": name}
		number = 6

		send(name, number)

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

/*Full List of Ram*/
func ListRam(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListRam.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "ram")
	var items []data.Ram

	if r.Method == http.MethodPost {

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("ssdbtn"))
		filter := bson.M{"_id": ObjID}

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
	}
}

/*Full List of Housing*/
func ListHousing(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListHousing.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "housing")
	var items []data.Housing

	if r.Method == http.MethodPost {

		name := r.FormValue("ssdbtn")
		filter := bson.M{"image": name}
		number = 5

		send(name, number)
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
	}
}

/*Full List of Hdd*/
func ListHdd(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListHdd.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "hdd")
	var items []data.Hdd

	if r.Method == http.MethodPost {

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("ssdbtn"))
		filter := bson.M{"_id": ObjID}

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
	}
}

/*Full List of Cooling*/
func ListCooling(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListCooling.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "cooling")
	var items []data.Cooling

	if r.Method == http.MethodPost {

		name := r.FormValue("ssdbtn")
		filter := bson.M{"image": name}
		number = 3

		send(name, number)

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
	}
}

/*Full List of Cpu*/
func ListCpu(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListCpu.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "cpu")
	var items []data.Cpu

	if r.Method == http.MethodPost {

		name := r.FormValue("ssdbtn")
		filter := bson.M{"image": name}
		number = 2

		send(name, number)
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
	}
}

/*Full List of Ssd*/
func ListSsd(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/fullListSsd.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "ssd")
	var items []data.Ssd

	if r.Method == http.MethodPost {

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("ssdbtn"))
		filter := bson.M{"_id": ObjID}

		cur, err := data.Collection.Find(context.TODO(), filter)
		if err != nil {
			logger.Infof("error:", err)
		}
		defer cur.Close(context.TODO())

		for cur.Next(context.TODO()) {
			var item data.Ssd

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

/*Function to list elements that we have*/
func List(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()

	if r.Method == http.MethodPost {
		value := r.FormValue("element")

		switch value {
		case "cpu":
			tmpl := template.Must(template.ParseFiles("html/listCpu.html"))
			tmpl.Execute(w, data.GetUser(w))
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
			tmpl := template.Must(template.ParseFiles("html/listCooling.html"))
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
			tmpl := template.Must(template.ParseFiles("html/listHdd.html"))
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
			tmpl := template.Must(template.ParseFiles("html/listHousing.html"))
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
			tmpl := template.Must(template.ParseFiles("html/listRam.html"))
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
			tmpl := template.Must(template.ParseFiles("html/listMotherboard.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "motherboards")
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

		case "ssd":
			tmpl := template.Must(template.ParseFiles("html/listSsd.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "ssd")
			var items []data.Ssd

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.Ssd

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

		case "powersupply":
			tmpl := template.Must(template.ParseFiles("html/listPowersupply.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "powersupplies")
			var items []data.PowerSupply

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.PowerSupply

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

		case "gpu":
			tmpl := template.Must(template.ParseFiles("html/listGpu.html"))
			tmpl.Execute(w, data.GetUser(w))
			data.Init("shop", "gpu")
			var items []data.Gpu

			filter := bson.M{}

			cur, err := data.Collection.Find(context.TODO(), filter)
			if err != nil {
				logger.Infof("error:", err)
			}
			defer cur.Close(context.TODO())

			for cur.Next(context.TODO()) {
				var item data.Gpu

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
