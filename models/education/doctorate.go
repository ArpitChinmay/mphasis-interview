package education

type Doctorate struct {
	CandidateId int    `json:"candidateId"`
	College     string `json:"college"`
	University  string `json:"university"`
	Start_year  int    `json:"startYear"`
	End_year    int    `json:"endYear"`
}
