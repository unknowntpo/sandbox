TARGET = gojq

all: $(TARGET)

$(TARGET):
	go build -o $(TARGET) ./cmd

.PHONY: test
test:
	@echo "testing with python script..."

.PHONY: clean
clean:
	rm -f $(TARGET)
