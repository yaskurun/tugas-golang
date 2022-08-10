package handler

import (
		"net/http"
		"html/template"	
		"log"	
		"path"	
		"database/sql"
		)
import _ "github.com/go-sql-driver/mysql"

func HandlerIndex(w http.ResponseWriter, r *http.Request){
	
	tmpl, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views","layout.html"))
	if(err != nil){
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, tampil("Berhasil Tampil"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}
}

func InsertTask(w http.ResponseWriter, r *http.Request){
	
	tmpl, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views","layout.html"))
	if(err != nil){
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}

	var taskInput = r.FormValue("task")
	var assigneeInput = r.FormValue("assignee")
	var deadlineInput = r.FormValue("deadline")

	var hasil = tambah(taskInput, assigneeInput,deadlineInput)

	err = tmpl.Execute(w, tampil(hasil.Pesan))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}
}

func EditTask(w http.ResponseWriter, r *http.Request){
	
	tmpl, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views","layout.html"))
	if(err != nil){
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}

	var idInput = r.FormValue("id")
	var taskInput = r.FormValue("task")
	var assigneeInput = r.FormValue("assignee")
	var deadlineInput = r.FormValue("deadline")

	var hasil = ubah(idInput, taskInput, assigneeInput, deadlineInput)

	err = tmpl.Execute(w, tampil(hasil.Pesan))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}
}

func MarkDone(w http.ResponseWriter, r *http.Request){
	
	tmpl, err := template.ParseFiles(path.Join("views","index.html"), path.Join("views","layout.html"))
	if(err != nil){
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}

	id := r.URL.Query()["id"]

	var hasil = setDone(id[0])

	err = tmpl.Execute(w, tampil(hasil.Pesan))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}
}

func CreatePage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles(path.Join("views","form.html"), path.Join("views","layout.html"))
	if(err != nil){
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":"Insert New Task",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}
}

func EditPage(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles(path.Join("views","edit.html"), path.Join("views","layout.html"))
	if(err != nil){
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}

	id := r.URL.Query()["id"]

	err = tmpl.Execute(w, getTask(id[0]))
	if err != nil {
		log.Println(err)
		http.Error(w, "Page Not Found", http.StatusInternalServerError)
		return
	}
}

func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/tugas_golang_task")
    if err != nil {
        return nil, err
    }

    return db, nil
}

type task struct {
	    Id    int
	    Task  string
	    Assignee  string
	    Deadline  string
	    Status   bool
}

type response struct {
	Status bool
	Pesan string
	Data []task
}


func tampil(message string) response{
	db, err := connect()

	if err != nil {
		return response{
			Status:false,
			Pesan:"Gagal Koneksi database: "+err.Error(),
			Data:[]task{},
		}
	}

	defer db.Close()
	dataTask, err := db.Query("select * from task order by id desc")

	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Query: "+err.Error(),
			Data:[]task{},
		}
	}

	defer dataTask.Close()
	var hasil []task
	for dataTask.Next(){
		var taskList = task{}
		var err = dataTask.Scan(&taskList.Id, &taskList.Task, &taskList.Assignee, &taskList.Deadline, &taskList.Status)
		if err != nil {
			return response{
				Status:false,
				Pesan:"Gagal baca:"+err.Error(),
				Data:[]task{},
			}
		}
		hasil = append(hasil,taskList)
	}

	if err != nil{
		return response{
			Status:false,
			Pesan:"Kesalahan: "+err.Error(),
			Data:[]task{},
		}
	}

	return response{
		Status:true,
		Pesan:message,
		Data:hasil,
	}


}

func getTask(id string) response{
	db, err := connect()

	if err != nil {
		return response{
			Status:false,
			Pesan:"Gagal Koneksi database: "+err.Error(),
			Data:[]task{},
		}
	}

	defer db.Close()
	dataTask, err := db.Query("select * from task where id=?",id)

	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Query: "+err.Error(),
			Data:[]task{},
		}
	}

	defer dataTask.Close()
	var hasil []task
	for dataTask.Next(){
		var taskList = task{}
		var err = dataTask.Scan(&taskList.Id, &taskList.Task, &taskList.Assignee, &taskList.Deadline, &taskList.Status)
		if err != nil {
			return response{
				Status:false,
				Pesan:"Gagal baca:"+err.Error(),
				Data:[]task{},
			}
		}
	hasil = append(hasil,taskList)
	}


	if err != nil{
		return response{
			Status:false,
			Pesan:"Kesalahan: "+err.Error(),
			Data:[]task{},
		}
	}

	return response{
		Status:true,
		Pesan:"Berhasil tampil",
		Data:hasil,
	}

}


func tambah(taskInput string, assignee string, deadline string) response{
	db, err := connect()

	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Koneksi:"+err.Error(),
			Data:[]task{},
		}
	}

	defer db.Close()
	_, err = db.Exec("insert into task values(?, ?, ?, ?, ?)",0,taskInput, assignee, deadline,true)
	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Query Insert:"+err.Error(),
			Data:[]task{},
		}
	}

	return response{
		Status:true,
		Pesan:"Berhasil tambah",
		Data:[]task{},
	}
}

func ubah(id string, taskInput string, assignee string, deadline string) response{
	db, err := connect()

	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Koneksi:"+err.Error(),
			Data:[]task{},
		}
	}

	defer db.Close()
	_, err = db.Exec("update task set task=?, assignee=?, deadline=? where id=?",taskInput,assignee,deadline,id)
	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Query Update:"+err.Error(),
			Data:[]task{},
		}
	}

	return response{
		Status:true,
		Pesan:"Berhasil update",
		Data:[]task{},
	}
}


func setDone(id string) response{
	db, err := connect()

	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Koneksi:"+err.Error(),
			Data:[]task{},
		}
	}

	defer db.Close()
	_, err = db.Exec("update task set status='false'  where id=?",id)
	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Query Update:"+err.Error(),
			Data:[]task{},
		}
	}

	return response{
		Status:true,
		Pesan:"Berhasil update",
		Data:[]task{},
	}
}

func hapus(id string) response{
	db, err := connect()

	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Koneksi:"+err.Error(),
			Data:[]task{},
		}
	}

	defer db.Close()
	_, err = db.Exec("delete from task swhere id=?")
	if err != nil{
		return response{
			Status:false,
			Pesan:"Gagal Query Update:"+err.Error(),
			Data:[]task{},
		}
	}

	return response{
		Status:true,
		Pesan:"Berhasil hapus",
		Data:[]task{},
	}
}

