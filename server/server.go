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

	"github.com/julienschmidt/httprouter"

	_ "github.com/mattn/go-sqlite3"
)

func post(res resource) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		dec := json.NewDecoder(r.Body.(io.Reader))
		var report report
		err := dec.Decode(&report)
		if err != nil {
			log.Printf("error decoding json from request: %v\n", err)
			http.Error(w, "Error. Sorry. Contact bcgraham@gmail.com. Error code A.", http.StatusBadRequest)
			return
		}

		report.RemoteAddr = r.RemoteAddr
		if report.Submitter != ps[0].Value {
			log.Printf("SearchReport and URL disagree about identity of submitter: SR says \"%v\", URL says \"%v\".\n", report.Submitter, ps[0].Value)
		}
		go func() {
			result, err := res.stmt.Exec(res.args(report))
			if err != nil {
				log.Printf("error writing to db: %v\n", err)
				return
			}
			affected, err := result.RowsAffected()
			if err != nil {
				log.Printf("error getting row count: %v\n", err)
				return
			}
			log.Printf("affected %d rows\n", affected)
		}()
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
	flag.Parse()
	db, err := sql.Open("sqlite3", "users.db")
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

const strangersSQL = `SELECT id FROM validIDs WHERE id NOT IN (SELECT id FROM validIDsSubmitters WHERE submitter = ?);`
const reportsSQL = `SELECT id FROM allIDsSubmitters WHERE submitter = ?`
const searchSQL = `INSERT INTO reports (submitter, userid, mid, tstp, ip) VALUES (?, ?, ?, ?, ?)`
const inviteSQL = `INSERT INTO invites (submitter, mid, ip) VALUES (?, ?, ?)`

func mustPrepare(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	return stmt
}
