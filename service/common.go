package service

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type Mod struct {
	Path  string
	Dir   string
	GoMod string
}

type CurDir struct {
	Dir string
}

func parseModInfo() (Mod, CurDir, error) {
	var mod Mod
	var dir CurDir

	m, err := modInfoJSON()
	if err != nil {
		return mod, dir, err
	}

	/*	fmt.Println("module", m)*/
	errParseMod := json.Unmarshal(m, &mod)
	if errParseMod != nil {
		return mod, dir, fmt.Errorf("parse module json errParseMod: %v", errParseMod)
	}

	// Unsure why, but if no module is present Path is set to this string.
	if mod.Path == "command-line-arguments" {
		cobra.CheckErr("Please run `go mod init <MODNAME>` before `cobra-cli init`")
	}

	e, err := modDirInfoJSON()
	if err != nil {
		return mod, dir, err
	}

	errParseDir := json.Unmarshal(e, &dir)
	if errParseDir != nil {
		return mod, dir, fmt.Errorf("parse module json errParseDir: %v", errParseDir)
	}

	return mod, dir, err
}

func getModImportPath() (string, error) {
	mod, cd, err := parseModInfo()
	if err != nil {
		return "", err
	}
	fmt.Printf(fileToURL(strings.TrimPrefix(cd.Dir, mod.Dir)))
	return path.Join(mod.Path, fileToURL(strings.TrimPrefix(cd.Dir, mod.Dir))), nil
}

func fileToURL(in string) string {
	i := strings.Split(in, string(filepath.Separator))
	return path.Join(i...)
}

/*func modInfoJSON(args ...string) ([]byte, error) {
	cmdArgs := append([]string{"list", "-json"}, args...)
	out, err := exec.Command("go", cmdArgs...).Output()

	return out, err
}
*/

func modDirInfoJSON() ([]byte, error) {
	cmdArgs := append([]string{"list", "-json", "-e"})
	out, err := exec.Command("go", cmdArgs...).Output()

	return out, err
}

func modInfoJSON() ([]byte, error) {
	cmdArgs := append([]string{"list", "-json", "-m"})
	out, err := exec.Command("go", cmdArgs...).Output()

	modInfo := gjson.Get(string(out), "..0")

	if !modInfo.Exists() {
		return out, fmt.Errorf("module info form 'go list -json -m' is empty")
	}

	return []byte(modInfo.String()), err
}
