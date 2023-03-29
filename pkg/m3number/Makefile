#RACE=
DEBUG=--tags debug

test:
	@godebug
	@#go test $(RACE) $(DEBUG) --run TestM3numberImporter
	@go test $(RACE) $(DEBUG) --run TestM3numberEvaluator
#	@rm *_debug.go *_nodebug.go

clean:
	@rm *_debug.go *_nodebug.go

