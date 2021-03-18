package main

import (
	"fmt"
	"reflect"
)

func main() {
	ch := make(chan int, 1)
	chVal := reflect.ValueOf(ch)
	succeeded := chVal.TrySend(reflect.ValueOf(123))
	fmt.Println(succeeded, chVal.Len(), chVal.Cap())
	fmt.Println()

	vSend := reflect.ValueOf(789)
	vRecv := reflect.Value{}

	fmt.Println(vSend)
	fmt.Println(vRecv)
	fmt.Println()

	branches := []reflect.SelectCase{
		{Dir: reflect.SelectDefault, Chan: vRecv, Send: vRecv},
		{Dir: reflect.SelectRecv, Chan: chVal, Send: vRecv},
		{Dir: reflect.SelectSend, Chan: chVal, Send: vSend},
	}

	selChosen, selRecv, selRecvOk := reflect.Select(branches)
	fmt.Println("chosen: ", selChosen)
	fmt.Println("recvOk: ", selRecvOk)
	fmt.Println("recv: ", selRecv.Int())
	fmt.Println()
	chVal.Close()

	selChosen, _, selRecvOk = reflect.Select(branches[:2])
	fmt.Println("chosen: ", selChosen)
	fmt.Println("recvOk: ", selRecvOk)
	fmt.Println()

}
