#!/usr/bin/python2

import urllib2
import urllib
import sys
import json
try:
	data = file(sys.argv[1], 'r').read()
except:
	print "please give a picture, and make sure your filename is correct"
	exit()
urldata = urllib.urlencode({"image" : data})
url = "https://api.imgur.com/3/upload"
r = urllib2.Request(url)
r.add_data(urldata)
r.add_header('Authorization','Client-ID 1c49486ec8e9565')
res = json.loads( urllib2.urlopen(r).read())
print res['data']['link']


#upload a picture to imgur, and return a link
