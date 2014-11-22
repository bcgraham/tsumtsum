package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/bcgraham/tsumtsum/external/line"
)

type PurgeCommand struct {
	Whitelist string `short:"w" long:"whitelist" required:"yes" description:"List of MIDs (NOT user IDs) from your contact list that are your 'real' friends. Removes all other contacts."`
}

var purgeCommand PurgeCommand

func (x *PurgeCommand) Execute(args []string) error {
	session := MustNewSession(common.User, common.Device, common.ReportingServer)
	t := time.Now()

	deleted, err := session.PurgeContactsExceptWhitelist(x.Whitelist)
	fmt.Printf("\n\nElapsed time: %v. Averaged %.0f contacts per minute.\n", time.Since(t), avg(deleted, time.Since(t)))
	if err != nil {
		log.Fatalf("Error purging contacts: %v", err)
	}
	return nil
}

func init() {
	if _, err := parser.AddCommand("purge",
		"Purge contacts",
		"Purge command removes all contacts from your list except those provided in a whitelist.",
		&purgeCommand); err != nil {
		log.Fatal(err)
	}
}

func (s *Session) DeleteContactByMid(mid string) error {
	err := s.client.UpdateContactSetting(<-s.reqSeq, mid, line.ContactSetting_CONTACT_SETTING_DELETE, "true")
	if err != nil {
		err = s.Rebuild()
		if err != nil {
			return err
		}
		err = s.client.UpdateContactSetting(<-s.reqSeq, mid, line.ContactSetting_CONTACT_SETTING_DELETE, "true")
		if err != nil {
			s.logger.Printf("error deleting contact \"%v\", second attempt: %v\n", mid, err)
			return err
		}
	}
	return nil
}

func (s *Session) PurgeContactsExceptWhitelist(input string) (deleted int, err error) {
	whitelist, err := GetIDsFile(input)
	if err != nil {
		return deleted, err
	}
	for mid := range s.friends {
		if _, ok := whitelist[mid]; ok {
			continue
		}
		err = s.DeleteContactByMid(mid)
		if err != nil {
			s.logger.Printf("error deleting contact: %v\n", err)
		}
		deleted++
		purgeProgress(deleted, len(s.friends)-len(whitelist))
	}
	return deleted, nil
}

func purgeProgress(count, total int) {
	printProgress(prog{
		str:  "Deleted %" + strconv.Itoa(len(strconv.Itoa(total))) + "d of %" + strconv.Itoa(len(strconv.Itoa(total))) + "d (%.2f%%)",
		args: []interface{}{count, total, 100 * float64(count) / float64(total)},
	})
}
