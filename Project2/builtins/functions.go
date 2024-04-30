package builtins

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// list all files in current directory
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

// deletes the specified file
func RemoveFile(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no file specified to remove")
	}
	filename := args[0]
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "Removed file: %s\n", filename)
	return err
}

// prints the current directory
func PrintWorkingDirectory(w io.Writer) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	_, err = fmt.Fprintln(w, wd)
	return err
}

// outputs the argument as a string
func Echo(w io.Writer, args ...string) error {
	_, err := fmt.Fprintln(w, strings.Join(args, " "))
	return err
}

// creates new directory with provided name
func MakeDirectory(w io.Writer, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("no directory name given")
	}
	dirName := args[0]
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "New directory: %s\n", dirName)
	return err
}
