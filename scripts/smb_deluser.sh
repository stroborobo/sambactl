#!/bin/sh

set -e

USER=$1
PROJEKTNAME=$USER


## SMB Sessions schließen (optional?)

#?? (sonst kommt beim löschen mit deluser: userdel: user testuser2 is currently used by process 4347)

## User löschen:

smbpasswd -x $USER
deluser $USER

## Share Ordner löschen:

rm -r /srv/samba/$PROJEKTNAME

## Share löschen:

rm /etc/samba/shares/$PROJEKTNAME

## Share austragen:

sed -i '/include = \/etc\/samba\/shares\/$PROJEKTNAME/d' /etc/samba/smbshares.conf
