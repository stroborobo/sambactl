#!/bin/sh

set -e

USER=$1
PASS=$2

PROJEKTNAME=$USER

useradd -K UID_MIN=2000 -s /bin/false $USER
mysmbpasswd -a "$USER" "$PASS"

## Share Ordner erstellen:

mkdir -m 700 "/srv/samba/$PROJEKTNAME"
chown $USER:$USER "/srv/samba/$PROJEKTNAME"

## Share anlegen:

echo "[$PROJEKTNAME]
path = /srv/samba/$PROJEKTNAME
writable = yes
valid users = $USER" > /etc/samba/shares/$PROJEKTNAME

## Share hinzufügen:

echo "include = /etc/samba/shares/$PROJEKTNAME" >> /etc/samba/smbshares.conf
