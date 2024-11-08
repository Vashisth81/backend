func getStudents(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    var list []Student
    for _, student := range students {
        list = append(list, student)
    }
    json.NewEncoder(w).Encode(list)
}
