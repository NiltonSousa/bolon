package domain

type Match struct {
	ID          string `json:"id"`
	Team1       Team   `json:"team1" validate:"required"`
	Team2       Team   `json:"team2" validate:"required"`
	StartTime   string `json:"start_time" validate:"required"`
	EndTime     string `json:"end_time" validate:"required"`
	Location    string `json:"location" validate:"required"`
	ScoreTeam1  int    `json:"score_team1" validate:"required"`
	ScoreTeam2  int    `json:"score_team2" validate:"required"`
	IsCompleted bool   `json:"is_completed" validate:"required"`
}

func NewMatch(id string, team1 Team, team2 Team, startTime, endTime, location string, scoreTeam1, scoreTeam2 int, isCompleted bool) *Match {
	return &Match{
		ID:          id,
		Team1:       team1,
		Team2:       team2,
		StartTime:   startTime,
		EndTime:     endTime,
		Location:    location,
		ScoreTeam1:  scoreTeam1,
		ScoreTeam2:  scoreTeam2,
		IsCompleted: isCompleted,
	}
}
