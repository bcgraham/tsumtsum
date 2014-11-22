package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/bcgraham/tsumtsum/external/line"
)

type AddCommand struct {
	SourceInput string `short:"s" long:"source" default:"[UseReportingServer]" description:"source of user IDs to be added to contact list. Tries to parse as a URL; if this fails, will assume it is a file."`
	Limit       int    `short:"l" long:"limit" default:"500" description:"how many contacts to add before stopping (defaults to 500; anything more than 800 will probably result in a tempban)"`
}

var addCommand AddCommand

func (x *AddCommand) Execute(args []string) error {
	session := MustNewSession(common.User, common.Device, common.ReportingServer)

	err := session.LoadStrangers(x.SourceInput)
	if err != nil {
		log.Fatalf("Could not load strangers: %v", err)
	}

	t := time.Now()
	found, err := session.AddStrangers(x.Limit)
	fmt.Printf("\n\nElapsed time: %v. Averaged %.0f contacts per minute.\n", time.Since(t), avg(found, time.Since(t)))
	if err != nil {
		log.Print(err.Error())
		return err
	}
	return nil
}

func init() {
	if _, err := parser.AddCommand("add",
		"Add strangers",
		"The add command adds contacts from a source to your contact list.",
		&addCommand); err != nil {
		log.Fatal(err)
	}
}

func (s *Session) AddContact(id string) (mid string, err error) {

	AddContactFunc := (*line.TalkServiceClient).FindAndAddContactsByUserid
	if midMatcher.MatchString(id) {
		AddContactFunc = (*line.TalkServiceClient).FindAndAddContactsByMid
	}
	r, err := AddContactFunc(s.client, <-s.reqSeq, id)
	if err != nil {
		// http response code 400 basically means
		// reconnect. not sure what causes it.
		if err.Error() == "HTTP Response code: 400" {
			s.logger.Printf("Got error \"%v\"\n.", err.Error())
			err = s.Rebuild()
			if err != nil {
				return mid, err
			}
			r, err = AddContactFunc(s.client, <-s.reqSeq, id)
		}
		if err != nil {
			errType := reflect.TypeOf(err)
			// if it's a connection issue, bump it up the stack
			if errType.String() == "*thrift.tTransportException" {
				return mid, err
			}
			// if it's an error from the line server, it's a
			// possible stop condition
			if errType.String() == "*line.TalkException" {
				switch err.(*line.TalkException).GetCode() {
				case line.ErrorCode_INVALID_STATE:
					msg := "\nCan't continue adding contacts. Your contact list is probably full (5000 contacts). Sleeping for ten minutes, then will resume.\n"
					s.logger.Print(msg)
					log.Print(msg)
					time.Sleep(10 * time.Minute)
				case line.ErrorCode_ABUSE_BLOCK:
					s.logger.Print("Your usage has been flagged as abuse and you can't presently add friends. This is almost certainly from trying to add too many friends. This is usually a temporary ban that lasts between 12 and 24 hours, but they last longer if you're a repeat offender.\n")
					panic(ErrAbuse(line.ErrorCode_ABUSE_BLOCK))
				}
			}
		}
	}
	var userID string
	if !isMid(id) {
		userID = id
		if contact, ok := r[id]; ok {
			mid = contact.GetMid()
			s.strangers[userID] = mid
		}
	} else {
		mid = id
	}
	err = s.SendReport(Report{
		Submitter: s.username,
		UserID:    userID,
		MID:       mid,
		Type:      search,
	})
	if err != nil {
		log.Printf("error sending search result: %v\n", err)
	}
	return mid, err
}

func (s *Session) AddStrangers(limit int) (found int, err error) {
	defer func() {
		if r := recover(); r != nil && reflect.TypeOf(r).String() == "ErrAbuse" {
			err = r.(ErrAbuse)
		}
	}()
	max := limit
	if len(s.strangers) < max {
		max = len(s.strangers)
	}

	var count int
	for id := range s.strangers {
		if count >= max {
			break
		}

		if !s.isNewID(id) {
			continue
		}

		mid, err := s.AddContact(id)
		if err != nil {
			s.logger.Printf("error adding contact: %v\n", err)
		}

		count++
		if mid != "" {
			found++
		}
		addProgress(count, max, found)
	}
	return found, nil
}

var midMatcher *regexp.Regexp

func init() {
	midMatcher = regexp.MustCompile("^u[a-fA-F0-9]{32}$")
}

func isMid(id string) bool {
	return midMatcher.MatchString(id)
}

func avg(found int, d time.Duration) float64 {
	return float64(found) * (float64(time.Minute) / float64(d))
}

type ErrAbuse line.ErrorCode

func (ErrAbuse) Error() string {
	return "Your usage has been flagged as abuse and you can't presently add friends. This is almost certainly from trying to add too many friends. This is usually a temporary ban that lasts between 12 and 24 hours, but they last longer if you're a repeat offender."
}

func addProgress(count, max, found int) {
	printProgress(prog{
		str: "%.2f%% completed. (%" + strconv.Itoa(len(strconv.Itoa(max))) + "d/%d) %" + strconv.Itoa(len(strconv.Itoa(max))) + "d of %" + strconv.Itoa(len(strconv.Itoa(max))) + "d found (%.2f%%).",
		args: []interface{}{
			100 * float64(count) / float64(max), count, max, found, count, 100 * float64(found) / float64(count),
		},
	})
}
