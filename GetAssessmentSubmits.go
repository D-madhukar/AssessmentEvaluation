package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"time"
	"github.com/gorilla/mux"
	"log"
	"bytes"
	"io/ioutil"
)

type AssessmentSubmit struct {
	SchoolId	   string					     `json:"schoolid,omitempty"`
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
	SchoolId 	 		string 								 `json:"schoolid"`
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
type SubmissionPost struct{
	Submittable		bool 		`json:"submittable"`
	Parent 			string 		`json:"parent"`
	Type			string 		`json:"type"`
	ResponseType	string		`json:"responseType"`
	School			string		`json:"school"`
	Class			string		`json:"class"`
	Author			string		`json:"author"`
	// Users			[]interface{}`json:"users"`
	// Poll            interface{}	`json:"poll"`
	Assessment		struct{
				Answered	int	 `json:"answered"`
				Correct		int	 `json:"correct"`
				Status		string`json:"status"`
				Timespent	int	 `json:"timespent"`
				Attempts	int  `json:"attempts"`
			}					`json:"assessment"`
	Stars		    int			`json:"stars"`
}
var assessments []Assessment
var assessmentsubmits []AssessmentSubmit
func main(){
	router:=mux.NewRouter()
	router.HandleFunc("/postassessmentsubmissions",getAssessmentSubmissions).Methods("POST")
	router.HandleFunc("/getassessments",getAssessments)
	log.Fatal(http.ListenAndServe(":3132",router))
}
func getAssessmentSubmissions(w http.ResponseWriter,r *http.Request){
	json.NewDecoder(r.Body).Decode(&assessmentsubmits)

	postid:=assessmentsubmits[0].Post
	// schoolid:=assessmentsubmits[0].SchoolId
	// response,_:=http.Get("http://localhost:3131/api/school/"+schoolid+"/assessments/"+postid)
	response,_:=http.Get("http://localhost:3131/getassessments/"+postid)
	body,_:=ioutil.ReadAll(response.Body)
	err1:=json.Unmarshal(body,&assessments)


	var flag string="0"
	var submissionpostid string=""
	if(assessmentsubmits[0].AssessmentPost!=""){
		flag="1"
		submissionpostid=assessmentsubmits[0].AssessmentPost
	}
	notanswered,correct,score:=evaluateSubmissions()
	submissionpost:=prepareSubmissionPost(notanswered,correct,score)
	submissionpostjson,_:=json.Marshal(submissionpost)
	var url string="http://localhost:3131/savesubmissionpost/"+flag+"/"+submissionpostid
	fmt.Println("url-->",url);
	res,err:=http.Post(url,"application/json",bytes.NewBuffer(submissionpostjson))
		if err!=nil{
			fmt.Println("error",err);
		}else{
			fmt.Println("submissionpost",submissionpost);
		}
	data,_:=ioutil.ReadAll(res.Body)
	var id string;
	err=json.Unmarshal(data,&id)
	fmt.Println(id);
	initAssessmentSubmits(id)
	jsondata,_:=json.Marshal(assessmentsubmits)
	_,err=http.Post("http://localhost:3131/savesubmissions","application/json",bytes.NewBuffer(jsondata))
	if err!=nil{
		fmt.Println("error in saving submissions");
	}else{
		fmt.Println("submissions sent successfully")
	}

}
func initAssessmentSubmits(submitpostid string){
	for i,_:=range assessmentsubmits{
		assessmentsubmits[i].AssessmentPost=submitpostid;
	}
}
func prepareSubmissionPost(notanswered,correct,post int) SubmissionPost{
	var submissionpost SubmissionPost
	answered:=len(assessments)-notanswered
	submissionpost.Submittable=true
	submissionpost.Parent=assessmentsubmits[0].Post
	submissionpost.Type="Assessment"
	submissionpost.ResponseType="submission"
	submissionpost.Class="5b0be3225dd45e2250f45e2f"
	submissionpost.School="5b0be2ee5dd45e2250f45e09"
	submissionpost.Author=assessmentsubmits[0].User
	// submissionpost.Users=make([]interface{},0)
	submissionpost.Assessment.Answered=answered
	submissionpost.Assessment.Correct=correct
	submissionpost.Assessment.Status="success"
	submissionpost.Assessment.Timespent=30
	submissionpost.Assessment.Attempts=5
	submissionpost.Stars=1
	return submissionpost
}
func getAssessments(w http.ResponseWriter,r *http.Request){
	fmt.Println("got the assessments")
	json.NewDecoder(r.Body).Decode(&assessments)
}

//Evaluating the submissions
func evaluateSubmissions()(int,int,int){
	var notanswered int=0
	var correct int =0
	var score int=0
	for i,submission:=range assessmentsubmits{
			switch submission.Type{
				case "TF":{
					if submission.Answer=="not_answered"{
						notanswered++
						assessmentsubmits[i].Status="not answered"
					}else if submission.Answer==assessments[i].Answer{
						correct++
						score=score+assessments[i].Marks
						assessmentsubmits[i].Status="correct"
					}else{
						assessmentsubmits[i].Status="not correct"
					}
				}
				case "SA":{
					if submission.Answer=="not_answered"{
						notanswered++
						assessmentsubmits[i].Status="not answerd"
					}else if checkShortAnswer(i){
						correct++
						score=score+assessments[i].Marks
						assessmentsubmits[i].Status="correct"
					}else{
						assessmentsubmits[i].Status="not correct"
					}
				}
				case "MC":{
					if checkMcAttempt(i){
						notanswered++
						assessmentsubmits[i].Status="not answered"
					}else if checkMultipleChoice(i){
						correct++
						score=score+assessments[i].Marks
						assessmentsubmits[i].Status="correct"
					}else{
						assessmentsubmits[i].Status="not correct"
					}
				}
			}

	}
	return notanswered,correct,score
}

//Checking whether Student attempted the MC Question or not.
func checkMcAttempt(i int) bool{
	fmt.Println("checkmc##",i)
	for _,option:=range assessmentsubmits[i].Options{
		if option.Answer==true{
			return false
		}
	}
	return true
}

//Checking SA correctness
func checkShortAnswer(i int ) bool{
	fmt.Println("checkshort",i)
	var studentanswer=assessmentsubmits[i].Answer
	for _,option:=range assessments[i].Options{
		if option.Text==studentanswer{
			return true
		}
	}
	return false
}

//checking MC correctness
func checkMultipleChoice(i int)bool{
	fmt.Println("checkMC@@",i)
	for j,option:=range assessmentsubmits[i].Options{
		if option.Answer!=assessments[i].Options[j].Answer{
			return false
		}
	}
	return true
}
