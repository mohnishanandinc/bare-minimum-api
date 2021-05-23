package routers

import (
	"encoding/json"
	"github.com/bare-minimum-api/apis"
	"log"
	"net/http"
	"time"
)

const JSON = "application/json"

func ReturnTimeStamp(w http.ResponseWriter, request *http.Request) {
	log.Println("call made to Timestamp API")
	returnTimestamp := apis.Timestamp{
		Message: "A basic time API",
		Timestamp: time.Now().Unix(),
	}
	w.Header().Set("Content-Type", JSON)
	err := json.NewEncoder(w).Encode(returnTimestamp)
	if err != nil{
		log.Println("error serving the timestamp apis ",err)
	}
}