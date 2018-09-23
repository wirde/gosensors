package main
import (
        "log"
        "encoding/json"
        "net/http"
        "strconv"
)

func main() {
        log.Println("Starting server")
        http.HandleFunc("/", HandleRequest)

        err := http.ListenAndServe(":" + strconv.Itoa(10000), nil)
        if err != nil {
                log.Panicln("Server failed starting. Error: %s", err)
        }
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
        log.Println("Incoming Request:", r.Method)
        switch r.Method {
        case http.MethodGet:
                log.Println("Get")
                break
        case http.MethodPost:
                log.Println("Post")
                var sensorData SensorData
                err := json.NewDecoder(r.Body).Decode(sensorData)
                if err != nil {
                        log.Println("Got data: ")
                        log.Println(sensorData)
                } else {
                        log.Println("Failed: ")
                        log.Println(err)
                }
                break
        case http.MethodDelete:
                log.Println("Delete")
                break
        default:
                HandleError(&w, 405, "Method not allowed", "Method not allowed", nil)
                break

        }
}

func HandleError(w *http.ResponseWriter, code int, responseText string, logMessage string, err error) {
        errorMessage := ""
        writer := *w

        if err != nil {
                errorMessage = err.Error()
        }

        log.Println(logMessage, errorMessage)
        writer.WriteHeader(code)
        writer.Write([]byte(responseText))
}

type SensorData struct {
        id string
        value string
}