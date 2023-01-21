package Handlers

import (
	"MongoDb/pkg/logging"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var client *mongo.Client
var users *mongo.Collection

var gmail string

type User struct {
	Email    string
	Password string
}

func init() {
	var err error
	logger := logging.GetLogger()
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	users = client.Database("test").Collection("users")
	logger.Info("Connected to a database")
}

func Register(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		var count int64
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		count, err := users.CountDocuments(ctx, bson.M{"email": email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if count > 0 {
			http.Error(w, "Email already in use", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
		_, err = users.InsertOne(ctx, bson.M{"email": email, "password": hashedPassword})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		logger.Infof("Create a new user %s", email)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("html/register.html"))
	tmpl.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		gmail = email
		password := r.FormValue("password")

		if email == "" || password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		var result User
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := users.FindOne(ctx, bson.M{"email": email}).Decode(&result)
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
		if err != nil {
			http.Error(w, "Invalid email or password", http.StatusUnauthorized)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session",
			Value:   email,
			Expires: time.Now().Add(24 * time.Hour),
		})

		http.Redirect(w, r, "/home", http.StatusSeeOther)
		logger.Infof("%s log in", email)
		return
	}

	tmpl := template.Must(template.ParseFiles("html/login.html"))
	tmpl.Execute(w, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   "",
		Expires: time.Unix(0, 0),
	})

	http.Redirect(w, r, "/login", http.StatusSeeOther)
	logger.Infof("%s do log out", gmail)
	gmail = ""
}
func getUser(w http.ResponseWriter) User {
	email := gmail
	var result User
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	err := users.FindOne(ctx, bson.M{"email": email}).Decode(&result)
	if err != nil {
		http.Error(w, "Error retrieving data from MongoDB", http.StatusInternalServerError)
		return result
	}
	return result
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/home.html"))
	tmpl.Execute(w, getUser(w))
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("session")
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		var count int64
		count, err = users.CountDocuments(ctx, bson.M{"email": cookie.Value})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if count == 0 {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func Shop(w http.ResponseWriter, r *http.Request) {

}
