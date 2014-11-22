package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/bcgraham/tsumtsum/external/line"
	"github.com/bgentry/speakeasy"

	"git.apache.org/thrift.git/lib/go/thrift"
)

type Session struct {
	client          *line.TalkServiceClient
	username        string
	password        string
	device          string
	friends         map[string]struct{}
	strangers       map[string]string
	profile         *line.Profile
	reportingServer *url.URL
	reqSeq          chan int32
	logger          *log.Logger
}

// NewSession returns a session that is authenticated
// and ready to use.
func NewSession(u, p, device, reportingServer string) (*Session, error) {
	rs, err := url.Parse(reportingServer)
	if err != nil {
		return &Session{}, err
	}
	os.Mkdir("logs", 0777)
	f, err := os.Create(filepath.Join("logs", "session"+time.Now().Format("20060102_150405")+".log"))
	if err != nil {
		log.Fatalf("could not create log file: %v", err)
	}
	logger := log.New(f, "", log.LstdFlags)

	rs.Path = u
	s := &Session{
		username:        u,
		password:        p,
		device:          device,
		reportingServer: rs,
		logger:          logger,
	}
	s.reqSeq = reqSeqGen()
	s.client, err = NewStandardClient("/S4")
	s.friends = make(map[string]struct{}, 0)
	s.strangers = make(map[string]string, 0)
	err = s.Login()
	if err != nil {
		return &Session{}, err
	}
	profile, err := s.client.GetProfile()
	if err != nil {
		return &Session{}, err
	}
	s.profile = profile
	err = s.LoadAllContacts()
	if err != nil {
		return &Session{}, err
	}

	return s, err
}

func (s *Session) Login() error {
	client, err := NewStandardClient("/api/v4/TalkService.do")
	if err != nil {
		fmt.Println("error creating standard client...")
		return err
	}
	defer client.Transport.Close()
	lr, err := client.LoginWithIdentityCredentialForCertificate(line.IdentityProvider(1), s.username, s.password, true, "127.0.0.1", s.device, "")
	if err != nil {
		return err
	}

	// experimental and untested
	if lr.GetTypeA1() == line.LoginResultType_REQUIRE_DEVICE_CONFIRM {
		fmt.Println("You must confirm this device with LINE. The PIN code is: ", lr.GetPinCode())
		fmt.Println("Waiting for PIN confirmation...")
		verifier, err := waitForDeviceVerification(lr.GetVerifier())
		if err != nil {
			log.Fatalf("Error logging in: %v", err)
		}
		lr, err = client.LoginWithVerifierForCertificate(verifier)
	}
	s.client.Transport.(*thrift.THttpClient).SetHeader("X-Line-Access", lr.GetAuthToken())
	return err
}

func (s *Session) Rebuild() error {
	s.client.Transport.Close()
	var err error
	s.client, err = NewStandardClient("S4")
	if err != nil {
		return fmt.Errorf("can't make new client for rebuild: %v", err.Error())
	}
	err = s.Login()
	if err != nil {
		s.logger.Printf("Rebuild A. Error: %v\n", err)
		err = s.Login()
		if err != nil {
			fmt.Println("rebuild failure: ", err.Error(), "\n")
			return fmt.Errorf("can't login for rebuild: %v", err.Error())
		}
	}
	return nil
}

func NewStandardClient(path string) (*line.TalkServiceClient, error) {
	u, err := url.Parse(apiServer)
	if err != nil {
		return &line.TalkServiceClient{}, err
	}
	u.Path = path
	trans, err := thrift.NewTHttpPostClient(u.String())
	if err != nil {
		return &line.TalkServiceClient{}, err
	}

	trans.(*thrift.THttpClient).SetHeader("X-Line-Application", "DESKTOPWIN\t3.2.1.83\tWINDOWS\t5.1.2600-XP-x64")

	client := line.NewTalkServiceClientFactory(trans, thrift.NewTCompactProtocolFactory())
	return client, nil
}

func MustNewSession(u, device, reportingServer string) *Session {
	password, err := speakeasy.Ask("Please enter your LINE password: ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	session, err := NewSession(u, password, device, reportingServer)
	if err != nil {
		log.Fatalf("Could not create new session: %v", err)
	}
	return session
}

func (s *Session) LoadAllContacts() error {
	mids, err := s.client.GetAllContactIds()
	if err != nil {
		return err
	}
	for _, mid := range mids {
		s.friends[mid] = struct{}{}
	}
	return nil
}

func waitForDeviceVerification(token string) (string, error) {
	req, err := http.NewRequest("GET", apiServer+"/Q", nil)
	if err != nil {
		log.Fatalf("Problem making verification request. Cannot verify. \n%v", err)
	}
	req.Header.Add("X-Line-Access", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error with response from verification request: %v\n", err)
	}
	jsonresp := make(map[string]interface{})
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&jsonresp)
	if err != nil {
		log.Fatalf("Problem decoding json response during verification: %v", err)
	}
	result, ok := jsonresp["result"]
	if !ok {
		if errM, ok := jsonresp["errorMessage"]; !ok {
			log.Fatalf("Unrecognized JSON response during verification: %v", jsonresp)
		} else {
			log.Fatalf("Error during verification: %v", errM)
		}
	}
	resMap := result.(map[string]string)
	_, ok = resMap["authPhase"]
	if !ok {
		log.Printf("expected authPhase; respMap: %v\n", resMap)
	}
	verifier, ok := resMap["verifier"]
	if !ok {
		log.Printf("expected verifier; respMap: %v\n", resMap)
	}
	trans, err := thrift.NewTHttpPostClient(apiServer + "/api/v4/TalkService.do")
	if err != nil {
		return "", err
	}
	defer trans.Close()

	trans.(*thrift.THttpClient).SetHeader("X-Line-Application", "DESKTOPWIN\t3.2.1.83\tWINDOWS\t5.1.2600-XP-x64")
	trans.(*thrift.THttpClient).SetHeader("X-Line-Access", verifier)

	if err = trans.Open(); err != nil {
		return "", err
	}
	client := line.NewTalkServiceClientFactory(trans, thrift.NewTCompactProtocolFactory())
	lr, err := client.LoginWithVerifierForCertificate(verifier)
	if err != nil {
		return "", err
	}
	return lr.GetAuthToken(), nil
}

func (s *Session) isNewID(id string) bool {
	// if it's a mid, check the list of mids we've submitted.
	// if it's a userID, check the list of userIDs we've submitted.
	_, ok := s.friends[id]
	return !ok
}

func reqSeqGen() chan int32 {
	ch := make(chan int32, 100)
	go func() {
		var i int32
		for {
			ch <- i
			i++
		}
	}()
	return ch
}
