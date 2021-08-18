package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// GetOrSetHomePath returns the existing value for env 'OSSICONES_SRC_HOME' if present.
// Otherwise, it sets and returns the given value.
func GetOrSetHomePath() (string, error) {
	homePath := os.Getenv("OSSICONES_SRC_HOME")
	if homePath == "" {
		cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
		if err != nil {
			log.Fatal(err.Error(), string(cmdOut))
			return "", err
		}
		homePath = strings.TrimSpace(string(cmdOut))
		fmt.Printf("HOME PATH : %s\n", homePath)
		os.Setenv("OSSICONES_SRC_HOME", homePath)
	}
	return homePath, nil
}

// IsRunning checks if the same process is already running.
func IsRunning() bool {
	err := filepath.Walk("/proc", findSameProcess)
	if err != nil {
		if err == io.EOF {
			// Not an error, just a signal when we are done
			return false
		} else {
			return false
		}
	}
	return true
}

// findProcess walks iterative through the /process directory tree
// looking up the process name found in each /proc/<pid>/status file. If
// the name matches the name in the argument the process with the corresponding
// this process will be stopped.
func findSameProcess(path string, info os.FileInfo, err error) error {
	// We just return in case of errors, as they are likely due to insufficient
	// privileges. We shouldn't get any errors for accessing the information we
	// are interested in. Run as root (sudo) and log the error, in case you want
	// this information.
	if err != nil {
		// log.Println(err)
		return nil
	}

	// We are only interested in files with a path looking like /proc/<pid>/status.
	if strings.Count(path, "/") == 3 {
		if strings.Contains(path, "/status") {

			// Let's extract the middle part of the path with the <pid> and
			// convert the <pid> into an integer. Log an error if it fails.
			pid, err := strconv.Atoi(path[6:strings.LastIndex(path, "/")])
			if err != nil {
				log.Println(err)
				return nil
			}

			if pid == os.Getpid() {
				return nil
			}

			// The status file contains the name of the process in its first line.
			// The line looks like "Name: theProcess".
			// Log an error in case we cant read the file.
			f, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println(err)
				return nil
			}

			// Extract the process name from within the first line in the buffer
			name := string(f[6:bytes.IndexByte(f, '\n')])

			if name == os.Args[0] {
				fmt.Printf("PID: %d, Name: %s is already running.\n", pid, name)
				_, err := os.FindProcess(pid)
				if err != nil {
					log.Println(err)
				}

				// Let's return a fake error to abort the walk through the
				// rest of the /proc directory tree
				return io.EOF
			}

		}
	}

	return nil
}

func HandleError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
