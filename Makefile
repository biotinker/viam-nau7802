viam-nau7802: *.go */*.go
	# the executable
	go build -o $@ -ldflags "-s -w" -tags osusergo,netgo
	file $@

module.tar.gz: viam-nau7802
	# the bundled module
	rm -f $@
	tar czf $@ $^
