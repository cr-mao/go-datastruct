package recurse
//递归遍历文件夹

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
)


//递归遍历文件夹
func RecurseFolders(path string, files []string) ([]string, error) {
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return files, errors.New("文件加不可读取")
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "/" + fi.Name() //构造新的路径
			files = append(files, fulldir)    //追加路径
			files, _ = RecurseFolders(fulldir, files) //文件夹递归处理
		} else {
			fulldir := path + "/" + fi.Name() //构造新的路径
			files = append(files, fulldir)    //追加路径
		}
	}
	return files, nil
}


//递归遍历文件夹打印层级
func RecurseFoldersWithLevelPrint(path string, level int) {
	levelstr := ""
	realLevel := level
	if level == 1 {
		levelstr = "|--"
	} else {
		levelstr += "|--"
		for ; level > 1; level-- {
			levelstr += "|--"
		}
	}
	read, err := ioutil.ReadDir(path) //读取文件夹
	if err != nil {
		return
	}
	for _, fi := range read { //循环每个文件或者文件夹
		if fi.IsDir() { //判断是否文件夹
			fulldir := path + "/" + fi.Name() //构造新的路径
			fmt.Println(strconv.Itoa(realLevel) + levelstr + fulldir)
			RecurseFoldersWithLevelPrint(fulldir, realLevel+1) //文件夹递归处理
		} else {
			fulldir := path + "/" + fi.Name() //构造新的路径
			fmt.Println(strconv.Itoa(realLevel) + levelstr + fulldir)
		}
	}
}

//递归文件夹显示层级并放到数组中
func RecursiveFolderWithLevel(path string, files []string, level int) ([]string, error) {
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
			files, _ = RecursiveFolderWithLevel(fulldir, files, realLevel+1) //文件夹递归处理
		} else {
			fulldir := path + "/" + fi.Name()                               //构造新的路径
			files = append(files, strconv.Itoa(realLevel)+levelstr+fulldir) //追加路径
		}
	}
	return files, nil
}
