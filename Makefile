GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

TARGET = motion-lib

build:
	$(GOBUILD) -o $(TARGET) -v

check:
	$(GOBUILD) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(TARGET)

test:
	$(GOTEST) -v ./...

run:	check build
	./$(TARGET)
