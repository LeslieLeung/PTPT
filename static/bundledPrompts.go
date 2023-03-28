package static

import "embed"

//go:embed *.yaml
var BundledPromptsStorage embed.FS
