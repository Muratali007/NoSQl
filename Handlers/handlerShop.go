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
	"strconv"
)

var images string
var number int

var itemsCpu []data.Cpu
var itemCpu data.Cpu

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

func ComparisonCpuMb(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listMotherboard.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "motherboards")
	var items []data.Motherboard

	if r.Method == http.MethodPost {

		filter := bson.M{"socket": r.FormValue("mb")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonCpuRam(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listRam.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "ram")
	var items []data.Ram

	if r.Method == http.MethodPost {

		filter := bson.M{"type": r.FormValue("ram")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonCpuCooling(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listCooling.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "cooling")
	var items []data.Cooling

	if r.Method == http.MethodPost {

		tdp, _ := strconv.ParseInt(r.FormValue("cooling"), 10, 0)
		filter := bson.M{"tdp": bson.M{"$gte": tdp}}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonMbCpu(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listCpu.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "cpu")
	var items []data.Cpu

	if r.Method == http.MethodPost {

		filter := bson.M{"socket": r.FormValue("cpu")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonMbRam(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listRam.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "ram")
	var items []data.Ram

	if r.Method == http.MethodPost {

		filter := bson.M{"type": r.FormValue("ram")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonMbHousing(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listHousing.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "housing")
	var items []data.Housing

	if r.Method == http.MethodPost {

		filter := bson.M{"mb_form_factor": r.FormValue("housing")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonMbHdd(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listHdd.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "hdd")
	var items []data.Hdd

	if r.Method == http.MethodPost {

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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonMbSsd(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listSsd.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "ssd")
	var items []data.Ssd

	if r.Method == http.MethodPost {

		filter := bson.M{"interface": "SATA3"}
		pciE := r.FormValue("ssd")
		if pciE == "3" {
			filter = bson.M{"interface": "PCI-E 4x 3.0"}
		} else {
			filter = bson.M{"interface": "PCI-E 4x 4.0"}
		}

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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonRamCpu(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listCpu.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "cpu")
	var items []data.Cpu

	if r.Method == http.MethodPost {

		filter := bson.M{"ram.type": r.FormValue("cpu")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonRamMb(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listMotherboard.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "motherboards")
	var items []data.Motherboard

	if r.Method == http.MethodPost {

		filter := bson.M{"ram.type": r.FormValue("mb")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonSsdMb(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listMotherboard.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "motherboards")
	var items []data.Motherboard

	if r.Method == http.MethodPost {

		var keyTag string

		formFactor := r.FormValue("mb")
		if formFactor != "M2" {
			keyTag = "interfaces.M2"
		} else {
			keyTag = "interfaces.SATA3"
		}

		filter := bson.M{keyTag: bson.M{"$gt": 0}}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonSsdHousing(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listHousing.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "housing")
	var items []data.Housing

	if r.Method == http.MethodPost {

		filter := bson.M{}
		formFactor := r.FormValue("housing")
		if formFactor != "M2" {
			switch formFactor {
			case "2.5":
				filter = bson.M{"drive_bays." + "2_5": bson.M{"$gte": 0}}
			case "3.5":
				filter = bson.M{"drive_bays." + "3_5": bson.M{"$gte": 0}}
			}
		}

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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonHddMb(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listMotherboard.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "motherboards")
	var items []data.Motherboard

	if r.Method == http.MethodPost {
		filter := bson.M{"interfaces.SATA3": bson.M{"$gt": 0}}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonHddHousing(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listHousing.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "housing")
	var items []data.Housing

	if r.Method == http.MethodPost {

		filter := bson.M{}
		formFactor := r.FormValue("housing")
		if formFactor != "M2" {
			switch formFactor {
			case "2.5":
				filter = bson.M{"drive_bays." + "2_5": bson.M{"$gte": 0}}
			case "3.5":
				filter = bson.M{"drive_bays." + "3_5": bson.M{"$gte": 0}}
			}
		}

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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonCoolingCpu(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listCpu.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "cpu")
	var items []data.Cpu

	if r.Method == http.MethodPost {

		tdp, _ := strconv.ParseInt(r.FormValue("cpu"), 10, 0)
		filter := bson.M{"tdp": bson.M{"$lte": tdp}}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonCoolingHousing(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listHousing.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "housing")
	var items []data.Housing

	if r.Method == http.MethodPost {

		height, _ := strconv.ParseInt(r.FormValue("housing"), 10, 0)
		filter := bson.M{"size.0": bson.M{"$gte": height}}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonHousingMb(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listMotherboard.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "motherboards")
	var items []data.Motherboard

	if r.Method == http.MethodPost {
		filter := bson.M{"form_factor": r.FormValue("mb")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonHousingPower(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listPowersupply.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "powersupplies")
	var items []data.PowerSupply

	if r.Method == http.MethodPost {
		filter := bson.M{"form_factor": r.FormValue("power")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonHousingGpu(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listGpu.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "gpu")
	var items []data.Gpu

	if r.Method == http.MethodPost {
		slots, err := strconv.ParseFloat(r.FormValue("gpu"), 64)
		filter := bson.M{"slots": bson.M{"$lte": slots}}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonHousingCooling(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listCooling.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "cooling")
	var items []data.Cooling

	if r.Method == http.MethodPost {
		height, err := strconv.Atoi(r.FormValue("cooling"))
		filter := bson.M{"height": bson.M{"$lte": height}}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonPowerMb(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listMotherboard.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "motherboards")
	var items []data.Motherboard

	if r.Method == http.MethodPost {
		mbPower, err := strconv.Atoi(r.FormValue("mb"))
		filter := bson.M{"mb_power": mbPower}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonPowerGpu(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listGpu.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "gpu")
	var items []data.Gpu

	if r.Method == http.MethodPost {
		power, err := strconv.Atoi(r.FormValue("gpu"))
		filter := bson.M{"tdpR": bson.M{"$gte": power}}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func ComparisonPowerHousing(w http.ResponseWriter, r *http.Request) {

	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/listHousing.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "housing")
	var items []data.Housing

	if r.Method == http.MethodPost {
		filter := bson.M{"ps_form_factor": r.FormValue("housing")}
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
		fmt.Println("Found multiple items:", len(items))

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
}

func AddCpuForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/addCpu.html"))
	tmpl.Execute(w, data.GetUser(w))
}

func AddCpu(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	data.Init("shop", "cpu")

	if r.Method == http.MethodPost {
		cores, _ := strconv.Atoi(r.FormValue("cores"))
		threads, _ := strconv.Atoi(r.FormValue("threads"))
		base, _ := strconv.ParseFloat(r.FormValue("base"), 64)
		turbo, _ := strconv.ParseFloat(r.FormValue("turbo"), 64)
		channels, _ := strconv.Atoi(r.FormValue("channels"))
		ramMaxFr, _ := strconv.Atoi(r.FormValue("maxFr"))
		ramMaxCap, _ := strconv.Atoi(r.FormValue("maxCap"))
		tdp, _ := strconv.Atoi(r.FormValue("tdp"))
		pcie, _ := strconv.Atoi(r.FormValue("pcie"))
		maxTemp, _ := strconv.Atoi(r.FormValue("maxTemp"))

		ram := data.RamCpu{Channels: channels, Type: r.FormValue("type"), MaxFrequency: ramMaxFr, MaxCapacity: ramMaxCap}

		recordCpu := data.Cpu{ID: primitive.NewObjectID(), Manufacturer: r.FormValue("man"), Model: r.FormValue("model"),
			Cores: cores, Threads: threads, ClockFrequency: []float64{base, turbo}, Socket: r.FormValue("socket"), Ram: ram,
			Tdp: tdp, PciE: pcie, MaxTemperature: maxTemp}

		_, err := data.Collection.InsertOne(context.TODO(), recordCpu)
		if err != nil {
			logger.Infof("A bulk write error occurred: %v", err)
		} else {
			logger.Infof("CPU record with ID: %s was CREATED!", recordCpu.ID)
		}
		data.Init("test", "users")
		http.Redirect(w, r, "/shop", http.StatusSeeOther)
	}
}

func ModifyCpuForm(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	tmpl := template.Must(template.ParseFiles("html/modifyCpu.html"))
	tmpl.Execute(w, data.GetUser(w))

	data.Init("shop", "cpu")
	var items []data.Cpu

	if r.Method == http.MethodPost {

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("modify")[10:34])
		filter := bson.M{"_id": ObjID}

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

		err = tmpl.Execute(w, items)
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")

	}
}

func ModifyCpu(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	data.Init("shop", "cpu")

	if r.Method == http.MethodPost {
		cores, _ := strconv.Atoi(r.FormValue("cores"))
		threads, _ := strconv.Atoi(r.FormValue("threads"))
		base, _ := strconv.ParseFloat(r.FormValue("base"), 64)
		turbo, _ := strconv.ParseFloat(r.FormValue("turbo"), 64)
		channels, _ := strconv.Atoi(r.FormValue("channels"))
		ramMaxFr, _ := strconv.Atoi(r.FormValue("maxFr"))
		ramMaxCap, _ := strconv.Atoi(r.FormValue("maxCap"))
		tdp, _ := strconv.Atoi(r.FormValue("tdp"))
		pcie, _ := strconv.Atoi(r.FormValue("pcie"))
		maxTemp, _ := strconv.Atoi(r.FormValue("maxTemp"))

		ram := data.RamCpu{Channels: channels, Type: r.FormValue("type"), MaxFrequency: ramMaxFr, MaxCapacity: ramMaxCap}

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("modifyCpu")[10:34])
		filter := bson.M{"_id": ObjID}

		recordCpu := data.Cpu{ID: ObjID, Manufacturer: r.FormValue("man"), Model: r.FormValue("model"),
			Cores: cores, Threads: threads, ClockFrequency: []float64{base, turbo}, Socket: r.FormValue("socket"), Ram: ram,
			Tdp: tdp, PciE: pcie, MaxTemperature: maxTemp}

		update := bson.M{"$set": bson.M{
			"manufacturer":    recordCpu.Manufacturer,
			"model":           recordCpu.Model,
			"cores":           recordCpu.Cores,
			"threads":         recordCpu.Threads,
			"clock_frequency": recordCpu.ClockFrequency,
			"socket":          recordCpu.Socket,
			"ram":             recordCpu.Ram,
			"tdp":             recordCpu.Tdp,
			"pci-e":           recordCpu.PciE,
			"max_temperature": recordCpu.MaxTemperature,
		}}

		_, err = data.Collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			logger.Infof("A bulk write error occurred: %v", err)
		} else {
			logger.Infof("CPU record with ID: %s was UPDATED!", ObjID)
		}
		data.Init("test", "users")
		http.Redirect(w, r, "/shop", http.StatusSeeOther)
	}
}

func DeleteCpu(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	data.Init("shop", "cpu")

	if r.Method == http.MethodPost {
		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("deleteCpu")[10:34])
		filter := bson.M{"_id": ObjID}

		_, err = data.Collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			logger.Infof("A bulk write error occurred: %v", err)
		} else {
			logger.Infof("CPU record with ID: %v was DELETED!", ObjID)
		}
		data.Init("test", "users")
		http.Redirect(w, r, "/shop", http.StatusSeeOther)

	}
}

func AddCpuToCart(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	data.Init("shop", "cpu")
	var items []data.Cpu

	if r.Method == http.MethodPost {

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("cartCpu")[10:34])
		filter := bson.M{"_id": ObjID}

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

			itemCpu = item
			itemsCpu = append(itemsCpu, itemCpu)
			items = append(items, item)
		}
		if err := cur.Err(); err != nil {
			logger.Infof("error :", err)
		}
		if err != nil {
			logger.Infof("error :", err)
		}
		data.Init("test", "users")
	}
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
		fmt.Println("Found multiple items:", len(items))

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
		fmt.Println("Found multiple items:", len(items))

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
		fmt.Println("Found multiple items:", len(items))

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

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("ssdbtn")[10:34])
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
		fmt.Println("Found multiple items:", len(items))

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
		fmt.Println("Found multiple items:", len(items))

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

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("ssdbtn")[10:34])
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
		fmt.Println("Found multiple items:", len(items))

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
		fmt.Println("Found multiple items:", len(items))

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

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("ssdbtn")[10:34])
		filter := bson.M{"_id": ObjID}

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
		fmt.Println("Found multiple items:", len(items))

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

		ObjID, err := primitive.ObjectIDFromHex(r.FormValue("ssdbtn")[10:34])
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
		fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

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
			fmt.Println("Found multiple items:", len(items))

			err = tmpl.Execute(w, items)
			if err != nil {
				logger.Infof("error :", err)
			}
			data.Init("test", "users")
		}
	}
}
