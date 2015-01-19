GOCMD = go
GORUN = $(GOCMD) run
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean

SOURCE = compile.go CheckIt.go template.go share.go util.go

make: $(SOURCE)
	$(GORUN) $(SOURCE)

build: $(SOURCE)
	$(GOBUILD) $(SOURCE)

clean: 
	$(GOCLEAN)
