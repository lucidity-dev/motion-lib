GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get
GOFMT = $(GOCMD) fmt

TARGET = motion-lib

.PHONY: test

build:
	$(GOBUILD) -o $(TARGET) -v

check:
	$(GOBUILD) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(TARGET) *.jpg

test:
	$(GOTEST) -v ./test

format:
	$(GOFMT) ./...

deps:
	$(GOGET) gonum.org/v1/plot/...

run:	check build
	./$(TARGET)
