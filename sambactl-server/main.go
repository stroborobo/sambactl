package main

import (
	"os"
	"sync"
	"html/template"
	"net/http"
	"net/http/fcgi"
	"log"
	"os/exec"
	"encoding/json"
	"bytes"
	"strings"
)

const (
	PS = string(os.PathSeparator)
	WWWDIR = "/usr/local/sambactl/"
	SRCURL = "https://github.com/Knorkebrot/sambactl"
)

var (
	m sync.Mutex = sync.Mutex{}
	fileHandler http.Handler
)

type TplData struct {
	SrcURL		string
	Errors		[]string
	Users		[]User
}

type User struct {
	Username	string
}

func main() {
	err := os.Chdir(WWWDIR)
	if err != nil {
		log.Fatalln("Chdir:", err)
	}
	err = os.Mkdir(WWWDIR + PS + "htdocs", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalln("Mkdir:", err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalln("Hostname:", err)
	}
	log.Println("Running in", WWWDIR, "- http://" + hostname + ":80/")

	fileHandler = http.FileServer(http.Dir(WWWDIR + PS + "htdocs"))
	http.HandleFunc("/", http.HandlerFunc(indexHandler))
	http.HandleFunc("/user/", http.HandlerFunc(handler))

	err = fcgi.Serve(nil, nil)
	/*err = http.ListenAndServe(":80", nil)*/
	if err != nil {
		log.Fatalln("Serve:", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if strings.HasSuffix(path, "/index.html") {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	if path != "/" {
		fileHandler.ServeHTTP(w, r)
		return
	}

	tpl, err := template.ParseFiles("htdocs/index.html")
	if err != nil {
		log.Println("template.ParseFiles:", err)
	}
	data := TplData{}
	data.SrcURL = SRCURL

	errors := make([]string, 0, 10)

	users, err := getUsers()
	if err != nil {
		errors = append(errors, "User konnten nicht geladen werden.")
	}
	data.Users = users

	// dummy
	//users = make([]User, 0, 1)
	//data.Users = append(users, User{"Bildbearbeitung"})

	data.Errors = errors
	tpl.Execute(w, data)
}

func handler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	passwd := r.FormValue("password")
	del := r.FormValue("del")
	found := false

	users, err := getUsers()
	if err != nil {
		http.Error(w, "User loading failed:" + err.Error(), http.StatusInternalServerError)
		return
	}
	for _, u := range users {
		if u.Username == username {
			found = true
			break
		}
	}
	if found {
		if len(del) != 0 {
			err = delUser(username)
		} else {
			err = changePassword(username, passwd)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err = addUser(username, passwd)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("\"ok\"\n"))
}

func getUsers() ([]User, error) {
	cmd := exec.Command("/usr/bin/sudo", "/usr/local/bin/sambactl-worker", "list")
	var out bytes.Buffer
	cmd.Stdout = &out
	m.Lock()
	err := cmd.Run()
	m.Unlock()
	if err != nil {
		return nil, err
	}
	var users []User
	err = json.Unmarshal(out.Bytes(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func changePassword(username, newpasswd string) error {
	cmd := exec.Command("/usr/bin/sudo", "/usr/local/bin/sambactl-worker", "passwd", username, newpasswd)
	var out bytes.Buffer
	cmd.Stdout = &out
	m.Lock()
	err := cmd.Run()
	m.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func addUser(username, passwd string) error {
	cmd := exec.Command("/usr/bin/sudo", "/usr/local/bin/sambactl-worker", "adduser", username, passwd)
	var out bytes.Buffer
	cmd.Stdout = &out
	m.Lock()
	err := cmd.Run()
	m.Unlock()
	if err != nil {
		return err
	}
	return nil
}

func delUser(username string) error {
	cmd := exec.Command("/usr/bin/sudo", "/usr/local/bin/sambactl-worker", "deluser", username)
	var out bytes.Buffer
	cmd.Stdout = &out
	m.Lock()
	err := cmd.Run()
	m.Unlock()
	if err != nil {
		return err
	}
	return nil
}
