package line

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func (p *TalkServiceClient) GetLastOpRevision() (r int64, err error) {
	if err = p.sendGetLastOpRevision(); err != nil {
		return
	}
	return p.recvGetLastOpRevision()
}

func (p *TalkServiceClient) sendGetLastOpRevision() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getLastOpRevision", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := GetLastOpRevisionArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvGetLastOpRevision() (value int64, err error) {
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
		error1057 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1058 error
		error1058, err = error1057.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1058
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getLastOpRevision failed: out of sequence response")
		return
	}
	result := GetLastOpRevisionResult{}
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

type GetLastOpRevisionArgs struct {
}

func NewGetLastOpRevisionArgs() *GetLastOpRevisionArgs {
	return &GetLastOpRevisionArgs{}
}

func (p *GetLastOpRevisionArgs) Read(iprot thrift.TProtocol) error {
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
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
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

func (p *GetLastOpRevisionArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getLastOpRevision_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *GetLastOpRevisionArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetLastOpRevisionArgs(%+v)", *p)
}

type GetLastOpRevisionResult struct {
	Success *int64         `thrift:"success,0" json:"success"`
	E       *TalkException `thrift:"e,1" json:"e"`
}

func NewGetLastOpRevisionResult() *GetLastOpRevisionResult {
	return &GetLastOpRevisionResult{}
}

var GetLastOpRevisionResult_Success_DEFAULT int64

func (p *GetLastOpRevisionResult) GetSuccess() int64 {
	if !p.IsSetSuccess() {
		return GetLastOpRevisionResult_Success_DEFAULT
	}
	return *p.Success
}

var GetLastOpRevisionResult_E_DEFAULT *TalkException

func (p *GetLastOpRevisionResult) GetE() *TalkException {
	if !p.IsSetE() {
		return GetLastOpRevisionResult_E_DEFAULT
	}
	return p.E
}
func (p *GetLastOpRevisionResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetLastOpRevisionResult) IsSetE() bool {
	return p.E != nil
}

func (p *GetLastOpRevisionResult) Read(iprot thrift.TProtocol) error {
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

func (p *GetLastOpRevisionResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 0: %s", err)
	} else {
		p.Success = &v
	}
	return nil
}

func (p *GetLastOpRevisionResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *GetLastOpRevisionResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getLastOpRevision_result"); err != nil {
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

func (p *GetLastOpRevisionResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.I64, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteI64(int64(*p.Success)); err != nil {
			return fmt.Errorf("%T.success (0) field write error: %s", p, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *GetLastOpRevisionResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *GetLastOpRevisionResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetLastOpRevisionResult(%+v)", *p)
}
