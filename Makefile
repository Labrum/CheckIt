GOCMD = go
GORUN = $(GOCMD) run
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean

SOURCE = main.go compile.go Server.go languages.go template.go share.go util.go init.go

make: $(SOURCE)
	$(GORUN) $(SOURCE)

build: $(SOURCE)
	$(GOBUILD) $(SOURCE)

clean: 
	$(GOCLEAN)
