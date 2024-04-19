package braidsBrokerTypes
import(
	"io"
	"math"
	"reflect"
	"unsafe"
	"ethos/llrb"
	"bytes"
	"encoding/hex"
	"time"
	_syscall "ethos/syscall" // This is for ETN generated source code, to separate ETN's syscall from user d e f i n e d syscall field
				 // Every line that has  d e f i n e d in it gets removed when building the go toolchain, sounds like a hack
				 // is a hack..
	"ethos/syscall"
	"ethos/defined"
	
)

// FIXME: this is a place holder for bytes. bytes is only used when there is an any type
var xxx = bytes.MinRead
var yyy hex.InvalidHexCharError
var sunday = time.Sunday // this needs to be remove after we fixed the sleep in Ipc and IpcWrite.
var the_import_is_not_used_error_is_not_helpful_It_depends_if_need_syscall_and_Ill_know_that_much_later syscall.Fd
var blah_blah_blah_MR_Freeman_blah_blah_blah defined.Rpc

type Int8 int8
type Uint8 uint8
type Int16 int16
type Uint16 uint16
type Int32 int32
type Uint32 uint32
type Int64 int64
type Uint64 uint64
type Float32 float32
type Float64 float64
type Bool bool
type Byte byte
type String string

type Puller struct {
     
     Username string
     
     Key string
     
}


type Message struct {
     
     Id string
     
     Data string
     
     CreatedAt int64
     
}


type Pusher struct {
     
     Username string
     
     Key string
     
}

const(
	
	
	methodIdBrokerPull = iota
	
	
	
	methodIdBrokerPullReply
	
	
	
	methodIdBrokerPush
	
	
	
	methodIdBrokerPushReply
	
	
)

func (e *Encoder) BrokerPull(user *Puller) (_returnStatus _syscall.Status){
     _returnStatus = e.uint64(methodIdBrokerPull)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Puller(user)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Flush()
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     return _returnStatus
}

func (e *Encoder) BrokerPullReply(msg *Message, status syscall.Status) (_returnStatus _syscall.Status){
     _returnStatus = e.uint64(methodIdBrokerPullReply)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Message(msg)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Status(status)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Flush()
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     return _returnStatus
}

func (e *Encoder) BrokerPush(user *Pusher, data string) (_returnStatus _syscall.Status){
     _returnStatus = e.uint64(methodIdBrokerPush)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Pusher(user)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.string(data)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Flush()
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     return _returnStatus
}

func (e *Encoder) BrokerPushReply(status syscall.Status) (_returnStatus _syscall.Status){
     _returnStatus = e.uint64(methodIdBrokerPushReply)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Status(status)
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     
     _returnStatus = e.Flush()
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     return _returnStatus
}

func (d *Decoder) HandleBroker(e *Encoder) (_returnStatus _syscall.Status){
     var retValue BrokerProcedure
     d.ReadAll()
     methodId, _returnStatus := d.uint64()
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus		
     }
     switch *methodId{
     	    
	    case methodIdBrokerPull:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		
		user, _returnStatus0 := d.Puller()
		if _returnStatus0 != _syscall.StatusOk {
		   return _returnStatus0
		}
		
		
		retValue = _fp_brokerPull(*user)
		
		
	    
	    case methodIdBrokerPullReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		
		msg, _returnStatus0 := d.Message()
		if _returnStatus0 != _syscall.StatusOk {
		   return _returnStatus0
		}
		
		status, _returnStatus1 := d.Status()
		if _returnStatus1 != _syscall.StatusOk {
		   return _returnStatus1
		}
		
		
		retValue = _fp_brokerPullReply(*msg,(syscall.Status)(*status))
		
		
	    
	    case methodIdBrokerPush:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		
		user, _returnStatus0 := d.Pusher()
		if _returnStatus0 != _syscall.StatusOk {
		   return _returnStatus0
		}
		
		data, _returnStatus1 := d.string()
		if _returnStatus1 != _syscall.StatusOk {
		   return _returnStatus1
		}
		
		
		retValue = _fp_brokerPush(*user,*data)
		
		
	    
	    case methodIdBrokerPushReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		
		status, _returnStatus0 := d.Status()
		if _returnStatus0 != _syscall.StatusOk {
		   return _returnStatus0
		}
		
		
		retValue = _fp_brokerPushReply((syscall.Status)(*status))
		
		
	    
	    default:
	    	 // e := fmt.Errorf("Wrong MethodID")
		 return _syscall.StatusFail
     }

     return e.BrokerProcedure(retValue)
    
}

func (t *Broker) RpcReadBuffer(buffer []byte) (defined.Rpc, _syscall.Status) {
     d := NewDecoder(bytes.NewBuffer(buffer))
     d.ReadAll()
     methodId, _returnStatus := d.uint64()
     if _returnStatus != _syscall.StatusOk {
     	return nil, _syscall.StatusFail
     }
     switch *methodId{
     	    
	    case methodIdBrokerPull:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		
		Arg0, _returnStatus0 := d.Puller()
		if _returnStatus0 != _syscall.StatusOk {
		   return nil, _syscall.StatusFail
		}
		
		

		if !d.IsEOF() {
			return nil, _syscall.StatusInvalidType
		}

		return &BrokerPull{ *Arg0,  }, _syscall.StatusOk
	    
	    case methodIdBrokerPullReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		
		Arg0, _returnStatus0 := d.Message()
		if _returnStatus0 != _syscall.StatusOk {
		   return nil, _syscall.StatusFail
		}
		
		Arg1, _returnStatus1 := d.Status()
		if _returnStatus1 != _syscall.StatusOk {
		   return nil, _syscall.StatusFail
		}
		
		

		if !d.IsEOF() {
			return nil, _syscall.StatusInvalidType
		}

		return &BrokerPullReply{ *Arg0, *Arg1,  }, _syscall.StatusOk
	    
	    case methodIdBrokerPush:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		
		Arg0, _returnStatus0 := d.Pusher()
		if _returnStatus0 != _syscall.StatusOk {
		   return nil, _syscall.StatusFail
		}
		
		Arg1, _returnStatus1 := d.string()
		if _returnStatus1 != _syscall.StatusOk {
		   return nil, _syscall.StatusFail
		}
		
		

		if !d.IsEOF() {
			return nil, _syscall.StatusInvalidType
		}

		return &BrokerPush{ *Arg0, *Arg1,  }, _syscall.StatusOk
	    
	    case methodIdBrokerPushReply:
	    	 /* FIXME: Temporarily place reset here. We also reset in ReadAll calls. This 
		  *        should probably go in one place: each type decoder, but we need
		  *        to be able to differentiate between an application-called decode
		  *        and a decode of internal elements. The latter case should not 
		  *        call reset. */
	    	 d.reset()
	    	 
		
		Arg0, _returnStatus0 := d.Status()
		if _returnStatus0 != _syscall.StatusOk {
		   return nil, _syscall.StatusFail
		}
		
		

		if !d.IsEOF() {
			return nil, _syscall.StatusInvalidType
		}

		return &BrokerPushReply{ *Arg0,  }, _syscall.StatusOk
	    
	    default:
		 return nil, _syscall.StatusFail
     }

	return nil, _syscall.StatusFail
}



func (t *BrokerPull) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.BrokerPull(&t.Var0, )
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *BrokerPull) read(d *Decoder) (*BrokerPull, _syscall.Status){
	var value BrokerPull
	d.reset()
	methodId, _returnStatus := d.uint64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}

	if *methodId != methodIdBrokerPull {
		return nil, _syscall.StatusFail
	}

	
	Var0, _returnStatus := d.Puller()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}
	value.Var0 = *Var0
 	
	return &value, _syscall.StatusOk
}


func (t *BrokerPull) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := t.read(d)
	if _returnStatus != _syscall.StatusOk {
		return _returnStatus
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *BrokerPull) GetHash() (_syscall.HashValue) {
       return _syscall.HashValue{ 0xe4,0x92,0x42,0x3f,0x0b,0xef,0x89,0xd8,0x37,0x67,0xc1,0x6b,0xe5,0x26,0xdc,0x71,0x0d,0xb1,0xd6,0x83,0x6b,0x51,0x06,0x52,0xab,0xd9,0xd6,0x73,0x37,0x18,0x5f,0x15,0x8b,0x3f,0x82,0xa1,0x92,0x83,0x02,0x3a,0x02,0xb4,0x96,0x4e,0xba,0x54,0xcb,0x34,0xf4,0x30,0x76,0x8b,0xd0,0x9d,0x6b,0x21,0x72,0xb1,0x67,0xb4,0xc2,0xe2,0x67,0x8b, }
}



func (t *BrokerPullReply) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.BrokerPullReply(&t.Var0, t.Var1, )
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *BrokerPullReply) read(d *Decoder) (*BrokerPullReply, _syscall.Status){
	var value BrokerPullReply
	d.reset()
	methodId, _returnStatus := d.uint64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}

	if *methodId != methodIdBrokerPullReply {
		return nil, _syscall.StatusFail
	}

	
	Var0, _returnStatus := d.Message()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}
	value.Var0 = *Var0
 	
	Var1, _returnStatus := d.Status()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}
	value.Var1 = (syscall.Status)(*Var1)
 	
	return &value, _syscall.StatusOk
}


func (t *BrokerPullReply) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := t.read(d)
	if _returnStatus != _syscall.StatusOk {
		return _returnStatus
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *BrokerPullReply) GetHash() (_syscall.HashValue) {
       return _syscall.HashValue{ 0x7a,0xd2,0x89,0xc2,0x07,0x75,0x5c,0x4d,0x45,0xa2,0xb0,0x91,0x2f,0xdc,0x42,0xf7,0x9d,0xb1,0x14,0x3d,0xff,0x4f,0x11,0x1e,0xc0,0x52,0x47,0xf5,0x6a,0x9b,0x1f,0x5b,0x8e,0xe7,0x68,0x6c,0x67,0x6d,0x96,0xef,0x9d,0xa5,0xe0,0x2e,0x9c,0x82,0x9a,0x64,0x62,0xfc,0xd1,0x85,0x4a,0x6f,0xf4,0x4a,0x19,0xca,0x5c,0x95,0x85,0x99,0x75,0x64, }
}



func (t *BrokerPush) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.BrokerPush(&t.Var0, t.Var1, )
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *BrokerPush) read(d *Decoder) (*BrokerPush, _syscall.Status){
	var value BrokerPush
	d.reset()
	methodId, _returnStatus := d.uint64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}

	if *methodId != methodIdBrokerPush {
		return nil, _syscall.StatusFail
	}

	
	Var0, _returnStatus := d.Pusher()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}
	value.Var0 = *Var0
 	
	Var1, _returnStatus := d.string()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}
	value.Var1 = *Var1
 	
	return &value, _syscall.StatusOk
}


func (t *BrokerPush) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := t.read(d)
	if _returnStatus != _syscall.StatusOk {
		return _returnStatus
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *BrokerPush) GetHash() (_syscall.HashValue) {
       return _syscall.HashValue{ 0x12,0x5d,0x1f,0x77,0x7f,0x87,0xad,0x93,0x87,0xa1,0x2e,0x55,0xf2,0xcb,0x11,0x14,0x4d,0x8f,0x9b,0x22,0x11,0x8d,0xbc,0xa5,0xd3,0x57,0x17,0x91,0xd8,0x44,0xb3,0xf9,0x02,0xec,0x79,0x16,0x2e,0xac,0xd8,0x92,0x78,0x3f,0x82,0x58,0x42,0xeb,0xd4,0x76,0x9b,0xe1,0xbf,0x5d,0x96,0x9d,0xd5,0x3d,0xf3,0xce,0x47,0x57,0xd8,0xc4,0xb2,0x30, }
}



func (t *BrokerPushReply) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.BrokerPushReply(t.Var0, )
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *BrokerPushReply) read(d *Decoder) (*BrokerPushReply, _syscall.Status){
	var value BrokerPushReply
	d.reset()
	methodId, _returnStatus := d.uint64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}

	if *methodId != methodIdBrokerPushReply {
		return nil, _syscall.StatusFail
	}

	
	Var0, _returnStatus := d.Status()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusFail
	}
	value.Var0 = (syscall.Status)(*Var0)
 	
	return &value, _syscall.StatusOk
}


func (t *BrokerPushReply) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := t.read(d)
	if _returnStatus != _syscall.StatusOk {
		return _returnStatus
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *BrokerPushReply) GetHash() (_syscall.HashValue) {
       return _syscall.HashValue{ 0x0f,0xe6,0x2a,0x71,0x04,0x59,0x75,0x14,0xfb,0xac,0x5c,0xea,0x22,0x78,0xd8,0x08,0x5b,0x4e,0x5d,0x42,0x0f,0x1e,0x76,0xfc,0x6f,0x91,0xff,0x2f,0xcd,0x57,0x35,0x8d,0x6c,0xdf,0xca,0x0c,0x2e,0xc5,0xde,0xbf,0x7f,0x37,0x11,0xc5,0x0f,0x43,0xd7,0xc0,0xb8,0x61,0x47,0xfd,0xa2,0x04,0x86,0x51,0x07,0x42,0xaf,0x34,0x3f,0x03,0x1c,0x96, }
}



type BrokerProcedure interface {
	//These interfaces are used to differ several rpc interfaces in one package
	//The compiler will complain if a rpc function returns a procedure from another
	//rpc interface
	GetHash() _syscall.HashValue
	WriteBuffer()             ([]byte, _syscall.Status)
	ReadBuffer(buffer []byte) (_syscall.Status)
	Invoke()                  defined.Rpc
	GetInterface()            defined.RpcInterface
	BrokerProcedure()
}

type Broker struct {
}



type BrokerPull struct {
	 Var0 Puller
	
}

func (t *BrokerPull) BrokerProcedure() {
	//oh look its nothing
}

func _default_brokerPull( Var0 Puller, ) (BrokerProcedure) {
	panic("brokerPull not implemented")
	return nil
}

var _fp_brokerPull func( Puller,)(BrokerProcedure) = _default_brokerPull


func SetupBrokerPull (fn func( Puller,)(BrokerProcedure)) {
	_fp_brokerPull = fn
}

func (t *BrokerPull) Invoke() defined.Rpc {
	return _fp_brokerPull( t.Var0, )
}

func (t *BrokerPull) GetInterface() defined.RpcInterface {
	return &Broker{}
}


type BrokerPullReply struct {
	 Var0 Message
	 Var1 syscall.Status
	
}

func (t *BrokerPullReply) BrokerProcedure() {
	//oh look its nothing
}

func _default_brokerPullReply( Var0 Message,  Var1 syscall.Status, ) (BrokerProcedure) {
	panic("brokerPullReply not implemented")
	return nil
}

var _fp_brokerPullReply func( Message, syscall.Status,)(BrokerProcedure) = _default_brokerPullReply


func SetupBrokerPullReply (fn func( Message, syscall.Status,)(BrokerProcedure)) {
	_fp_brokerPullReply = fn
}

func (t *BrokerPullReply) Invoke() defined.Rpc {
	return _fp_brokerPullReply( t.Var0,  t.Var1, )
}

func (t *BrokerPullReply) GetInterface() defined.RpcInterface {
	return &Broker{}
}


type BrokerPush struct {
	 Var0 Pusher
	 Var1 string
	
}

func (t *BrokerPush) BrokerProcedure() {
	//oh look its nothing
}

func _default_brokerPush( Var0 Pusher,  Var1 string, ) (BrokerProcedure) {
	panic("brokerPush not implemented")
	return nil
}

var _fp_brokerPush func( Pusher, string,)(BrokerProcedure) = _default_brokerPush


func SetupBrokerPush (fn func( Pusher, string,)(BrokerProcedure)) {
	_fp_brokerPush = fn
}

func (t *BrokerPush) Invoke() defined.Rpc {
	return _fp_brokerPush( t.Var0,  t.Var1, )
}

func (t *BrokerPush) GetInterface() defined.RpcInterface {
	return &Broker{}
}


type BrokerPushReply struct {
	 Var0 syscall.Status
	
}

func (t *BrokerPushReply) BrokerProcedure() {
	//oh look its nothing
}

func _default_brokerPushReply( Var0 syscall.Status, ) (BrokerProcedure) {
	panic("brokerPushReply not implemented")
	return nil
}

var _fp_brokerPushReply func( syscall.Status,)(BrokerProcedure) = _default_brokerPushReply


func SetupBrokerPushReply (fn func( syscall.Status,)(BrokerProcedure)) {
	_fp_brokerPushReply = fn
}

func (t *BrokerPushReply) Invoke() defined.Rpc {
	return _fp_brokerPushReply( t.Var0, )
}

func (t *BrokerPushReply) GetInterface() defined.RpcInterface {
	return &Broker{}
}



func InterfaceHandleBroker(params []interface{}) (_returnStatus _syscall.Status){
     var retValue BrokerProcedure
     e := params[0].(*Encoder)
     params = params[1:]
     methodId := params[0].(uint64)
     params = params[1:]
     switch methodId {
     	    
	    case methodIdBrokerPull:
	    	 
			
			     retValue = _fp_brokerPull(*(params[0].(*Puller)), )
		 	
	 	 
	    
	    case methodIdBrokerPullReply:
	    	 
			
			     retValue = _fp_brokerPullReply(*(params[0].(*Message)), *(params[1].(*syscall.Status)), )
		 	
	 	 
	    
	    case methodIdBrokerPush:
	    	 
			
			     retValue = _fp_brokerPush(*(params[0].(*Pusher)), *(params[1].(*string)), )
		 	
	 	 
	    
	    case methodIdBrokerPushReply:
	    	 
			
			     retValue = _fp_brokerPushReply(*(params[0].(*syscall.Status)), )
		 	
	 	 
	    
	    default:
	    	 // e := fmt.Errorf("Wrong MethodID")
		 return _syscall.StatusFail
     }
     return e.BrokerProcedure(retValue)
}

func (e *Encoder) BrokerProcedure(retValue BrokerProcedure) (_returnStatus _syscall.Status) {
    if retValue == nil {
        return _syscall.StatusOk
    }

    switch retValue.(type) {

        case *BrokerPull:
             
             t := retValue.(*BrokerPull)
             
             _returnStatus := e.BrokerPull(&t.Var0, )
             if _returnStatus != _syscall.StatusOk {
                  return _returnStatus
             }
             return e.Flush()

        case *BrokerPullReply:
             
             t := retValue.(*BrokerPullReply)
             
             _returnStatus := e.BrokerPullReply(&t.Var0, t.Var1, )
             if _returnStatus != _syscall.StatusOk {
                  return _returnStatus
             }
             return e.Flush()

        case *BrokerPush:
             
             t := retValue.(*BrokerPush)
             
             _returnStatus := e.BrokerPush(&t.Var0, t.Var1, )
             if _returnStatus != _syscall.StatusOk {
                  return _returnStatus
             }
             return e.Flush()

        case *BrokerPushReply:
             
             t := retValue.(*BrokerPushReply)
             
             _returnStatus := e.BrokerPushReply(t.Var0, )
             if _returnStatus != _syscall.StatusOk {
                  return _returnStatus
             }
             return e.Flush()

       default:
	 return _syscall.StatusInvalidType //fmt.Errorf("Invalid Type in return value")
    }

   return _syscall.StatusOk
}
func DecodeBroker(decoder interface{}) (params []interface{}, _returnStatus _syscall.Status) {
     d := decoder.(*Decoder)
     d.ReadAll()
     methodId, _returnStatus := d.uint64()
     if _returnStatus != _syscall.StatusOk {
     	return []interface{}{}, _returnStatus		
     }

     params = append(params, *methodId)
     switch *methodId {
     	
     	  case methodIdBrokerPull:
               d.reset()
     	       
     	       user, _returnStatus0 := d.Puller()
               if _returnStatus0 != _syscall.StatusOk {
	       	  return []interface{}{}, _returnStatus0
     	       }
     	       params = append(params, user)
     	       
     	
     	  case methodIdBrokerPullReply:
               d.reset()
     	       
     	       msg, _returnStatus0 := d.Message()
               if _returnStatus0 != _syscall.StatusOk {
	       	  return []interface{}{}, _returnStatus0
     	       }
     	       params = append(params, msg)
     	       
     	       status, _returnStatus1 := d.Status()
               if _returnStatus1 != _syscall.StatusOk {
	       	  return []interface{}{}, _returnStatus1
     	       }
     	       params = append(params, status)
     	       
     	
     	  case methodIdBrokerPush:
               d.reset()
     	       
     	       user, _returnStatus0 := d.Pusher()
               if _returnStatus0 != _syscall.StatusOk {
	       	  return []interface{}{}, _returnStatus0
     	       }
     	       params = append(params, user)
     	       
     	       data, _returnStatus1 := d.string()
               if _returnStatus1 != _syscall.StatusOk {
	       	  return []interface{}{}, _returnStatus1
     	       }
     	       params = append(params, data)
     	       
     	
     	  case methodIdBrokerPushReply:
               d.reset()
     	       
     	       status, _returnStatus0 := d.Status()
               if _returnStatus0 != _syscall.StatusOk {
	       	  return []interface{}{}, _returnStatus0
     	       }
     	       params = append(params, status)
     	       
          
         default:
		// e := fmt.Errorf("Wrong MethodID")
		return []interface{}{}, _syscall.StatusFail
     }    
     return params, _syscall.StatusOk
}
func (e *Encoder) Puller(v *Puller) (_returnStatus _syscall.Status){
     
     if _, _, _returnStatusP := e.t.PointerCheck(unsafe.Pointer(v), "18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358", uint64(unsafe.Sizeof(*v))); _returnStatusP == _syscall.StatusOk {
     	
       _returnStatus = e.string(v.Username)
       if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
       
       _returnStatus = e.string(v.Key)
       if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
        
     } else {
       return _returnStatusP
     }
     return _syscall.StatusOk
}
func (d *Decoder) Puller() (v *Puller, _returnStatus _syscall.Status){
     
     
     var valv Puller
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, _returnStatus0 := d.string()
     if _returnStatus0 != _syscall.StatusOk {
     	return &valv, _returnStatus0
     }
     
     valv.Username = *p0
     
     
     		
     p1, _returnStatus1 := d.string()
     if _returnStatus1 != _syscall.StatusOk {
     	return &valv, _returnStatus1
     }
     
     valv.Key = *p1
     
     
     v = &valv
     return v, _syscall.StatusOk
}
func (e *Encoder) Pusher(v *Pusher) (_returnStatus _syscall.Status){
     
     if _, _, _returnStatusP := e.t.PointerCheck(unsafe.Pointer(v), "6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e", uint64(unsafe.Sizeof(*v))); _returnStatusP == _syscall.StatusOk {
     	
       _returnStatus = e.string(v.Username)
       if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
       
       _returnStatus = e.string(v.Key)
       if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
        
     } else {
       return _returnStatusP
     }
     return _syscall.StatusOk
}
func (d *Decoder) Pusher() (v *Pusher, _returnStatus _syscall.Status){
     
     
     var valv Pusher
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, _returnStatus0 := d.string()
     if _returnStatus0 != _syscall.StatusOk {
     	return &valv, _returnStatus0
     }
     
     valv.Username = *p0
     
     
     		
     p1, _returnStatus1 := d.string()
     if _returnStatus1 != _syscall.StatusOk {
     	return &valv, _returnStatus1
     }
     
     valv.Key = *p1
     
     
     v = &valv
     return v, _syscall.StatusOk
}
func (e *Encoder) Message(v *Message) (_returnStatus _syscall.Status){
     
     if _, _, _returnStatusP := e.t.PointerCheck(unsafe.Pointer(v), "37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7", uint64(unsafe.Sizeof(*v))); _returnStatusP == _syscall.StatusOk {
     	
       _returnStatus = e.string(v.Id)
       if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
       
       _returnStatus = e.string(v.Data)
       if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
       
       _returnStatus = e.int64(v.CreatedAt)
       if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
        
     } else {
       return _returnStatusP
     }
     return _syscall.StatusOk
}
func (d *Decoder) Message() (v *Message, _returnStatus _syscall.Status){
     
     
     var valv Message
     d.indexToValue = append(d.indexToValue, &valv)
     
     		
     p0, _returnStatus0 := d.string()
     if _returnStatus0 != _syscall.StatusOk {
     	return &valv, _returnStatus0
     }
     
     valv.Id = *p0
     
     
     		
     p1, _returnStatus1 := d.string()
     if _returnStatus1 != _syscall.StatusOk {
     	return &valv, _returnStatus1
     }
     
     valv.Data = *p1
     
     
     		
     p2, _returnStatus2 := d.int64()
     if _returnStatus2 != _syscall.StatusOk {
     	return &valv, _returnStatus2
     }
     
     valv.CreatedAt = *p2
     
     
     v = &valv
     return v, _syscall.StatusOk
}
func (e *Encoder) PullerInternal(v *Puller) (_returnStatus _syscall.Status){
     
     
     	     _returnStatus = e.string(v.Username)
	     if _returnStatus != _syscall.StatusOk {
	     	return _returnStatus
	     }
     
     	     _returnStatus = e.string(v.Key)
	     if _returnStatus != _syscall.StatusOk {
	     	return _returnStatus
	     }
     
     return _returnStatus
}


func (d *Decoder) PullerInternal() (v *Puller, _returnStatus _syscall.Status){
     
      
     var valv Puller
     
          
     p0, _returnStatus0 := d.string()
     if _returnStatus0 != _syscall.StatusOk {
     	return &valv, _returnStatus0
     }
     
     valv.Username = *p0
     
     
          
     p1, _returnStatus1 := d.string()
     if _returnStatus1 != _syscall.StatusOk {
     	return &valv, _returnStatus1
     }
     
     valv.Key = *p1
     
     
     v = &valv     
     return v, _syscall.StatusOk
}

func (e *Encoder) PusherInternal(v *Pusher) (_returnStatus _syscall.Status){
     
     
     	     _returnStatus = e.string(v.Username)
	     if _returnStatus != _syscall.StatusOk {
	     	return _returnStatus
	     }
     
     	     _returnStatus = e.string(v.Key)
	     if _returnStatus != _syscall.StatusOk {
	     	return _returnStatus
	     }
     
     return _returnStatus
}


func (d *Decoder) PusherInternal() (v *Pusher, _returnStatus _syscall.Status){
     
      
     var valv Pusher
     
          
     p0, _returnStatus0 := d.string()
     if _returnStatus0 != _syscall.StatusOk {
     	return &valv, _returnStatus0
     }
     
     valv.Username = *p0
     
     
          
     p1, _returnStatus1 := d.string()
     if _returnStatus1 != _syscall.StatusOk {
     	return &valv, _returnStatus1
     }
     
     valv.Key = *p1
     
     
     v = &valv     
     return v, _syscall.StatusOk
}

func (e *Encoder) MessageInternal(v *Message) (_returnStatus _syscall.Status){
     
     
     	     _returnStatus = e.string(v.Id)
	     if _returnStatus != _syscall.StatusOk {
	     	return _returnStatus
	     }
     
     	     _returnStatus = e.string(v.Data)
	     if _returnStatus != _syscall.StatusOk {
	     	return _returnStatus
	     }
     
     	     _returnStatus = e.int64(v.CreatedAt)
	     if _returnStatus != _syscall.StatusOk {
	     	return _returnStatus
	     }
     
     return _returnStatus
}


func (d *Decoder) MessageInternal() (v *Message, _returnStatus _syscall.Status){
     
      
     var valv Message
     
          
     p0, _returnStatus0 := d.string()
     if _returnStatus0 != _syscall.StatusOk {
     	return &valv, _returnStatus0
     }
     
     valv.Id = *p0
     
     
          
     p1, _returnStatus1 := d.string()
     if _returnStatus1 != _syscall.StatusOk {
     	return &valv, _returnStatus1
     }
     
     valv.Data = *p1
     
     
          
     p2, _returnStatus2 := d.int64()
     if _returnStatus2 != _syscall.StatusOk {
     	return &valv, _returnStatus2
     }
     
     valv.CreatedAt = *p2
     
     
     v = &valv     
     return v, _syscall.StatusOk
}

func (e *Encoder) Any(a Any) (_returnStatus _syscall.Status){
     switch a.Value.(type) {
     
     case Puller:
     	hashByte, err := hex.DecodeString("18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358")
	if err != nil {
	   return _syscall.StatusFail
	}
     	// hashByte := []byte("18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358")
     	// e.string("18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358")
	_returnStatus = e.SliceOfBytes(hashByte)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     	t := a.Value.(Puller)
	_returnStatus = e.Puller(&t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     
     case Pusher:
     	hashByte, err := hex.DecodeString("6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e")
	if err != nil {
	   return _syscall.StatusFail
	}
     	// hashByte := []byte("6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e")
     	// e.string("6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e")
	_returnStatus = e.SliceOfBytes(hashByte)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     	t := a.Value.(Pusher)
	_returnStatus = e.Pusher(&t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     
     case Message:
     	hashByte, err := hex.DecodeString("37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7")
	if err != nil {
	   return _syscall.StatusFail
	}
     	// hashByte := []byte("37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7")
     	// e.string("37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7")
	_returnStatus = e.SliceOfBytes(hashByte)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     	t := a.Value.(Message)
	_returnStatus = e.Message(&t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     
     case nil:
     	_returnStatus = e.SliceOfBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	_returnStatus = e.uint8(pNIL)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case int8:
     	_returnStatus = e.SliceOfBytes([]byte{0xc2, 0x71, 0x6e, 0x3c, 0x34, 0x13, 0xc0, 0xbe, 0xb3, 0x9e, 0x2e, 0xbc, 0xc7, 0x99, 0x62, 0x3f, 0x28, 0xe9, 0x9d, 0x19, 0x71, 0x6a, 0x0, 0x5b, 0x69, 0x4b, 0x1c, 0xbb, 0x3d, 0x8d, 0xcc, 0x45, 0xbd, 0x51, 0xca, 0x50, 0xc7, 0x3f, 0x8a, 0x64, 0xc8, 0xa0, 0xf0, 0xb, 0x46, 0x9a, 0x87, 0x5, 0xe9, 0x3c, 0xb9, 0x27, 0x91, 0xdf, 0x95, 0x88, 0x1b, 0x2e, 0xfa, 0x9f, 0xc0, 0x1f, 0xcb, 0xad}) 
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(int8)
	_returnStatus = e.int8(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case uint8:
     	_returnStatus = e.SliceOfBytes([]byte{0x9, 0x66, 0x78, 0xbb, 0x4a, 0x86, 0x63, 0x8d, 0x8e, 0xc6, 0x58, 0x8e, 0xcc, 0x2d, 0x89, 0x5a, 0x5c, 0x17, 0xb4, 0x86, 0x37, 0x8c, 0x81, 0xc2, 0xf2, 0xac, 0xf2, 0x48, 0x67, 0x57, 0x82, 0x13, 0x3e, 0x2e, 0x7d, 0x80, 0xb6, 0x66, 0x8d, 0x84, 0xc4, 0x55, 0xf4, 0xc9, 0xe1, 0x33, 0xfc, 0x71, 0xb, 0x77, 0x43, 0x63, 0x6, 0x82, 0x76, 0x63, 0x7, 0xbc, 0xf1, 0x64, 0xdb, 0xd9, 0x5c, 0x93})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(uint8)
	_returnStatus = e.uint8(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case int16:
     	_returnStatus = e.SliceOfBytes([]byte{0x99, 0x7f, 0x69, 0x9c, 0x17, 0xfd, 0x6, 0x74, 0x8c, 0x2d, 0xba, 0xc4, 0x61, 0x1, 0x35, 0x21, 0xf1, 0x97, 0xd8, 0x1e, 0x74, 0x3c, 0x2f, 0x96, 0x56, 0xe6, 0xdc, 0xfc, 0x14, 0x1e, 0xd, 0x83, 0x36, 0xdc, 0x73, 0x36, 0xb0, 0xf4, 0x9e, 0x40, 0x2b, 0xfe, 0x97, 0x6f, 0xfa, 0xa5, 0x27, 0xac, 0xe1, 0xa2, 0x57, 0x2a, 0xae, 0x6d, 0x18, 0x22, 0xe2, 0xdc, 0xd8, 0x79, 0xe0, 0xb6, 0xf6, 0x7e})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(int16)
	_returnStatus = e.int16(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case uint16:
     	_returnStatus = e.SliceOfBytes([]byte{0xa9, 0x1f, 0x92, 0xf4, 0xe9, 0x96, 0xb6, 0xd4, 0xa4, 0xcb, 0x85, 0x8d, 0x11, 0x80, 0x20, 0x10, 0xd6, 0x29, 0xba, 0x29, 0xe1, 0x89, 0x50, 0x2c, 0xa0, 0xf2, 0xcb, 0x1d, 0x86, 0xb8, 0x27, 0x3b, 0x67, 0xf3, 0x35, 0x72, 0xd1, 0x78, 0x15, 0xcc, 0xb0, 0x94, 0x94, 0x6f, 0x2, 0xab, 0x2e, 0x46, 0xcd, 0x74, 0xea, 0xf5, 0x15, 0xa, 0x26, 0xdc, 0x4e, 0xf7, 0xd, 0x9f, 0x3a, 0x9c, 0x6e, 0x55})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(uint16)
	_returnStatus = e.uint16(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case int32:
     	_returnStatus = e.SliceOfBytes([]byte{0xeb, 0xc1, 0x67, 0x8b, 0x6, 0x82, 0x70, 0x21, 0x38, 0xc2, 0xd9, 0x9e, 0x33, 0x22, 0xd1, 0xa8, 0xc7, 0x2e, 0x9b, 0x68, 0xe9, 0x41, 0x12, 0x0, 0x1e, 0x3e, 0x51, 0xa8, 0xf5, 0xd9, 0xfa, 0x34, 0xc, 0x44, 0x9c, 0x6, 0x6d, 0x9d, 0x4c, 0xe7, 0x2a, 0x6, 0xab, 0x75, 0x77, 0x5d, 0xdf, 0x28, 0x34, 0x88, 0x7c, 0x7e, 0x96, 0x97, 0xbb, 0x8a, 0x95, 0xfe, 0x7, 0x65, 0xf7, 0x7c, 0x7e, 0x4c})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(int32)
	_returnStatus = e.int32(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case uint32:
     	_returnStatus = e.SliceOfBytes([]byte{0xce, 0xf4, 0x3a, 0x5, 0xae, 0x67, 0xd9, 0x73, 0xc2, 0xa2, 0x1d, 0xf8, 0xcd, 0xf9, 0xd2, 0xde, 0x69, 0x8d, 0xd, 0xb7, 0x61, 0xb9, 0x51, 0x22, 0x58, 0xed, 0x8f, 0xb1, 0x83, 0xf1, 0x5c, 0xff, 0x5b, 0x84, 0xe2, 0x14, 0xe, 0x10, 0x68, 0x3f, 0x7a, 0xd9, 0xa7, 0x8f, 0x5b, 0xe4, 0x9e, 0x4e, 0x0, 0x7d, 0xcb, 0xfb, 0xd1, 0x69, 0x59, 0x9d, 0xbf, 0x9b, 0x75, 0x65, 0x15, 0x9e, 0x8b, 0x82})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(uint32)
	_returnStatus = e.uint32(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case int64:
     	_returnStatus = e.SliceOfBytes([]byte{0x56, 0x16, 0x70, 0x35, 0xd0, 0x9, 0x18, 0x69, 0xe, 0xae, 0xad, 0x60, 0xd1, 0xee, 0x39, 0xa8, 0x61, 0x45, 0x58, 0x5b, 0x99, 0x20, 0x94, 0x57, 0x1f, 0xb0, 0x48, 0xeb, 0xb2, 0xcf, 0x5c, 0xa5, 0x8d, 0xc7, 0x8e, 0x7e, 0x3c, 0x89, 0xcd, 0x2f, 0xdc, 0xf2, 0x1c, 0x2a, 0xe3, 0xd2, 0x7f, 0x98, 0xc2, 0xad, 0x1c, 0x3d, 0x4e, 0x62, 0xd9, 0xdb, 0xc8, 0xc8, 0x59, 0xc5, 0xd5, 0xc6, 0xed, 0x7a})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(int64)
	_returnStatus = e.int64(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case uint64:
     	_returnStatus = e.SliceOfBytes([]byte{0x11, 0x26, 0xb3, 0xd, 0x51, 0x59, 0x87, 0x5e, 0xd, 0x5b, 0x93, 0xfc, 0x92, 0xf0, 0x78, 0xaa, 0x12, 0xac, 0x93, 0xb8, 0x30, 0x1f, 0x48, 0xe, 0x13, 0x4d, 0x8b, 0xfb, 0x4c, 0x58, 0xfa, 0x3a, 0x69, 0x6a, 0x81, 0x1, 0xc5, 0x47, 0xc1, 0x55, 0x43, 0x95, 0x41, 0xdf, 0x3c, 0x8e, 0xb6, 0x96, 0x4a, 0x3c, 0x88, 0xab, 0x3f, 0x88, 0xed, 0x37, 0x5f, 0x8, 0x4a, 0x41, 0x8e, 0xd5, 0xda, 0x1e})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(uint64)
	_returnStatus = e.uint64(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case float32:
     	_returnStatus = e.SliceOfBytes([]byte{0x42, 0x36, 0xae, 0xd3, 0x62, 0xca, 0x34, 0x75, 0x94, 0x52, 0xf0, 0x5f, 0x44, 0x83, 0x61, 0x75, 0x69, 0x39, 0xcf, 0x69, 0x74, 0x91, 0xee, 0x8d, 0x35, 0x8c, 0xd7, 0xa1, 0x63, 0xf, 0x88, 0x86, 0x6b, 0x52, 0xdd, 0x6d, 0xe1, 0xb2, 0x26, 0xf4, 0x3a, 0x9c, 0x9e, 0xf1, 0x56, 0xd, 0xf1, 0x48, 0x7, 0x39, 0x46, 0xf8, 0xe9, 0xd3, 0xab, 0x86, 0xe0, 0x1c, 0x98, 0xd, 0x17, 0x6b, 0x2, 0x63})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(float32)
	_returnStatus = e.float32(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case float64:
     	_returnStatus = e.SliceOfBytes([]byte{0x65, 0x9b, 0xb2, 0x59, 0x85, 0xe2, 0x60, 0xe7, 0x1e, 0x12, 0x17, 0x3f, 0xc3, 0x1f, 0x20, 0x45, 0x8, 0x9e, 0x7e, 0x11, 0x6b, 0xaa, 0xb3, 0x1e, 0x6d, 0x7d, 0x7a, 0x5b, 0xe3, 0x3d, 0x40, 0xb5, 0x40, 0x6, 0x52, 0x85, 0x37, 0x80, 0x2c, 0xd8, 0x7d, 0x48, 0x67, 0xe3, 0x9a, 0xdd, 0xc9, 0x13, 0x11, 0x2c, 0xa5, 0xcc, 0x5a, 0x33, 0xbc, 0x35, 0x6b, 0x3e, 0xa8, 0x75, 0x93, 0x84, 0xcf, 0x1b})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(float64)
	_returnStatus = e.float64(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case bool:
     	_returnStatus = e.SliceOfBytes([]byte{0x3e, 0x76, 0x6, 0x81, 0x73, 0x27, 0x61, 0xed, 0x16, 0x8e, 0xa3, 0x8, 0xe1, 0x10, 0x11, 0x85, 0xe1, 0xbd, 0x39, 0x2, 0xee, 0x67, 0x60, 0x21, 0x6a, 0x59, 0xc4, 0x7, 0x5a, 0x99, 0xc1, 0x46, 0xb7, 0xcd, 0x98, 0x14, 0xce, 0x14, 0x47, 0xe, 0xb0, 0x80, 0x6d, 0x91, 0x66, 0x50, 0xb0, 0xe5, 0xe, 0x77, 0x6f, 0x53, 0xe5, 0xd1, 0x72, 0x28, 0x1d, 0xd0, 0xe1, 0x70, 0x43, 0xc8, 0x65, 0x39})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(bool)
	_returnStatus = e.bool(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case string:
     	_returnStatus = e.SliceOfBytes([]byte{0x27, 0x57, 0xcb, 0x3c, 0xaf, 0xc3, 0x9a, 0xf4, 0x51, 0xab, 0xb2, 0x69, 0x7b, 0xe7, 0x9b, 0x4a, 0xb6, 0x1d, 0x63, 0xd7, 0x4d, 0x85, 0xb0, 0x41, 0x86, 0x29, 0xde, 0x8c, 0x26, 0x81, 0x1b, 0x52, 0x9f, 0x3f, 0x37, 0x80, 0xd0, 0x15, 0x0, 0x63, 0xff, 0x55, 0xa2, 0xbe, 0xee, 0x74, 0xc4, 0xec, 0x10, 0x2a, 0x2a, 0x27, 0x31, 0xa1, 0xf1, 0xf7, 0xf1, 0xd, 0x47, 0x3a, 0xd1, 0x8a, 0x6a, 0x87})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(string)
	_returnStatus = e.string(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     default:
	//v := reflect.ValueOf(a)
     	//e := fmt.Errorf("Wrong type used as Any " + v.Type().Name())
	return _syscall.StatusInvalidType
     }
     return _syscall.StatusOk
}

func (e *Encoder) AnyInternal(a Any) (_returnStatus _syscall.Status){
     switch a.Value.(type) {
     
     case Puller:
     	hashByte, err := hex.DecodeString("18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358")
	if err != nil {
	   return _syscall.StatusFail
	}
     	// hashByte := []byte("18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358")
     	// e.string("18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358")
	_returnStatus = e.SliceOfBytes(hashByte)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     	t := a.Value.(Puller)
	_returnStatus = e.PullerInternal(&t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     
     case Pusher:
     	hashByte, err := hex.DecodeString("6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e")
	if err != nil {
	   return _syscall.StatusFail
	}
     	// hashByte := []byte("6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e")
     	// e.string("6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e")
	_returnStatus = e.SliceOfBytes(hashByte)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     	t := a.Value.(Pusher)
	_returnStatus = e.PusherInternal(&t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     
     case Message:
     	hashByte, err := hex.DecodeString("37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7")
	if err != nil {
	   return _syscall.StatusFail
	}
     	// hashByte := []byte("37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7")
     	// e.string("37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7")
	_returnStatus = e.SliceOfBytes(hashByte)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     	t := a.Value.(Message)
	_returnStatus = e.MessageInternal(&t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     
     case nil:
     	_returnStatus = e.SliceOfBytes([]byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	_returnStatus = e.uint8(pNIL)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case int8:
     	_returnStatus = e.SliceOfBytes([]byte{0xc2, 0x71, 0x6e, 0x3c, 0x34, 0x13, 0xc0, 0xbe, 0xb3, 0x9e, 0x2e, 0xbc, 0xc7, 0x99, 0x62, 0x3f, 0x28, 0xe9, 0x9d, 0x19, 0x71, 0x6a, 0x0, 0x5b, 0x69, 0x4b, 0x1c, 0xbb, 0x3d, 0x8d, 0xcc, 0x45, 0xbd, 0x51, 0xca, 0x50, 0xc7, 0x3f, 0x8a, 0x64, 0xc8, 0xa0, 0xf0, 0xb, 0x46, 0x9a, 0x87, 0x5, 0xe9, 0x3c, 0xb9, 0x27, 0x91, 0xdf, 0x95, 0x88, 0x1b, 0x2e, 0xfa, 0x9f, 0xc0, 0x1f, 0xcb, 0xad}) 
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(int8)
	_returnStatus = e.int8(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case uint8:
     	_returnStatus = e.SliceOfBytes([]byte{0x9, 0x66, 0x78, 0xbb, 0x4a, 0x86, 0x63, 0x8d, 0x8e, 0xc6, 0x58, 0x8e, 0xcc, 0x2d, 0x89, 0x5a, 0x5c, 0x17, 0xb4, 0x86, 0x37, 0x8c, 0x81, 0xc2, 0xf2, 0xac, 0xf2, 0x48, 0x67, 0x57, 0x82, 0x13, 0x3e, 0x2e, 0x7d, 0x80, 0xb6, 0x66, 0x8d, 0x84, 0xc4, 0x55, 0xf4, 0xc9, 0xe1, 0x33, 0xfc, 0x71, 0xb, 0x77, 0x43, 0x63, 0x6, 0x82, 0x76, 0x63, 0x7, 0xbc, 0xf1, 0x64, 0xdb, 0xd9, 0x5c, 0x93})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(uint8)
	_returnStatus = e.uint8(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case int16:
     	_returnStatus = e.SliceOfBytes([]byte{0x99, 0x7f, 0x69, 0x9c, 0x17, 0xfd, 0x6, 0x74, 0x8c, 0x2d, 0xba, 0xc4, 0x61, 0x1, 0x35, 0x21, 0xf1, 0x97, 0xd8, 0x1e, 0x74, 0x3c, 0x2f, 0x96, 0x56, 0xe6, 0xdc, 0xfc, 0x14, 0x1e, 0xd, 0x83, 0x36, 0xdc, 0x73, 0x36, 0xb0, 0xf4, 0x9e, 0x40, 0x2b, 0xfe, 0x97, 0x6f, 0xfa, 0xa5, 0x27, 0xac, 0xe1, 0xa2, 0x57, 0x2a, 0xae, 0x6d, 0x18, 0x22, 0xe2, 0xdc, 0xd8, 0x79, 0xe0, 0xb6, 0xf6, 0x7e})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(int16)
	_returnStatus = e.int16(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case uint16:
     	_returnStatus = e.SliceOfBytes([]byte{0xa9, 0x1f, 0x92, 0xf4, 0xe9, 0x96, 0xb6, 0xd4, 0xa4, 0xcb, 0x85, 0x8d, 0x11, 0x80, 0x20, 0x10, 0xd6, 0x29, 0xba, 0x29, 0xe1, 0x89, 0x50, 0x2c, 0xa0, 0xf2, 0xcb, 0x1d, 0x86, 0xb8, 0x27, 0x3b, 0x67, 0xf3, 0x35, 0x72, 0xd1, 0x78, 0x15, 0xcc, 0xb0, 0x94, 0x94, 0x6f, 0x2, 0xab, 0x2e, 0x46, 0xcd, 0x74, 0xea, 0xf5, 0x15, 0xa, 0x26, 0xdc, 0x4e, 0xf7, 0xd, 0x9f, 0x3a, 0x9c, 0x6e, 0x55})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(uint16)
	_returnStatus = e.uint16(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case int32:
     	_returnStatus = e.SliceOfBytes([]byte{0xeb, 0xc1, 0x67, 0x8b, 0x6, 0x82, 0x70, 0x21, 0x38, 0xc2, 0xd9, 0x9e, 0x33, 0x22, 0xd1, 0xa8, 0xc7, 0x2e, 0x9b, 0x68, 0xe9, 0x41, 0x12, 0x0, 0x1e, 0x3e, 0x51, 0xa8, 0xf5, 0xd9, 0xfa, 0x34, 0xc, 0x44, 0x9c, 0x6, 0x6d, 0x9d, 0x4c, 0xe7, 0x2a, 0x6, 0xab, 0x75, 0x77, 0x5d, 0xdf, 0x28, 0x34, 0x88, 0x7c, 0x7e, 0x96, 0x97, 0xbb, 0x8a, 0x95, 0xfe, 0x7, 0x65, 0xf7, 0x7c, 0x7e, 0x4c})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(int32)
	_returnStatus = e.int32(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case uint32:
     	_returnStatus = e.SliceOfBytes([]byte{0xce, 0xf4, 0x3a, 0x5, 0xae, 0x67, 0xd9, 0x73, 0xc2, 0xa2, 0x1d, 0xf8, 0xcd, 0xf9, 0xd2, 0xde, 0x69, 0x8d, 0xd, 0xb7, 0x61, 0xb9, 0x51, 0x22, 0x58, 0xed, 0x8f, 0xb1, 0x83, 0xf1, 0x5c, 0xff, 0x5b, 0x84, 0xe2, 0x14, 0xe, 0x10, 0x68, 0x3f, 0x7a, 0xd9, 0xa7, 0x8f, 0x5b, 0xe4, 0x9e, 0x4e, 0x0, 0x7d, 0xcb, 0xfb, 0xd1, 0x69, 0x59, 0x9d, 0xbf, 0x9b, 0x75, 0x65, 0x15, 0x9e, 0x8b, 0x82})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(uint32)
	_returnStatus = e.uint32(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case int64:
     	_returnStatus = e.SliceOfBytes([]byte{0x56, 0x16, 0x70, 0x35, 0xd0, 0x9, 0x18, 0x69, 0xe, 0xae, 0xad, 0x60, 0xd1, 0xee, 0x39, 0xa8, 0x61, 0x45, 0x58, 0x5b, 0x99, 0x20, 0x94, 0x57, 0x1f, 0xb0, 0x48, 0xeb, 0xb2, 0xcf, 0x5c, 0xa5, 0x8d, 0xc7, 0x8e, 0x7e, 0x3c, 0x89, 0xcd, 0x2f, 0xdc, 0xf2, 0x1c, 0x2a, 0xe3, 0xd2, 0x7f, 0x98, 0xc2, 0xad, 0x1c, 0x3d, 0x4e, 0x62, 0xd9, 0xdb, 0xc8, 0xc8, 0x59, 0xc5, 0xd5, 0xc6, 0xed, 0x7a})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(int64)
	_returnStatus = e.int64(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case uint64:
     	_returnStatus = e.SliceOfBytes([]byte{0x11, 0x26, 0xb3, 0xd, 0x51, 0x59, 0x87, 0x5e, 0xd, 0x5b, 0x93, 0xfc, 0x92, 0xf0, 0x78, 0xaa, 0x12, 0xac, 0x93, 0xb8, 0x30, 0x1f, 0x48, 0xe, 0x13, 0x4d, 0x8b, 0xfb, 0x4c, 0x58, 0xfa, 0x3a, 0x69, 0x6a, 0x81, 0x1, 0xc5, 0x47, 0xc1, 0x55, 0x43, 0x95, 0x41, 0xdf, 0x3c, 0x8e, 0xb6, 0x96, 0x4a, 0x3c, 0x88, 0xab, 0x3f, 0x88, 0xed, 0x37, 0x5f, 0x8, 0x4a, 0x41, 0x8e, 0xd5, 0xda, 0x1e})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(uint64)
	_returnStatus = e.uint64(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case float32:
     	_returnStatus = e.SliceOfBytes([]byte{0x42, 0x36, 0xae, 0xd3, 0x62, 0xca, 0x34, 0x75, 0x94, 0x52, 0xf0, 0x5f, 0x44, 0x83, 0x61, 0x75, 0x69, 0x39, 0xcf, 0x69, 0x74, 0x91, 0xee, 0x8d, 0x35, 0x8c, 0xd7, 0xa1, 0x63, 0xf, 0x88, 0x86, 0x6b, 0x52, 0xdd, 0x6d, 0xe1, 0xb2, 0x26, 0xf4, 0x3a, 0x9c, 0x9e, 0xf1, 0x56, 0xd, 0xf1, 0x48, 0x7, 0x39, 0x46, 0xf8, 0xe9, 0xd3, 0xab, 0x86, 0xe0, 0x1c, 0x98, 0xd, 0x17, 0x6b, 0x2, 0x63})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(float32)
	_returnStatus = e.float32(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case float64:
     	_returnStatus = e.SliceOfBytes([]byte{0x65, 0x9b, 0xb2, 0x59, 0x85, 0xe2, 0x60, 0xe7, 0x1e, 0x12, 0x17, 0x3f, 0xc3, 0x1f, 0x20, 0x45, 0x8, 0x9e, 0x7e, 0x11, 0x6b, 0xaa, 0xb3, 0x1e, 0x6d, 0x7d, 0x7a, 0x5b, 0xe3, 0x3d, 0x40, 0xb5, 0x40, 0x6, 0x52, 0x85, 0x37, 0x80, 0x2c, 0xd8, 0x7d, 0x48, 0x67, 0xe3, 0x9a, 0xdd, 0xc9, 0x13, 0x11, 0x2c, 0xa5, 0xcc, 0x5a, 0x33, 0xbc, 0x35, 0x6b, 0x3e, 0xa8, 0x75, 0x93, 0x84, 0xcf, 0x1b})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(float64)
	_returnStatus = e.float64(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case bool:
     	_returnStatus = e.SliceOfBytes([]byte{0x3e, 0x76, 0x6, 0x81, 0x73, 0x27, 0x61, 0xed, 0x16, 0x8e, 0xa3, 0x8, 0xe1, 0x10, 0x11, 0x85, 0xe1, 0xbd, 0x39, 0x2, 0xee, 0x67, 0x60, 0x21, 0x6a, 0x59, 0xc4, 0x7, 0x5a, 0x99, 0xc1, 0x46, 0xb7, 0xcd, 0x98, 0x14, 0xce, 0x14, 0x47, 0xe, 0xb0, 0x80, 0x6d, 0x91, 0x66, 0x50, 0xb0, 0xe5, 0xe, 0x77, 0x6f, 0x53, 0xe5, 0xd1, 0x72, 0x28, 0x1d, 0xd0, 0xe1, 0x70, 0x43, 0xc8, 0x65, 0x39})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(bool)
	_returnStatus = e.bool(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     case string:
     	_returnStatus = e.SliceOfBytes([]byte{0x27, 0x57, 0xcb, 0x3c, 0xaf, 0xc3, 0x9a, 0xf4, 0x51, 0xab, 0xb2, 0x69, 0x7b, 0xe7, 0x9b, 0x4a, 0xb6, 0x1d, 0x63, 0xd7, 0x4d, 0x85, 0xb0, 0x41, 0x86, 0x29, 0xde, 0x8c, 0x26, 0x81, 0x1b, 0x52, 0x9f, 0x3f, 0x37, 0x80, 0xd0, 0x15, 0x0, 0x63, 0xff, 0x55, 0xa2, 0xbe, 0xee, 0x74, 0xc4, 0xec, 0x10, 0x2a, 0x2a, 0x27, 0x31, 0xa1, 0xf1, 0xf7, 0xf1, 0xd, 0x47, 0x3a, 0xd1, 0x8a, 0x6a, 0x87})
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	t := a.Value.(string)
	_returnStatus = e.string(t)
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
     default:
	// v := reflect.ValueOf(a)
	// e := fmt.Errorf("Wrong type used as Any " + v.Type().Name())
	return _syscall.StatusInvalidType
     }
     return _syscall.StatusOk
}
func (d *Decoder) Any() (retValue *Any, _returnStatus _syscall.Status) {
     l, _returnStatus := d.uint32()
     if _returnStatus != _syscall.StatusOk {
     	return nil, _returnStatus
     }
     hashValue := make([]byte, *l)
     _returnStatus = d.SliceOfBytes(hashValue, *l)
     if _returnStatus != _syscall.StatusOk {
     	return nil, _returnStatus
     }
     d.indexToValue = append(d.indexToValue, retValue)	
     index := len(d.indexToValue) - 1
     
     encodedHash := []byte(hex.EncodeToString(hashValue))
     
     switch {
     
     case bytes.Equal(encodedHash, []byte("18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358")):
     	
	p, _returnStatus := d.Puller()	
	if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(encodedHash, []byte("6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e")):
     	
	p, _returnStatus := d.Pusher()	
	if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(encodedHash, []byte("37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7")):
     	
	p, _returnStatus := d.Message()	
	if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(hashValue, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}):
     	  _, _returnStatus := d.uint8()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	   }
     	  retValue.Value = nil
     case bytes.Equal(hashValue, []byte{0xc2, 0x71, 0x6e, 0x3c, 0x34, 0x13, 0xc0, 0xbe, 0xb3, 0x9e, 0x2e, 0xbc, 0xc7, 0x99, 0x62, 0x3f, 0x28, 0xe9, 0x9d, 0x19, 0x71, 0x6a, 0x0, 0x5b, 0x69, 0x4b, 0x1c, 0xbb, 0x3d, 0x8d, 0xcc, 0x45, 0xbd, 0x51, 0xca, 0x50, 0xc7, 0x3f, 0x8a, 0x64, 0xc8, 0xa0, 0xf0, 0xb, 0x46, 0x9a, 0x87, 0x5, 0xe9, 0x3c, 0xb9, 0x27, 0x91, 0xdf, 0x95, 0x88, 0x1b, 0x2e, 0xfa, 0x9f, 0xc0, 0x1f, 0xcb, 0xad}):
     	  p, _returnStatus := d.int8()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x9, 0x66, 0x78, 0xbb, 0x4a, 0x86, 0x63, 0x8d, 0x8e, 0xc6, 0x58, 0x8e, 0xcc, 0x2d, 0x89, 0x5a, 0x5c, 0x17, 0xb4, 0x86, 0x37, 0x8c, 0x81, 0xc2, 0xf2, 0xac, 0xf2, 0x48, 0x67, 0x57, 0x82, 0x13, 0x3e, 0x2e, 0x7d, 0x80, 0xb6, 0x66, 0x8d, 0x84, 0xc4, 0x55, 0xf4, 0xc9, 0xe1, 0x33, 0xfc, 0x71, 0xb, 0x77, 0x43, 0x63, 0x6, 0x82, 0x76, 0x63, 0x7, 0xbc, 0xf1, 0x64, 0xdb, 0xd9, 0x5c, 0x93}):
     	  p, _returnStatus := d.uint8()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x99, 0x7f, 0x69, 0x9c, 0x17, 0xfd, 0x6, 0x74, 0x8c, 0x2d, 0xba, 0xc4, 0x61, 0x1, 0x35, 0x21, 0xf1, 0x97, 0xd8, 0x1e, 0x74, 0x3c, 0x2f, 0x96, 0x56, 0xe6, 0xdc, 0xfc, 0x14, 0x1e, 0xd, 0x83, 0x36, 0xdc, 0x73, 0x36, 0xb0, 0xf4, 0x9e, 0x40, 0x2b, 0xfe, 0x97, 0x6f, 0xfa, 0xa5, 0x27, 0xac, 0xe1, 0xa2, 0x57, 0x2a, 0xae, 0x6d, 0x18, 0x22, 0xe2, 0xdc, 0xd8, 0x79, 0xe0, 0xb6, 0xf6, 0x7e}):
     	  p, _returnStatus := d.int16()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xa9, 0x1f, 0x92, 0xf4, 0xe9, 0x96, 0xb6, 0xd4, 0xa4, 0xcb, 0x85, 0x8d, 0x11, 0x80, 0x20, 0x10, 0xd6, 0x29, 0xba, 0x29, 0xe1, 0x89, 0x50, 0x2c, 0xa0, 0xf2, 0xcb, 0x1d, 0x86, 0xb8, 0x27, 0x3b, 0x67, 0xf3, 0x35, 0x72, 0xd1, 0x78, 0x15, 0xcc, 0xb0, 0x94, 0x94, 0x6f, 0x2, 0xab, 0x2e, 0x46, 0xcd, 0x74, 0xea, 0xf5, 0x15, 0xa, 0x26, 0xdc, 0x4e, 0xf7, 0xd, 0x9f, 0x3a, 0x9c, 0x6e, 0x55}):
     	  p, _returnStatus := d.uint16()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xeb, 0xc1, 0x67, 0x8b, 0x6, 0x82, 0x70, 0x21, 0x38, 0xc2, 0xd9, 0x9e, 0x33, 0x22, 0xd1, 0xa8, 0xc7, 0x2e, 0x9b, 0x68, 0xe9, 0x41, 0x12, 0x0, 0x1e, 0x3e, 0x51, 0xa8, 0xf5, 0xd9, 0xfa, 0x34, 0xc, 0x44, 0x9c, 0x6, 0x6d, 0x9d, 0x4c, 0xe7, 0x2a, 0x6, 0xab, 0x75, 0x77, 0x5d, 0xdf, 0x28, 0x34, 0x88, 0x7c, 0x7e, 0x96, 0x97, 0xbb, 0x8a, 0x95, 0xfe, 0x7, 0x65, 0xf7, 0x7c, 0x7e, 0x4c}):
     	  p, _returnStatus := d.int32()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xce, 0xf4, 0x3a, 0x5, 0xae, 0x67, 0xd9, 0x73, 0xc2, 0xa2, 0x1d, 0xf8, 0xcd, 0xf9, 0xd2, 0xde, 0x69, 0x8d, 0xd, 0xb7, 0x61, 0xb9, 0x51, 0x22, 0x58, 0xed, 0x8f, 0xb1, 0x83, 0xf1, 0x5c, 0xff, 0x5b, 0x84, 0xe2, 0x14, 0xe, 0x10, 0x68, 0x3f, 0x7a, 0xd9, 0xa7, 0x8f, 0x5b, 0xe4, 0x9e, 0x4e, 0x0, 0x7d, 0xcb, 0xfb, 0xd1, 0x69, 0x59, 0x9d, 0xbf, 0x9b, 0x75, 0x65, 0x15, 0x9e, 0x8b, 0x82}):
     	  p, _returnStatus := d.uint32()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x56, 0x16, 0x70, 0x35, 0xd0, 0x9, 0x18, 0x69, 0xe, 0xae, 0xad, 0x60, 0xd1, 0xee, 0x39, 0xa8, 0x61, 0x45, 0x58, 0x5b, 0x99, 0x20, 0x94, 0x57, 0x1f, 0xb0, 0x48, 0xeb, 0xb2, 0xcf, 0x5c, 0xa5, 0x8d, 0xc7, 0x8e, 0x7e, 0x3c, 0x89, 0xcd, 0x2f, 0xdc, 0xf2, 0x1c, 0x2a, 0xe3, 0xd2, 0x7f, 0x98, 0xc2, 0xad, 0x1c, 0x3d, 0x4e, 0x62, 0xd9, 0xdb, 0xc8, 0xc8, 0x59, 0xc5, 0xd5, 0xc6, 0xed, 0x7a}):
     	  p, _returnStatus := d.int64()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x11, 0x26, 0xb3, 0xd, 0x51, 0x59, 0x87, 0x5e, 0xd, 0x5b, 0x93, 0xfc, 0x92, 0xf0, 0x78, 0xaa, 0x12, 0xac, 0x93, 0xb8, 0x30, 0x1f, 0x48, 0xe, 0x13, 0x4d, 0x8b, 0xfb, 0x4c, 0x58, 0xfa, 0x3a, 0x69, 0x6a, 0x81, 0x1, 0xc5, 0x47, 0xc1, 0x55, 0x43, 0x95, 0x41, 0xdf, 0x3c, 0x8e, 0xb6, 0x96, 0x4a, 0x3c, 0x88, 0xab, 0x3f, 0x88, 0xed, 0x37, 0x5f, 0x8, 0x4a, 0x41, 0x8e, 0xd5, 0xda, 0x1e}):
     	  p, _returnStatus := d.uint64()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x42, 0x36, 0xae, 0xd3, 0x62, 0xca, 0x34, 0x75, 0x94, 0x52, 0xf0, 0x5f, 0x44, 0x83, 0x61, 0x75, 0x69, 0x39, 0xcf, 0x69, 0x74, 0x91, 0xee, 0x8d, 0x35, 0x8c, 0xd7, 0xa1, 0x63, 0xf, 0x88, 0x86, 0x6b, 0x52, 0xdd, 0x6d, 0xe1, 0xb2, 0x26, 0xf4, 0x3a, 0x9c, 0x9e, 0xf1, 0x56, 0xd, 0xf1, 0x48, 0x7, 0x39, 0x46, 0xf8, 0xe9, 0xd3, 0xab, 0x86, 0xe0, 0x1c, 0x98, 0xd, 0x17, 0x6b, 0x2, 0x63}):
     	  p, _returnStatus := d.float32()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x65, 0x9b, 0xb2, 0x59, 0x85, 0xe2, 0x60, 0xe7, 0x1e, 0x12, 0x17, 0x3f, 0xc3, 0x1f, 0x20, 0x45, 0x8, 0x9e, 0x7e, 0x11, 0x6b, 0xaa, 0xb3, 0x1e, 0x6d, 0x7d, 0x7a, 0x5b, 0xe3, 0x3d, 0x40, 0xb5, 0x40, 0x6, 0x52, 0x85, 0x37, 0x80, 0x2c, 0xd8, 0x7d, 0x48, 0x67, 0xe3, 0x9a, 0xdd, 0xc9, 0x13, 0x11, 0x2c, 0xa5, 0xcc, 0x5a, 0x33, 0xbc, 0x35, 0x6b, 0x3e, 0xa8, 0x75, 0x93, 0x84, 0xcf, 0x1b}):
     	  p, _returnStatus := d.float64()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x3e, 0x76, 0x6, 0x81, 0x73, 0x27, 0x61, 0xed, 0x16, 0x8e, 0xa3, 0x8, 0xe1, 0x10, 0x11, 0x85, 0xe1, 0xbd, 0x39, 0x2, 0xee, 0x67, 0x60, 0x21, 0x6a, 0x59, 0xc4, 0x7, 0x5a, 0x99, 0xc1, 0x46, 0xb7, 0xcd, 0x98, 0x14, 0xce, 0x14, 0x47, 0xe, 0xb0, 0x80, 0x6d, 0x91, 0x66, 0x50, 0xb0, 0xe5, 0xe, 0x77, 0x6f, 0x53, 0xe5, 0xd1, 0x72, 0x28, 0x1d, 0xd0, 0xe1, 0x70, 0x43, 0xc8, 0x65, 0x39}):
     	  p, _returnStatus := d.bool()
	  if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x27, 0x57, 0xcb, 0x3c, 0xaf, 0xc3, 0x9a, 0xf4, 0x51, 0xab, 0xb2, 0x69, 0x7b, 0xe7, 0x9b, 0x4a, 0xb6, 0x1d, 0x63, 0xd7, 0x4d, 0x85, 0xb0, 0x41, 0x86, 0x29, 0xde, 0x8c, 0x26, 0x81, 0x1b, 0x52, 0x9f, 0x3f, 0x37, 0x80, 0xd0, 0x15, 0x0, 0x63, 0xff, 0x55, 0xa2, 0xbe, 0xee, 0x74, 0xc4, 0xec, 0x10, 0x2a, 0x2a, 0x27, 0x31, 0xa1, 0xf1, 0xf7, 0xf1, 0xd, 0x47, 0x3a, 0xd1, 0x8a, 0x6a, 0x87}):
     	  p, _returnStatus := d.string()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
     	  }
     	retValue.Value = *(p)
     default:
     //	e := fmt.Errorf("Decoder::Any(): Wrong type used as Any %s", hashValue)
	return retValue, _syscall.StatusInvalidType
     }
     d.indexToValue[index] = retValue		
     return retValue, _syscall.StatusOk
}

func (d *Decoder) AnyInternal() (retValue *Any, _returnStatus _syscall.Status) {
     // hashValue := d.string()
     l, _returnStatus := d.uint32()
     if _returnStatus != _syscall.StatusOk {
     	return nil, _returnStatus
     }
     hashValue := make([]byte, *l)
     _returnStatus = d.SliceOfBytes(hashValue, *l)
     if _returnStatus != _syscall.StatusOk {
     	return nil, _returnStatus
     }

     
     encodedHash := []byte(hex.EncodeToString(hashValue))     
     

     switch {
     
     case bytes.Equal(encodedHash, []byte("18089530c02145cd8e3fe65d186aa98749ffdc0805cd1536972c9af2a14cea5114792a8b31b95198e537178f1c215e335799486054d25f471fb2a30d07e47358")):
     	
	p, _returnStatus := d.PullerInternal()
	if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(encodedHash, []byte("6f563f886c954072f7603e58a16fe55357b46258b51042f147c1a8a7c7091bbbfe789a8aadd295d40d49b6b6d3e09609bc5597a5a8dc9337ef48fc3a936c6f3e")):
     	
	p, _returnStatus := d.PusherInternal()
	if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(encodedHash, []byte("37083d2293625074bba48abf70ee781e2503aae74e88513a996c8cf480d82c5cafc5be8bf2664e3c1cb4b7e8eb8e93382b86df55b78ecf7f19e1c321e95d77c7")):
     	
	p, _returnStatus := d.MessageInternal()
	if _returnStatus != _syscall.StatusOk {
     	   return nil, _returnStatus
     	}
     	retValue.Value = *(p)
     
     case bytes.Equal(hashValue, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}):
     _, _returnStatus := d.uint8()
     if _returnStatus != _syscall.StatusOk {
     	return nil, _returnStatus
     }
     retValue.Value = nil
     case bytes.Equal(hashValue, []byte{0xc2, 0x71, 0x6e, 0x3c, 0x34, 0x13, 0xc0, 0xbe, 0xb3, 0x9e, 0x2e, 0xbc, 0xc7, 0x99, 0x62, 0x3f, 0x28, 0xe9, 0x9d, 0x19, 0x71, 0x6a, 0x0, 0x5b, 0x69, 0x4b, 0x1c, 0xbb, 0x3d, 0x8d, 0xcc, 0x45, 0xbd, 0x51, 0xca, 0x50, 0xc7, 0x3f, 0x8a, 0x64, 0xc8, 0xa0, 0xf0, 0xb, 0x46, 0x9a, 0x87, 0x5, 0xe9, 0x3c, 0xb9, 0x27, 0x91, 0xdf, 0x95, 0x88, 0x1b, 0x2e, 0xfa, 0x9f, 0xc0, 0x1f, 0xcb, 0xad}):
     	  p, _returnStatus := d.int8()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x9, 0x66, 0x78, 0xbb, 0x4a, 0x86, 0x63, 0x8d, 0x8e, 0xc6, 0x58, 0x8e, 0xcc, 0x2d, 0x89, 0x5a, 0x5c, 0x17, 0xb4, 0x86, 0x37, 0x8c, 0x81, 0xc2, 0xf2, 0xac, 0xf2, 0x48, 0x67, 0x57, 0x82, 0x13, 0x3e, 0x2e, 0x7d, 0x80, 0xb6, 0x66, 0x8d, 0x84, 0xc4, 0x55, 0xf4, 0xc9, 0xe1, 0x33, 0xfc, 0x71, 0xb, 0x77, 0x43, 0x63, 0x6, 0x82, 0x76, 0x63, 0x7, 0xbc, 0xf1, 0x64, 0xdb, 0xd9, 0x5c, 0x93}):
     	  p, _returnStatus := d.uint8()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x99, 0x7f, 0x69, 0x9c, 0x17, 0xfd, 0x6, 0x74, 0x8c, 0x2d, 0xba, 0xc4, 0x61, 0x1, 0x35, 0x21, 0xf1, 0x97, 0xd8, 0x1e, 0x74, 0x3c, 0x2f, 0x96, 0x56, 0xe6, 0xdc, 0xfc, 0x14, 0x1e, 0xd, 0x83, 0x36, 0xdc, 0x73, 0x36, 0xb0, 0xf4, 0x9e, 0x40, 0x2b, 0xfe, 0x97, 0x6f, 0xfa, 0xa5, 0x27, 0xac, 0xe1, 0xa2, 0x57, 0x2a, 0xae, 0x6d, 0x18, 0x22, 0xe2, 0xdc, 0xd8, 0x79, 0xe0, 0xb6, 0xf6, 0x7e}):
     	  p, _returnStatus := d.int16()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xa9, 0x1f, 0x92, 0xf4, 0xe9, 0x96, 0xb6, 0xd4, 0xa4, 0xcb, 0x85, 0x8d, 0x11, 0x80, 0x20, 0x10, 0xd6, 0x29, 0xba, 0x29, 0xe1, 0x89, 0x50, 0x2c, 0xa0, 0xf2, 0xcb, 0x1d, 0x86, 0xb8, 0x27, 0x3b, 0x67, 0xf3, 0x35, 0x72, 0xd1, 0x78, 0x15, 0xcc, 0xb0, 0x94, 0x94, 0x6f, 0x2, 0xab, 0x2e, 0x46, 0xcd, 0x74, 0xea, 0xf5, 0x15, 0xa, 0x26, 0xdc, 0x4e, 0xf7, 0xd, 0x9f, 0x3a, 0x9c, 0x6e, 0x55}):
     	  p, _returnStatus := d.uint16()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xeb, 0xc1, 0x67, 0x8b, 0x6, 0x82, 0x70, 0x21, 0x38, 0xc2, 0xd9, 0x9e, 0x33, 0x22, 0xd1, 0xa8, 0xc7, 0x2e, 0x9b, 0x68, 0xe9, 0x41, 0x12, 0x0, 0x1e, 0x3e, 0x51, 0xa8, 0xf5, 0xd9, 0xfa, 0x34, 0xc, 0x44, 0x9c, 0x6, 0x6d, 0x9d, 0x4c, 0xe7, 0x2a, 0x6, 0xab, 0x75, 0x77, 0x5d, 0xdf, 0x28, 0x34, 0x88, 0x7c, 0x7e, 0x96, 0x97, 0xbb, 0x8a, 0x95, 0xfe, 0x7, 0x65, 0xf7, 0x7c, 0x7e, 0x4c}):
     	  p, _returnStatus := d.int32()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0xce, 0xf4, 0x3a, 0x5, 0xae, 0x67, 0xd9, 0x73, 0xc2, 0xa2, 0x1d, 0xf8, 0xcd, 0xf9, 0xd2, 0xde, 0x69, 0x8d, 0xd, 0xb7, 0x61, 0xb9, 0x51, 0x22, 0x58, 0xed, 0x8f, 0xb1, 0x83, 0xf1, 0x5c, 0xff, 0x5b, 0x84, 0xe2, 0x14, 0xe, 0x10, 0x68, 0x3f, 0x7a, 0xd9, 0xa7, 0x8f, 0x5b, 0xe4, 0x9e, 0x4e, 0x0, 0x7d, 0xcb, 0xfb, 0xd1, 0x69, 0x59, 0x9d, 0xbf, 0x9b, 0x75, 0x65, 0x15, 0x9e, 0x8b, 0x82}):
     	  p, _returnStatus := d.uint32()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x56, 0x16, 0x70, 0x35, 0xd0, 0x9, 0x18, 0x69, 0xe, 0xae, 0xad, 0x60, 0xd1, 0xee, 0x39, 0xa8, 0x61, 0x45, 0x58, 0x5b, 0x99, 0x20, 0x94, 0x57, 0x1f, 0xb0, 0x48, 0xeb, 0xb2, 0xcf, 0x5c, 0xa5, 0x8d, 0xc7, 0x8e, 0x7e, 0x3c, 0x89, 0xcd, 0x2f, 0xdc, 0xf2, 0x1c, 0x2a, 0xe3, 0xd2, 0x7f, 0x98, 0xc2, 0xad, 0x1c, 0x3d, 0x4e, 0x62, 0xd9, 0xdb, 0xc8, 0xc8, 0x59, 0xc5, 0xd5, 0xc6, 0xed, 0x7a}):
     	  p, _returnStatus := d.int64()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x11, 0x26, 0xb3, 0xd, 0x51, 0x59, 0x87, 0x5e, 0xd, 0x5b, 0x93, 0xfc, 0x92, 0xf0, 0x78, 0xaa, 0x12, 0xac, 0x93, 0xb8, 0x30, 0x1f, 0x48, 0xe, 0x13, 0x4d, 0x8b, 0xfb, 0x4c, 0x58, 0xfa, 0x3a, 0x69, 0x6a, 0x81, 0x1, 0xc5, 0x47, 0xc1, 0x55, 0x43, 0x95, 0x41, 0xdf, 0x3c, 0x8e, 0xb6, 0x96, 0x4a, 0x3c, 0x88, 0xab, 0x3f, 0x88, 0xed, 0x37, 0x5f, 0x8, 0x4a, 0x41, 0x8e, 0xd5, 0xda, 0x1e}):
     	  p, _returnStatus := d.uint64()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x42, 0x36, 0xae, 0xd3, 0x62, 0xca, 0x34, 0x75, 0x94, 0x52, 0xf0, 0x5f, 0x44, 0x83, 0x61, 0x75, 0x69, 0x39, 0xcf, 0x69, 0x74, 0x91, 0xee, 0x8d, 0x35, 0x8c, 0xd7, 0xa1, 0x63, 0xf, 0x88, 0x86, 0x6b, 0x52, 0xdd, 0x6d, 0xe1, 0xb2, 0x26, 0xf4, 0x3a, 0x9c, 0x9e, 0xf1, 0x56, 0xd, 0xf1, 0x48, 0x7, 0x39, 0x46, 0xf8, 0xe9, 0xd3, 0xab, 0x86, 0xe0, 0x1c, 0x98, 0xd, 0x17, 0x6b, 0x2, 0x63}):
     	  p, _returnStatus := d.float32()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x65, 0x9b, 0xb2, 0x59, 0x85, 0xe2, 0x60, 0xe7, 0x1e, 0x12, 0x17, 0x3f, 0xc3, 0x1f, 0x20, 0x45, 0x8, 0x9e, 0x7e, 0x11, 0x6b, 0xaa, 0xb3, 0x1e, 0x6d, 0x7d, 0x7a, 0x5b, 0xe3, 0x3d, 0x40, 0xb5, 0x40, 0x6, 0x52, 0x85, 0x37, 0x80, 0x2c, 0xd8, 0x7d, 0x48, 0x67, 0xe3, 0x9a, 0xdd, 0xc9, 0x13, 0x11, 0x2c, 0xa5, 0xcc, 0x5a, 0x33, 0xbc, 0x35, 0x6b, 0x3e, 0xa8, 0x75, 0x93, 0x84, 0xcf, 0x1b}):
     	  p, _returnStatus := d.float64()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x3e, 0x76, 0x6, 0x81, 0x73, 0x27, 0x61, 0xed, 0x16, 0x8e, 0xa3, 0x8, 0xe1, 0x10, 0x11, 0x85, 0xe1, 0xbd, 0x39, 0x2, 0xee, 0x67, 0x60, 0x21, 0x6a, 0x59, 0xc4, 0x7, 0x5a, 0x99, 0xc1, 0x46, 0xb7, 0xcd, 0x98, 0x14, 0xce, 0x14, 0x47, 0xe, 0xb0, 0x80, 0x6d, 0x91, 0x66, 0x50, 0xb0, 0xe5, 0xe, 0x77, 0x6f, 0x53, 0xe5, 0xd1, 0x72, 0x28, 0x1d, 0xd0, 0xe1, 0x70, 0x43, 0xc8, 0x65, 0x39}):
     	  p, _returnStatus := d.bool()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     case bytes.Equal(hashValue, []byte{0x27, 0x57, 0xcb, 0x3c, 0xaf, 0xc3, 0x9a, 0xf4, 0x51, 0xab, 0xb2, 0x69, 0x7b, 0xe7, 0x9b, 0x4a, 0xb6, 0x1d, 0x63, 0xd7, 0x4d, 0x85, 0xb0, 0x41, 0x86, 0x29, 0xde, 0x8c, 0x26, 0x81, 0x1b, 0x52, 0x9f, 0x3f, 0x37, 0x80, 0xd0, 0x15, 0x0, 0x63, 0xff, 0x55, 0xa2, 0xbe, 0xee, 0x74, 0xc4, 0xec, 0x10, 0x2a, 0x2a, 0x27, 0x31, 0xa1, 0xf1, 0xf7, 0xf1, 0xd, 0x47, 0x3a, 0xd1, 0x8a, 0x6a, 0x87}):
     	  p, _returnStatus := d.string()
	  if _returnStatus != _syscall.StatusOk {
     	     return nil, _returnStatus
	  }
     	retValue.Value = *(p)
     default:
     	// e := fmt.Errorf("Decoder::AnyInternal(): Wrong type used as Any %s", hashValue)
	return retValue, _syscall.StatusInvalidType
     }
     return retValue, _syscall.StatusOk
}


type Any struct{Value interface{}}

const(
	pNIL = 0
	pIDX = 1
	pVAL = 2
	BufSize = 1024 * 50
)




type Encoder struct {
	w io.Writer
	buf []byte
	t *TypeTree
	m []interface{} // this is for storing maps
	curPos int
	bufSpace uint32
	bufStart uint32
	bufLen uint32
	count uint32
}
type Decoder struct {
	r io.Reader
	buf []byte
	indexToValue []interface{}
	m []interface{} // this is for storing maps
	curPos uint32
	bufStart uint32
	bufLen uint32
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w:w, buf:make([]byte, BufSize), t: NewTypeTree(), curPos:4, bufSpace:BufSize, bufStart:0, bufLen:0, count:0}
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r:r, buf:make([]byte, BufSize), curPos:0, bufStart:0, bufLen:0}
}

func (e *Encoder) MapCheck(t interface{}) (int, bool) {
	for index, entry := range e.m {
		if entry == t {
			return index, true
		}
	}
	return -1, false
}

func (e *Encoder) isEnoughSpace(length uint32) bool {
     if length <= e.bufSpace {
     	return true
     }    
     return false
}

func (d *Decoder) isEnoughData(length uint32) bool {
     if length <= d.bufLen {
     	return true
     }
     return false
}


func (d *Decoder) readAtLeast(length uint32) _syscall.Status {
     if d.bufLen > 0 {
     	copy(d.buf, d.buf[d.bufStart:(d.bufStart + d.bufLen)])
     }
     n, err := io.ReadAtLeast(d.r, d.buf[d.bufLen:], int(length))
     if err !=  nil {
     	return _syscall.StatusFail
     }
     d.bufLen += uint32(n)
     d.bufStart = 0
     return _syscall.StatusOk
}


// IsEOF tries to read a byte from the stream
// and returns true if this fails. If it reads successful
// the byte read will be put into the decoder buffer and false
// is returned
func (d *Decoder) IsEOF() bool {
	return d.readAtLeast(1) != _syscall.StatusOk
}


func (e *Encoder) byte(b byte) (_syscall.Status){
     return e.uint8(uint8(b))
     
}

func (e *Encoder) Byte(b Byte) (_syscall.Status) {
	return e.byte(byte(b))
}

func (d *Decoder) byte() (b *byte, _returnStatus _syscall.Status) {
     v, _returnStatus := d.uint8()
     if _returnStatus != _syscall.StatusOk {
        return nil, _returnStatus
     }
     value := byte(*v)
     return &value, _returnStatus
}

func (d *Decoder) Byte() (v *Byte, _returnStatus _syscall.Status) {
     b, _returnStatus := d.byte()
     if _returnStatus != _syscall.StatusOk {
          return nil, _returnStatus
     }
     B := Byte(*b)
     return &B, _returnStatus
}

func (e *Encoder) uint8(u uint8) _syscall.Status{
	if !e.isEnoughSpace(1) {
	   _returnStatus := e.Flush() 
	   if _returnStatus != _syscall.StatusOk {
	      return _returnStatus
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen++
	e.bufSpace--
	return _syscall.StatusOk
}

func (e *Encoder) Uint8(u Uint8) (_syscall.Status){
     return e.uint8(uint8(u))
}

func (e *Encoder) int8(u int8) (_syscall.Status){
     return e.int8(u)
}

func (e *Encoder) Int8(u Int8) (_syscall.Status){
     return e.uint8(uint8(u))
}

func (e *Encoder) Int8Boxed(u Int8) (_syscall.Status){
     return e.int8(int8(u))
}

func (d *Decoder) uint8() (w *uint8, _returnStatus _syscall.Status) {
	if !d.isEnoughData(1) {
	   _returnStatus = d.readAtLeast(1)	   
	   if _returnStatus != _syscall.StatusOk {
	      return nil, _returnStatus
	   }
	}
	v := uint8(d.buf[d.bufStart])
	d.bufStart = (d.bufStart + 1) % BufSize
	d.bufLen -= 1
	return &v, _syscall.StatusOk
}

func (d *Decoder) Uint8() (v *Uint8, _returnStatus _syscall.Status) {
     w, _returnStatus := d.uint8()
     if _returnStatus != _syscall.StatusOk {
          return nil, _returnStatus
     }
     W := Uint8(*w)
     return &W, _returnStatus
}

func (d *Decoder) int8() (w *int8, _returnStatus _syscall.Status) {
     	v, _returnStatus := d.uint8()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	r := int8(*v)
	return &r, _returnStatus
}

func (d *Decoder) Int8() (v *Int8, _returnStatus _syscall.Status) {
     i, _returnStatus := d.int8()
     if _returnStatus != _syscall.StatusOk {
         return nil, _returnStatus
     }
     I := Int8(*i)
     return &I, _returnStatus
}

func (e *Encoder) uint16(u uint16) _syscall.Status {
	if !e.isEnoughSpace(2) {
	   _returnStatus := e.Flush() 
	   if _returnStatus != _syscall.StatusOk {
	      return _returnStatus
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 2
	e.bufSpace -= 2
	return _syscall.StatusOk
}

func (e *Encoder) Uint16(u Uint16) (_syscall.Status) {
     return e.uint16(uint16(u))
}

func (e *Encoder) int16(u int16) (_syscall.Status) {
	return e.uint16(uint16(u))
}

func (e *Encoder) Int16(u Int16) (_syscall.Status) {
     return e.int16(int16(u))
}

func (d *Decoder) uint16() (w *uint16, _returnStatus _syscall.Status) {
	if !d.isEnoughData(2) {
	   _returnStatus = d.readAtLeast(2)	   
	   if _returnStatus != _syscall.StatusOk {
	      return nil, _returnStatus
	   }
	}
	v := uint16(d.buf[d.bufStart]) | uint16(d.buf[d.bufStart + 1]) << 8
	d.bufStart = (d.bufStart + 2) % BufSize
	d.bufLen -= 2
	return &v, _syscall.StatusOk
}

func (d *Decoder) Uint16() (v *Uint16, _returnStatus _syscall.Status) {
     u, _returnStatus := d.uint16()
     if _returnStatus != _syscall.StatusOk {
          return nil, _returnStatus
     }
     U := Uint16(*u)
     return &U, _returnStatus
}

func (d *Decoder) int16() (w *int16, _returnStatus _syscall.Status) {
     	v, _returnStatus := d.uint16()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	r := int16(*v)
	return &r, _returnStatus
}

func (d *Decoder) Int16() (w *Int16, _returnStatus _syscall.Status) {
	u, _returnStatus := d.uint16()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	U := Int16(*u)
	return &U, _returnStatus
}

func (e *Encoder) uint32(u uint32) _syscall.Status {
	if !e.isEnoughSpace(4) {
	   _returnStatus := e.Flush()   		      
	   if _returnStatus != _syscall.StatusOk {
	      return _returnStatus
	   }
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 16)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 24)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 4
	e.bufSpace -= 4
	return _syscall.StatusOk
}

func (e *Encoder) Uint32(u Uint32) (_syscall.Status){
     return e.uint32(uint32(u))
}
func (e *Encoder) int32(u int32) (_syscall.Status){
	return e.uint32(uint32(u))
}

func (e *Encoder) Int32(u Int32) (_syscall.Status){
     return e.int32(int32(u))
}

func (d *Decoder) uint32() (w *uint32, _returnStatus _syscall.Status) {
	if !d.isEnoughData(4) {
	   _returnStatus = d.readAtLeast(4)	   
	   if _returnStatus != _syscall.StatusOk {
	      return nil, _returnStatus
	   }
	}
	v := uint32(d.buf[d.bufStart]) | uint32(d.buf[d.bufStart + 1]) << 8 | uint32(d.buf[d.bufStart + 2]) << 16 | uint32(d.buf[d.bufStart + 3]) << 24
	d.bufStart = (d.bufStart + 4) % BufSize
	d.bufLen -= 4
	return &v, _syscall.StatusOk
}

func (d *Decoder) Uint32() (v *Uint32, _returnStatus _syscall.Status) {
	w, _returnStatus := d.uint32()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	W := Uint32(*w)
	return &W, _returnStatus
}

func (d *Decoder) int32() (w *int32, _returnStatus _syscall.Status) {
     	v, _returnStatus := d.uint32()
	r := int32(*v)
	return &r, _returnStatus
}

func (d *Decoder) Int32() (v *Int32, _returnStatus _syscall.Status) {
	w, _returnStatus := d.int32()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	W := Int32(*w)
	return &W, _returnStatus
}

func (e *Encoder) Status(s _syscall.Status) (_syscall.Status){
	return e.uint32(uint32(s))
}


func (d *Decoder) Status() (*_syscall.Status, _syscall.Status) {
	i, _returnStatus := d.uint32()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	S := _syscall.Status(*i)
	return &S, _returnStatus
}

func (e *Encoder) uint64(u uint64) _syscall.Status{
	if !e.isEnoughSpace(8) {
	   	_returnStatus := e.Flush()    
		if _returnStatus != _syscall.StatusOk {
	      	   return _returnStatus
	      	}  
	}
	e.buf[e.bufStart] = byte(u)
	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 8)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 16)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 24)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 32)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 40)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 48)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.buf[e.bufStart] = byte(u >> 56)
   	e.bufStart = (e.bufStart + 1) % BufSize
	e.bufLen += 8
	e.bufSpace -= 8
	return _syscall.StatusOk
}

func (e *Encoder) Uint64(u Uint64) (_syscall.Status){
     return e.uint64(uint64(u))
}

func (e *Encoder) int64(u int64) (_syscall.Status){
	return e.uint64(uint64(u))
}

func (e *Encoder) Int64(u Int64) (_syscall.Status){
     return e.int64(int64(u))
}

func (d *Decoder) uint64() (w *uint64, _returnStatus _syscall.Status) {
	if !d.isEnoughData(8) {
	   _returnStatus = d.readAtLeast(8)	   
	   if _returnStatus != _syscall.StatusOk {
	      return nil, _returnStatus
	   }
	}
	v := uint64(d.buf[d.bufStart]) | uint64(d.buf[d.bufStart + 1]) << 8 | uint64(d.buf[d.bufStart + 2]) << 16 | uint64(d.buf[d.bufStart + 3]) << 24 | uint64(d.buf[d.bufStart + 4]) << 32 | uint64(d.buf[d.bufStart + 5]) << 40 | uint64(d.buf[d.bufStart + 6]) << 48 | uint64(d.buf[d.bufStart + 7]) << 56
	d.bufStart = (d.bufStart + 8) % BufSize
	d.bufLen -= 8
	return &v, _syscall.StatusOk
}

func (d *Decoder) Uint64() (v *Uint64, _returnStatus _syscall.Status) {
	u, _returnStatus := d.uint64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	U := Uint64(*u)
	return &U, _returnStatus
}

func (d *Decoder) int64() (w *int64, _returnStatus _syscall.Status) {
	v, _returnStatus := d.uint64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	r := int64(*v)
	return &r, _returnStatus
}

func (d *Decoder) Int64() (w *Int64, _returnStatus _syscall.Status) {
	i, _returnStatus := d.int64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	I := Int64(*i)
	return &I, _returnStatus
}

func (e *Encoder) float32(u float32) (_syscall.Status){
	return e.uint32(math.Float32bits(u))
}

func (e *Encoder) Float32(u Float32) (_syscall.Status){
 	return e.float32(float32(u))
}

func (d *Decoder) float32() (w *float32, _returnStatus _syscall.Status) {
     	v, _returnStatus := d.uint32()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	r := math.Float32frombits(*v)
	return &r, _returnStatus
}

func (d *Decoder) Float32() (v *Float32, _returnStatus _syscall.Status) {
	f, _returnStatus := d.float32()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	F := Float32(*f)
	return &F, _returnStatus
}

func (e *Encoder) float64(u float64) (_syscall.Status){
	return e.uint64(math.Float64bits(u))
}

func (e *Encoder) Float64(u Float64) (_syscall.Status){
	return e.float64(float64(u))
}

func (d *Decoder) float64() (w *float64, _returnStatus _syscall.Status) {
	v, _returnStatus := d.uint64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	r := math.Float64frombits(*v)
	return &r, _returnStatus
}

func (d *Decoder) Float64() (v *Float64, _returnStatus _syscall.Status) {
	f, _returnStatus := d.float64()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	F := Float64(*f)
	return &F, _returnStatus
}

func (e *Encoder) bool(u bool) (_returnStatus _syscall.Status){
	if u {
		_returnStatus = e.uint8(1)
	} else {
		_returnStatus = e.uint8(0)
	}
	return _returnStatus
}

func (e *Encoder) Bool(u Bool) (_syscall.Status){
	return e.bool(bool(u))
}

func (d *Decoder) bool() (w *bool, _returnStatus _syscall.Status) {
	v, _returnStatus := d.uint8()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	var u bool
	if *v == 1 {
		u = true
	} else {
		u = false
	}
	return &u, _returnStatus
}

func (d *Decoder) Bool() (w *Bool, _returnStatus _syscall.Status) {
	b, _returnStatus := d.bool()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	B := Bool(*b)
	return &B, _returnStatus
}

func (e *Encoder) SliceOfBytes(u []byte) (_returnStatus _syscall.Status){
     _returnStatus = e.length(uint32(len(u)))
     if _returnStatus != _syscall.StatusOk {
     	return _returnStatus
     }
     sliceStartPos := uint32(0)
     for ;!e.isEnoughSpace(uint32(len(u[sliceStartPos:]))); {
	copy(e.buf[e.bufStart:], u[sliceStartPos:(sliceStartPos + e.bufSpace)])
	sliceStartPos += e.bufSpace
	e.bufLen += e.bufSpace
	e.bufSpace = 0
	if e.bufLen != 0 {
	   _returnStatus := e.Flush()
	   if _returnStatus != _syscall.StatusOk {
	      return _returnStatus
	   }
	}
     } 
     if len(u[sliceStartPos:]) > 0 {
       copy(e.buf[e.bufStart:], u[sliceStartPos:])
       e.bufStart += uint32(len(u[sliceStartPos:]))	
       e.bufLen += uint32(len(u[sliceStartPos:]))
       e.bufSpace -= uint32(len(u[sliceStartPos:]))
     }
     return _returnStatus
     
}

func (d *Decoder) SliceOfBytes(v []byte, length uint32) (_returnStatus _syscall.Status){
     if length > d.bufLen {
     	copy(v, d.buf[d.bufStart:(d.bufStart + d.bufLen)])
	io.ReadFull(d.r, v[d.bufLen:])
	d.bufStart = 0
	d.bufLen = 0
	return _syscall.StatusOk
     }
     if !d.isEnoughData(length) {
	   _returnStatus = d.readAtLeast(length)
	   if _returnStatus != _syscall.StatusOk {
	      return _returnStatus
	   }
	}
     copy(v, d.buf[d.bufStart: (d.bufStart + length)])
     d.bufStart = (d.bufStart + length) % BufSize
     d.bufLen -= length
     return _returnStatus
}

func (e *Encoder) string(u string) (_returnStatus _syscall.Status){
	_returnStatus = e.length(uint32(len(u)))
	if _returnStatus != _syscall.StatusOk {
	   return _returnStatus
	}
	stringStartPos := uint32(0)
	for ;!e.isEnoughSpace(uint32(len(u[stringStartPos:]))); {
	   copy(e.buf[e.bufStart:], u[stringStartPos:(stringStartPos + e.bufSpace)])
	   stringStartPos += e.bufSpace
	   e.bufLen += e.bufSpace
	   e.bufSpace = 0
	   if e.bufSpace == 0 {
	      _returnStatus = e.Flush()
	      if _returnStatus != _syscall.StatusOk {
	      	 return _returnStatus
	      }
	   }	  
	} 
	if len(u[stringStartPos:]) > 0 {
		copy(e.buf[e.bufStart:], u[stringStartPos:])
		e.bufStart += uint32(len(u[stringStartPos:]))
		e.bufLen += uint32(len(u[stringStartPos:]))
		e.bufSpace -= uint32(len(u[stringStartPos:]))
	}
	return _returnStatus
	
}

func (e *Encoder) String(u String) (_syscall.Status){
     return e.string(string(u))
}

func (d *Decoder) string() (w *string, _returnStatus _syscall.Status) {
	len, _returnStatus := d.length()
	if _returnStatus != _syscall.StatusOk {
	   return nil, _returnStatus
	}

	if len > d.bufLen {
 	   b := make([]byte, len) 
	   copy(b[0:], d.buf[d.bufStart:(d.bufStart + d.bufLen)])
	   _, err := io.ReadFull(d.r, b[d.bufLen:])
	   if err != nil {
	      return nil, _syscall.StatusFail
	   }
	   d.bufStart = 0
	   d.bufLen = 0	   
	   str := string(b)
	   return &str, _syscall.StatusOk
	}
	if !d.isEnoughData(uint32(len)) {
	  	   _returnStatus = d.readAtLeast(uint32(len))
		   if _returnStatus != _syscall.StatusOk {
	      	      return nil, _returnStatus
		   }  
	}
	b := d.buf[d.bufStart:(d.bufStart + len)]
	d.bufStart = (d.bufStart + len) % BufSize
	d.bufLen -= len
	str := string(b)
	return &str, _returnStatus
}

func (d *Decoder) String() (v *String, _returnStatus _syscall.Status) {
	s, _returnStatus := d.string()
	if _returnStatus != _syscall.StatusOk {
		return nil, _returnStatus
	}
	S := String(*s)
	return &S, _returnStatus
}

func (e *Encoder) length(l uint32) (_syscall.Status){
	return e.uint32(l)
}

func (d *Decoder) length() (l uint32, _returnStatus _syscall.Status) {
     	v, _returnStatus := d.uint32()
	if _returnStatus != _syscall.StatusOk {
		return 0, _returnStatus
	}
	return *v, _returnStatus
}

func (e *Encoder) Flush() _syscall.Status{
	if e.bufLen == 0 {
	   return _syscall.StatusFail
	}
	if _, err := e.w.Write(e.buf[:e.bufLen]); err != nil {
	   return _syscall.StatusFail
	}
	e.bufStart = 0
	e.bufLen = 0
	e.bufSpace = BufSize
	e.reset()
	return _syscall.StatusOk
}

func (d *Decoder) ReadAll() {
     /* FIXME: Temporarily place reset here. We also reset in RPC calls. This 
      *        should probably go in one place: each type decoder, but we need
      *        to be able to differentiate between an application-called decode
      *        and a decode of internal elements. The latter case should not 
      *        call reset. */
     d.reset() 
}

func (e *Encoder) reset() {
	//e.t = NewTypeTree()
	e.t.Reset()
	e.m = make([]interface{}, 0)
}

func (d *Decoder) reset() {
	d.indexToValue = make([]interface{}, 0)
	d.m = make([]interface{}, 0)
	d.curPos = 0
}


func Hash(v interface{}) reflect.Type{
     return reflect.ValueOf(v).Type()
}

func Sizeof(v interface{}) uint64 {
     return uint64(unsafe.Sizeof(v))
}


var emptyPtrNode *llrb.Item

type TypeTree struct{
	tree *llrb.Tree
	Index uint32
	reusedItem llrb.Item
	min *llrb.Item
	max *llrb.Item
}

func lessPtr(a, b llrb.Item) bool {
	return uintptr(a.Ptr) < uintptr(b.Ptr)
}

func NewTypeTree() *TypeTree{
	t := TypeTree{}
	t.tree = llrb.New(lessPtr)
	t.Index = 0
	emptyPtrNode = &llrb.Item{0, "", 0, 0}
	return &t
}

func (t *TypeTree) Reset() {
     t.tree.Reset()
}

func (t *TypeTree) closestPtr(ptr uint64, typ string, size uint64, index uint32) (*llrb.Item, *llrb.Item){
	t.reusedItem.Ptr = ptr 
	t.reusedItem.EleType = typ
	t.reusedItem.EleSize = size
	t.reusedItem.Index = index
	minItem, maxItem := t.tree.FindAdjacentNodes(t.reusedItem)
	//minItem, maxItem = t.tree.FindAdjacentNodesLog(t.reusedItem)
	return minItem, maxItem
	
}

func (t *TypeTree) addToTree(elePtr uint64, eleType string, eleSize uint64, eleIndex uint32) {
	t.reusedItem.Ptr = elePtr 
	t.reusedItem.EleType = eleType
	t.reusedItem.EleSize = eleSize
	t.reusedItem.Index = eleIndex
	t.tree.InsertNoReplace(&t.reusedItem)
}

func (t *TypeTree) PointerCheck(ptr_unsafe unsafe.Pointer, typ string, size uint64) (index uint32, encoded bool, _returnStatus _syscall.Status) {
       	
	ptr := uint64(uintptr(ptr_unsafe))
	t.reusedItem.Ptr = ptr 
	t.reusedItem.EleType = typ
	t.reusedItem.EleSize = size
	t.reusedItem.Index = 0
	sameItem := t.tree.Get(t.reusedItem)
	t.min, t.max = t.closestPtr(ptr, sameItem.EleType, sameItem.EleSize, sameItem.Index) 
	switch {
	case !sameItem.Equal(emptyPtrNode)  && sameItem.EleType == typ:
		// already in the tree
		return sameItem.Index, true, _syscall.StatusOk
	case !sameItem.Equal(emptyPtrNode)  && sameItem.EleType != typ:
		t.addToTree(ptr, typ, size, t.Index)
		t.Index++
		return t.Index, false, _syscall.StatusOk
	case (t.min.Equal(emptyPtrNode) && !t.max.Equal(emptyPtrNode) && (ptr + size) <= t.max.Ptr) ||
	     (!t.min.Equal(emptyPtrNode) && t.max.Equal(emptyPtrNode) && ptr >= (t.min.Ptr + t.min.EleSize)) ||
	     (!t.min.Equal(emptyPtrNode) && !t.max.Equal(emptyPtrNode) && (ptr + size) <= t.max.Ptr && ptr >= (t.min.Ptr + t.min.EleSize)) ||
	     (t.min.Equal(emptyPtrNode) && t.max.Equal(emptyPtrNode)):
		t.addToTree(ptr, typ, size, t.Index)
		t.Index++
		return t.Index, false, _syscall.StatusOk
	default:
		// e := fmt.Errorf("Illegal pointer")
		return 0, false, _syscall.StatusFail 
	}
	// e := fmt.Errorf("Illegal pointer")
	return 0, false, _syscall.StatusFail
}
func (t *Puller) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Puller(t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}


	return buffer.Bytes(), _syscall.StatusOk
}


func (t *Puller) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Puller()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}


func (t *Puller) GetHash() (_syscall.HashValue) {
       return _syscall.HashValue{ 0x18,0x08,0x95,0x30,0xc0,0x21,0x45,0xcd,0x8e,0x3f,0xe6,0x5d,0x18,0x6a,0xa9,0x87,0x49,0xff,0xdc,0x08,0x05,0xcd,0x15,0x36,0x97,0x2c,0x9a,0xf2,0xa1,0x4c,0xea,0x51,0x14,0x79,0x2a,0x8b,0x31,0xb9,0x51,0x98,0xe5,0x37,0x17,0x8f,0x1c,0x21,0x5e,0x33,0x57,0x99,0x48,0x60,0x54,0xd2,0x5f,0x47,0x1f,0xb2,0xa3,0x0d,0x07,0xe4,0x73,0x58, }
}
func (t *Message) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Message(t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}


	return buffer.Bytes(), _syscall.StatusOk
}


func (t *Message) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Message()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}


func (t *Message) GetHash() (_syscall.HashValue) {
       return _syscall.HashValue{ 0x37,0x08,0x3d,0x22,0x93,0x62,0x50,0x74,0xbb,0xa4,0x8a,0xbf,0x70,0xee,0x78,0x1e,0x25,0x03,0xaa,0xe7,0x4e,0x88,0x51,0x3a,0x99,0x6c,0x8c,0xf4,0x80,0xd8,0x2c,0x5c,0xaf,0xc5,0xbe,0x8b,0xf2,0x66,0x4e,0x3c,0x1c,0xb4,0xb7,0xe8,0xeb,0x8e,0x93,0x38,0x2b,0x86,0xdf,0x55,0xb7,0x8e,0xcf,0x7f,0x19,0xe1,0xc3,0x21,0xe9,0x5d,0x77,0xc7, }
}
func (t *Pusher) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Pusher(t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}


	return buffer.Bytes(), _syscall.StatusOk
}


func (t *Pusher) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Pusher()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}


func (t *Pusher) GetHash() (_syscall.HashValue) {
       return _syscall.HashValue{ 0x6f,0x56,0x3f,0x88,0x6c,0x95,0x40,0x72,0xf7,0x60,0x3e,0x58,0xa1,0x6f,0xe5,0x53,0x57,0xb4,0x62,0x58,0xb5,0x10,0x42,0xf1,0x47,0xc1,0xa8,0xa7,0xc7,0x09,0x1b,0xbb,0xfe,0x78,0x9a,0x8a,0xad,0xd2,0x95,0xd4,0x0d,0x49,0xb6,0xb6,0xd3,0xe0,0x96,0x09,0xbc,0x55,0x97,0xa5,0xa8,0xdc,0x93,0x37,0xef,0x48,0xfc,0x3a,0x93,0x6c,0x6f,0x3e, }
}
func (t *Int8) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Int8(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Int8) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Int8()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Int8) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0xc2,0x71,0x6e,0x3c,0x34,0x13,0xc0,0xbe,0xb3,0x9e,0x2e,0xbc,0xc7,0x99,0x62,0x3f,0x28,0xe9,0x9d,0x19,0x71,0x6a,0x00,0x5b,0x69,0x4b,0x1c,0xbb,0x3d,0x8d,0xcc,0x45,0xbd,0x51,0xca,0x50,0xc7,0x3f,0x8a,0x64,0xc8,0xa0,0xf0,0x0b,0x46,0x9a,0x87,0x05,0xe9,0x3c,0xb9,0x27,0x91,0xdf,0x95,0x88,0x1b,0x2e,0xfa,0x9f,0xc0,0x1f,0xcb,0xad, }
}
func (t *Uint8) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Uint8(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Uint8) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Uint8()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Uint8) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0x09,0x66,0x78,0xbb,0x4a,0x86,0x63,0x8d,0x8e,0xc6,0x58,0x8e,0xcc,0x2d,0x89,0x5a,0x5c,0x17,0xb4,0x86,0x37,0x8c,0x81,0xc2,0xf2,0xac,0xf2,0x48,0x67,0x57,0x82,0x13,0x3e,0x2e,0x7d,0x80,0xb6,0x66,0x8d,0x84,0xc4,0x55,0xf4,0xc9,0xe1,0x33,0xfc,0x71,0x0b,0x77,0x43,0x63,0x06,0x82,0x76,0x63,0x07,0xbc,0xf1,0x64,0xdb,0xd9,0x5c,0x93, }
}
func (t *Bool) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Bool(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Bool) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Bool()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Bool) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0x3e,0x76,0x06,0x81,0x73,0x27,0x61,0xed,0x16,0x8e,0xa3,0x08,0xe1,0x10,0x11,0x85,0xe1,0xbd,0x39,0x02,0xee,0x67,0x60,0x21,0x6a,0x59,0xc4,0x07,0x5a,0x99,0xc1,0x46,0xb7,0xcd,0x98,0x14,0xce,0x14,0x47,0x0e,0xb0,0x80,0x6d,0x91,0x66,0x50,0xb0,0xe5,0x0e,0x77,0x6f,0x53,0xe5,0xd1,0x72,0x28,0x1d,0xd0,0xe1,0x70,0x43,0xc8,0x65,0x39, }
}
func (t *Int16) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Int16(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Int16) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Int16()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Int16) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0x99,0x7f,0x69,0x9c,0x17,0xfd,0x06,0x74,0x8c,0x2d,0xba,0xc4,0x61,0x01,0x35,0x21,0xf1,0x97,0xd8,0x1e,0x74,0x3c,0x2f,0x96,0x56,0xe6,0xdc,0xfc,0x14,0x1e,0x0d,0x83,0x36,0xdc,0x73,0x36,0xb0,0xf4,0x9e,0x40,0x2b,0xfe,0x97,0x6f,0xfa,0xa5,0x27,0xac,0xe1,0xa2,0x57,0x2a,0xae,0x6d,0x18,0x22,0xe2,0xdc,0xd8,0x79,0xe0,0xb6,0xf6,0x7e, }
}
func (t *Uint16) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Uint16(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Uint16) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Uint16()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Uint16) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0xa9,0x1f,0x92,0xf4,0xe9,0x96,0xb6,0xd4,0xa4,0xcb,0x85,0x8d,0x11,0x80,0x20,0x10,0xd6,0x29,0xba,0x29,0xe1,0x89,0x50,0x2c,0xa0,0xf2,0xcb,0x1d,0x86,0xb8,0x27,0x3b,0x67,0xf3,0x35,0x72,0xd1,0x78,0x15,0xcc,0xb0,0x94,0x94,0x6f,0x02,0xab,0x2e,0x46,0xcd,0x74,0xea,0xf5,0x15,0x0a,0x26,0xdc,0x4e,0xf7,0x0d,0x9f,0x3a,0x9c,0x6e,0x55, }
}
func (t *Int32) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Int32(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Int32) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Int32()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Int32) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0xeb,0xc1,0x67,0x8b,0x06,0x82,0x70,0x21,0x38,0xc2,0xd9,0x9e,0x33,0x22,0xd1,0xa8,0xc7,0x2e,0x9b,0x68,0xe9,0x41,0x12,0x00,0x1e,0x3e,0x51,0xa8,0xf5,0xd9,0xfa,0x34,0x0c,0x44,0x9c,0x06,0x6d,0x9d,0x4c,0xe7,0x2a,0x06,0xab,0x75,0x77,0x5d,0xdf,0x28,0x34,0x88,0x7c,0x7e,0x96,0x97,0xbb,0x8a,0x95,0xfe,0x07,0x65,0xf7,0x7c,0x7e,0x4c, }
}
func (t *Uint32) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Uint32(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Uint32) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Uint32()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Uint32) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0xce,0xf4,0x3a,0x05,0xae,0x67,0xd9,0x73,0xc2,0xa2,0x1d,0xf8,0xcd,0xf9,0xd2,0xde,0x69,0x8d,0x0d,0xb7,0x61,0xb9,0x51,0x22,0x58,0xed,0x8f,0xb1,0x83,0xf1,0x5c,0xff,0x5b,0x84,0xe2,0x14,0x0e,0x10,0x68,0x3f,0x7a,0xd9,0xa7,0x8f,0x5b,0xe4,0x9e,0x4e,0x00,0x7d,0xcb,0xfb,0xd1,0x69,0x59,0x9d,0xbf,0x9b,0x75,0x65,0x15,0x9e,0x8b,0x82, }
}
func (t *Int64) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Int64(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Int64) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Int64()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Int64) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0x56,0x16,0x70,0x35,0xd0,0x09,0x18,0x69,0x0e,0xae,0xad,0x60,0xd1,0xee,0x39,0xa8,0x61,0x45,0x58,0x5b,0x99,0x20,0x94,0x57,0x1f,0xb0,0x48,0xeb,0xb2,0xcf,0x5c,0xa5,0x8d,0xc7,0x8e,0x7e,0x3c,0x89,0xcd,0x2f,0xdc,0xf2,0x1c,0x2a,0xe3,0xd2,0x7f,0x98,0xc2,0xad,0x1c,0x3d,0x4e,0x62,0xd9,0xdb,0xc8,0xc8,0x59,0xc5,0xd5,0xc6,0xed,0x7a, }
}
func (t *Uint64) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Uint64(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Uint64) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Uint64()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Uint64) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0x11,0x26,0xb3,0x0d,0x51,0x59,0x87,0x5e,0x0d,0x5b,0x93,0xfc,0x92,0xf0,0x78,0xaa,0x12,0xac,0x93,0xb8,0x30,0x1f,0x48,0x0e,0x13,0x4d,0x8b,0xfb,0x4c,0x58,0xfa,0x3a,0x69,0x6a,0x81,0x01,0xc5,0x47,0xc1,0x55,0x43,0x95,0x41,0xdf,0x3c,0x8e,0xb6,0x96,0x4a,0x3c,0x88,0xab,0x3f,0x88,0xed,0x37,0x5f,0x08,0x4a,0x41,0x8e,0xd5,0xda,0x1e, }
}
func (t *Float32) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Float32(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Float32) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Float32()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Float32) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0x42,0x36,0xae,0xd3,0x62,0xca,0x34,0x75,0x94,0x52,0xf0,0x5f,0x44,0x83,0x61,0x75,0x69,0x39,0xcf,0x69,0x74,0x91,0xee,0x8d,0x35,0x8c,0xd7,0xa1,0x63,0x0f,0x88,0x86,0x6b,0x52,0xdd,0x6d,0xe1,0xb2,0x26,0xf4,0x3a,0x9c,0x9e,0xf1,0x56,0x0d,0xf1,0x48,0x07,0x39,0x46,0xf8,0xe9,0xd3,0xab,0x86,0xe0,0x1c,0x98,0x0d,0x17,0x6b,0x02,0x63, }
}
func (t *Float64) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Float64(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Float64) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Float64()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *Float64) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0x65,0x9b,0xb2,0x59,0x85,0xe2,0x60,0xe7,0x1e,0x12,0x17,0x3f,0xc3,0x1f,0x20,0x45,0x08,0x9e,0x7e,0x11,0x6b,0xaa,0xb3,0x1e,0x6d,0x7d,0x7a,0x5b,0xe3,0x3d,0x40,0xb5,0x40,0x06,0x52,0x85,0x37,0x80,0x2c,0xd8,0x7d,0x48,0x67,0xe3,0x9a,0xdd,0xc9,0x13,0x11,0x2c,0xa5,0xcc,0x5a,0x33,0xbc,0x35,0x6b,0x3e,0xa8,0x75,0x93,0x84,0xcf,0x1b, }
}
func (t *String) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.String(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *String) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.String()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	*t = *value
	return _syscall.StatusOk
}

func (t *String) GetHash() (_syscall.HashValue) {
	return  _syscall.HashValue{ 0x27,0x57,0xcb,0x3c,0xaf,0xc3,0x9a,0xf4,0x51,0xab,0xb2,0x69,0x7b,0xe7,0x9b,0x4a,0xb6,0x1d,0x63,0xd7,0x4d,0x85,0xb0,0x41,0x86,0x29,0xde,0x8c,0x26,0x81,0x1b,0x52,0x9f,0x3f,0x37,0x80,0xd0,0x15,0x00,0x63,0xff,0x55,0xa2,0xbe,0xee,0x74,0xc4,0xec,0x10,0x2a,0x2a,0x27,0x31,0xa1,0xf1,0xf7,0xf1,0x0d,0x47,0x3a,0xd1,0x8a,0x6a,0x87, }
}
func (t *Any) WriteBuffer() ([]byte, _syscall.Status) {
	var buffer bytes.Buffer
	e := NewEncoder(&buffer)
	_returnStatus := e.Any(*t)
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	_returnStatus = e.Flush()
	if _returnStatus != _syscall.StatusOk {
		return nil, _syscall.StatusWriteFail
	}

	return buffer.Bytes(), _syscall.StatusOk
}

func (t *Any) ReadBuffer(buffer []byte) (_syscall.Status) {
	d := NewDecoder(bytes.NewBuffer(buffer))
	value, _returnStatus := d.Any()
	if _returnStatus != _syscall.StatusOk {
		return _syscall.StatusFail
	}

	if !d.IsEOF() {
		return _syscall.StatusInvalidType
	}

	t.Value = value.Value
	return _syscall.StatusOk
}
func (t *Any) GetHash() (_syscall.HashValue) {
	return _syscall.HashValue{0x40, 0xd3, 0x80, 0xd9, 0x2f, 0x53, 0xad, 0x12, 0xcf, 0x21, 0x94, 0x59, 0x68, 0x74, 0xa9, 0x17, 0x9f, 0x1e, 0xe3, 0xf9, 0x2e, 0x8f, 0x7c, 0x99, 0x4b, 0xf9, 0x4d, 0xb3, 0x29, 0x1a, 0xbb, 0x89, 0xc3, 0xff, 0x35, 0x1e, 0xd2, 0xb9, 0x11, 0x30, 0x15, 0x7f, 0xc7, 0xd3, 0x2f, 0x84, 0x2c, 0xed, 0x4b, 0x99, 0x8a, 0x9d, 0xe6, 0xe0, 0xe0, 0x1d, 0x98, 0x7a, 0x28, 0xd9, 0x69, 0x34, 0xe6, 0xcc}
}

