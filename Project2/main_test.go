package main

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"testing/iotest"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_runLoop(t *testing.T) {
	t.Parallel()
	exitCmd := strings.NewReader("exit\n")
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name     string
		args     args
		wantW    string
		wantErrW string
	}{
		{
			name: "no error",
			args: args{
				r: exitCmd,
			},
		},
		{
			name: "read error should have no effect",
			args: args{
				r: iotest.ErrReader(io.EOF),
			},
			wantErrW: "EOF",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			w := &bytes.Buffer{}
			errW := &bytes.Buffer{}

			// Create a new exit channel for this specific test case
			exit := make(chan struct{})

			// Run `runLoop` in a goroutine
			go func() {
				defer close(exit) // Close the exit channel only after the loop finishes
				runLoop(tt.args.r, w, errW, exit)
			}()

			// Give some time for `runLoop` to start, then signal it to exit
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			// Allow the loop to finish before checking outputs
			time.Sleep(10 * time.Millisecond)

			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}
