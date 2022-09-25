#!/bin/bash

# customize these
WGET=wget
ICS2ORG=/home/meta/script/ical2org
ICSFILE=/tmp/basic.ics
ORGFILE=/home/meta/org/calendar.org
URL=https://calendar.google.com/calendar/ical/tmgfaction%40gmail.com/private-db4c40eaa2604e68046fc6c83e782285/basic.ics

# no customization needed below

$WGET -O $ICSFILE $URL
$ICS2ORG $ICSFILE > $ORGFILE
