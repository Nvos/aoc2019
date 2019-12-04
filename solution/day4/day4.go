package day4

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func Standard(from int, to int) int {
	logger := logrus.New()
	count := 0
	for i := from; i <= to; i++ {
		if isMatchingBonusRules(i) {
			count += 1

			logger.Info("Value = ", i, " is matching")
		}
	}

	return count
}

func isMatchingBonusRules(value int) bool {
	itoa := strconv.Itoa(value)
	var prev uint8 = 0
	var curr uint8 = 0
	var blacklist uint8 = 0
	hasAdjecent := false
	l := len(itoa)

	for i := 0; i < l; i++ {
		if curr != 0 {
			prev = curr
		}

		curr = itoa[i]
		if prev == 0 {
			continue
		}

		if curr < prev {
			return false
		}

		if curr == blacklist {
			continue
		}

		if prev == curr {
			if i+1 < l {
				if itoa[i + 1] != curr {
					hasAdjecent = true
				} else {
					blacklist = curr
				}
			} else {
				hasAdjecent = true
			}
		}
	}

	return hasAdjecent
	//for i, value := range itoa {
	//	if curr != -1 {
	//		prev = curr
	//	}
	//
	//	curr = value
	//	if prev == -1 {
	//		continue
	//	}
	//
	//	if curr < prev {
	//		return false
	//	}
	//
	//	if hasAdjascent {
	//		continue
	//	}
	//
	//	if adjacent == curr {
	//		adjacent = -1
	//		blacklist = curr
	//		continue
	//	}
	//
	//	if curr == prev && curr != blacklist {
	//		adjacent = curr
	//		println(int32(itoa[i]))
	//		if i < len(itoa) && int32(itoa[i]) != adjacent  {
	//			println("!")
	//			hasAdjascent = true
	//		}
	//	}
	//}
	//
	//return adjacent != -1
}

func isMatchingStandardRules(value int) bool {
	itoa := strconv.Itoa(value)
	var prev int32 = -1
	var curr int32 = -1

	hasAdjacent := false

	for _, value := range itoa {
		if curr != -1 {
			prev = curr
		}

		curr = value
		if prev == -1 {
			continue
		}

		if curr < prev {
			return false
		}

		if curr == prev {
			hasAdjacent = true
			continue
		}
	}

	return hasAdjacent
}
