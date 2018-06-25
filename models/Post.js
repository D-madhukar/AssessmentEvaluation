var mongoose=require("mongoose");
mongoose.connect("mongodb://localhost/test");
var Schema=mongoose.Schema;
var PostSchema = new Schema(
  {
    // _id:String,  
    type: {
      type: String,
      enum: [
        "Discussion",
        "Message",
        "Assignment",
        "Notification",
        "Blog",
        "Material",
        "Newsletter",
        "Task",
        "Assessment",
        "Poll"
      ]
    },
    author: { type: Schema.Types.ObjectId, ref: "User", required: true },
    school: { type: Schema.Types.ObjectId, ref: "School", required: true },
    class: { type: Schema.Types.ObjectId, ref: "Class", required: false },
    parent: { type: Schema.Types.ObjectId, ref: "Post" },
    // files: [AttachmentSchema],
    stars: { type: Number, default: 0 },

    text: String,
    section: String,
    replies: [{ type: Schema.Types.ObjectId, ref: "Post" }],
    locked: Boolean,
    pinned: Boolean,

    //assignment, task, polls, assessments - tasks that need response from students
    //responses can be viewed by trainer only
    submittable: Boolean,
    task: {
      title: String,
      start: Date,
      dueby: Date,
      graded: Boolean,
      category: String,
      grade: Number, //Actual score
      gradeCategory: String, //A, B, C
      comment: String,
      submissions: Number,
      pending: Number,
      maxmarks: Number
    },

    assessment: {
      duration: Number,
      count: Number,
      minmarks: Number,
      perItemMarks: Number,
      allocationType: { type: String, enum: ["spread", "custom"] },
      answered: Number,
      unanswered: Number,
      correct: Number,
      incorrect: Number,
      //score: Number,
      status: { type: String, enum: ["fail", "success"] },
      timeSpent: Number, //Seconds,
      attempts: { type: Number, default: 1 }
    },

    poll: {
      options: [
        {
          text: String,
          seqno: String,
          noOfSelections: Number
        }
      ],
      selected: String
    },

    //for private messages - default trainers can view all messages
    //no need to includes trainers and admins in this list
    users: [{ type: Schema.Types.ObjectId, ref: "User" }],
    responseType: { type: String, enum: ["submission", "reply", "solution"] },

    //Rich text preview
    richtext: {
      id: { type: Schema.Types.ObjectId, ref: "RichText" },
      previewImg: String,
      previewText: String,
      title: String
    },

    linkPreview: {
      title: String,
      description: String,
      image: String,
      logo: String,
      publisher: String,
      url: String,
      html: String
    },

    updated_by: String,
    created_at: Date,
    updated_at: Date
  },
  { usePushEach: true }
);
module.exports=mongoose.model("post",PostSchema);
