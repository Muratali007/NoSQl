package Handlers

import (
	"MongoDb/internal/data"
	"MongoDb/pkg/logging"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
	"time"
)

func Register(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" || name == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		var count int64
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		count, err := data.Collection.CountDocuments(ctx, bson.M{"email": email})
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
		_, err = data.Collection.InsertOne(ctx, bson.M{"email": email, "name": name, "password": hashedPassword})
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
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" || name == "" {
			http.Error(w, "Name, email, and password are required", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		result := data.Collection.FindOneAndUpdate(
			ctx,
			bson.M{"email": email},
			bson.M{
				"$set": bson.M{
					"name":     name,
					"password": hashedPassword,
				},
			},
			options.FindOneAndUpdate().SetReturnDocument(options.After),
		)
		if result.Err() != nil {
			http.Error(w, result.Err().Error(), http.StatusInternalServerError)
			return
		}

		var user data.User
		err = result.Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "session",
			Value:   user.Email,
			Expires: time.Now().Add(24 * time.Hour),
		})

		http.Redirect(w, r, "/home", http.StatusSeeOther)
		logger.Infof("%s updated profile", email)
		return
	}

	email, err := getEmailFromCookie(r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	var user data.User
	err = data.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("html/update.html"))
	tmpl.Execute(w, user)
}

func getEmailFromCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()

	if r.Method == http.MethodPost {
		email := r.FormValue("email")

		if email == "" {
			http.Error(w, "Email is required", http.StatusBadRequest)
			return
		}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		// Check if the user exists
		var result data.User
		err := data.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
		if err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		// Delete the user
		_, err = data.Collection.DeleteOne(ctx, bson.M{"email": email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		logger.Infof("%s deleted their account", email)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	tmpl := template.Must(template.ParseFiles("html/delete.html"))
	tmpl.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger()
	if r.Method == http.MethodPost {

		name := r.FormValue("name")
		data.Gmail = name
		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			http.Error(w, "Email and password are required", http.StatusBadRequest)
			return
		}

		var result data.User
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		err := data.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&result)
		if err != nil {
			http.Error(w, "Invalid email", http.StatusUnauthorized)
			return
		}

		err = data.Collection.FindOne(ctx, bson.M{"name": name}).Decode(&result)
		if err != nil {
			http.Error(w, "Invalid name", http.StatusUnauthorized)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
		if err != nil {
			http.Error(w, "Invalid password", http.StatusUnauthorized)
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
	logger.Infof("%s do log out", data.Gmail)
	data.Gmail = ""

}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/home.html"))
	tmpl.Execute(w, data.GetUser(w))
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
		count, err = data.Collection.CountDocuments(ctx, bson.M{"email": cookie.Value})
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
