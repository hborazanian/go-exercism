package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0

	for _, value := range birdsPerDay {
		count += value
	}

	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	weekEndIndex := week * 7
	weekStartIndex := weekEndIndex - 7
	weekArray := birdsPerDay[weekStartIndex:weekEndIndex]

	return TotalBirdCount(weekArray)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for index := range birdsPerDay {
		if index == 0 || index%2 == 0 {
			birdsPerDay[index]++
		}
	}

	return birdsPerDay
}
