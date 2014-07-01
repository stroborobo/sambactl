#!/bin/sh

set -e

BACKUP=/var/backup
USER=$1
PROJEKTNAME=$USER


## SMB Sessions schließen (optional?)

#?? (sonst kommt beim löschen mit deluser: userdel: user testuser2 is currently used by process 4347)

## User löschen:

smbpasswd -x $USER
deluser $USER

## Share Ordner löschen:

mkdir -p $BACKUP

ts=`date "+%s"`
targetdir=${BACKUP}/${PROJEKTNAME}.${ts}
sharedir=/srv/samba/$PROJEKTNAME
i=0
while [ $i < 100 ]; do
	if [ ! -d ${targetdir}.$i ]; then
		mv $sharedir ${targetdir}.$i
		break
	fi
	i=$(($i + 1))
done
if [ -d $sharedir ]; then
	rm -r $sharedir
fi

## Share löschen:

rm /etc/samba/shares/$PROJEKTNAME

## Share austragen:

sed -i '/include = \/etc\/samba\/shares\/$PROJEKTNAME/d' /etc/samba/smbshares.conf
