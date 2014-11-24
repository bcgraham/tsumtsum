package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"

	"github.com/bgentry/speakeasy"

	"github.com/julienschmidt/httprouter"

	_ "github.com/lib/pq"
)

func post(res resource) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		dec := json.NewDecoder(r.Body.(io.Reader))
		var rep report
		err := dec.Decode(&rep)
		if err != nil {
			log.Printf("error decoding json from request: %v\n", err)
			http.Error(w, "Error. Sorry. Contact bcgraham@gmail.com. Error code A.", http.StatusBadRequest)
			return
		}

		rep.RemoteAddr = r.RemoteAddr
		if rep.Submitter != ps[0].Value {
			log.Printf("SearchReport and URL disagree about identity of submitter: SR says \"%v\", URL says \"%v\".\n", rep.Submitter, ps[0].Value)
		}
		go func(r report) {
			_, err := res.stmt.Exec(res.args(r)...)
			if err != nil {
				log.Printf("error writing to db: %v\n", err)
				return
			}
		}(rep)
	}
}

func get(stmt *sql.Stmt) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		users, err := getUserIDs(stmt, ps[0].Value)
		if err != nil {
			http.Error(w, "Error. Sorry. Contact bcgraham@gmail.com. Error code A.", http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		err = enc.Encode(users)
		if err != nil {
			log.Printf("error encoding json and sending response: %v\n", err)
		}
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	port := flag.String("port", ":8080", "port on which to run server")
	dbuser := flag.String("dbuser", "", "user for db connection")
	dbhost := flag.String("dbhost", "localhost:5432", "host and port of database server, formatted as 'host:port'")
	flag.Parse()
	dbpass, err := speakeasy.Ask("Enter password for database user '" + *dbuser + "':")
	if err != nil {
		log.Fatal(err)
	}
	credentials := mustParse(*dbuser, dbpass)
	db, err := sql.Open("postgres", "postgres://"+credentials+"@"+*dbhost+"/tsum?sslmode=disable")
	//	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatalf("Could not open db: %v", err)
	}

	strangers := mustPrepare(db, strangersSQL)
	reports := mustPrepare(db, reportsSQL)

	searches := resource{
		stmt: mustPrepare(db, searchSQL),
		args: report.searchArgs,
	}

	invites := resource{
		stmt: mustPrepare(db, inviteSQL),
		args: report.inviteArgs,
	}

	router := httprouter.New()

	router.GET("/tsum/:user/strangers/", get(strangers))
	router.GET("/tsum/:user/reports", get(reports))
	router.POST("/tsum/:user/reports/searches", post(searches))
	router.POST("/tsum/:user/reports/invites", post(invites))
	err = http.ListenAndServe(*port, router)
	if err != nil {
		fmt.Println(err)
	}
}

type report struct {
	Submitter  string
	UserID     string
	MID        string
	RemoteAddr string
}

type resource struct {
	stmt *sql.Stmt
	args func(report) []interface{}
}

func (r report) searchArgs() []interface{} {
	return []interface{}{
		r.Submitter,
		r.UserID,
		r.MID,
		r.RemoteAddr,
	}
}

func (r report) inviteArgs() []interface{} {
	return []interface{}{
		r.Submitter,
		r.MID,
		r.RemoteAddr,
	}
}

func getUserIDs(stmt *sql.Stmt, submitter string) (ids map[string]string, err error) {
	rows, err := stmt.Query(submitter)
	if err != nil {
		log.Printf("query error with submitter %v: %v\n", submitter, err)
		return make(map[string]string), err
	}
	defer rows.Close()

	ids = make(map[string]string)
	for rows.Next() {
		var id string
		err = rows.Scan(&id)
		if err != nil {
			log.Printf("error scanning rows: %v\n", err)
		}
		ids[id] = ""
	}
	return ids, nil
}

const strangersSQL = `SELECT id FROM validIDsCache WHERE id NOT IN (SELECT id FROM validIDsSubmittersCache WHERE submitter = $1) LIMIT 1000;`
const reportsSQL = `SELECT id FROM allIDsSubmittersCache WHERE submitter = $1`
const searchSQL = `INSERT INTO reports (submitter, userid, mid, ip) VALUES ($1, $2, $3, $4)`
const inviteSQL = `INSERT INTO invites (submitter, mid, ip) VALUES ($1, $2, $3)`

func mustPrepare(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	return stmt
}

func mustParse(user, pass string) string {
	if user == "" {
		log.Fatal("Username for DB access cannot be blank.")
	}
	if pass == "" {
		log.Fatal("Password for DB access cannot be blank.")
	}
	if strings.Contains(user, ":") {
		log.Fatal("Username cannot contain colon.")
	}
	repl := strings.NewReplacer("\\", "\\\\", "'", "\\'")
	escPass := repl.Replace(pass)
	return user + ":" + escPass
}
