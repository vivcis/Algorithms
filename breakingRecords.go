package main

func breakingRecords(scores []int32) []int32 {

	var minCount int32
	var maxCount int32
	var maxScore int32 = scores[0]
	var minScore int32 = scores[0]
	var result []int32

	if len(scores) > 1 {
		for i := 0; i < len(scores); i++ {
			if scores[i] < minScore {
				minScore = scores[i]
				minCount++

			}
			if scores[i] < maxScore {
				maxScore = scores[i]
				maxCount++

			}
		}
	}
	result = append(result, minCount, maxCount)
	return result
}
