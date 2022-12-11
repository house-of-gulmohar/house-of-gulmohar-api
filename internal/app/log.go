package app

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {

	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		DisableLevelTruncation: false, // log level field configuration
		DisableTimestamp:       true,
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmet everything before the file name
			// and added a line number in the end
			now := time.Now()
			return "", fmt.Sprintf(
				": %d/%d/%d %d:%d:%d %s:%d ",
				now.Year(), now.Month(), now.Day(),
				now.Hour(), now.Minute(), now.Second(),
				formatFilePath(f.File), f.Line,
			)
		},
	}
	logrus.SetFormatter(formatter)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
