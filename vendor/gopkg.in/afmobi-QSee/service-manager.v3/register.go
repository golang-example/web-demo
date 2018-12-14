package service_manager

import (
	"time"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"fmt"
	"log"
	"strconv"
)

var kHeartBeatInterval = time.Second * 2
var kTTL = time.Second * 5

var kRoot = "service"

type Register struct {
	kapi    client.KeysAPI
	serviceName     string
	serviceId	string
	protocol	string
	ip		string
	port 		int
	active  	bool
}

func NewRegister(serviceName string,ip string, port int, endpoints []string) (*Register, error) {
	//new etcd client
	cfg := client.Config{
		Endpoints:               endpoints,
		HeaderTimeoutPerRequest: time.Second * 2,
	}
	cli, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	//create service info
	if(ip == ""){
		ip = GetIP()
	}
	serviceID := serviceName + "_" + ip + "_" + strconv.Itoa(port)
	Register := &Register{
		kapi:    	client.NewKeysAPI(cli),
		serviceName:	serviceName,
		serviceId:	serviceID,
		protocol:	"http",
		ip:		ip,
		port:		port,
		active:  	true,
	}
	return Register, nil
}

func (reg *Register) Register() {
	reg.heartbeat()
	go reg.heartbeat()
}

func (reg *Register) Unregister() {
	reg.active = false
}

func (reg *Register) heartbeat() {
	etcdKey := reg.serviceId + ":" + reg.protocol + ":" + reg.ip + ":" + strconv.Itoa(reg.port)
	etcdValue := reg.active
	for{
		response, err := reg.kapi.Set(context.Background(),
			fmt.Sprintf("%s/%s/%s", kRoot, reg.serviceName, etcdKey),
			strconv.FormatBool(etcdValue),
			&client.SetOptions{
				TTL: kTTL,
			})

		log.Println(response)
		if(err != nil){
			log.Println(err)
		}

		if(!reg.active){
			break
		}
		time.Sleep(kHeartBeatInterval)
	}
}

