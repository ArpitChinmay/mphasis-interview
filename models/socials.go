package models

type Socials struct {
	SocialId      int    `json:"socialId"`
	Linkedin      string `json:"linkedinURL"`
	Github        string `json:"githubURL"`
	StackOverflow string `json:"stack-overflowURL"`
}
