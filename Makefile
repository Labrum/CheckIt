GOCMD = go
GORUN = $(GOCMD) run
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean

SOURCE = compile.go main.go languages.go template.go share.go

make: $(SOURCE)
	$(GORUN) $(SOURCE)

build: $(SOURCE)
	$(GOBUILD) $(SOURCE)

clean: 
	$(GOCLEAN)