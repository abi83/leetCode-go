package main

import (
	"leetCode/GoRoutine/payload"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(8)
	payload.RunHeavyComputationsSync()
	//payload.RunHeavyComputationsInRoutites()

	//payload.ConcurrentStringRotation()
	//payload.RunTimer()
	//payload.AgeRunner()
}
