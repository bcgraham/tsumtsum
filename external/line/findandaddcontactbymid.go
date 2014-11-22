package line

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func (p *TalkServiceClient) FindAndAddContactsByMid(reqSeq int32, mid string) (r map[string]*Contact, err error) {
	if err = p.sendFindAndAddContactsByMid(reqSeq, mid); err != nil {
		return
	}
	return p.recvFindAndAddContactsByMid()
}

func (p *TalkServiceClient) sendFindAndAddContactsByMid(reqSeq int32, mid string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("findAndAddContactsByMid", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FindAndAddContactsByMidArgs{
		ReqSeq: reqSeq,
		Mid:    mid,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvFindAndAddContactsByMid() (value map[string]*Contact, err error) {
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
		error991 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error992 error
		error992, err = error991.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error992
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "findAndAddContactsByMid failed: out of sequence response")
		return
	}
	result := FindAndAddContactsByMidResult{}
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

type FindAndAddContactsByMidArgs struct {
	ReqSeq int32  `thrift:"reqSeq,1" json:"reqSeq"`
	Mid    string `thrift:"mid,2" json:"mid"`
}

func NewFindAndAddContactsByMidArgs() *FindAndAddContactsByMidArgs {
	return &FindAndAddContactsByMidArgs{}
}

func (p *FindAndAddContactsByMidArgs) GetReqSeq() int32 {
	return p.ReqSeq
}

func (p *FindAndAddContactsByMidArgs) GetMid() string {
	return p.Mid
}
func (p *FindAndAddContactsByMidArgs) Read(iprot thrift.TProtocol) error {
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
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
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

func (p *FindAndAddContactsByMidArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.ReqSeq = v
	}
	return nil
}

func (p *FindAndAddContactsByMidArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Mid = v
	}
	return nil
}

func (p *FindAndAddContactsByMidArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("findAndAddContactsByMid_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
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

func (p *FindAndAddContactsByMidArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("reqSeq", thrift.I32, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:reqSeq: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.ReqSeq)); err != nil {
		return fmt.Errorf("%T.reqSeq (1) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:reqSeq: %s", p, err)
	}
	return err
}

func (p *FindAndAddContactsByMidArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("mid", thrift.STRING, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:mid: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Mid)); err != nil {
		return fmt.Errorf("%T.mid (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:mid: %s", p, err)
	}
	return err
}

func (p *FindAndAddContactsByMidArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FindAndAddContactsByMidArgs(%+v)", *p)
}

type FindAndAddContactsByMidResult struct {
	Success map[string]*Contact `thrift:"success,0" json:"success"`
	E       *TalkException      `thrift:"e,1" json:"e"`
}

func NewFindAndAddContactsByMidResult() *FindAndAddContactsByMidResult {
	return &FindAndAddContactsByMidResult{}
}

var FindAndAddContactsByMidResult_Success_DEFAULT map[string]*Contact

func (p *FindAndAddContactsByMidResult) GetSuccess() map[string]*Contact {
	return p.Success
}

var FindAndAddContactsByMidResult_E_DEFAULT *TalkException

func (p *FindAndAddContactsByMidResult) GetE() *TalkException {
	if !p.IsSetE() {
		return FindAndAddContactsByMidResult_E_DEFAULT
	}
	return p.E
}
func (p *FindAndAddContactsByMidResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FindAndAddContactsByMidResult) IsSetE() bool {
	return p.E != nil
}

func (p *FindAndAddContactsByMidResult) Read(iprot thrift.TProtocol) error {
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

func (p *FindAndAddContactsByMidResult) ReadField0(iprot thrift.TProtocol) error {
	_, _, size, err := iprot.ReadMapBegin()
	if err != nil {
		return fmt.Errorf("error reading map begin: %s", err)
	}
	tMap := make(map[string]*Contact, size)
	p.Success = tMap
	for i := 0; i < size; i++ {
		var _key1336 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_key1336 = v
		}
		_val1337 := &Contact{}
		if err := _val1337.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _val1337, err)
		}
		p.Success[_key1336] = _val1337
	}
	if err := iprot.ReadMapEnd(); err != nil {
		return fmt.Errorf("error reading map end: %s", err)
	}
	return nil
}

func (p *FindAndAddContactsByMidResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *FindAndAddContactsByMidResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("findAndAddContactsByMid_result"); err != nil {
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

func (p *FindAndAddContactsByMidResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.MAP, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteMapBegin(thrift.STRING, thrift.STRUCT, len(p.Success)); err != nil {
			return fmt.Errorf("error writing map begin: %s", err)
		}
		for k, v := range p.Success {
			if err := oprot.WriteString(string(k)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v, err)
			}
		}
		if err := oprot.WriteMapEnd(); err != nil {
			return fmt.Errorf("error writing map end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *FindAndAddContactsByMidResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *FindAndAddContactsByMidResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FindAndAddContactsByMidResult(%+v)", *p)
}
