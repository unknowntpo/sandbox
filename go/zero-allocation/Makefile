all: prof

.PHONY: prof
prof:
	go build -o ./bin .
	./bin
	go tool pprof cpu.pprof

.PHONY: clean
clean:
	-rm -f ./bin
