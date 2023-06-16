
version =

.PHONY: gox
gox:
	gox -osarch="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64" -output="build/ptpt_{{.OS}}_{{.Arch}}"

.PHONY: clean
clean:
	rm -f build/ptpt_*

.PHONY: priv
priv:
	gox -osarch="darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64" -output="build/ptpt_private_{{.OS}}_{{.Arch}}_${version}" -tags=private

release: gox