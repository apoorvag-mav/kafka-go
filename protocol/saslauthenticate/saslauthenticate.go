package saslauthenticate

import "github.com/apoorvag-mav/kafka-go/protocol"

func init() {
	protocol.Register(&Request{}, &Response{})
}

type Request struct {
	AuthBytes []byte `kafka:"min=v0,max=v1"`
}

func (r *Request) ApiKey() protocol.ApiKey { return protocol.SaslAuthenticate }

type Response struct {
	ErrorCode         int16  `kafka:"min=v0,max=v1"`
	ErrorMessage      string `kafka:"min=v0,max=v1,nullable"`
	AuthBytes         []byte `kafka:"min=v0,max=v1"`
	SessionLifetimeMs int64  `kafka:"min=v1,max=v1"`
}

func (r *Response) ApiKey() protocol.ApiKey { return protocol.SaslAuthenticate }
