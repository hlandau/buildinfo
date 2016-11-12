package buildinfo

import (
	"bufio"
	"strings"
)

// Takes a golang repository name like "github.com/foo/bar" and returns the
// branch name such as "master" or "v1.0.0". The prefix "tags/" and "heads/"
// are stripped if present. Returns "" if repository not found in build
// information or build information not available. This parses the build
// information so you should cache the results rather than calling this
// frequently.
func RepositoryVersion(repositoryName string) string {
	r := bufio.NewReader(strings.NewReader(Full()))
	for {
		l, err := r.ReadString('\n')
		if err != nil {
			return ""
		}

		if !strings.HasPrefix(l, "git ") {
			continue
		}

		l = strings.Trim(l, "\r\n\t ")
		parts := strings.Split(l, " ")
		if len(parts) < 4 {
			continue
		}

		if parts[1] != repositoryName {
			continue
		}

		v := parts[3]
		if strings.HasPrefix(v, "tags/") {
			v = v[5:]
		} else if strings.HasPrefix(v, "heads/") {
			v = v[6:]
		}

		return v
	}

	return ""
}

func IsVersionName(name string) bool {
	return len(name) >= 2 && name[0] == 'v' && name[1] >= '0' && name[1] <= '9'
}

func VersionSummary(repositoryName, shortName string) string {
	v := RepositoryVersion(repositoryName)
	if IsVersionName(v) {
		v = v[1:]
	} else if v == "" {
		v = "unknown"
	}
	return shortName + "/" + v
}

func GoVersionSummary() string {
	return goVersionSummary
}
