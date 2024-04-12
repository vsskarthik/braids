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
all: expenseReportClient expenseReportServer

ethos:
	mkdir ethos
	cp -pr /usr/lib64/go/pkg/ethos_$(GOARCH)/* ethos

typeDefs.go: typeDefs.t
	$(ETN2GO) . typeDefs $^

typeDefs.goo.ethos : typeDefs.go ethos
	ethosGoPackage  typeDefs ethos typeDefs.go

expenseReportServer: expenseReportServer.go typeDefs.goo.ethos
	ethosGo expenseReportServer.go

expenseReportClient: expenseReportClient.go typeDefs.goo.ethos
	ethosGo expenseReportClient.go

# install typeDefs, service,
install: all
	sudo rm -rf client
	(ethosParams client && cd client && ethosMinimaltdBuilder)
	(cd client/rootfs && ethosUserRecord user1 "User 1" "user1@example.com" "" && ethosUserRecord user2 "User 2" "user2@example.com" "")
	echo 60 > client/param/sleepTime 
	ethosTypeInstall typeDefs
	ethosDirCreate $(ETHOSROOT)/services/typeDefs   $(ETHOSROOT)/types/spec/typeDefs/ExpenseReport all
	install -D  expenseReportClient expenseReportServer                   $(ETHOSROOT)/programs
	#ethosStringEncode /programs/expenseReportServer    > $(ETHOSROOT)/etc/init/services/expenseReportServer
	#ethosStringEncode /programs/expenseReportClient       > $(ETHOSROOT)/etc/init/services/expenseReportClient

# remove build artifacts
clean:
	rm -rf typeDefs/ typeDefsIndex/ ethos clent
	rm -f typeDefs.go
	rm -f expenseReportClient
	rm -f expenseReportServer
	rm -f typeDefs.goo.ethos
	rm -f expenseReportServer.goo.ethos
	rm -f expenseReportClient.goo.ethos
