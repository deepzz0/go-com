package etcd

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/deepzz0/logd"
	"gopkg.in/coreos/go-etcd.v2/etcd"
)

type EtcdClient struct {
	*etcd.Client
}

func Connect(host string) *EtcdClient {
	return &EtcdClient{etcd.NewClient([]string{fmt.Sprintf("http://%s", host)})}
}

func (c *EtcdClient) WatchDo(key string, f func(*etcd.Response)) {
	go func() {
		for {
			ch := make(chan *etcd.Response)
			fmt.Println("watch...")
			if _, err := c.Watch(key, 0, true, ch, nil); err != nil {
				fmt.Println("watch error")
				logd.Error(err)
			} else {
				go processWatchResult(ch, f)
			}
			time.Sleep(time.Second)
		}
	}()
}

func (c *EtcdClient) GetJson(key string, result interface{}) error {
	if resp, err := c.Get(key, false, false); err != nil {
		return err
	} else {
		return json.Unmarshal([]byte(resp.Node.Value), result)
	}
}

func (c *EtcdClient) GetString(key string) (string, error) {
	if resp, err := c.Get(key, false, false); err != nil {
		return "", err
	} else {
		return resp.Node.Value, nil
	}
}

func (c *EtcdClient) SetJson(key string, data interface{}) error {
	if bytes, err := json.Marshal(data); err != nil {
		return err
	} else {
		_, e := c.Set(key, string(bytes[:]), 0)
		return e
	}
}

func processWatchResult(ch chan *etcd.Response, f func(*etcd.Response)) {
	fmt.Println("proccess")
	for {
		if resp, ok := <-ch; ok {
			f(resp)
		} else {
			fmt.Println("!ok")
			break
		}
	}
}
