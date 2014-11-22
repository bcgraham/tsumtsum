package line

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func (p *TalkServiceClient) LoginWithIdentityCredentialForCertificate(identityProvider IdentityProvider, identifier string, password string, keepLoggedIn bool, accessLocation string, systemName string, certificate string) (r *LoginResult_, err error) {
	if err = p.sendLoginWithIdentityCredentialForCertificate(identityProvider, identifier, password, keepLoggedIn, accessLocation, systemName, certificate); err != nil {
		return
	}
	return p.recvLoginWithIdentityCredentialForCertificate()
}

func (p *TalkServiceClient) sendLoginWithIdentityCredentialForCertificate(identityProvider IdentityProvider, identifier string, password string, keepLoggedIn bool, accessLocation string, systemName string, certificate string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("loginWithIdentityCredentialForCertificate", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := LoginWithIdentityCredentialForCertificateArgs{
		IdentityProvider: identityProvider,
		Identifier:       identifier,
		Password:         password,
		KeepLoggedIn:     keepLoggedIn,
		AccessLocation:   accessLocation,
		SystemName:       systemName,
		Certificate:      certificate,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvLoginWithIdentityCredentialForCertificate() (value *LoginResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error1131 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1132 error
		error1132, err = error1131.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1132
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "loginWithIdentityCredentialForCertificate failed: out of sequence response")
		return
	}
	result := LoginWithIdentityCredentialForCertificateResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.E != nil {
		err = result.E
		return
	}
	value = result.GetSuccess()
	return
}

type LoginWithIdentityCredentialForCertificateArgs struct {
	// unused fields # 1 to 2
	Identifier       string           `thrift:"identifier,3" json:"identifier"`
	Password         string           `thrift:"password,4" json:"password"`
	KeepLoggedIn     bool             `thrift:"keepLoggedIn,5" json:"keepLoggedIn"`
	AccessLocation   string           `thrift:"accessLocation,6" json:"accessLocation"`
	SystemName       string           `thrift:"systemName,7" json:"systemName"`
	IdentityProvider IdentityProvider `thrift:"identityProvider,8" json:"identityProvider"`
	Certificate      string           `thrift:"certificate,9" json:"certificate"`
}

func NewLoginWithIdentityCredentialForCertificateArgs() *LoginWithIdentityCredentialForCertificateArgs {
	return &LoginWithIdentityCredentialForCertificateArgs{}
}

func (p *LoginWithIdentityCredentialForCertificateArgs) GetIdentityProvider() IdentityProvider {
	return p.IdentityProvider
}

func (p *LoginWithIdentityCredentialForCertificateArgs) GetIdentifier() string {
	return p.Identifier
}

func (p *LoginWithIdentityCredentialForCertificateArgs) GetPassword() string {
	return p.Password
}

func (p *LoginWithIdentityCredentialForCertificateArgs) GetKeepLoggedIn() bool {
	return p.KeepLoggedIn
}

func (p *LoginWithIdentityCredentialForCertificateArgs) GetAccessLocation() string {
	return p.AccessLocation
}

func (p *LoginWithIdentityCredentialForCertificateArgs) GetSystemName() string {
	return p.SystemName
}

func (p *LoginWithIdentityCredentialForCertificateArgs) GetCertificate() string {
	return p.Certificate
}
func (p *LoginWithIdentityCredentialForCertificateArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 8:
			if err := p.ReadField8(iprot); err != nil {
				return err
			}
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
				return err
			}
		case 5:
			if err := p.ReadField5(iprot); err != nil {
				return err
			}
		case 6:
			if err := p.ReadField6(iprot); err != nil {
				return err
			}
		case 7:
			if err := p.ReadField7(iprot); err != nil {
				return err
			}
		case 9:
			if err := p.ReadField9(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) ReadField8(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 8: %s", err)
	} else {
		temp := IdentityProvider(v)
		p.IdentityProvider = temp
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Identifier = v
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Password = v
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) ReadField5(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadBool(); err != nil {
		return fmt.Errorf("error reading field 5: %s", err)
	} else {
		p.KeepLoggedIn = v
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) ReadField6(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 6: %s", err)
	} else {
		p.AccessLocation = v
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) ReadField7(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 7: %s", err)
	} else {
		p.SystemName = v
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) ReadField9(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 9: %s", err)
	} else {
		p.Certificate = v
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("loginWithIdentityCredentialForCertificate_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
		return err
	}
	if err := p.writeField5(oprot); err != nil {
		return err
	}
	if err := p.writeField6(oprot); err != nil {
		return err
	}
	if err := p.writeField7(oprot); err != nil {
		return err
	}
	if err := p.writeField8(oprot); err != nil {
		return err
	}
	if err := p.writeField9(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("identifier", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:identifier: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Identifier)); err != nil {
		return fmt.Errorf("%T.identifier (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:identifier: %s", p, err)
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateArgs) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("password", thrift.STRING, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:password: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Password)); err != nil {
		return fmt.Errorf("%T.password (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:password: %s", p, err)
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateArgs) writeField5(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("keepLoggedIn", thrift.BOOL, 5); err != nil {
		return fmt.Errorf("%T write field begin error 5:keepLoggedIn: %s", p, err)
	}
	if err := oprot.WriteBool(bool(p.KeepLoggedIn)); err != nil {
		return fmt.Errorf("%T.keepLoggedIn (5) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 5:keepLoggedIn: %s", p, err)
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateArgs) writeField6(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("accessLocation", thrift.STRING, 6); err != nil {
		return fmt.Errorf("%T write field begin error 6:accessLocation: %s", p, err)
	}
	if err := oprot.WriteString(string(p.AccessLocation)); err != nil {
		return fmt.Errorf("%T.accessLocation (6) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 6:accessLocation: %s", p, err)
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateArgs) writeField7(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("systemName", thrift.STRING, 7); err != nil {
		return fmt.Errorf("%T write field begin error 7:systemName: %s", p, err)
	}
	if err := oprot.WriteString(string(p.SystemName)); err != nil {
		return fmt.Errorf("%T.systemName (7) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 7:systemName: %s", p, err)
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateArgs) writeField8(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("identityProvider", thrift.I32, 8); err != nil {
		return fmt.Errorf("%T write field begin error 8:identityProvider: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.IdentityProvider)); err != nil {
		return fmt.Errorf("%T.identityProvider (8) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 8:identityProvider: %s", p, err)
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateArgs) writeField9(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("certificate", thrift.STRING, 9); err != nil {
		return fmt.Errorf("%T write field begin error 9:certificate: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Certificate)); err != nil {
		return fmt.Errorf("%T.certificate (9) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 9:certificate: %s", p, err)
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LoginWithIdentityCredentialForCertificateArgs(%+v)", *p)
}

type LoginWithIdentityCredentialForCertificateResult struct {
	Success *LoginResult_  `thrift:"success,0" json:"success"`
	E       *TalkException `thrift:"e,1" json:"e"`
}

func NewLoginWithIdentityCredentialForCertificateResult() *LoginWithIdentityCredentialForCertificateResult {
	return &LoginWithIdentityCredentialForCertificateResult{}
}

var LoginWithIdentityCredentialForCertificateResult_Success_DEFAULT *LoginResult_

func (p *LoginWithIdentityCredentialForCertificateResult) GetSuccess() *LoginResult_ {
	if !p.IsSetSuccess() {
		return LoginWithIdentityCredentialForCertificateResult_Success_DEFAULT
	}
	return p.Success
}

var LoginWithIdentityCredentialForCertificateResult_E_DEFAULT *TalkException

func (p *LoginWithIdentityCredentialForCertificateResult) GetE() *TalkException {
	if !p.IsSetE() {
		return LoginWithIdentityCredentialForCertificateResult_E_DEFAULT
	}
	return p.E
}
func (p *LoginWithIdentityCredentialForCertificateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginWithIdentityCredentialForCertificateResult) IsSetE() bool {
	return p.E != nil
}

func (p *LoginWithIdentityCredentialForCertificateResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &LoginResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("loginWithIdentityCredentialForCertificate_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *LoginWithIdentityCredentialForCertificateResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetE() {
		if err := oprot.WriteFieldBegin("e", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:e: %s", p, err)
		}
		if err := p.E.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.E, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:e: %s", p, err)
		}
	}
	return err
}

func (p *LoginWithIdentityCredentialForCertificateResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LoginWithIdentityCredentialForCertificateResult(%+v)", *p)
}
