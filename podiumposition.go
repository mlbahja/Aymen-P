package piscine

func PodiumPosition(podium [][]string) [][]string {
	res := [][]string(podium)

	for range podium {
		for i := 1; i < len(podium)-1; i++ {
			if podium[i][0] > podium[i+1][0] {
				res[i], res[i+1] = res[i+1], res[i]
			}
		}
	}
	return res
}
