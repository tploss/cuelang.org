// Code generated by internal/cmd/genpreprocessorembed; DO NOT EDIT.

package preprocessembed

import "embed"

//go:embed go.mod
//go:embed go.sum
//go:embed internal/fsnotify/*.go
//go:embed internal/functions/gerritstatusupdater/*.go
//go:embed internal/functions/snippets/*.go
//go:embed internal/parse/*.go
//go:embed internal/cmd/preprocessor/cmd/*.go
//go:embed internal/cmd/preprocessor/*.go
var Files embed.FS
