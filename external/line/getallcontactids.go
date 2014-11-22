package line

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func (p *TalkServiceClient) GetAllContactIds() (r []string, err error) {
	if err = p.sendGetAllContactIds(); err != nil {
		return
	}
	return p.recvGetAllContactIds()
}

func (p *TalkServiceClient) sendGetAllContactIds() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getAllContactIds", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := GetAllContactIdsArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvGetAllContactIds() (value []string, err error) {
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
		error1015 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1016 error
		error1016, err = error1015.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1016
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getAllContactIds failed: out of sequence response")
		return
	}
	result := GetAllContactIdsResult{}
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

type GetAllContactIdsArgs struct {
}

func NewGetAllContactIdsArgs() *GetAllContactIdsArgs {
	return &GetAllContactIdsArgs{}
}

func (p *GetAllContactIdsArgs) Read(iprot thrift.TProtocol) error {
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

func (p *GetAllContactIdsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getAllContactIds_args"); err != nil {
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

func (p *GetAllContactIdsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetAllContactIdsArgs(%+v)", *p)
}

type GetAllContactIdsResult struct {
	Success []string       `thrift:"success,0" json:"success"`
	E       *TalkException `thrift:"e,1" json:"e"`
}

func NewGetAllContactIdsResult() *GetAllContactIdsResult {
	return &GetAllContactIdsResult{}
}

var GetAllContactIdsResult_Success_DEFAULT []string

func (p *GetAllContactIdsResult) GetSuccess() []string {
	return p.Success
}

var GetAllContactIdsResult_E_DEFAULT *TalkException

func (p *GetAllContactIdsResult) GetE() *TalkException {
	if !p.IsSetE() {
		return GetAllContactIdsResult_E_DEFAULT
	}
	return p.E
}
func (p *GetAllContactIdsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetAllContactIdsResult) IsSetE() bool {
	return p.E != nil
}

func (p *GetAllContactIdsResult) Read(iprot thrift.TProtocol) error {
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

func (p *GetAllContactIdsResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]string, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		var _elem1351 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_elem1351 = v
		}
		p.Success = append(p.Success, _elem1351)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *GetAllContactIdsResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *GetAllContactIdsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getAllContactIds_result"); err != nil {
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

func (p *GetAllContactIdsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRING, len(p.Success)); err != nil {
			return fmt.Errorf("error writing list begin: %s", err)
		}
		for _, v := range p.Success {
			if err := oprot.WriteString(string(v)); err != nil {
				return fmt.Errorf("%T. (0) field write error: %s", p, err)
			}
		}
		if err := oprot.WriteListEnd(); err != nil {
			return fmt.Errorf("error writing list end: %s", err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *GetAllContactIdsResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *GetAllContactIdsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetAllContactIdsResult(%+v)", *p)
}
