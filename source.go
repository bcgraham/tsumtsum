package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func (s *Session) SendReport(r Report) error {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	err := enc.Encode(r)
	u := *s.reportingServer
	u.Path = path.Join("tsum", s.username, "reports", string(r.Type))
	resp, err := http.Post(u.String(), "application/json", &buf)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("status code: %v\n", resp.StatusCode)
	}
	return nil
}

type Report struct {
	Submitter string
	UserID    string
	MID       string
	Type      ReportType
}

type ReportType string

const (
	invite ReportType = "invites"
	search ReportType = "searches"
)

func (s *Session) LoadStrangers(sourceraw string) error {
	var strangers map[string]string

	// Is it a URL?
	_, err := url.ParseRequestURI(sourceraw)
	if err == nil {
		fmt.Println("Parsed source as URL...")
		strangers, err = GetIDsURL(Resource(s.reportingServer, "strangers"))
		if err != nil {
			return err
		}
	} else {
		fmt.Println("Parsed source as file...")
		strangers, err = GetIDsFile(sourceraw)
		if err != nil {
			return err
		}
	}
	// remove people from this file we've already added
	alreadyAdded, err := GetIDsURL(Resource(s.reportingServer, "reports"))
	if err != nil {
		return nil
	}
	for id := range strangers {
		if _, ok := alreadyAdded[id]; ok {
			delete(strangers, id)
		}
		if _, ok := s.friends[id]; ok {
			delete(strangers, id)
		}
	}
	s.strangers = strangers
	return nil
}

func GetIDsURL(url *url.URL) (map[string]string, error) {
	ids := make(map[string]string)
	resp, err := http.Get(url.String())
	if err != nil {
		return ids, err
	}
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&ids)
	if err != nil {
		return make(map[string]string), err
	}
	return ids, nil
}

// GetIDsFile loads a text file of with a list of newline-
// separated IDs (either userIDs or MIDs).
func GetIDsFile(file string) (IDs map[string]string, err error) {
	IDs = make(map[string]string)

	f, err := os.Open(file)
	if err != nil {
		return IDs, err
	}

	scnr := bufio.NewScanner(f)
	for scnr.Scan() {
		IDs[scnr.Text()] = ""
	}

	if err := scnr.Err(); err != nil {
		return make(map[string]string), err
	}
	return IDs, nil
}

func Resource(rs *url.URL, elements ...string) *url.URL {
	u := *rs
	for _, el := range elements {
		u.Path = path.Join(u.Path, el)
	}
	return &u
}
