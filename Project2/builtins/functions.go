package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// ListFiles lists the files in the current or a specified directory.
func ListFiles(w io.Writer, args ...string) error {
	dir := "."
	if len(args) > 0 {
		dir = args[0]
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		_, err := fmt.Fprintln(w, file.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

// PrintWorkingDirectory prints the current working directory.
func PrintWorkingDirectory(w io.Writer) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, wd)
	return err
}

// Echo outputs the arguments as a single string.
func Echo(w io.Writer, args ...string) error {
	_, err := fmt.Fprintln(w, strings.Join(args, " "))
	return err
}
