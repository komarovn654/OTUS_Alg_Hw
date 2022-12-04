package hw06simplesorts

import (
	"fmt"
	"time"
)

type sortAnimation struct {
	cmp      []int // compared index
	swaped   bool  // swaped index
	stepNum  int
	cmpCount int
	asgCount int
}

func (sa *sortAnimation) transmitCondition(ch chan sortAnimation, waitResponse bool) {
	ch <- *sa
	if waitResponse {
		<-ch
	}
}

func (sa *sortAnimation) printHeader() {
	fmt.Printf("Step number: %v; Comparisons count: %v; Assignments count: %v\n\r",
		sa.stepNum, sa.cmpCount, sa.asgCount)
}

func (sa *sortAnimation) printCmpStageBubble(array *[]Item) {
	for i := 0; i < len(*array); i++ {
		if sa.cmp[0] == i {
			fmt.Printf("\033[0;30m\033[43m%v> %v\033[0m ", (*array)[i], (*array)[i+1])
			i++
			continue
		}
		fmt.Printf("%v ", (*array)[i])
	}
	fmt.Printf("\r")
	time.Sleep(time.Millisecond * 500)
}

func (sa *sortAnimation) printSwapStageBubble(array *[]Item) {
	for i := 0; i < len(*array); i++ {
		if sa.cmp[0] == i {
			if sa.swaped {
				fmt.Printf("\033[0;30m\033[42m%v->%v\033[0m ", (*array)[i], (*array)[i+1])
			} else {
				fmt.Printf("\033[0;30m\033[41m%v< %v\033[0m ", (*array)[i], (*array)[i+1])
			}
			i++
			continue
		}
		fmt.Printf("%v ", (*array)[i])
	}
	fmt.Printf("\r\033[1A")
	time.Sleep(time.Millisecond * 500)
}

func (sa *sortAnimation) initBubbleAnimation(array *[]Item) chan sortAnimation {
	cntrl := make(chan sortAnimation)

	go func() {
		for {
			sa := <-cntrl
			sa.printHeader()
			sa.printCmpStageBubble(array)
			cntrl <- sortAnimation{}

			sa = <-cntrl
			sa.printSwapStageBubble(array)
			cntrl <- sortAnimation{}
		}
	}()

	return cntrl
}
