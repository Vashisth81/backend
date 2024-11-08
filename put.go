func updateStudent(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.URL.Path[len("/students/"):])
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    var updatedStudent Student
    if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    mu.Lock()
    student, exists := students[id]
    if exists {
        updatedStudent.ID = student.ID
        students[id] = updatedStudent
    }
    mu.Unlock()

    if !exists {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(updatedStudent)
}
