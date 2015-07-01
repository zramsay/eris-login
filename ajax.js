

function new_request_obj(){
	if (window.XMLHttpRequest)
		return new XMLHttpRequest();
	else
		return new ActiveXObject("Microsoft.XMLHTTP");
}

function request_callback(xObj, _func, args){
	xObj.onreadystatechange=function(){
		if (xObj.readyState==4 && xObj.status==200){
			args.unshift(xObj);
			_func.apply(this, args);
		}
	}
}

function make_request(xmlhttp, method, path, async, params){
	xmlhttp.open(method, path, async);
	xmlhttp.setRequestHeader("Content-Type", "application/json");
	xmlhttp.send(JSON.stringify(params));
}

function sendinfo() {
	var name = document.getElementById("name").value;
	var auth = document.getElementById("password").value;
	console.log("name & pass ", name, auth);
	
	xmlhttp = new_request_obj();
	request_callback(xmlhttp, callback, []);
	make_request(xmlhttp, "POST", "http://localhost:4767/gen", true, {"auth":auth, "name":name});

}

function callback() {
	console.log("callback");
}

