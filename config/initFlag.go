package config

import "flag"

var EtcdAddr = flag.String("etcd", "http://192.168.1.52:2379,http://192.168.1.52:2379,http://192.168.1.52:2379", "Etcd address")
