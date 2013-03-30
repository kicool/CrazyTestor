/* 
 * 表单提交工具,基于jQuery
 * Author:	Liangdi
 * Email:	wu@liangdi.me
 */

function SmartForm(id,els,option){
	this.form = $("#" + id);
	this.els = [];
	this.url=option.url;
	this.action = option.action?option.action:"GET";
	
	
	this.debug = option.debug?true:false;
	
	this.successCallbacks = [];
	this.failueCallbacks = [];
	this.completeCallbacks = [];
	this.validCallbacks = [];
	
	for (var i = 0; i < els.length; i++) {
		var id = els[i];
		var el = this.form.find("#"+id);
		var elObj = {
			id:id,
			el:el,
			notNull:el.attr("data-not-null")?true:false
		};
		this.els.push(elObj);
	}
	
	if(this.debug){
		console.log("form",this.form);
		console.log("els",this.els);
	}
	var that = this;
	this.form.submit(function(){
		that.submit();
		return false;
	});
}
SmartForm.prototype.submit = function(){
	
	var data = {};
	for (var i = 0; i < this.els.length; i++) {
		var elObj = this.els[i];
		
		var v = elObj.el.val();
		var notNull = elObj.notNull;
		if(this.debug){
			console.log("value:",elObj.id,notNull,v,v == null || v == "")	
		}
		if(notNull && (v == null || v == "")){
			for (var j = 0; j < this.validCallbacks.length; j++) {
				this.validCallbacks[j](elObj);
			}
			return false;
		}
		if(v instanceof Array){
			v=v.join();
		}
		data[elObj.id] = v;
	}
	if(this.debug){
		console.log("data:",data);
	}
	var url = this.url;
	var type = this.type;
	var successCallbacks = this.successCallbacks;
	var completeCallbacks = this.completeCallbacks;
	$.ajax({
		url:url,
		type:type,
		data:data,
		success:function(data){
			for (var j = 0; j < successCallbacks.length; j++) {
				successCallbacks[j](data);
			}
		},
		complete:function(){
			for (var j = 0; j < completeCallbacks.length; j++) {
				completeCallbacks[j]();
			}
		}
	});
	
	return false;
};
SmartForm.prototype.onSuccess = function(callback){
	typeof callback === "function" && this.successCallbacks.push(callback);
};
SmartForm.prototype.onComplete = function(callback){
	typeof callback === "function" && this.completeCallbacks.push(callback);
};
SmartForm.prototype.onVerificationError = function(callback){
	typeof callback === "function" && this.validCallbacks.push(callback);
};


