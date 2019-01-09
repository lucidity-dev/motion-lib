GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

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

deps:
	$(GOGET) -u gonum.org/v1/plot/...

run:	check build
	./$(TARGET)
