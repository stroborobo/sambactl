sambactl
========

A simple web interface for managing users and samba shares. Samba authentication
uses unix users, therefor this tool creates a unix user, adds it to smbpasswd
and creates a share directory.

Each user has a single share, groups are not supported yet. At the time of
writing no additional features are planned.

Dependencies
------------

* Samba
* Apache (you may need to change stuff for other httpds)
* mod_fastcgi
* Sudo (So sambactl-worker has the rights to manage samba)
* expect
* Go[1]

Installation
------------

	make
	make install

Check the config directory after installation, so the other services will be
configured correctly.

[1]: http://golang.org/
