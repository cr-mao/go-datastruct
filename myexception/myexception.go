package myexception

import (
	"fmt"
	"time"
)

type InvalidRadiusError struct {
	//非法半径
	InValidRadius  float64
	//合法的最小半径
	MinValidRadius float64
	//合法的最大半径
	MaxValidRadius float64
	//异常的发生时间
	errTime        time.Time
}

func GetCircleArea(radius float64) (area float64, err error) {

	if radius < 10 || radius > 50 {
		return 0, NeWInvalidRadiusError(radius,10,50)
	}

	return 3.14 * radius * radius, nil
}

//实现error 接口
func (ire *InvalidRadiusError) Error() string {
	return ire.String()
}
//异常被打印的方式
func (ire InvalidRadiusError)String() string {
	return fmt.Sprintf("invalidRadius{%.2f是非法半径，合法半径[%.2f,%.2f],错误发生时间是%v}",ire.InValidRadius,ire.MinValidRadius,ire.MaxValidRadius,ire.errTime)
}

//创建自定义异常的工厂方法
func NeWInvalidRadiusError(inValidRadius,minRadius,maxRadius float64) *InvalidRadiusError{
	ire :=new(InvalidRadiusError)
	ire.InValidRadius=inValidRadius
	ire.MinValidRadius=minRadius
	ire.MaxValidRadius=maxRadius
	ire.errTime=time.Now()
	return ire
}

func MyexceptionTest(){
	area,err :=GetCircleArea(-5)
	fmt.Println(err)
	fmt.Println(area)
}

