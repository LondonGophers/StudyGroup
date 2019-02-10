package popcount

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestAllPopCountsAreEqual(t *testing.T) {
	log.Println(PopCount(0x1234567890ABCDEF))
	log.Println(PopCountLoop(0x1234567890ABCDEF))
	log.Println(PopCountShift(0x1234567890ABCDEF))
	log.Println(PopCountClear(0x1234567890ABCDEF))

	if PopCount(0x1234567890ABCDEF) != PopCountLoop(0x1234567890ABCDEF) {
		t.Errorf("%d should equal %d", PopCount(0x1234567890ABCDEF), PopCountLoop(0x1234567890ABCDEF))
	}

	if PopCount(0x1234567890ABCDEF) != PopCountShift(0x1234567890ABCDEF) {
		t.Errorf("%d should equal (PopCountShift) %d", PopCount(0x1234567890ABCDEF), PopCountShift(0x1234567890ABCDEF))
	}

	if PopCount(0x1234567890ABCDEF) != PopCountClear(0x1234567890ABCDEF) {
		t.Errorf("%d should equal (PopCountClear) %d", PopCount(0x1234567890ABCDEF), PopCountClear(0x1234567890ABCDEF))
	}

}
func timeMe(name string, f func(uint64) int) int64 {
	start := time.Now()
	f(0x1234567890ABCDEF)
	finished := time.Now()
	elapsed := finished.Sub(start)
	return elapsed.Nanoseconds()
}

func TestTimings(t *testing.T) {
	var pcTime, pcLoopTime, pcLoopShift, pcLoopClear int64
	for i := 0; i < 1000000; i++ {
		pcTime += timeMe("PopCount", PopCount)
		pcLoopTime += timeMe("PopCountLoop", PopCountLoop)
		pcLoopShift += timeMe("PopCountShift", PopCountShift)
		pcLoopClear += timeMe("PopCountClear", PopCountClear)
	}

	fmt.Printf("PopCount took an average of %dns\n", pcTime/1000000)
	fmt.Printf("PopCountLoop took an average of %dns\n", pcLoopTime/1000000)
	fmt.Printf("PopCountShift took an average of %dns\n", pcLoopShift/1000000)
	fmt.Printf("PopCountClear took an average of %dns\n", pcLoopClear/1000000)

}