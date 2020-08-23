
package main
import (
    "log"
    "encoding/json"
    "math/rand"
    "net/http"
    "github.com/gorilla/mux"
)
// Visitor Struct (Model)
type Visitor struct {
    Value int  `json:"value"`
    MinutesLastActive int   `json:"minuteslastactive`

}


//Init visitors var as a slice Visitor struct

var visitors []Visitor 
var result int
var sum int
//get all Visitor

func getVisitor(w http.ResponseWriter, r *http.Request) {


    w.Header().Set("Content-Type","application/json")
      //Loop through minuteslastactive

    for _, item := range visitors {          
        
        if item.MinutesLastActive < 60 {
            result += item.Value   
                     }     
       }
    
    json.NewEncoder(w).Encode(result)  
 
}

func createVisitor(w http.ResponseWriter, r *http.Request) {
   
    w.Header().Set("Content-Type","application/json")
    var visitor Visitor     
    _ = json.NewDecoder(r.Body).Decode(&visitor)
    visitor.Value = rand.Intn(100000)
    visitor.MinutesLastActive = rand.Intn(500)
    visitors = append(visitors, visitor)
    json.NewEncoder(w).Encode(&visitor)  
}

func main() {
    //Init Router
    r:= mux.NewRouter()


   //Mock Data
   visitors = append(visitors, Visitor{Value: 10, MinutesLastActive: 80})
   visitors = append(visitors, Visitor{Value: 20, MinutesLastActive: 90})
   visitors = append(visitors, Visitor{Value: 30, MinutesLastActive: 30})
   visitors = append(visitors, Visitor{Value: 40, MinutesLastActive: 20})

      


   //Route Handlers / Endpoint
    r.HandleFunc("/api/visitor", getVisitor).Methods("GET")
    r.HandleFunc("/api/visitor", createVisitor).Methods("POST")

    log.Fatal(http.ListenAndServe(":9000",r))
}