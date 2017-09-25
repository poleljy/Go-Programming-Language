package executor

import (
	"log"
	"os"
	"os/exec"
	_ "path/filepath"
	_ "strings"
	"sync"
	"time"
)

// Executor is the interface which allows a driver to launch and supervise
// a process
type Executor interface {
	LaunchCmd(command *ExecCommand) (*ProcessState, error)
	Wait() (*ProcessState, error)
	ShutDown() error
	Exit() error
	Signal(s os.Signal) error
	Exec(deadline time.Time, cmd string, args []string) ([]byte, int, error)
}

// UniversalExecutor is an implementation of the Executor which launches and
// supervises processes. In addition to process supervision it provides resource
// and file system isolation
type UniversalExecutor struct {
	cmd     exec.Cmd
	command *ExecCommand

	pids                map[int]int
	pidLock             sync.RWMutex
	exitState           *ProcessState
	processExited       chan interface{}
	fsIsolationEnforced bool

	shutdownCh chan struct{}

	logger *log.Logger

	// TaskDir is the host path to the task's root
	TaskDir string

	// LogDir is the host path where logs should be written
	LogDir string
}

// NewExecutor returns an Executor
//func NewExecutor(logger *log.Logger) Executor {
//	exec := &UniversalExecutor{
//		logger:        logger,
//		processExited: make(chan interface{}),
//	}
//
//	return exec
//}

// ExecCommand holds the user command, args, and other isolation related
// settings.
type ExecCommand struct {
	// Cmd is the command that the user wants to run.
	Cmd string

	// Args is the args of the command that the user wants to run.
	Args []string

	// FSIsolation determines whether the command would be run in a chroot.
	FSIsolation bool

	// User is the user which the executor uses to run the command.
	User string
}

// ProcessState holds information about the state of a user process.
type ProcessState struct {
	Pid      int
	ExitCode int
	Signal   int
	Time     time.Time
}

//////////////////////////////////////////////////////////////////////////////////////

/*
// LaunchCmd launches the main process and returns its state. It also
// configures an applies isolation on certain platforms.
func (e *UniversalExecutor) LaunchCmd(command *ExecCommand) (*ProcessState, error) {
	e.logger.Printf("[DEBUG] executor: launching command %v %v", command.Cmd, strings.Join(command.Args, " "))
	e.command = command

	// set the task dir as the working directory for the command
	e.cmd.Dir = e.TaskDir

	// Set the commands arguments
	//e.cmd.Path = path
	//e.cmd.Args = append([]string{e.cmd.Path}, e.ctx.TaskEnv.ParseAndReplace(command.Args)...)

	// Look up the binary path and make it executable
	if err := makeExecutable(e.cmd.Path); err != nil {
		return nil, err
	}

	path := e.cmd.Path

	// Determine the path to run as it may have to be relative to the chroot.
	if e.fsIsolationEnforced {
		rel, err := filepath.Rel(e.ctx.TaskDir, path)
		if err != nil {
			return nil, fmt.Errorf("failed to determine relative path base=%q target=%q: %v", e.ctx.TaskDir, path, err)
		}
		path = rel
	}

	// Set the commands arguments
	e.cmd.Path = path
	e.cmd.Args = append([]string{e.cmd.Path}, e.ctx.TaskEnv.ParseAndReplace(command.Args)...)
	e.cmd.Env = e.ctx.TaskEnv.List()

	// Start the process
	if err := e.cmd.Start(); err != nil {
		return nil, fmt.Errorf("failed to start command path=%q --- args=%q: %v", path, e.cmd.Args, err)
	}
	go e.collectPids()
	go e.wait()
	ic := e.resConCtx.getIsolationConfig()
	return &ProcessState{Pid: e.cmd.Process.Pid, ExitCode: -1, IsolationConfig: ic, Time: time.Now()}, nil
}
*/
