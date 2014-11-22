package line

import (
	"fmt"

	"git.apache.org/thrift.git/lib/go/thrift"
)

func (p *TalkServiceClient) GetContacts(ids []string) (r []*Contact, err error) {
	if err = p.sendGetContacts(ids); err != nil {
		return
	}
	return p.recvGetContacts()
}

func (p *TalkServiceClient) sendGetContacts(ids []string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("getContacts", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := GetContactsArgs{
		Ids: ids,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvGetContacts() (value []*Contact, err error) {
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
		error1037 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1038 error
		error1038, err = error1037.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1038
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "getContacts failed: out of sequence response")
		return
	}
	result := GetContactsResult{}
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

type GetContactsArgs struct {
	// unused field # 1
	Ids []string `thrift:"ids,2" json:"ids"`
}

func NewGetContactsArgs() *GetContactsArgs {
	return &GetContactsArgs{}
}

func (p *GetContactsArgs) GetIds() []string {
	return p.Ids
}
func (p *GetContactsArgs) Read(iprot thrift.TProtocol) error {
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

func (p *GetContactsArgs) ReadField2(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]string, 0, size)
	p.Ids = tSlice
	for i := 0; i < size; i++ {
		var _elem1357 string
		if v, err := iprot.ReadString(); err != nil {
			return fmt.Errorf("error reading field 0: %s", err)
		} else {
			_elem1357 = v
		}
		p.Ids = append(p.Ids, _elem1357)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *GetContactsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getContacts_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
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

func (p *GetContactsArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("ids", thrift.LIST, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:ids: %s", p, err)
	}
	if err := oprot.WriteListBegin(thrift.STRING, len(p.Ids)); err != nil {
		return fmt.Errorf("error writing list begin: %s", err)
	}
	for _, v := range p.Ids {
		if err := oprot.WriteString(string(v)); err != nil {
			return fmt.Errorf("%T. (0) field write error: %s", p, err)
		}
	}
	if err := oprot.WriteListEnd(); err != nil {
		return fmt.Errorf("error writing list end: %s", err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:ids: %s", p, err)
	}
	return err
}

func (p *GetContactsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetContactsArgs(%+v)", *p)
}

type GetContactsResult struct {
	Success []*Contact     `thrift:"success,0" json:"success"`
	E       *TalkException `thrift:"e,1" json:"e"`
}

func NewGetContactsResult() *GetContactsResult {
	return &GetContactsResult{}
}

var GetContactsResult_Success_DEFAULT []*Contact

func (p *GetContactsResult) GetSuccess() []*Contact {
	return p.Success
}

var GetContactsResult_E_DEFAULT *TalkException

func (p *GetContactsResult) GetE() *TalkException {
	if !p.IsSetE() {
		return GetContactsResult_E_DEFAULT
	}
	return p.E
}
func (p *GetContactsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *GetContactsResult) IsSetE() bool {
	return p.E != nil
}

func (p *GetContactsResult) Read(iprot thrift.TProtocol) error {
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

func (p *GetContactsResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*Contact, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem1358 := &Contact{}
		if err := _elem1358.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem1358, err)
		}
		p.Success = append(p.Success, _elem1358)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *GetContactsResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *GetContactsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("getContacts_result"); err != nil {
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

func (p *GetContactsResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.LIST, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := oprot.WriteListBegin(thrift.STRUCT, len(p.Success)); err != nil {
			return fmt.Errorf("error writing list begin: %s", err)
		}
		for _, v := range p.Success {
			if err := v.Write(oprot); err != nil {
				return fmt.Errorf("%T error writing struct: %s", v, err)
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

func (p *GetContactsResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *GetContactsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetContactsResult(%+v)", *p)
}
