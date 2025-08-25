package match

import (
	"domain/team"
)

type Match struct {
	id          string
	team1       team.Team
	team2       team.Team
	startTime   string
	endTime     string
	location    string
	scoreTeam1  int
	scoreTeam2  int
	isCompleted bool
}

func NewMatch(id string, team1 team.Team, team2 team.Team, startTime, endTime, location string, scoreTeam1, scoreTeam2 int, isCompleted bool) *Match {
	return &Match{
		id:          id,
		team1:       team1,
		team2:       team2,
		startTime:   startTime,
		endTime:     endTime,
		location:    location,
		scoreTeam1:  scoreTeam1,
		scoreTeam2:  scoreTeam2,
		isCompleted: isCompleted,
	}
}
