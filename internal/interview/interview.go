package interview

import "fmt"

func OptimalPoint() int32 {
	magic := []int32{2, 4, 5, 2}
	dist := []int32{4, 3, 1, 2}

	initPosition := int32(0)
	currPosition := int32(0)
	length := int32(len(magic))

	for magic[initPosition] < dist[initPosition] {
		initPosition++
		currPosition++
	}

	calc := magic[initPosition]
	for {
		fmt.Printf("calc: %v, curr: %v, init: %v, magic: %v, dist: %v \n", calc, currPosition,
			initPosition, magic[(currPosition+int32(1))%length], dist[currPosition])
		// fmt.Printf("curr: %v \t", currPosition)

		// subtract magic by distance
		calc = calc - dist[currPosition]
		if calc < 0 {
			// if there is no more magic left, move onto the next iteration
			for calc < 0 && initPosition != currPosition {
				calc = calc - magic[initPosition] + dist[initPosition]
				initPosition++

				// no more indexes left to try
				if initPosition == length {
					return -1
				}
			}
		}

		// add magic
		calc += magic[(currPosition+int32(1))%length]
		currPosition++

		// this is to avoid index out of bounds errors
		if currPosition == length {
			currPosition = 0
		}

		if currPosition == initPosition {
			return initPosition
		}
	}
}

/* Naive solution is to brute force each calculation. For each element in magic array, calculate
*  Aladdin's journey. Then return the minimum viable index. Sounds like a DP problem.
 */

/* Optimized solution uses the previous solution to calculate the next solution.
 */
