package buildinfo

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/hlandau/easyconfig.v1/cflag"
	"os"
	"runtime"
	"strings"
)

// Full build info.
var BuildInfo string

// Set via go build.
var RawBuildInfo string

// Program-settable extra version information to print.
var Extra string

var goVersionSummary string

// You should never need to call this.
func Update() {
	if RawBuildInfo == "" || BuildInfo != "" {
		return
	}

	b, err := base64.RawStdEncoding.DecodeString(strings.TrimRight(RawBuildInfo, "="))
	if err != nil {
		return
	}

	BuildInfo = string(b)
}

func init() {
	goVersionSummary = runtime.Version() + "/" + runtime.GOARCH + "/" + runtime.GOOS + "/" + runtime.Compiler
	if Cgo {
		goVersionSummary += "+cgo"
	}

	versionFlag := cflag.Bool(nil, "version", false, "Print version information")
	versionFlag.RegisterOnChange(func(bf *cflag.BoolFlag) {
		if !bf.Value() {
			return
		}

		fmt.Print(Full())
		os.Exit(2)
	})

	Update()
}

func Full() string {
	bi := BuildInfo
	if bi == "" {
		bi = "build unknown"
	}
	return fmt.Sprintf("%sgo version %s %s/%s %s cgo=%v\n%s\n", Extra, runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.Compiler, Cgo, bi)
}
