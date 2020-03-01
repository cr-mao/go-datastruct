package myutils
//递归遍历文件夹

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)


//递归遍历文件夹
func GetALL(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件加不可读取")
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "/" + fi.Name() //构造新的路径
			files = append(files, fulldir)    //追加路径
			files, _ = GetALL(fulldir, files) //文件夹递归处理
		} else {
			fulldir := path + "/" + fi.Name() //构造新的路径
			files = append(files, fulldir)    //追加路径
		}
	}
	return files, nil
}


//打印层级
func GetALLX(path string, level int) {
	levelstr := ""
	realLevel := level
	if level == 1 {
		levelstr = "|--"
	} else {
		for ; level > 1; level-- {
			levelstr += "|--"
		}
		levelstr += "|--"
	}
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "/" + fi.Name() //构造新的路径
			fmt.Println(strconv.Itoa(realLevel) + levelstr + fulldir)
			GetALLX(fulldir, realLevel+1) //文件夹递归处理

		} else {
			fulldir := path + "/" + fi.Name() //构造新的路径
			fmt.Println(strconv.Itoa(realLevel) + levelstr + fulldir)
		}
	}
}

//栈模拟递归文件并打印层级
func GetALLY(path string, files []string, level int) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件加不可读取")
	}
	levelstr := ""
	realLevel := level
	if level == 1 {
		levelstr = "|--"
	} else {
		for ; level > 1; level-- {
			levelstr += "|--"
		}
		levelstr += "|--"
	}

	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "/" + fi.Name()                               //构造新的路径
			files = append(files, strconv.Itoa(realLevel)+levelstr+fulldir) //追加路径
			fmt.Println(fulldir)
			files, _ = GetALLY(fulldir, files, realLevel+1) //文件夹递归处理
		} else {
			fulldir := path + "/" + fi.Name()                               //构造新的路径
			files = append(files, strconv.Itoa(realLevel)+levelstr+fulldir) //追加路径
		}
	}
	return files, nil
}
