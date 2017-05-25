package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func weekDaysCount(years []int, months [12]time.Month) error {
	f, err := os.Create("./days.csv")
	if err != nil {
		return err
	}
	defer f.Close()
	w := csv.NewWriter(f)
	defer w.Flush()
	for _, year := range years {
		for _, month := range months {
			count := 0
			for d := time.Date(year, month, 1, 0, 0, 0, 0, &time.Location{}); time.Date(year, month+time.Month(1), 0, 0, 0, 0, 0, &time.Location{}).Sub(d) >= 0; d = d.Add(24 * time.Hour) {
				if d.Weekday() != time.Sunday && d.Weekday() != time.Saturday {
					count++
				}
			}
			w.Write([]string{fmt.Sprintf("%d", year), fmt.Sprintf("%d", month), fmt.Sprintf("%d", count)})
		}
	}

	return nil
}

func main() {
	months := [12]time.Month{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	years := []int{2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024}
	if err := weekDaysCount(years, months); err != nil {
		panic(err)
	}
}
