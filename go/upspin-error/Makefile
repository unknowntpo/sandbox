build/prod:
	go build -o ./cmd .
build/debug:
	go build -tags debug -o ./cmd .
clean:
	rm -f ./cmd

.PHONY: build/prod build/debug clean

