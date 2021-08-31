package main

import (
	"log"
	"sort"
)

type teams struct{
	Name string
	Wins int
	Lose int
	Draw int
	GoalFor int
	GoalAgainst int
	ScoreTotal int
}
func main()  {
	var i int = 0
	rawData := []struct{
		Name string
		Wins int
		Lose int
		Draw int
		GoalFor int
		GoalAgainst int
		Rank *int
		ScoreTotal *int
	}{
		{"A",5,1,2,23,20, &i,&i},
		{"B",6,1,1,24,20,&i,&i},
		{"C",7,1,0,24,10,&i,&i},
		{"D",6,1,1,15,11,&i,&i},
		{"E",8,0,1,23,12,&i,&i},
	}

	for i, t := range rawData{
		scoreTotal := (t.Wins*3)+(t.Draw*1)+(t.Lose*0)
		rawData[i].ScoreTotal = &scoreTotal
		//fmt.Println("Score total ", scoreTotal)
	}

	// Sort by total, keeping original order or equal elements.
	sort.SliceStable(rawData, func(i, j int) bool {
		if *rawData[i].ScoreTotal > *rawData[j].ScoreTotal {
			return true
		}else if *rawData[i].ScoreTotal == *rawData[j].ScoreTotal {
			if rawData[i].GoalFor-rawData[i].GoalAgainst > rawData[j].GoalFor-rawData[j].GoalAgainst {
				return true
			}else if rawData[i].GoalFor-rawData[i].GoalAgainst == rawData[j].GoalFor-rawData[j].GoalAgainst{
				if rawData[i].GoalFor > rawData[j].GoalFor {
					return true
				}
			}
		}

		return false
	})

	var resp []teams
	for _, v := range rawData {
		a := teams{
			Name: v.Name,
			Wins: v.Wins,
			Lose: v.Lose,
			Draw: v.Draw,
			GoalFor: v.GoalFor,
			GoalAgainst: v.GoalAgainst,
			ScoreTotal: *v.ScoreTotal,
		}

		resp = append(resp, a)
	}
	log.Printf("Rank %+v", resp)


}
