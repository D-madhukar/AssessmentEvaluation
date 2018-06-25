package main
import (
	"time"
)
type AssessmentSubmit struct {
	ID             string       				 `json:"_id"`
	AssessmentPost string      					 `json:"assessmentPost"`
	User           string        				 `json:"user"`
	Assessment     string       				 `json:"assessment"`
	Type           string       				 `json:"type"`
	Status         string       				 `json:"status"`
	TimeSpent      int         				     `json:"timeSpent"`
	Answer         string       				 `json:"answer,omitempty"`
	Post           string       				 `json:"post"`
	CreatedBy      string        				 `json:"created_by"`
	Options        []struct{
							Seqno    string `json:"seqno"`
							Text     string `json:"text"`
							Answer   bool   `json:"answer"`
							Isanswer bool   `json:"isanswer"`
					} 							 `json:"options,omitempty"`
}
type Assessment struct {
	ID        string        		`json:"_id"`
	CreatedAt time.Time     		`json:"created_at"`
	UpdatedAt time.Time     		`json:"updated_at"`
	Seqno     int           		`json:"seqno"`
	Type      string        		`json:"type"`
	Answer    string        		`json:"answer,omitempty"`
	Marks     int           		`json:"marks"`
	Question  string        		`json:"question"`
	CreatedBy string        		`json:"created_by"`
	V         int           		`json:"__v"`
	Post      string        		`json:"post"`
	Options   []struct{
					Text   string `json:"text"`
					Seqno  string `json:"seqno"`
					Answer bool   `json:"answer"`
					ID     string `json:"_id"`
				}				 	`json:"options,omitempty"`
}