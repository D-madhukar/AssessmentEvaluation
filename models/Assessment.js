var mongoose = require('mongoose');
var Schema = mongoose.Schema;
var AssessmentSchema = new Schema({

	//TF - True False, SA - Simple Answer, MC - Multi Choiceq
	post : {type:Schema.Types.ObjectId, ref: 'Post', required: false},
	seqno: Number,
	type: { type: String, enum: ['TF', 'SA', 'MC']},
	question: String,
	answer: String,
	isSingeOption: Boolean,
	options: [{
		seqno: String,
		text: String,
		answer: Boolean
	}],
	explanation: String,
	difficulty: { type: String, enum: ['S', 'M', 'L']},
	marks: Number,

	created_by: {type:Schema.Types.ObjectId, ref: 'User', required: false},
	updated_by: String,
	created_at: Date,
  	updated_at: Date

}, { usePushEach: true });
module.exports = mongoose.model('Assessment', AssessmentSchema);
