func deleteStudent(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/students/"):])
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    mu.Lock()
    _, exists := students[id]
    if exists {
        delete(students, id)
    }
    mu.Unlock()

    if !exists {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
