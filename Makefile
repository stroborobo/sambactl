INSTALLDIR=/usr/local/bin/
WEBROOT=/usr/local/sambactl/

all: build

build:
	cd sambactl-worker && go build
	cd sambactl-server && go build

clean:
	rm -f sambactl-worker/sambactl-worker
	rm -f sambactl-server/sambactl-server

install: build
	mkdir -p ${INSTALLDIR}
	install sambactl-worker/sambactl-worker ${INSTALLDIR}
	install sambactl-server/sambactl-server ${INSTALLDIR}
	install scripts/mysmbpasswd ${INSTALLDIR}/mysmbpasswd
	install scripts/smb_adduser.sh ${INSTALLDIR}/smb_adduser
	install scripts/smb_deluser.sh ${INSTALLDIR}/smb_deluser
	mkdir -p ${WEBROOT}
	cp -r webroot/* ${WEBROOT}
