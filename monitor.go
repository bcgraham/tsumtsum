package main

import (
	"log"
	"strings"

	"github.com/bcgraham/tsumtsum/external/line"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type MonitorCommand struct{}

var monitorCommand MonitorCommand

func (x *MonitorCommand) Execute(args []string) error {
	session := MustNewSession(common.User, common.Device, common.ReportingServer)
	session.DeleteAfterInviting()
	return nil
}

func init() {
	if _, err := parser.AddCommand("monitor",
		"Monitor invites",
		"Monitors invites as they are sent and removes those friends from your contact list (presumably they are no longer needed).",
		&monitorCommand); err != nil {
		log.Fatal(err)
	}
}

func (s *Session) DeleteAfterInviting() {
	//  TODO: figure out why this doesn't work
	//	localRev, err := client.GetLastOpRevision()
	var localRev int64
	var invited, deleted int
	for {
		client, err := NewStandardClient("P4")
		if err != nil {
			log.Fatalf("Could not make standard client for delete after invite process: %v", err)
		}
		copyAuthToken(client, s.client)
		operations, err := client.FetchOperations(localRev, 10)
		if err != nil {
			s.logger.Printf("Couldn't get messages: %v\n", err)
		}
		for _, op := range operations {
			if isInviteFrom(op, s.profile.GetMid()) {
				invited++
				// send to central service
				err := s.SendReport(Report{
					Submitter: s.username,
					MID:       op.GetMessage().GetTo(),
					Type:      invite,
				})
				if err != nil {
					s.logger.Printf("C: Error sending invite record: %v\n", err)
				}
				// delete contact
				err = s.DeleteContactByMid(op.GetMessage().GetTo())
				if err != nil {
					s.logger.Printf("D: Error deleting contact: %v\n", err)
				} else {
					deleted++
				}
				monitorProgress(deleted, invited)
			}
			if op.GetRevision() > localRev {
				localRev = op.GetRevision()
			}
		}
	}
}

func monitorProgress(deleted, invited int) {
	printProgress(prog{
		str: "Witnessed %d invites; deleted %d invites (%.2f%%).",
		args: []interface{}{
			invited, deleted, 100 * float64(deleted) / float64(invited),
		},
	})
}

func isInviteFrom(op *line.Operation, userMID string) bool {
	if op.GetTypeA1() == line.OpType_SEND_MESSAGE {
		msg := op.GetMessage()
		from := msg.GetFrom()
		return strings.Contains(msg.GetText(), "LINE: Disney Tsum Tsum") && strings.Contains(msg.GetText(), "has invited you to play") && from == userMID
	}
	return false
}

func copyAuthToken(dst *line.TalkServiceClient, src *line.TalkServiceClient) {
	authToken := src.Transport.(*thrift.THttpClient).GetHeader("X-Line-Access")
	dst.Transport.(*thrift.THttpClient).SetHeader("X-Line-Access", authToken)
}
