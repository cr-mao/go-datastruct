package searchUtils

// 顺序查找
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

const fileLinesNum int = 18333811

//用户结构体
type userStruct struct {
	name     string
	md5      string
	email    string
	password string
}

//统计文件行数
func CountFileLines(filepath string) int {
	var file *os.File
	var err error
	var bufioReader *bufio.Reader
	file, err = os.Open(filepath)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	bufioReader = bufio.NewReader(file)
	var i = 0
	for {
		_, _, err := bufioReader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		i++
	}
	return i
}

func SequentialSearch(filePath string) {
	var dataStock []userStruct = make([]userStruct, fileLinesNum, fileLinesNum)
	var file *os.File
	var err error
	var bufioReader *bufio.Reader
	file, err = os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	bufioReader = bufio.NewReader(file)
	starttime:=time.Now()

	i :=0
	for {
		lineBytes, _, err := bufioReader.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		lineStr := string(lineBytes)
		var stringSlice = strings.Split(lineStr, " | ")
		if (len(stringSlice) == 4) {
			var user userStruct
			user.name = stringSlice[0]
			user.md5 = stringSlice[1]
			user.email = stringSlice[2]
			user.password = stringSlice[3]
			dataStock[i]=user
		}
		if i == fileLinesNum{
			break
		}
		i++

	}
	fmt.Println("内存载入完成")
	fmt.Println("内存载入完成用了",time.Since(starttime))
	for ;;{
		fmt.Println("请输入要查询的用户名")
		var inputstr string
		fmt.Scanln(&inputstr) //用户输入
		starttime:=time.Now()
		for i:=0;i<fileLinesNum;i++ {
			if dataStock[i].name == inputstr{
				fmt.Println("找到",inputstr)
			}
		}
		fmt.Println("本次查询用了",time.Since(starttime))
	}
}

func CountFileLinesTest() {
	src := "/Users/mac/Downloads/uuu9.com.sql"
	countNum := CountFileLines(src) //18333811  行数
	fmt.Println(countNum)
}
func SequentialSearchTest(){
	src := "/Users/mac/Downloads/uuu9.com.sql"
	SequentialSearch(src)
}
