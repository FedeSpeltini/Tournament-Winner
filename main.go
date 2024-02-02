package main

const HOME_TEAM_WON = 1

type Competitor struct {
	Name  string
	Score int
}

func main() {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	println(TournamentWinner(competitions, results))
}

func TournamentWinner(competitions [][]string, results []int) string {

	competitors := []Competitor{}
	for i := 0; i < len(results); i++ {

		if results[i] == HOME_TEAM_WON {

			competitors = addCompetitors(competitors, competitions[i][0], competitions[i][1])
		} else {

			competitors = addCompetitors(competitors, competitions[i][1], competitions[i][0])
		}
	}
	return checkFinalWinner(competitors)
}

func addCompetitors(competitors []Competitor, competitorWinner string, competitorLosser string) []Competitor {
	if len(competitors) == 0 {
		return append(competitors, Competitor{Name: competitorWinner, Score: 3}, Competitor{Name: competitorLosser, Score: 0})
	} else {

		foundWinner := false
		foundLosser := false
		for i, competitor := range competitors {
			if competitor.Name == competitorWinner {
				competitors[i].Score += 3
				foundWinner = true
			}
			if competitor.Name == competitorLosser {
				foundLosser = true
			}

		}
		if !foundWinner {
			return append(competitors, Competitor{Name: competitorWinner, Score: 3})
		}
		if !foundLosser {
			return append(competitors, Competitor{Name: competitorLosser, Score: 0})
		}
		return competitors
	}
}

func checkFinalWinner(competitors []Competitor) string {

	maxPoint := 0
	winner := ""
	for _, competitor := range competitors {
		if competitor.Score > maxPoint {
			maxPoint = competitor.Score
			winner = competitor.Name
		}
	}
	return winner
}
