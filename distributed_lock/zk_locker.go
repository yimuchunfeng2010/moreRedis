package distributed_lock

import (
	"github.com/samuel/go-zookeeper/zk"
	"time"
	"strconv"
	"strings"
	"github.com/sirupsen/logrus"
	"more-for-redis/global"
	"errors"
)

func ReConnect()(err error){
	var hosts = []string{global.Config.ZkIPaddr}
	conn, _, err := zk.Connect(hosts, 100000*time.Minute)
	if err != nil {
		logrus.Errorf("Connect %s", err.Error())
		return
	}
	global.Config.ZkConn = conn
	return
}
// 写锁
func Lock() (lockName string, err error) {
	// 创建临时写节点
	if nil == global.Config.ZkConn{
		if err = ReConnect();err != nil{
			logrus.Errorf("zk connect fail err: %s",err.Error())
		}
	}
	// 获取当前子节点
	children, _, err := global.Config.ZkConn.Children("/Lock")
	if err != nil {
		logrus.Errorf("Children %s", err.Error())
		return "", err
	}

	maxChild := GetMaxchild(children)

	// 创建当前节点
	var wLockPath = "/Lock/w"
	var wLockData []byte = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	var wLockFlags int32 = 2 // 永久序列增长节点
	var acl = zk.WorldACL(zk.PermAll)

	lockPath, err := global.Config.ZkConn.Create(wLockPath, wLockData, wLockFlags, acl)
	if err != nil {
		logrus.Errorf("Create %s", err.Error())
		return "", err
	}

	if "" != maxChild {
		// 对最大子节点设置观察点
		_, _, ech, err := global.Config.ZkConn.ExistsW("/Lock/"+maxChild)
		if err != nil {
			logrus.Errorf("ExistsW maxChild: %s, err: %s ", maxChild, err.Error())
			return "", err
		}
		timeout := global.Config.Timeout // 超时时间10s
		for timeout > 0 {
			select {
			case _, ok := <-ech:
				if ok {
					return lockPath, nil
				}
			default:
				time.Sleep(time.Millisecond)
				timeout--
			}
		}
		// 失败则删除当前节点
		DeleteNode(lockName)
		return "", errors.New("Get Lock Fail timeout")
	} else {
		return lockPath, nil
	}

	// 失败则删除当前节点
	DeleteNode(lockName)
	return "", errors.New("Get Lock Fail")
}

// 释放写锁
func Unlock(lockName string) (err error) {
	logrus.Infof("Unlock lockName %s",lockName)
	if nil == global.Config.ZkConn{
		if err = ReConnect();err != nil{
			logrus.Errorf("zk connect fail err: %s",err.Error())
		}
	}
	// 删除节点
	err = global.Config.ZkConn.Delete(lockName, 0)
	if err != nil {
		logrus.Errorf("Unlock Delete lockName: %s, err: %s", lockName, err.Error())
		return err
	}

	return nil
}

// 读锁
func RLock() (lockName string, err error) {
	// 创建临时写节点
	if nil == global.Config.ZkConn{
		if err = ReConnect();err != nil{
			logrus.Errorf("zk connect fail err: %s",err.Error())
		}
	}
	// 获取当前子节点
	children, _, err := global.Config.ZkConn.Children("/Lock")
	if err != nil {
		logrus.Errorf("RLock Children %s", err.Error())
		return "", err
	}

	maxChild := GetMaxWritechild(children)

	// 创建子节点
	var wLockPath = "/Lock/r"
	var wLockData []byte = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	var wLockFlags int32 = 2 // 永久序列增长节点
	var acl = zk.WorldACL(zk.PermAll)

	lockPath, err := global.Config.ZkConn.Create(wLockPath, wLockData, wLockFlags, acl)
	if err != nil {
		logrus.Errorf("RLock Create %s", err.Error())
		return "", err
	}

	logrus.Infof("maxChild %s",maxChild)
	if "" != maxChild {
		// 对最大子节点设置观察点
		_, _, ech, err := global.Config.ZkConn.ExistsW("/Lock/"+maxChild)
		if err != nil {
			logrus.Errorf("RLock ExistsW %s", err.Error())
			return "", err
		}
		timeout := global.Config.Timeout // 超时时间10s
		for timeout > 0 {
			select {
			case _, ok := <-ech:
				if ok {
					logrus.Infof("success reture lockPath",lockPath)
					return lockPath, nil
				}
			default:
				time.Sleep(time.Millisecond)
				timeout--
			}
		}
		DeleteNode(lockName)
		return "", errors.New("Get Lock Fail timeout")
	} else {
		return lockPath, nil
	}

	// 失败则删除当前节点
	DeleteNode(lockName)
	return "", errors.New("Get Lock Fail")
}

// 释放读锁
func RUnlock(lockName string) (err error ){
	if nil == global.Config.ZkConn{
		if err = ReConnect();err != nil{
			logrus.Errorf("zk connect fail err: %s",err.Error())
		}
	}
	// 删除节点
	err = global.Config.ZkConn.Delete(lockName, 0)
	if err != nil {
		logrus.Errorf("RUnlock Delete %s", err.Error())
		return err
	}

	return nil
}

func GetMaxchild(children []string) (child string) {
	if 0 == len(children) {
		return ""
	}

	var maxChild = children[0]
	maxIndex := maxChild[1:]
	for _, value := range children {
		curIndex := value[1:]
		if curIndex > maxIndex {
			maxIndex = curIndex
		}
		maxChild = value
	}

	return maxChild
}

func GetMaxWritechild(children []string) (child string) {
	//过滤所有写节点
	writeChildren := make([]string, 0)
	for _, value := range children {
		if strings.HasPrefix(value, "w") {
			writeChildren = append(writeChildren, value)
		}
	}

	if 0 == len(writeChildren) {
		return ""
	}

	var maxChild = children[0]
	maxIndex := maxChild[1:]
	for _, value := range children {
		curIndex := value[1:]
		if curIndex > maxIndex {
			maxIndex = curIndex
		}
		maxChild = value
	}

	return maxChild
}

func RegisterNode()(error){
	var hosts = []string{global.Config.ZkIPaddr}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return err
	}
	defer conn.Close()

	var nodePath = "/Cluster/"
	var nodeData []byte = []byte(strconv.FormatInt(time.Now().Unix(), 10))
	var nodeFlags int32 = 4 // 临时增长序列节点
	var acl = zk.WorldACL(zk.PermAll)

	_, err = conn.Create(nodePath, nodeData, nodeFlags, acl)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return  err
	}

	return nil
}


func GetWorkingNode()(int,error){
	var hosts = []string{global.Config.ZkIPaddr}
	conn, _, err := zk.Connect(hosts, time.Second*5)
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return -1, err
	}
	defer conn.Close()

	children, _, err := conn.Children("/Lock")
	if err != nil {
		logrus.Errorf("%s", err.Error())
		return -1, err
	}

	return len(children), nil
}

func DeleteNode(nodeName string)(err error){
	// 删除节点
	err = global.Config.ZkConn.Delete(nodeName, 0)
	if err != nil {
		logrus.Errorf("deleteNode %s", err.Error())
		return err
	}
	return
}


func DeleteAllChildren(dir string)(err error){
	// 获取当前子节点
	children, _, err := global.Config.ZkConn.Children(dir)
	if err != nil {
		logrus.Errorf("RLock Children %s", err.Error())
		return err
	}

	for _, child := range children{
		nodeName := dir + "/" + child
		err = DeleteNode(nodeName)
		if err != nil {
			logrus.Errorf("RLock Children %s", err.Error())
			return err
		}
	}

	return
}