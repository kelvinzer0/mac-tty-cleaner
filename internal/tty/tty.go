package tty

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// GetTTYDevices returns a list of TTY devices in the /dev directory
func GetTTYDevices() ([]string, error) {
	var ttyDevices []string

	devDir := "/dev"
	files, err := os.ReadDir(devDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), "ttys0") {
			ttyDevices = append(ttyDevices, filepath.Join(devDir, file.Name()))
		}
	}

	return ttyDevices, nil
}

// RunCommandOnTTY runs the specified command on the given TTY device
func RunCommandOnTTY(command, ttyDevice string) error {
	// Create a context with a 3-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Open TTY device with a timeout
	tty, err := openTTYWithTimeout(ctx, ttyDevice)
	if err != nil {
		return fmt.Errorf("error opening TTY: %v", err)
	}
	defer tty.Close()

	// Create command
	cmd := exec.CommandContext(ctx, command)

	// Set TTY file descriptors
	cmd.Stdin = os.Stdin // Set standard input to the original stdin
	cmd.Stdout = tty
	cmd.Stderr = tty

	// Set TTY as controlling terminal
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	// Redirect standard input from /dev/null to avoid waiting for input
	cmd.Stdin, err = os.Open("/dev/null")
	if err != nil {
		return fmt.Errorf("error opening /dev/null: %v", err)
	}

	// Log the command being executed
	log.Printf("Executing command '%s' on %s\n", command, ttyDevice)

	return cmd.Run()
}


// OpenTTYWithTimeout opens the TTY device with a timeout
func openTTYWithTimeout(ctx context.Context, ttyDevice string) (*os.File, error) {
	var tty *os.File
	var err error

	done := make(chan struct{})
	go func() {
		defer close(done)
		tty, err = os.OpenFile(ttyDevice, os.O_RDWR, 0)
	}()

	select {
	case <-done:
		return tty, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
