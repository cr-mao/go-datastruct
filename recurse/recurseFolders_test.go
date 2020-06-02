package recurse

import (
	"datastruct/stack"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// 递归遍历文件夹
func TestRecurseFolders(t *testing.T) {
	//当前目录
	path, _ := os.Getwd()
	//上一级目录
	path = filepath.Dir(path)
	t.Log("遍历的文件夹:", path)
	files := []string{}                    //数组字符串
	files, _ = RecurseFolders(path, files) //抓取所有文件
	for i := 0; i < len(files); i++ {      //打印路径
		t.Log(files[i])
	}
}

//栈模拟递归 遍历文件夹
func TestRecursiveFolderWithStack(t *testing.T) {
	//当前目录
	path, _ := os.Getwd()
	//上一级目录
	path = filepath.Dir(path)
	t.Log("遍历的文件夹:", path)
	var files = make([]string, 0)
	mystack := stack.NewStack()
	mystack.Push(path)
	for !mystack.IsEmpty() {
		path := mystack.Pop().(string)
		files = append(files, path)     //加入列表
		read, _ := ioutil.ReadDir(path) //读取文件夹下面所有的路径
		for _, fi := range read {
			if fi.IsDir() {
				fulldir := path + "/" + fi.Name() //构造新的路径
				mystack.Push(fulldir)
			} else {
				fulldir := path + "/" + fi.Name() //构造新的路径
				files = append(files, fulldir)    //追加路径
			}
		}

	}
	for i := 0; i < len(files); i++ { //打印
		t.Log(files[i])
	}
}

//递归文件夹并打印层级1
func TestRecurseFoldersWithLevelPrint(t *testing.T) {
	//当前目录
	path, _ := os.Getwd()
	//上一级目录
	path = filepath.Dir(path)
	t.Log("遍历的文件夹:", path)
	RecurseFoldersWithLevelPrint(path, 1)
}

//递归文件夹并打印层级2
func TestRecursiveFolderWithLevel(t *testing.T) {
	//当前目录
	path, _ := os.Getwd()
	//上一级目录
	path = filepath.Dir(path)
	t.Log("遍历的文件夹:", path)
	var files = make([]string, 0)             //数组字符串
	files, _ = RecursiveFolderWithLevel(path, files, 1) //抓取所有文件
	for i := 0; i < len(files); i++ {  //打印路径
		t.Log(files[i])
	}
}
