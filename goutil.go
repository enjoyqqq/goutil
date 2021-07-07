package goutil

import (
	"fmt"
	"math/rand"
	"strconv"
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

func MyPrintf(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is an int value.")
		case string:
			fmt.Println(arg, "is a string value.")
		case int64:
			fmt.Println(arg, "is an int64 value.")
		default:
			fmt.Println(arg, "is an unknown type.")
		}
	}
}

func GetRandomWelcome() {
	sWelcomeList := []string{"come on qzp", "happy everyday qzp", "take easy qzp", "you are best qzp", "relax yourself qzp", "please laugh haha qzp"}
	iWelcomeListLen := len(sWelcomeList)
	// for idx := 0; idx < sWelcomeLen; idx++ {
	// 	fmt.Println(sWelcomeList[idx])
	// }

	t := time.Now()
	dateFmt := t.Format("2006-01-02 15:04:05") //go诞生的时间
	logStr := "[当前的时间是:" + dateFmt + "]" + "欢迎语长度为" + strconv.Itoa(iWelcomeListLen) + "，欢迎语为"

	rand.Seed(time.Now().UnixNano())
	iRandNum := rand.Intn(iWelcomeListLen)
	logStr += sWelcomeList[iRandNum]
	fmt.Println(logStr)
}
