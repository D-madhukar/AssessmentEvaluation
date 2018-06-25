var mongoose=require("mongoose");
mongoose.connect("mongodb://localhost/test");
var Schema=mongoose.Schema;
var AssessmentSubmitSchema = new Schema({

	//TF - True False, SA - Simple Answer, MC - Multi Choiceq
	// _id:String,
	assessmentPost : {type:Schema.Types.ObjectId, ref: 'Post', required: true},
	post : {type:Schema.Types.ObjectId, ref: 'Post', required: false},
	user: {type:Schema.Types.ObjectId, ref: 'User', required: true},
	assessment: {type:Schema.Types.ObjectId, ref: 'Assessment', required: true},
	type: { type: String, enum: ['TF', 'SA', 'MC']},

	answer: String, //TF, SA
	options: [{
		seqno: String, //MC
	}],
	// result: { type: String, enum: ['correct', 'wrong', 'tbd']},
	timeSpent: Number,
	status: { type: String, enum: ['correct', 'not correct', 'not answered']},

	created_by: {type:Schema.Types.ObjectId, ref: 'User', required: false},
	updated_by: String,
	created_at: Date,
  	updated_at: Date

}, { usePushEach: true });



module.exports = mongoose.model('AssessmentSubmit', AssessmentSubmitSchema);
