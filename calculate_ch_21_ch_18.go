package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	dtRef := time.Date(2022, time.August, 11, 0, 0, 0, 0, time.UTC)
	//dtNow := time.Now()
	//dtNow := time.Date(2021, time.August, 30, 0, 0, 0, 0, time.UTC)
	//dtNow := time.Date(2023, time.July, 23, 0, 0, 0, 0, time.UTC)
	dtNow := dtRef.Add(-55 * 24 * time.Hour)
	dtNowStr := dtNow.Format("January 02, 2006")

	dtDiff := dtNow.Sub(dtRef)
	diff := math.Round(dtDiff.Hours() / 24)
	//diff := float64(int(dtDiff.Hours() / 24))
	sign := diff / math.Abs(diff)
	diff = math.Abs(diff)

	diff_mod_21 := math.Mod(diff, 21)
	diff_mod_18 := math.Mod(diff, 18)
	if sign < 0 {
		diff_mod_21 = 21 - diff_mod_21
		diff_mod_18 = 18 - diff_mod_18
	}

	ch_21 := int(diff_mod_21 + 1)
	ch_18 := int(diff_mod_18 + 1)

	if ch_21 == 22 {
		ch_21 = 1
	}
	if ch_18 == 19 {
		ch_18 = 1
	}

	fmt.Println(dtNowStr)
	fmt.Println("ch_21 = ", ch_21, ", ch_18 = ", ch_18)

}
