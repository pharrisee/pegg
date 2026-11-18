package shell

import (
	"github.com/pharrisee/pegg/internal/agent/tools"
	"github.com/pharrisee/pegg/internal/vfs"
)

func ShellTools(fs *vfs.VFS) {
	tools.Register(bashTool(fs))
}
