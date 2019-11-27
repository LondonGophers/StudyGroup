package main

import (
	"crypto/sha256"
	"fmt"

	popcount "github.com/LondonGophers/StudyGroup/workspaces/jlucktay/ch2/ex2_5"
)

func main() {
	bdcv := BitDiffCountVerbose(sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X")))

	fmt.Println("Per page 83 and as seen above, 'approximately half the bits are different' between 'x' and 'X'.")
	fmt.Printf("The actual result is '%d'.\n\n", bdcv)

	fmt.Printf("The count for 'hello' vs 'world' is '%d'.\n",
		BitDiffCount(sha256.Sum256([]byte("hello")), sha256.Sum256([]byte("world"))))
}

func BitDiffCountVerbose(alpha, bravo [sha256.Size]byte) int {
	return bitDiffCount(alpha, bravo, true)
}

func BitDiffCount(alpha, bravo [sha256.Size]byte) int {
	return bitDiffCount(alpha, bravo, false)
}

func bitDiffCount(alpha, bravo [sha256.Size]byte, verbose bool) (result int) {
	for index := 0; index < sha256.Size; index++ {
		pc := popcount.PopCount(uint64(alpha[index] ^ bravo[index]))
		result += pc

		if verbose {
			fmt.Printf("%3[1]d^%-3[2]d: %8[1]b\n%17[2]b\n%17[3]b = %3[3]d (%d bits differ, %d accumulated)\n\n",
				alpha[index],
				bravo[index],
				alpha[index]^bravo[index],
				pc,
				result)
		}
	}

	return
}
