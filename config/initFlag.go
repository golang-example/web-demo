package config

import "flag"

var EtcdAddr = flag.String("etcd", "http://127.0.0.1:2379,http://127.0.0.1:2379,http://127.0.0.1:2379", "Etcd address")
