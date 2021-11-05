direct_topfunc_call.s: direct_topfunc_call.go
	GOOS=linux GOARCH=amd64 go tool compile -S $< > $@

.PHONY: clean
clean:
	-rm -f direct_topfunc_call.s
