package static

import "embed"

//go:embed hello.yaml role_play.yaml translate.yaml
var BundledPromptsStorage embed.FS
