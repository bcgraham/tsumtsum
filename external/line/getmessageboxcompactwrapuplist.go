package line

import (
	"fmt"

	"github.com/bcgraham/tsumtsum/external/thrift"
)

func (p *TalkServiceClient) GetMessageBoxCompactWrapUpList(start int32, messageBoxCount int32) (r *TMessageBoxWrapUpResponse, err error) {
	if err = p.sendGetMessageBoxCompactWrapUpList(start, messageBoxCount); err != nil {
		return
	}
	return p.recvGetMessageBoxCompactWrapUpList()
}

func (p *TalkServiceClient) sendGetMessageBoxCompactWrapUpList(start int32, messageBoxCount int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getMessageBoxCompactWrapUpList", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := GetMessageBoxCompactWrapUpListArgs{
		Start:           start,
		MessageBoxCount: messageBoxCount,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvGetMessageBoxCompactWrapUpList() (value *TMessageBoxWrapUpResponse, err error) {
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
		error1063 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1064 error
		error1064, err = error1063.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1064
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getMessageBoxCompactWrapUpList failed: out of sequence response")
		return
	}
	result := GetMessageBoxCompactWrapUpListResult{}
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

type GetMessageBoxCompactWrapUpListArgs struct {
	// unused field # 1
	Start           int32 `thrift:"start,2" json:"start"`
	MessageBoxCount int32 `thrift:"messageBoxCount,3" json:"messageBoxCount"`
}

func NewGetMessageBoxCompactWrapUpListArgs() *GetMessageBoxCompactWrapUpListArgs {
	return &GetMessageBoxCompactWrapUpListArgs{}
}

func (p *GetMessageBoxCompactWrapUpListArgs) GetStart() int32 {
	return p.Start
}

func (p *GetMessageBoxCompactWrapUpListArgs) GetMessageBoxCount() int32 {
	return p.MessageBoxCount
}
func (p *GetMessageBoxCompactWrapUpListArgs) Read(iprot thrift.TProtocol) error {
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
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
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

func (p *GetMessageBoxCompactWrapUpListArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Start = v
	}
	return nil
}

func (p *GetMessageBoxCompactWrapUpListArgs) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.MessageBoxCount = v
	}
	return nil
}

func (p *GetMessageBoxCompactWrapUpListArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getMessageBoxCompactWrapUpList_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField2(oprot); err != nil {
		return err
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

func (p *GetMessageBoxCompactWrapUpListArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("start", thrift.I32, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:start: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Start)); err != nil {
		return fmt.Errorf("%T.start (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:start: %s", p, err)
	}
	return err
}

func (p *GetMessageBoxCompactWrapUpListArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("messageBoxCount", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:messageBoxCount: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.MessageBoxCount)); err != nil {
		return fmt.Errorf("%T.messageBoxCount (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:messageBoxCount: %s", p, err)
	}
	return err
}

func (p *GetMessageBoxCompactWrapUpListArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetMessageBoxCompactWrapUpListArgs(%+v)", *p)
}

type GetMessageBoxCompactWrapUpListResult struct {
	Success *TMessageBoxWrapUpResponse `thrift:"success,0" json:"success"`
	E       *TalkException             `thrift:"e,1" json:"e"`
}

func NewGetMessageBoxCompactWrapUpListResult() *GetMessageBoxCompactWrapUpListResult {
	return &GetMessageBoxCompactWrapUpListResult{}
}

var GetMessageBoxCompactWrapUpListResult_Success_DEFAULT *TMessageBoxWrapUpResponse

func (p *GetMessageBoxCompactWrapUpListResult) GetSuccess() *TMessageBoxWrapUpResponse {
	if !p.IsSetSuccess() {
		return GetMessageBoxCompactWrapUpListResult_Success_DEFAULT
	}
	return p.Success
}

var GetMessageBoxCompactWrapUpListResult_E_DEFAULT *TalkException

func (p *GetMessageBoxCompactWrapUpListResult) GetE() *TalkException {
	if !p.IsSetE() {
		return GetMessageBoxCompactWrapUpListResult_E_DEFAULT
	}
	return p.E
}
func (p *GetMessageBoxCompactWrapUpListResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetMessageBoxCompactWrapUpListResult) IsSetE() bool {
	return p.E != nil
}

func (p *GetMessageBoxCompactWrapUpListResult) Read(iprot thrift.TProtocol) error {
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

func (p *GetMessageBoxCompactWrapUpListResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &TMessageBoxWrapUpResponse{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *GetMessageBoxCompactWrapUpListResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *GetMessageBoxCompactWrapUpListResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getMessageBoxCompactWrapUpList_result"); err != nil {
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

func (p *GetMessageBoxCompactWrapUpListResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *GetMessageBoxCompactWrapUpListResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *GetMessageBoxCompactWrapUpListResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetMessageBoxCompactWrapUpListResult(%+v)", *p)
}
