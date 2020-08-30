package queue

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestQueue(t *testing.T) {
	myq := NewQueue()
	myq.EnQueue(1)
	myq.EnQueue(2)
	myq.EnQueue(3)
	if myq.Size() != 3 {
		t.Error("enqueue error")
	}
	front := myq.Front()
	if front != 1 {
		t.Error("front error")
	}
	end := myq.End()
	if end != 3 {
		t.Error("end error")
	}
	data := myq.DeQueue()
	if data != 1 {
		t.Error("dequeue error")
	}
	myq.DeQueue()
	myq.Clear()
	if myq.Size() != 0 {
		t.Error("clear error")
	}
	if !myq.IsEmpty() {
		t.Error("IsEmpty error")
	}
}


// 队列广度遍历文件夹
func TestQueueFolders(t *testing.T) {
	//当前目录
	path, _ := os.Getwd()
	//上一级目录
	path = filepath.Dir(path)
	t.Log("遍历文件夹", path)
	var files = make([]string, 0) //数组字符串
	myqueue := NewQueue()
	myqueue.EnQueue(path)
	for {
		path := myqueue.DeQueue() //不断从队列取出数据
		if path == nil {
			break
		}
		files = append(files, path.(string))     //加入列表
		read, _ := ioutil.ReadDir(path.(string)) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path.(string) + "/" + fi.Name() //构造新的路径
				myqueue.EnQueue(fulldir)
			} else {
				fulldir := path.(string) + "/" + fi.Name() //构造新的路径
				files = append(files, fulldir)             //追加路径
			}
		}
	}
	for i := 0; i < len(files); i++ { //打印
		t.Log(files[i])
	}
}
