export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=ethos
export GOLINUXINCLUDE=linux
export BUILD=ethos

export ETHOSROOT=client/rootfs
export MINIMALTDROOT=client/minimaltdfs


.PHONY: all install clean
all:  braidsAuth client1

ethos:
	mkdir ethos
	cp -pr /usr/lib64/go/pkg/ethos_$(GOARCH)/* ethos

braidsAuthTypes.go: braidsAuthTypes.t
	$(ETN2GO) . braidsAuthTypes $^

braidsAuthTypes.goo.ethos : braidsAuthTypes.go ethos
	ethosGoPackage  braidsAuthTypes ethos braidsAuthTypes.go

braidsAuth: braidsAuth.go braidsAuthTypes.goo.ethos
	ethosGo braidsAuth.go

client1: client1.go braidsAuthTypes.goo.ethos
	ethosGo client1.go

# install braidsAuthTypes, service,
install: all
	sudo rm -rf client
	(ethosParams client && cd client && ethosMinimaltdBuilder)
	echo 80 > client/param/sleepTime 
	ethosTypeInstall braidsAuthTypes
	ethosDirCreate $(ETHOSROOT)/services/braidsAuthTypes   $(ETHOSROOT)/types/spec/braidsAuthTypes/Auth all
	install -D braidsAuth                   $(ETHOSROOT)/programs
	install -D client1                   $(ETHOSROOT)/programs
	ethosStringEncode /programs/braidsAuth    > $(ETHOSROOT)/etc/init/services/braidsAuth
	ethosStringEncode /programs/client1    > $(ETHOSROOT)/etc/init/services/client1

# remove build artifacts
clean:
	rm -rf braidsAuthTypes/ braidsAuthTypesIndex/ ethos clent
	rm -f braidsAuthTypes.go
	rm -f braidsAuth
	rm -f braidsAuthTypes.goo.ethos
	rm -f braidsAuth.goo.ethos

run: clean install
	(cd client && sudo -E ethosRun -t)
	cat client/rootfs/log/application/braidsAuth/* > serverLog
	cat client/rootfs/log/test/braidsClient/* > clientLog

