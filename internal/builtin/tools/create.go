package tools

import (
	"github.com/pharrisee/pegg/internal/builtin/tools/ask"
	"github.com/pharrisee/pegg/internal/builtin/tools/file"
	"github.com/pharrisee/pegg/internal/builtin/tools/loadskill"
	"github.com/pharrisee/pegg/internal/builtin/tools/plan"
	"github.com/pharrisee/pegg/internal/builtin/tools/searchskill"
	"github.com/pharrisee/pegg/internal/builtin/tools/sendfile"
	"github.com/pharrisee/pegg/internal/builtin/tools/shell"
	"github.com/pharrisee/pegg/internal/builtin/tools/subagent"
	"github.com/pharrisee/pegg/internal/builtin/tools/todo"
	"github.com/pharrisee/pegg/internal/builtin/tools/web"
	"github.com/pharrisee/pegg/internal/session"
	"github.com/pharrisee/pegg/internal/vfs"
)

func Create(fs *vfs.VFS, sess *session.Manager) {
	file.FileTools(fs)
	sendfile.SendFileTools(fs)
	todo.TodoTools(fs, sess.Bus())
	shell.ShellTools(fs)
	web.WebTools(fs)
	subagent.SubAgentTools(sess)
	ask.AskTool()
	loadskill.LoadSkillTool(fs, sess)
	searchskill.SearchSkillTool()
	plan.PlanTools(fs)
}
