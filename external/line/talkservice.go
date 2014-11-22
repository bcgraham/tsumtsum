package line

import (
	"bytes"
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type TalkService interface {
	FetchOperations(localRev int64, count int32) (r []*Operation, err error)
	// Parameters:
	//  - ReqSeq
	//  - Emails
	FindAndAddContactsByMid(reqSeq int32, mid string) (r map[string]*Contact, err error)
	// Parameters:
	//  - ReqSeq
	//  - Phones
	FindAndAddContactsByUserid(reqSeq int32, userid string) (r map[string]*Contact, err error)
	// Parameters:
	//  - Userid
	GetAllContactIds() (r []string, err error)
	// Parameters:
	//  - KeepLoggedIn
	//  - SystemName
	GetContacts(ids []string) (r []*Contact, err error)
	GetLastOpRevision() (r int64, err error)
	// Parameters:
	//  - IdentityProvider
	//  - Identifier
	//  - Password
	//  - KeepLoggedIn
	//  - AccessLocation
	//  - SystemName
	//  - Certificate
	GetProfile() (r *Profile, err error)
	LoginWithIdentityCredentialForCertificate(identityProvider IdentityProvider, identifier string, password string, keepLoggedIn bool, accessLocation string, systemName string, certificate string) (r *LoginResult_, err error)
	// Parameters:
	//  - Verifier
	LoginWithVerifierForCertificate(verifier string) (r *LoginResult_, err error)
	UpdateContactSetting(reqSeq int32, mid string, flag ContactSetting, value string) (err error)
}

type TalkServiceClient struct {
	Transport       thrift.TTransport
	ProtocolFactory thrift.TProtocolFactory
	InputProtocol   thrift.TProtocol
	OutputProtocol  thrift.TProtocol
	SeqId           int32
}

func NewTalkServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *TalkServiceClient {
	return &TalkServiceClient{Transport: t,
		ProtocolFactory: f,
		InputProtocol:   f.GetProtocol(t),
		OutputProtocol:  f.GetProtocol(t),
		SeqId:           0,
	}
}
