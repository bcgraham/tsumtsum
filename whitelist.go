package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type WhitelistCommand struct {
	Output string `short:"o" long:"output" default:"" description:"Name of file where whitelist will be saved."`
}

var whitelistCommand WhitelistCommand

func (x *WhitelistCommand) Execute(args []string) error {
	session := MustNewSession(common.User, common.Device, common.ReportingServer)
	count, err := session.CreateWhitelist(x.Output)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Recorded %v IDs to %v.\n", count, x.Output)
	return nil
}

func init() {
	if _, err := parser.AddCommand("whitelist",
		"Create a whitelist of current contacts.",
		"Creates a file with the IDs of your current contact list, usually to be used as a whitelist for later purging. If this file exists already, it will be overwritten.",
		&whitelistCommand); err != nil {
		log.Fatal(err)
	}
}

func (s *Session) CreateWhitelist(filename string) (int, error) {
	mids, err := s.client.GetAllContactIds()
	if err != nil {
		return 0, err
	}
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		return 0, err
	}
	enc := json.NewEncoder(f)
	IDs := make(map[string]string, len(mids))
	for _, mid := range mids {
		IDs[mid] = ""
	}
	err = enc.Encode(&IDs)
	if err != nil {
		return 0, err
	}
	return len(mids), nil
}
