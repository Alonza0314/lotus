package lotus

import "crypto/tls"

type lotusServer struct {
	tlsCert		tls.Certificate
	serviceMap	map[string]interface{}
}

func NewLotusServer(pemPath string) *lotusServer {

}