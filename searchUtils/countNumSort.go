package searchUtils



import "fmt"
//计数排序，有序情况下
func CountNumSortTest(){
	mylist :=[]string{"1111","222","222","222","222","3333","3333","4444","4444","55"}
	length:=len(mylist)
	i:=0
	for i<length{
		times:=1
		password:=mylist[i]
		for i+1<=length-1 && mylist[i]==mylist[i+1]{
			i++
			times++
		}
		fmt.Println(password,times)
		i++
	}
}
