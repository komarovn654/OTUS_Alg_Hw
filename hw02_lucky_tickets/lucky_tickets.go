package hw02_lucky_tickets

const (
	lastDigit = 9
)

func calculateSum(refSums []int, n int, target int) int {
	digitSums, ticketsCount := calculateSumWrap(refSums, n)
	if n == target {
		return ticketsCount
	}
	return calculateSum(digitSums, n+1, target)
}

func calculateSumWrap(refSums []int, n int) ([]int, int) {
	refSums = append(refSums, make([]int, 9)...)
	digitSums := make([]int, 10+(n-1)*9)
	digitSums[0] = 1
	ticketsCount := 0
	for i := 1; i < len(digitSums); i++ {
		startIndex := i - 9
		if i-9 < 0 {
			startIndex = 0
		}
		for _, v := range refSums[startIndex : i+1] {
			digitSums[i] += v
		}
		ticketsCount += digitSums[i] * digitSums[i]
	}
	return digitSums, ticketsCount + 1
}

func LuckyTickets(n int) int {
	if n == 1 {
		return 10
	}

	firstSlice := make([]int, 10)
	for i := 0; i < 10; i++ {
		firstSlice[i] = 1
	}

	return calculateSum(firstSlice, 2, n)
}
