package service_manager

import (
	"log"
	"sync"
	"time"

	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
	"strings"
	"errors"
	"strconv"
)

const CONFIG_ROOT = "/config/"

type Config struct {
	sync.RWMutex
	kapi          client.KeysAPI
	serviceName   string
	serviceStruct interface{}
	Items         map[string]interface{}
}

/**
 * init conifg
 */
func InitConfig(serviceName string, serviceStruct interface{}, endpoints []string) (*Config, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		HeaderTimeoutPerRequest: time.Second * 2,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	Config := &Config{
		kapi:          client.NewKeysAPI(c),
		serviceName:   serviceName,
		serviceStruct: serviceStruct,
		Items:         make(map[string]interface{}),
	}

	Config.fetch(CONFIG_ROOT + serviceName, Config.Items)
	Config.reload()

	/// `fetch` Timer may work well too?
	go Config.watch(CONFIG_ROOT + serviceName)

	return Config, err
}

/**
 * init conifg function
 */
func InitConfigCallFunction(serviceName string, serviceStruct interface{}, endpoints []string, f func(interface{})) (*Config, error) {
	cfg := client.Config{
		Endpoints:               endpoints,
		HeaderTimeoutPerRequest: time.Second * 2,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	Config := &Config{
		kapi:          client.NewKeysAPI(c),
		serviceName:   serviceName,
		serviceStruct: serviceStruct,
		Items:         make(map[string]interface{}),
	}

	Config.fetch(CONFIG_ROOT + serviceName, Config.Items)
	Config.reload()

	/// `fetch` Timer may work well too?
	go Config.watchCallFunction(CONFIG_ROOT + serviceName, f)
	f(Config.Items)

	return Config, err
}


/**
 * log config into struct
 */
func (cfg *Config) reload() {
	jsonb, _ := json.Marshal(cfg.Items)
	fmt.Println(string(jsonb))
	err := json.Unmarshal(jsonb,cfg.serviceStruct)
	if(err != nil){
		fmt.Println(err.Error())
	}
	fmt.Println(cfg.serviceStruct)
}

/**
 * fetch root items
 */
func (cfg *Config) fetch(path string,items map[string]interface{}) error {
	resp, err := cfg.kapi.Get(context.Background(), path, &client.GetOptions{
		Recursive: true,
	})
	if err != nil {
		return err
	}
	return extract(resp.Node, cfg.Items)
}

/**
 * fill the value cantiner
 */
func extract(node *client.Node,nodeValues map[string]interface{}) error {
	if(node.Dir){
		for _, childNode := range node.Nodes {
			if(childNode.Dir){
				newitem := make(map[string]interface{})

				vKey := getKey(childNode)
				nodeValues[vKey] = newitem

				extract(childNode,newitem)
			}else {
				if(!createOrUpdateList(nodeValues,childNode)){
					vKey := getKey(childNode)
					nodeValues[vKey] = childNode.Value
				}
			}
		}
	}else {
		return errors.New("Node is not a Dir Node")
	}
	return nil
}

/**
 * get node`s value key
 */
func getKey(node *client.Node)string  {
	subs := strings.Split(node.Key,"/")
	if(len(subs) < 1){
		return ""
	}
	return subs[len(subs)-1]
}


/**
 * process list items
 */
func createOrUpdateList(nodeValues map[string]interface{}, node *client.Node) bool {
	nodeKey := getKey(node)
	isListKey := strings.Index(nodeKey,"list_")

	if(isListKey == 0){//list key
		idx_key := strings.Split(nodeKey,"_")
		idx, err := strconv.Atoi(idx_key[1])
		fmt.Println(idx)

		if(err != nil){
			fmt.Println(err.Error())
		}
		key := idx_key[2]
		if(key != ""){
			//if list inited
			list,ok:= nodeValues[key].([]interface{})
			if(ok){
				//fmt.Println("ok")
			}
			if(list != nil){
				for i:=0; i<len(list); i++  {
					//update
					if(list[i] == node.Value){
						list[i] = node.Value
						nodeValues[key] = list
						break
					//insert
					}else if(i+1 == len(list)){
						list = append(list,node.Value)
						nodeValues[key] = list
					}
				}
			}else {//create
				list := make([]interface{},0)
				list = append(list,node.Value)
				nodeValues[key] = list
			}
		}
		return true

	}else {// not list key
		return false
	}
}

/**
 * process list items
 */
func deleteList(nodeValues map[string]interface{}, node *client.Node,preNode *client.Node) bool {
	nodeKey := getKey(node)
	isListKey := strings.Index(nodeKey,"list_")

	if(isListKey == 0){//list key
		idx_key := strings.Split(nodeKey,"_")
		idx, err := strconv.Atoi(idx_key[1])
		fmt.Println(idx)

		if(err != nil){
			fmt.Println(err.Error())
		}
		key := idx_key[2]
		if(key != ""){
			//if list inited
			list,ok:= nodeValues[key].([]interface{})
			if(ok){
				//fmt.Println("ok")
			}
			if(list != nil){
				//delete
				if(node.Value == ""){
					for i:=0; i<len(list); i++  {
						if(list[i] == preNode.Value){
							list = append(list[:i], list[i+1:]...)
							if(len(list) <= 0){
								delete(nodeValues,key)
							}else {
								nodeValues[key] = list
							}
							break
						}
					}
				}
			}
		}
		return true

	}else {// not list key
		return false
	}
}

/**
 * watch changing
 */
func (cfg *Config) watch(path string) {
	watcher := cfg.kapi.Watcher(path, &client.WatcherOptions{
		Recursive: true,
	})
	for {
		resp, err := watcher.Next(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		switch resp.Action {
		case "set", "update":
			parentNodeValues,err := cfg.getParentNodeValues(resp.Node)
			if(err != nil){
				fmt.Println(err.Error())
			}

			if(!createOrUpdateList(parentNodeValues,resp.Node)){
				nodeKey := getKey(resp.Node)
				parentNodeValues[nodeKey] = resp.Node.Value
			}
			fmt.Println("update-config--- key: "+ resp.Node.Key + "- value: " + resp.Node.Value)
			break
		case "expire", "delete":
			parentNodeValues,err := cfg.getParentNodeValues(resp.Node)
			if(err != nil){
				fmt.Println(err.Error())
			}

			if(!deleteList(parentNodeValues, resp.Node, resp.PrevNode)){
				nodeKey := getKey(resp.Node)
				delete(parentNodeValues,nodeKey)
			}
			fmt.Println("delete-config--- key: "+ resp.Node.Key + "- value: " + resp.Node.Value)
			break
		default:
			log.Println("watch me!!!", "resp ->", resp)
		}

		cfg.reload()
	}
}

/**
 * watch changing call function
 */
func (cfg *Config) watchCallFunction(path string, f func(interface{})) {
	watcher := cfg.kapi.Watcher(path, &client.WatcherOptions{
		Recursive: true,
	})
	for {
		resp, err := watcher.Next(context.Background())
		if err != nil {
			log.Println(err)
			continue
		}

		switch resp.Action {
		case "set", "update":
			parentNodeValues,err := cfg.getParentNodeValues(resp.Node)
			if(err != nil){
				fmt.Println(err.Error())
			}

			if(!createOrUpdateList(parentNodeValues,resp.Node)){
				nodeKey := getKey(resp.Node)
				if resp.Node.Dir {
					parentNodeValues[nodeKey] = make(map[string]interface{})
				} else {
					parentNodeValues[nodeKey] = resp.Node.Value
				}
			}
			fmt.Println("update-config--- key: "+ resp.Node.Key + "- value: " + resp.Node.Value)
			break
		case "expire", "delete":
			parentNodeValues,err := cfg.getParentNodeValues(resp.Node)
			if(err != nil){
				fmt.Println(err.Error())
			}

			if(!deleteList(parentNodeValues, resp.Node, resp.PrevNode)){
				nodeKey := getKey(resp.Node)
				delete(parentNodeValues,nodeKey)
			}
			fmt.Println("delete-config--- key: "+ resp.Node.Key + "- value: " + resp.Node.Value)
			break
		default:
			log.Println("watch me!!!", "resp ->", resp)
		}

		cfg.reload()
		f(cfg.Items)
	}
}

/**
 * get node values
 */
func (cfg *Config) getParentNodeValues(node *client.Node) (map[string]interface{}, error)  {
	var result  map[string]interface{}

	subs := strings.Split(node.Key,"/")
	//check
	if(len(subs) < 3){
		return nil,errors.New("error path")
	}
	if(subs[1] != "config"){
		return nil,errors.New("error path, root path is not [/conifg].")
	}
	if(subs[2] != cfg.serviceName){
		return nil,errors.New("error path, serviceName is error.")
	}

	for i, item := range subs{
		if(i == 2){
			result = cfg.Items
		} else if (i > 2){
			if(i == (len(subs)-1)){
				break
			}
			if(result[item] == nil){
				tempItem := make(map[string]interface{})
				result[item] = tempItem
				result = tempItem
			}else {
				res,ok := result[item].(map[string]interface{})
				if(ok){
					result = res
				}
			}
		}
	}
	return result, nil
}


