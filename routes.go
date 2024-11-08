func main() {
    http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            createStudent(w, r)
        case http.MethodGet:
            getStudents(w, r)
        }
    })
    http.HandleFunc("/students/", func(w http.ResponseWriter, r *http.Request) {
        id := r.URL.Path[len("/students/"):]

        switch r.Method {
        case http.MethodGet:
            if len(id) > 0 && id[len(id)-8:] == "/summary" {
                generateSummary(w, r)
            } else {
                getStudentByID(w, r)
            }
        case http.MethodPut:
            updateStudent(w, r)
        case http.MethodDelete:
            deleteStudent(w, r)
        }
    })

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
