var hourMap = map[int][]string{
	0: []string{"0"},
	1: []string{"1", "2", "4", "8"},
	2: []string{"3", "5", "9", "6", "10"},
	3: []string{"7", "11"},
}
var mintueMap = map[int][]string{
	0: []string{"00"},
	1: []string{"01", "02", "04", "08", "16", "32"},
	2: []string{"03", "05", "09", "17", "33", "06", "10", "18", "34", "12", "20", "36", "24", "40", "48"},
	3: []string{"07", "11", "19", "35", "13", "21", "37", "25", "41", "49", "14", "22", "38", "26", "42", "50", "28", "44", "52", "56"},
	//用63减去2那一行就行
	4: []string{"58", "54", "46", "30", "57", "53", "45", "29", "51", "43", "27", "39", "23", "15"},
	5: []string{"59", "47", "31", "55"},
}

func readBinaryWatch(turnedOn int) (ret []string) {

	if turnedOn > 8 {
		return ret
	}
	hour := 0
	for hour <= turnedOn && hour <= 3 {
		hourString := hourMap[hour]
		mintueString := mintueMap[turnedOn-hour]
		for i := 0; i < len(hourString); i++ {
			for j := 0; j < len(mintueString); j++ {
				ret = append(ret, hourString[i]+":"+mintueString[j])
			}
		}

		hour++
	}

	return
}



