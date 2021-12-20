package bininfo

import (
	"bytes"
	"fmt"
	"runtime"
)

var (
	GitTag    = "Unknown"
	GitCommit = "Unknown"
	GitBranch = "Unknown"
	GitStatus = "Unknown"
	BuildTime = "Unknown"
)

func PrettyVersion() []byte {
	if GitStatus == "" {
		GitStatus = "cleanly"
	}

	buf := new(bytes.Buffer)
	_, _ = fmt.Fprintf(buf, "API Version:  %s\n", GitTag)
	_, _ = fmt.Fprintf(buf, "\n")

	_, _ = fmt.Fprintf(buf, "Build Info:\n")
	_, _ = fmt.Fprintf(buf, "    Go version - %s %s/%s\n", runtime.Version(), runtime.GOARCH, runtime.GOOS)
	_, _ = fmt.Fprintf(buf, "    time - %s\n", BuildTime)
	_, _ = fmt.Fprintf(buf, "\n")

	_, _ = fmt.Fprintf(buf, "Git Info:\n")
	_, _ = fmt.Fprintf(buf, "    branch - %s\n", GitBranch)
	_, _ = fmt.Fprintf(buf, "    commit - %s\n", GitCommit)
	_, _ = fmt.Fprintf(buf, "    status - %s\n", GitStatus)

	return buf.Bytes()
}
