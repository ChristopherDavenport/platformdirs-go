package platformdirs

import (
	"regexp"
	"testing"
)

func TestPlatformDirsSimple(t *testing.T) {
	var appName = "test"
	var appAuthor = "davencorp"
	dirs := New(appName, appAuthor)
	dir, err := dirs.UserDataDir()
	want := regexp.MustCompile("/home/([A-z]*)/.local/share/" + appName)
	if (!want.MatchString(dir)) || err != nil {
		t.Fatalf(`New("test", "davencorp").UserDataDir() = %q, %v, want match for %#q, nil`, dir, err, want)
	}
}
