package main
import "sync"
import "fmt"

type Counter struct{
	count int
}

var instance *Counter
var once sync.Once


func GetCounterInterface()*Counter{
	once.Do(func(){
		instance = $Counter{}
	})
	return instance
}

func(c *Counter) Increment{
	c.count++
}

func (c *Counter) GetCount{
	return c.count
}

func main(){
	count1 := GetCounterInterface()
	count2 := GetCounterInterface()

	count1.Increment()
	fmt.Println(count1.GetCount)
	fmt.Println(count2.GetCount)

	count2.Increment()
	fmt.Println(count1.GetCount)
	fmt.Println(count2.GetCount) // So it increases for 2 don't matter if they are not same variables, but they refer to one object


}