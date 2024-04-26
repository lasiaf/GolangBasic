package main
import (
    "fmt"
    "html/template"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        var data = map[string]string{
            "Name":    "M. Faisal",
            "Message": "Faisal Golang Project",
        }
        var t, err = template.ParseFiles("index.html")
        if err != nil {
            fmt.Println(err.Error())
            return
        }
        t.Execute(w, data)
    })
    fmt.Println("starting web server at    http://localhost:1927/")
    http.ListenAndServe(":1927", nil)
}