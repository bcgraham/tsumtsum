package line

import (
	"fmt"

	"github.com/bcgraham/tsumtsum/external/thrift"
)

func (p *TalkServiceClient) LoginWithVerifierForCertificate(verifier string) (r *LoginResult_, err error) {
	if err = p.sendLoginWithVerifierForCertificate(verifier); err != nil {
		return
	}
	return p.recvLoginWithVerifierForCertificate()
}

func (p *TalkServiceClient) sendLoginWithVerifierForCertificate(verifier string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("loginWithVerifierForCertificate", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := LoginWithVerifierForCertificateArgs{
		Verifier: verifier,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvLoginWithVerifierForCertificate() (value *LoginResult_, err error) {
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
		error1137 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1138 error
		error1138, err = error1137.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1138
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "loginWithVerifierForCertificate failed: out of sequence response")
		return
	}
	result := LoginWithVerifierForCertificateResult{}
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

type LoginWithVerifierForCertificateArgs struct {
	// unused fields # 1 to 2
	Verifier string `thrift:"verifier,3" json:"verifier"`
}

func NewLoginWithVerifierForCertificateArgs() *LoginWithVerifierForCertificateArgs {
	return &LoginWithVerifierForCertificateArgs{}
}

func (p *LoginWithVerifierForCertificateArgs) GetVerifier() string {
	return p.Verifier
}
func (p *LoginWithVerifierForCertificateArgs) Read(iprot thrift.TProtocol) error {
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
		case 3:
			if err := p.ReadField3(iprot); err != nil {
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

func (p *LoginWithVerifierForCertificateArgs) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Verifier = v
	}
	return nil
}

func (p *LoginWithVerifierForCertificateArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("loginWithVerifierForCertificate_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField3(oprot); err != nil {
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

func (p *LoginWithVerifierForCertificateArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("verifier", thrift.STRING, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:verifier: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Verifier)); err != nil {
		return fmt.Errorf("%T.verifier (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:verifier: %s", p, err)
	}
	return err
}

func (p *LoginWithVerifierForCertificateArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LoginWithVerifierForCertificateArgs(%+v)", *p)
}

type LoginWithVerifierForCertificateResult struct {
	Success *LoginResult_  `thrift:"success,0" json:"success"`
	E       *TalkException `thrift:"e,1" json:"e"`
}

func NewLoginWithVerifierForCertificateResult() *LoginWithVerifierForCertificateResult {
	return &LoginWithVerifierForCertificateResult{}
}

var LoginWithVerifierForCertificateResult_Success_DEFAULT *LoginResult_

func (p *LoginWithVerifierForCertificateResult) GetSuccess() *LoginResult_ {
	if !p.IsSetSuccess() {
		return LoginWithVerifierForCertificateResult_Success_DEFAULT
	}
	return p.Success
}

var LoginWithVerifierForCertificateResult_E_DEFAULT *TalkException

func (p *LoginWithVerifierForCertificateResult) GetE() *TalkException {
	if !p.IsSetE() {
		return LoginWithVerifierForCertificateResult_E_DEFAULT
	}
	return p.E
}
func (p *LoginWithVerifierForCertificateResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *LoginWithVerifierForCertificateResult) IsSetE() bool {
	return p.E != nil
}

func (p *LoginWithVerifierForCertificateResult) Read(iprot thrift.TProtocol) error {
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

func (p *LoginWithVerifierForCertificateResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &LoginResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *LoginWithVerifierForCertificateResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *LoginWithVerifierForCertificateResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("loginWithVerifierForCertificate_result"); err != nil {
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

func (p *LoginWithVerifierForCertificateResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *LoginWithVerifierForCertificateResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *LoginWithVerifierForCertificateResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("LoginWithVerifierForCertificateResult(%+v)", *p)
}
