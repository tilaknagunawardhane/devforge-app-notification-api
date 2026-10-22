package main
import (\n    "context"
    "database/sql"
    "encoding/json"
    _ "github.com/go-chi/chi/v5"
    _ "github.com/go-chi/httpfwd"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/notify", func(w http.ResponseWriter, r *http.Request) {
        db, err := sql.Open("postgres", "user:password@localhost/notification_api")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer db.Close()
        rows, err := db.Query("SELECT * FROM notifications")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()
        for rows.Next() {
            var notification struct{
                ID int
                Message string
            }
            err = rows.Scan(&notification.ID, &notification.Message)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            json.NewEncoder(w).Encode(notification)
        }
    })
    http.ListenAndServe("localhost:8080", nil)
}
