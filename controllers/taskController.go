package controllers

import (
	"GoTask/config"
	"GoTask/models"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


var (
	id          int
	name        string
	description string
	employee    string
	deadline    time.Time
	status      string
	completed   int
	created_at  time.Time
	updated_at  time.Time
	btn_status	string
	view	    = template.Must(template.ParseFiles("./views/index.html"))
	database    = config.Database()
)

const DDMMYYYY = "02-01-2006"

func TaskList(w http.ResponseWriter, r *http.Request) {
	statement, err := database.Query(`SELECT * FROM tasks`)

	if err != nil {
		fmt.Println(err)
	}

	var tasks []models.TaskModel

	for statement.Next() {
		err = statement.Scan(&id, &name, &description, &employee, &deadline, &status, &completed, &created_at, &updated_at)

		if err != nil {
			fmt.Println(err)
		}

		// deadline, _ := time.Parse(DDMMYYYY, deadline)

		switch status {
		case "In Progress":
			btn_status = "btn-warning"
		case "Pending":
			btn_status = "btn-danger text-white"
		case "Completed":
			btn_status = "btn-success"
		default:
			btn_status = "btn-warning"	
		}
		// fmt.Println(status)

		task := models.TaskModel{
			Id:          id,
			Name:        name,
			Description: description,
			Employee:    employee,
			Deadline:    deadline,
			Status:      status,
			BtnStatus:	 btn_status,
			Completed:   completed,
			CreatedAt:   created_at,
			UpdatedAt:   updated_at,
		}

		tasks = append(tasks, task)
	}
	data := models.ViewModel{
		Tasks: tasks,
	}

	_ = view.Execute(w, data)

}

func TaskAdd(w http.ResponseWriter, r *http.Request) {

	taskName	:= r.FormValue("task_name")
	taskDesc	:= r.FormValue("task_desc")
	taskEmp		:= r.FormValue("task_emp")
	taskDate	:= r.FormValue("task_date")
	taskCreate	:= time.Now()
	taskUpdate	:= time.Now()

	_, err := database.Exec(`INSERT INTO tasks (name, description, employee, deadline, created_at, updated_at) VALUE (?,?,?,?,?,?)`, taskName,taskDesc,taskEmp,taskDate,taskCreate,taskUpdate)
	
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func TaskDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`DELETE FROM tasks WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func TaskUpdate(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]

	taskId		:= r.FormValue("edit_id")
	taskName	:= r.FormValue("edit_name")
	taskDesc	:= r.FormValue("edit_desc")
	taskEmp		:= r.FormValue("edit_emp")
	taskStat	:= r.FormValue("edit_stat")
	taskDate	:= r.FormValue("edit_date")

	if taskStat == "Completed" {
		completed = 1
	} else {
		completed = 0
	}

	_, err := database.Exec(`
	UPDATE tasks SET 
	name = ?, 
	description = ?, 
	employee = ?, 
	status = ?, 
	deadline = ?, 
	completed = ?, 
	updated_at = CURRENT_TIMESTAMP 
	WHERE id = ?
	`, taskName, taskDesc, taskEmp, taskStat, taskDate, completed,
	taskId)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func TaskComplete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	_, err := database.Exec(`UPDATE tasks SET status = "Completed", completed = 1, updated_at = CURRENT_TIMESTAMP WHERE id = ?`, id)

	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}