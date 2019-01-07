GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

TARGET = motion-lib

build:
	$(GOBUILD) -o $(TARGET) -v

clean:
	$(GOCLEAN)
	rm -f $(TARGET)

run: build
	./$(TARGET)
