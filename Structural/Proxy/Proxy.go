package Proxy

import "fmt"

// an object on remote server, assuming the creation of it costs a lot
type RemoteObj struct {
	name string
}

func (r RemoteObj) GetName() string {
	return r.name
}

// APIGate
type APIGate map[string]RemoteObj

// an local proxy, by which we can call methods of the remote object
type ProxyObj struct {
	remoteAddr string
	apiGate    APIGate
}

func (p ProxyObj) GetName() string {
	fmt.Println("visit remote")
	o := p.apiGate[p.remoteAddr]
	fmt.Println("call remote method")
	r := o.GetName()
	fmt.Println("down")
	return r
}

func NewProxyObj(remoteAddr string, gate APIGate) *ProxyObj {

	return &ProxyObj{remoteAddr, gate}
}

func Test() {
	var apiGate = APIGate{"1.1.1.1": RemoteObj{"onServer"}}
	proxyObj := NewProxyObj("1.1.1.1", apiGate)
	fmt.Println(proxyObj.GetName())
}

/*
refer: https://www.runoob.com/design-pattern/proxy-pattern.html
使对远端的访问像对本地的访问那样表现出来
*/
