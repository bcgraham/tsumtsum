package line

import (
	"fmt"

	"github.com/bcgraham/tsumtsum/external/thrift"
)

func (p *TalkServiceClient) FetchOperations(localRev int64, count int32) (r []*Operation, err error) {
	if err = p.sendFetchOperations(localRev, count); err != nil {
		return
	}
	return p.recvFetchOperations()
}

func (p *TalkServiceClient) sendFetchOperations(localRev int64, count int32) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("fetchOperations", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := FetchOperationsArgs{
		LocalRev: localRev,
		Count:    count,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvFetchOperations() (value []*Operation, err error) {
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
		error985 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error986 error
		error986, err = error985.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error986
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "fetchOperations failed: out of sequence response")
		return
	}
	result := FetchOperationsResult{}
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

type FetchOperationsArgs struct {
	// unused field # 1
	LocalRev int64 `thrift:"localRev,2" json:"localRev"`
	Count    int32 `thrift:"count,3" json:"count"`
}

func NewFetchOperationsArgs() *FetchOperationsArgs {
	return &FetchOperationsArgs{}
}

func (p *FetchOperationsArgs) GetLocalRev() int64 {
	return p.LocalRev
}

func (p *FetchOperationsArgs) GetCount() int32 {
	return p.Count
}
func (p *FetchOperationsArgs) Read(iprot thrift.TProtocol) error {
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

func (p *FetchOperationsArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI64(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.LocalRev = v
	}
	return nil
}

func (p *FetchOperationsArgs) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		p.Count = v
	}
	return nil
}

func (p *FetchOperationsArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("fetchOperations_args"); err != nil {
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

func (p *FetchOperationsArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("localRev", thrift.I64, 2); err != nil {
		return fmt.Errorf("%T write field begin error 2:localRev: %s", p, err)
	}
	if err := oprot.WriteI64(int64(p.LocalRev)); err != nil {
		return fmt.Errorf("%T.localRev (2) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 2:localRev: %s", p, err)
	}
	return err
}

func (p *FetchOperationsArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("count", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:count: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Count)); err != nil {
		return fmt.Errorf("%T.count (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:count: %s", p, err)
	}
	return err
}

func (p *FetchOperationsArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FetchOperationsArgs(%+v)", *p)
}

type FetchOperationsResult struct {
	Success []*Operation   `thrift:"success,0" json:"success"`
	E       *TalkException `thrift:"e,1" json:"e"`
}

func NewFetchOperationsResult() *FetchOperationsResult {
	return &FetchOperationsResult{}
}

var FetchOperationsResult_Success_DEFAULT []*Operation

func (p *FetchOperationsResult) GetSuccess() []*Operation {
	return p.Success
}

var FetchOperationsResult_E_DEFAULT *TalkException

func (p *FetchOperationsResult) GetE() *TalkException {
	if !p.IsSetE() {
		return FetchOperationsResult_E_DEFAULT
	}
	return p.E
}
func (p *FetchOperationsResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *FetchOperationsResult) IsSetE() bool {
	return p.E != nil
}

func (p *FetchOperationsResult) Read(iprot thrift.TProtocol) error {
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

func (p *FetchOperationsResult) ReadField0(iprot thrift.TProtocol) error {
	_, size, err := iprot.ReadListBegin()
	if err != nil {
		return fmt.Errorf("error reading list begin: %s", err)
	}
	tSlice := make([]*Operation, 0, size)
	p.Success = tSlice
	for i := 0; i < size; i++ {
		_elem1331 := &Operation{}
		if err := _elem1331.Read(iprot); err != nil {
			return fmt.Errorf("%T error reading struct: %s", _elem1331, err)
		}
		p.Success = append(p.Success, _elem1331)
	}
	if err := iprot.ReadListEnd(); err != nil {
		return fmt.Errorf("error reading list end: %s", err)
	}
	return nil
}

func (p *FetchOperationsResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *FetchOperationsResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("fetchOperations_result"); err != nil {
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

func (p *FetchOperationsResult) writeField0(oprot thrift.TProtocol) (err error) {
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

func (p *FetchOperationsResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *FetchOperationsResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("FetchOperationsResult(%+v)", *p)
}
