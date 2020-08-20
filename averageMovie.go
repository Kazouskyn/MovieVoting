package main

func averageMovies(movies []Movie) {
	Data = make(map[string][]float64)

	for _, mov := range movies {
		name := mov.MovieTitle
		vals := mov.VotingValues
		length := len(vals)
		total := 0
		for _, num := range vals {
			total = total + num
		}
		average := float64(total) / float64(length)
		Data[name] = append(Data[name], average)
	}
}
