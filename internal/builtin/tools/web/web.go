package web

import (
	"github.com/pharrisee/pegg/internal/agent/tools"
	"github.com/pharrisee/pegg/internal/vfs"
)

func WebTools(fs *vfs.VFS) {
	tools.Register(webfetchTool(fs))
	tools.Register(websearchTool(fs))
}
