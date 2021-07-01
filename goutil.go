package goutil

import (
	"fmt"
	"math/rand"
	"time"
)

func SayHello() {
	fmt.Println("hello qzp")
}

func FindType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

func GetRandomWelcome() {
	t := time.Now()
	dateFmt := t.Format("2006-01-02 15:04:05") //go诞生的时间
	logStr := "[当前的时间是:" + dateFmt + "]"

	sWelcomeList := []string{"come on qzp", "happy everyday qzp", "take easy qzp", "you are best qzp", "relax yourself qzp"}
	iWelcomeListLen := len(sWelcomeList)
	// for idx := 0; idx < sWelcomeLen; idx++ {
	// 	fmt.Println(sWelcomeList[idx])
	// }

	rand.Seed(time.Now().UnixNano())
	iRandNum := rand.Intn(iWelcomeListLen)
	logStr += sWelcomeList[iRandNum]
	fmt.Println(logStr)
}
