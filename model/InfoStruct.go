package model

import "fmt"

type InfoStruct struct {
	Id                     int    `json:"id"`
	DisplayOs              string `json:"displayOs"`
	Category               string `json:"category"`
	Subject                string `json:"subject"`
	Text                   string `json:"text"`
	HtmlPath               string `json:"htmlPath"`
	ImgPath                string `json:"imgPath"`
	HiddenUnderMaintenance bool   `json:"hiddenUnderMaintenance"`
	StartAt                string `json:"startAt"`
	EndAt                  string `json:"endAt"`
	SortKey                int    `json:"sortKey"`
}

func String(s *InfoStruct) string {
	return fmt.Sprintf("[subject=%v\n,text=%v\n,startAt=%v]", s.Subject, s.Text, s.StartAt)
}

type HtmlStruct struct {
	SubText string
	Text    string
}
