package line

import (
	"fmt"

	"github.com/bcgraham/tsumtsum/external/thrift"
)

func (p *TalkServiceClient) UpdateContactSetting(reqSeq int32, mid string, flag ContactSetting, value string) (err error) {
	if err = p.sendUpdateContactSetting(reqSeq, mid, flag, value); err != nil {
		return
	}
	return p.recvUpdateContactSetting()
}

func (p *TalkServiceClient) sendUpdateContactSetting(reqSeq int32, mid string, flag ContactSetting, value string) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("updateContactSetting", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := UpdateContactSettingArgs{
		ReqSeq: reqSeq,
		Mid:    mid,
		Flag:   flag,
		Value:  value,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *TalkServiceClient) recvUpdateContactSetting() (err error) {
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
		error1277 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error1278 error
		error1278, err = error1277.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error1278
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "updateContactSetting failed: out of sequence response")
		return
	}
	result := UpdateContactSettingResult{}
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
	return
}

type UpdateContactSettingArgs struct {
	ReqSeq int32          `thrift:"reqSeq,1" json:"reqSeq"`
	Mid    string         `thrift:"mid,2" json:"mid"`
	Flag   ContactSetting `thrift:"flag,3" json:"flag"`
	Value  string         `thrift:"value,4" json:"value"`
}

func NewUpdateContactSettingArgs() *UpdateContactSettingArgs {
	return &UpdateContactSettingArgs{}
}

func (p *UpdateContactSettingArgs) GetReqSeq() int32 {
	return p.ReqSeq
}

func (p *UpdateContactSettingArgs) GetMid() string {
	return p.Mid
}

func (p *UpdateContactSettingArgs) GetFlag() ContactSetting {
	return p.Flag
}

func (p *UpdateContactSettingArgs) GetValue() string {
	return p.Value
}
func (p *UpdateContactSettingArgs) Read(iprot thrift.TProtocol) error {
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
		case 3:
			if err := p.ReadField3(iprot); err != nil {
				return err
			}
		case 4:
			if err := p.ReadField4(iprot); err != nil {
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

func (p *UpdateContactSettingArgs) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 1: %s", err)
	} else {
		p.ReqSeq = v
	}
	return nil
}

func (p *UpdateContactSettingArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 2: %s", err)
	} else {
		p.Mid = v
	}
	return nil
}

func (p *UpdateContactSettingArgs) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return fmt.Errorf("error reading field 3: %s", err)
	} else {
		temp := ContactSetting(v)
		p.Flag = temp
	}
	return nil
}

func (p *UpdateContactSettingArgs) ReadField4(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return fmt.Errorf("error reading field 4: %s", err)
	} else {
		p.Value = v
	}
	return nil
}

func (p *UpdateContactSettingArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("updateContactSetting_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := p.writeField3(oprot); err != nil {
		return err
	}
	if err := p.writeField4(oprot); err != nil {
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

func (p *UpdateContactSettingArgs) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *UpdateContactSettingArgs) writeField2(oprot thrift.TProtocol) (err error) {
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

func (p *UpdateContactSettingArgs) writeField3(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("flag", thrift.I32, 3); err != nil {
		return fmt.Errorf("%T write field begin error 3:flag: %s", p, err)
	}
	if err := oprot.WriteI32(int32(p.Flag)); err != nil {
		return fmt.Errorf("%T.flag (3) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 3:flag: %s", p, err)
	}
	return err
}

func (p *UpdateContactSettingArgs) writeField4(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("value", thrift.STRING, 4); err != nil {
		return fmt.Errorf("%T write field begin error 4:value: %s", p, err)
	}
	if err := oprot.WriteString(string(p.Value)); err != nil {
		return fmt.Errorf("%T.value (4) field write error: %s", p, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 4:value: %s", p, err)
	}
	return err
}

func (p *UpdateContactSettingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateContactSettingArgs(%+v)", *p)
}

type UpdateContactSettingResult struct {
	E *TalkException `thrift:"e,1" json:"e"`
}

func NewUpdateContactSettingResult() *UpdateContactSettingResult {
	return &UpdateContactSettingResult{}
}

var UpdateContactSettingResult_E_DEFAULT *TalkException

func (p *UpdateContactSettingResult) GetE() *TalkException {
	if !p.IsSetE() {
		return UpdateContactSettingResult_E_DEFAULT
	}
	return p.E
}
func (p *UpdateContactSettingResult) IsSetE() bool {
	return p.E != nil
}

func (p *UpdateContactSettingResult) Read(iprot thrift.TProtocol) error {
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

func (p *UpdateContactSettingResult) ReadField1(iprot thrift.TProtocol) error {
	p.E = &TalkException{}
	if err := p.E.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.E, err)
	}
	return nil
}

func (p *UpdateContactSettingResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("updateContactSetting_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
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

func (p *UpdateContactSettingResult) writeField1(oprot thrift.TProtocol) (err error) {
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

func (p *UpdateContactSettingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateContactSettingResult(%+v)", *p)
}
