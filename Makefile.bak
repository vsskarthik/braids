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
all:  server

ethos:
	mkdir ethos
	cp -pr /usr/lib64/go/pkg/ethos_$(GOARCH)/* ethos

typeDefs.go: typeDefs.t
	$(ETN2GO) . typeDefs $^

typeDefs.goo.ethos : typeDefs.go ethos
	ethosGoPackage  typeDefs ethos typeDefs.go

server: server.go typeDefs.goo.ethos
	ethosGo server.go

# install typeDefs, service,
install: all
	sudo rm -rf client
	(ethosParams client && cd client && ethosMinimaltdBuilder)
	echo 80 > client/param/sleepTime 
	ethosTypeInstall typeDefs
	ethosDirCreate $(ETHOSROOT)/services/typeDefs   $(ETHOSROOT)/types/spec/typeDefs/Auth all
	install -D server                   $(ETHOSROOT)/programs
	ethosStringEncode /programs/server    > $(ETHOSROOT)/etc/init/services/server

# remove build artifacts
clean:
	rm -rf typeDefs/ typeDefsIndex/ ethos clent
	rm -f typeDefs.go
	rm -f server
	rm -f typeDefs.goo.ethos
	rm -f server.goo.ethos

run: clean install
	(cd client && sudo -E ethosRun -t)
	ethosLog client > log
