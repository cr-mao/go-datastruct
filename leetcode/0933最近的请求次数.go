package leetcode

// import "fmt"

// func main() {
// 	var a = []int{1, 2, 3}
// 	a = a[1:]
// 	fmt.Println(a[0])
// }

type RecentCounter struct {
	list []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		list: make([]int, 0),
	}
}

func (this *RecentCounter) Ping(t int) int {
	for len(this.list) > 0 && this.list[0] < t-3000 {
		this.list = this.list[1:]
	}
	this.list = append(this.list, t)
	return len(this.list)
}
