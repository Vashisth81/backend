func createStudent(w http.ResponseWriter, r *http.Request) {
    var student Student
    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    mu.Lock()
    student.ID = nextID
    nextID++
    students[student.ID] = student
    mu.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(student)
}
