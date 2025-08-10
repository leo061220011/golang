package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
	"time"

	_ "modernc.org/sqlite"
)

type Agent struct {
	Id             int
	Name           string
	LastTimeOnline string
}

type Task struct {
	Id               int
	AgentName        string
	TaskCreationTime string
	Text             string
	Result           string
	ResultTime       string
}

type AgentNameRequest struct {
	AgentName string
	/*
		{
			"AgentName": "John Smith"
		}
	*/
}

func dbConnect() *sql.DB {
	db, err := sql.Open("sqlite", "agents.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func selectAllAgents(db *sql.DB) []Agent {
	rows, err := db.Query("SELECT * FROM agents")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	agents := []Agent{}
	for rows.Next() {
		a := Agent{}
		err := rows.Scan(&a.Id, &a.Name, &a.LastTimeOnline)
		if err != nil {
			fmt.Println(err)
			continue
		}
		agents = append(agents, a)
	}
	return agents
}

func selectAllTasks(db *sql.DB) []Task {
	rows, err := db.Query("SELECT * FROM tasks")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	tasks := []Task{}
	for rows.Next() {
		t := Task{}
		err := rows.Scan(&t.Id, &t.AgentName, &t.TaskCreationTime, &t.Text, &t.Result, &t.ResultTime)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tasks = append(tasks, t)
	}
	return tasks
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	defer db.Close()
	if r.Method == "GET" {
		agents := selectAllAgents(db)
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Println(err)
		}
		tmpl.Execute(w, agents)
	}
	if r.Method == "POST" {
		var Agent AgentNameRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&Agent)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Agent: %s online", Agent.AgentName)
		agents := selectAllAgents(db)
		for _, v := range agents {
			if v.Name == Agent.AgentName {
				_, err := db.Exec("UPDATE agents SET LastTimeOnline=? WHERE Name = ?", time.Now().Format("2006-Jan-02 15:04"), Agent.AgentName)
				if err != nil {
					fmt.Println(err)
				}
				return
			}
			_, err := db.Exec("INSERT INTO agents (Name, LastTimeOnline) VALUES (?,?)", Agent.AgentName, time.Now().Format("2006-Jan-02 15:04"))
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		db := dbConnect()
		defer db.Close()
		tasks := selectAllTasks(db)
		tmpl, err := template.ParseFiles("templates/tasks.html")
		if err != nil {
			fmt.Println(err)
		}
		tmpl.Execute(w, tasks)
	}
}
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	defer db.Close()
	if r.Method == "GET" {
		agents := selectAllAgents(db)
		tmpl, err := template.ParseFiles("templates/createTask.html")
		if err != nil {
			fmt.Println(err)
		}
		tmpl.Execute(w, agents)
	}
	if r.Method == "POST" {
		var newTask Task
		newTask.AgentName = r.FormValue("TaskAgent")
		newTask.Text = r.FormValue("TaskText")
		newTask.TaskCreationTime = time.Now().Format("2006-Jan-02 15:04")
		_, err := db.Exec("INSERT INTO tasks (AgentName, TaskCreationTime, Text, Result, ResultCreationTime) VALUES (?,?,?,?,?)", newTask.AgentName, newTask.TaskCreationTime, newTask.Text, " ", " ")
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/tasks", http.StatusFound)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)                 //просмотр таблицы агентов, создание нового агента
	http.HandleFunc("/tasks", tasksHandler)            //просмотр таблицы задач, обновление задачи
	http.HandleFunc("/create_task", createTaskHandler) //просмотр формы создания задачи, создание новой задачи
	fmt.Println("server is starting...")
	http.ListenAndServe(":8080", nil)
}
