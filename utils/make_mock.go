package main

import (
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

const ShellToUse = "bash"

var logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelDebug,
}))

func main() {
	dir := os.DirFS("./")
	createdDirs := map[string]bool{}
	makeMocks := func(path string, d fs.DirEntry, err error) error {
		return handelFile(path, d, createdDirs, err)
	}
	err := fs.WalkDir(dir, ".", makeMocks)
	if err != nil {
		logger.Error("Generating mocks failed", "error", err)
	}
}

func handelFile(path string, d fs.DirEntry, createdDirs map[string]bool, err error) error {
	logger.Debug("Looking for source", "path", path)
	if filepath.Ext(path) == ".go" {
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		match, err := regexp.Match("type .* interface", content)
		if err != nil {
			return err
		}
		if match && !strings.HasSuffix(path, "make_mock.go") {
			logger.Debug("Contains interface", "file", path)
			err := generateMock(path, createdDirs)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func generateMock(file string, created map[string]bool) error {
	logger.Debug("Generating mock", "file", file)
	dir, err := createMockDir(filepath.Dir(file), created)
	if err != nil {
		return err
	}
	fileInfo, err := os.Stat(file)
	if err != nil {
		return err
	}
	name := fileInfo.Name()
	name = name[:len(name)-3] + "_mock.go"
	target := filepath.Join(dir, name)
	command := fmt.Sprintf("mockgen -source=%s -destination=%s", file, target)
	cmd := exec.Command(ShellToUse, "-c", command)

	logger.Info("Executing generate command", "command", command)
	return cmd.Run()
}
func createMockDir(path string, created map[string]bool) (string, error) {
	logger.Debug("Checking mock directory")
	mockDir := filepath.Join(path, "mock")
	if _, ok := created[mockDir]; !ok {
		err := os.RemoveAll(mockDir)
		if err != nil {
			logger.Debug("Mock directory doesn't exists", "path", mockDir)
		} else {
			logger.Debug("Removing mock directory", "path", mockDir)
		}
		err = os.Mkdir(mockDir, os.ModePerm)
		if err != nil {
			return mockDir, err
		}
		created[mockDir] = true
	}
	return mockDir, nil
}
