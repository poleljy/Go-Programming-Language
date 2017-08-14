## Go语言面向对象

> 参考源码：github.com/mitchellh/cli

### 接口的定义
```go
// A command is a runnable sub-command of a CLI.
type Command interface {
	// Help should return long-form help text that includes the command-line
	// usage, a brief few sentences explaining the function of the command,
	// and the complete list of flags the command accepts.
	Help() string

	// Run should run the actual command with the given CLI instance and
	// command-line arguments. It should return the exit status when it is
	// finished.
	//
	// There are a handful of special exit codes this can return documented
	// above that change behavior.
	Run(args []string) int

	// Synopsis should return a one-line, short synopsis of the command.
	// This should be less than 50 characters ideally.
	Synopsis() string
}

// CommandHelpTemplate is an extension of Command that also has a function
// for returning a template for the help rather than the help itself. In
// this scenario, both Help and HelpTemplate should be implemented.
//
// If CommandHelpTemplate isn't implemented, the Help is output as-is.
type CommandHelpTemplate interface {
	// HelpTemplate is the template in text/template format to use for
	// displaying the Help. The keys available are:
	//
	//   * ".Help" - The help text itself
	//   * ".Subcommands"
	//
	HelpTemplate() string
}

// CommandFactory is a type of function that is a factory for commands.
// We need a factory because we may need to setup some state on the
// struct that implements the command itself.
type CommandFactory func() (Command, error)



```

### 函数定义
```go
// HelpFunc is the type of the function that is responsible for generating
// the help output when the CLI must show the general help text.
type HelpFunc func(map[string]CommandFactory) string

// BasicHelpFunc generates some basic help output that is usually good enough
// for most CLI applications.

func BasicHelpFunc(app string) HelpFunc {
}


```


### 类的定义

