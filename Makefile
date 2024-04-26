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
all:  braidsAuth braidsBroker client1 client2

ethos:
	mkdir ethos
	cp -pr /usr/lib64/go/pkg/ethos_$(GOARCH)/* ethos

braidsAuthTypes.go: braidsAuthTypes.t
	$(ETN2GO) . braidsAuthTypes $^

braidsAuthTypes.goo.ethos : braidsAuthTypes.go ethos
	ethosGoPackage  braidsAuthTypes ethos braidsAuthTypes.go

braidsBrokerTypes.go: braidsBrokerTypes.t
	$(ETN2GO) . braidsBrokerTypes $^

braidsBrokerTypes.goo.ethos : braidsBrokerTypes.go ethos
	ethosGoPackage  braidsBrokerTypes ethos braidsBrokerTypes.go

braidsAuth: braidsAuth.go braidsAuthTypes.goo.ethos
	ethosGo braidsAuth.go

braidsBroker: braidsBroker.go braidsBrokerTypes.goo.ethos
	ethosGo braidsBroker.go

client1: client1.go braidsAuthTypes.goo.ethos
	ethosGo client1.go

client2: client2.go braidsAuthTypes.goo.ethos
	ethosGo client2.go

# install braidsAuthTypes, service,
install: all
	sudo rm -rf client
	(ethosParams client && cd client && ethosMinimaltdBuilder)
	echo 80 > client/param/sleepTime 
	ethosTypeInstall braidsAuthTypes
	ethosTypeInstall braidsBrokerTypes
	ethosDirCreate $(ETHOSROOT)/services/braidsAuthTypes   $(ETHOSROOT)/types/spec/braidsAuthTypes/Auth all
	ethosDirCreate $(ETHOSROOT)/services/braidsBrokerTypes   $(ETHOSROOT)/types/spec/braidsBrokerTypes/Broker all
	install -D braidsAuth                   $(ETHOSROOT)/programs
	install -D braidsBroker                   $(ETHOSROOT)/programs
	install -D client1                   $(ETHOSROOT)/programs
	install -D client2                   $(ETHOSROOT)/programs
	ethosStringEncode /programs/braidsAuth    > $(ETHOSROOT)/etc/init/services/braidsAuth
	ethosStringEncode /programs/braidsBroker    > $(ETHOSROOT)/etc/init/services/braidsBroker
	ethosStringEncode /programs/client1    > $(ETHOSROOT)/etc/init/services/client1
	ethosStringEncode /programs/client2    > $(ETHOSROOT)/etc/init/services/client2

# remove build artifacts
clean:
	rm -rf braidsAuthTypes/ braidsAuthTypesIndex/ ethos client
	rm -rf braidsBrokerTypes/ braidsBrokerTypesIndex/ ethos client
	rm -f braidsAuthTypes.go
	rm -f braidsBrokerTypes.go
	rm -f braidsAuth
	rm -f braidsBroker
	rm -f client1
	rm -f client2
	rm -f *.goo.*

run: clean install
	(cd client && sudo -E ethosRun -t)
	cat client/rootfs/log/application/braidsAuth/* > authLog
	cat client/rootfs/log/application/braidsBroker/* > brokerLog
	cat client/rootfs/log/test/braidsClient/* > clientLog
	cat authLog
	cat brokerLog
	cat clientLog
