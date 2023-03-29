
.PHONY: gox
gox:
	gox -osarch="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64" -output="build/ptpt_{{.OS}}_{{.Arch}}"

.PHONY: clean
clean:
	rm -f build/ptpt_*

release: gox