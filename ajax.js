

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
console.log("paramss", params);
	xmlhttp.send(JSON.stringify(params));
}

function sendinfo() {
	var name = document.getElementById("name").value;
	var auth = document.getElementById("password").value;
	console.log("name & pass ", name, auth);
	
	nonce = makeNonce();
	console.log("nonce", nonce);

	xmlhttp = new_request_obj();
	request_callback(xmlhttp, callback, []);
	make_request(xmlhttp, "POST", "http://localhost:4767/name", true, {"name":name});

}

function callback() {
	var response = JSON.parse(xmlhttp.responseText);
	console.log("repsonse is here", response);
	
	addr = response.Response;
	var err = response.Error;

	var hash = nonce;
	console.log("hahs", hash);

	console.log("addr & err", addr, err);

	xmlhttp = new_request_obj();
	request_callback(xmlhttp, callback2, []);
	make_request(xmlhttp, "POST", "http://localhost:4767/sign", true, {"addr":addr, "hash":hash });
}

function callback2(){
	console.log("callback2");
	var response = JSON.parse(xmlhttp.responseText);
	console.log("signature", response);

	var hash = nonce;
	
	var sig = response.Response;
	var err = response.Error;

	xmlhttp = new_request_obj();
	request_callback(xmlhttp, callback3, []);
	make_request(xmlhttp, "POST", "http://localhost:4767/verify", true, {"addr":addr, "hash":hash, "sig":sig });

}

function callback3() {
	console.log("callback3");
	var response = JSON.parse(xmlhttp.responseText);
	console.log("verify", response);

}

function makeNonce() {
	var text = "";
	//hex only for now
	var possible = "ABCDEFabcdef0123456789";
	
	for( var i=0; i < 10 ; i++ )
		text += possible.charAt(Math.floor(Math.random() * possible.length));
	return text;
}
