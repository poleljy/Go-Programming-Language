package executor

import (
	"fmt"
	"os"
	"runtime"
)

// makeExecutable makes the given file executable for root,group,others.
func makeExecutable(binPath string) error {
	if runtime.GOOS == "windows" {
		return nil
	}

	fi, err := os.Stat(binPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("binary %q does not exist", binPath)
		}
		return fmt.Errorf("specified binary is invalid: %v", err)
	}

	// If it is not executable, make it so.
	perm := fi.Mode().Perm()
	req := os.FileMode(0555)
	if perm&req != req {
		if err := os.Chmod(binPath, perm|req); err != nil {
			return fmt.Errorf("error making %q executable: %s", binPath, err)
		}
	}
	return nil
}

// Path
func CombineEnvPath(paths []string) error {
	var split, env string
	if runtime.GOOS == "windows" {
		env = "PATH"
		split = ";"
	} else if runtime.GOOS == "linux" {
		env = "LD_LIBRARY_PATH"
		split = ":"
	}
	envPath := os.Getenv(env)
	for _, str := range paths {
		envPath = envPath + split + str
	}

	if err := os.Setenv(env, envPath); err != nil {
		return err
	}
	return nil
}
