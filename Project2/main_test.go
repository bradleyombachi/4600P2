package main

import (
	"bytes"
	"io"
	"strings"
	"sync"
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

			// Create a new exit channel and a WaitGroup for this specific test case
			exit := make(chan struct{}, 1)
			var wg sync.WaitGroup
			wg.Add(1)

			// Run `runLoop` in a goroutine
			go func() {
				defer wg.Done()
				runLoop(tt.args.r, w, errW, exit)
			}()

			// Give some time for `runLoop` to start, then signal it to exit
			time.Sleep(10 * time.Millisecond)
			exit <- struct{}{}

			// Wait for the goroutine to finish before checking outputs
			wg.Wait()

			// Check output and error output
			require.NotEmpty(t, w.String())
			if tt.wantErrW != "" {
				require.Contains(t, errW.String(), tt.wantErrW)
			} else {
				require.Empty(t, errW.String())
			}
		})
	}
}
