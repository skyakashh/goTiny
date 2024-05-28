package mongohelpers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controllergo "github.com/skyakashh/UrlShortner/controller.go"
	"github.com/skyakashh/UrlShortner/hashing"
	"github.com/skyakashh/UrlShortner/models"
	"go.mongodb.org/mongo-driver/bson"
)

// crete and insert url

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Methods", "POST")
	var str models.OriginalUrl
	if err := json.NewDecoder(r.Body).Decode(&str); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s := createShortUrl(str)
	json.NewEncoder(w).Encode(s)
}

func createShortUrl(str models.OriginalUrl) models.OriginalUrl {
	var inserturl models.Url
	inserturl.LongUrl = str.Url
	inserturl.ShortUrl = hashing.Hash()
	inserted, err := controllergo.Collection.InsertOne(context.Background(), inserturl)
	if err != nil {
		log.Fatal(err)
	}
	str.Url = inserturl.ShortUrl
	fmt.Println("inserted url with id : ", inserted.InsertedID)
	return str

}

// fetch a url

func GetLongURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	shortURL := params["shortURL"]

	var urlMapping models.Url

	err := controllergo.Collection.FindOne(context.Background(), bson.M{"shorturl": shortURL}).Decode(&urlMapping)
	fmt.Println(urlMapping)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, urlMapping.LongUrl, http.StatusTemporaryRedirect)
}
