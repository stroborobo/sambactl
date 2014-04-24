INSTALLDIR=/usr/local/bin/
WEBROOT=/usr/local/sambactl/

SERVER_SRCS+=	sambactl-server/main.go
WORKER_SRCS+=	sambactl-worker/main.go

all: build

build:	sambactl-worker/sambactl-worker sambactl-server/sambactl-server

sambactl-server/sambactl-server: ${SERVER_SRCS}
	cd sambactl-server && go build

sambactl-worker/sambactl-worker: ${WORKER_SRCS}
	cd sambactl-worker && go build

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
