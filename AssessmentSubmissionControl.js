//Importing the modules

var Assessment=require('./models/Assessment.js');
var Post = require('./models/Post.js');
var AssessmentSubmit = require('./models/AssessmetSubmit.js');
var mongoose = require('mongoose');
var express=require("express");
var request=require("request");
var bodyParser=require("body-parser");
var router=express();


    //ROUTE
    mongoose.connect("mongodb://192.168.0.9:27017/test");
    router.use(bodyParser.urlencoded({extended:true}));
    router.use(bodyParser.json());
    router.post("/postsubmissions/:userid/:postid",(req,res)=>{
        Post.find({"author":req.params.userid,"parent":req.params.postid})
        .exec(function(error,submissionpost){
            if(error){console.log("database error"+error);}
            else{
                   var submissions=req.body;
                   var submissionpostid="";
                   if(!submissionpost){
                          var assessmentpost=new Post(submissionpost);
                          submissionpostid=submissionpost[0]._id;
                    }
                    if(submissionpostid!=""){
                          submissions[0].assessmentpost=submissionpostid;
                    }
                    request.post({
                        "headers":{"content-type":"application/json"},
                        "url":"http://localhost:3132/postassessmentsubmissions",
                        "body":JSON.stringify(submissions)
                    },(error,response,body)=>{
                        if(error){
                            console.log("'Failure'-submissions are not sent to Golang Server"+error);
                        }
                        else{
                            console.log("'Success'-submissions are sent to Golang server")
                            // res.status(200).send("ok");
                        }
                    });
            }
        });

    });
    router.post('/savesubmissionpost/:flag/:submissionpostid/',function(req,res){
        var submissionpost=req.body;
        console.log(submissionpost);
        if(req.params.flag === "1"){
            Post.findOneAndUpdate({"_id":req.params.submissionpostid},submissionpost)
            .exec(function (error,submission){
                if(error){console.log("error in updating the submissionpost");}
                else{
                  if(submission!=null && submission._id!=null)
                    console.log("submissionpost updated successfully");
                    res.status(200).send(submission._id);
                }
            });
        }
        else{
            var SaveSubmissionPost=new Post(submissionpost);
            SaveSubmissionPost.save(function(error,submission){
                if(error){
                    console.log("server error"+error);
                    res.status(500).send(error);
                }
                else{
                    console.log("submission_id"+submission._id);
                    res.status(200).send(submission._id);
                }
            });
        }
    });
    router.post('/savesubmissions',function(req,res){
        var assessmentsubmissions=req.body;
        var submissionids=[];
        console.log(assessmentsubmissions);
        // assessmentsubmissions.forEach(function(submission){submissionids.push(submission.post);});
        AssessmentSubmit.deleteMany({user:assessmentsubmissions[0].user,post:assessmentsubmissions[0].post})
        .exec((error,submissions)=>{
            if(error)
                console.log("ERROR!");
            else{
                if(!submissions)
                  console.log("Old Submissions are removed");
                else
                  console.log("No submissions are present for this post");
            }
        });
        assessmentsubmissions.forEach((submission)=>{
                var newSubmission=new AssessmentSubmit(submission);
                newSubmission.save((error,nsubmission)=>{
                  if(error)
                    console.log("error in storing"+error);
                  else {
                    console.log("storing");
                  }
                });
        });

    });

    router.get('/getassessmentpost/:userid/:parentid',(req,res)=>{
        Post.find({"author":req.params.userid,"parent":req.params.parentid})
        .exec((err,submissionpost)=>{
          if(!submissionpost)
              res.status(200).send(assessmentpost)
          else
              res.send("No data found");
        });
    });
    router.get('/getassessments/:postid',(req,res)=>{
        Assessment.find({"post":req.params.postid})
        .exec((error,assessments)=>{
            if(error){
              console.log("error in getting the  asessments");
              res.status(500).send(error);
            }
            else{
              console.log("assessments getting"+assessments);
                res.send(assessments);
            }
        });
    });
    router.get("/insertdata",(req,res)=>{
        var newSubmit=new AssessmentSubmit(
          {
          "assessment":"12bca1121",
          "user":"232242421",
          "assessmentpost":"24234234234bc",
          "created_at":"2018-05-28T11:17:52.926Z",
          "updated_at":"2018-05-28T11:19:03.331Z",
          "seqno":1,"type":"TF","answer":"true","marks":1,"question":"My name is madhukar",
          "explanation":"<p>True</p>",
          "created_by":"5b0be2d65dd45e2250f45dfe",
          "options":[],
          "__v":0,
          "post":"5b0be5f35dd45e2250f45eaf"
         }
        );
        newSubmit.save((err)=>{
          if(err){console.log("error"+err);}
          else{console.log("inserted success");}
        });
    });


//Assigning the port
router.listen(3131,()=>{
  console.log("server listening at port 3131");
});
