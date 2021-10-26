package cmd

import (
	"bytes"
	"encoding/base32"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zeebo/blake3"
)

func HashContents(contents []byte) string {
	hash := blake3.Sum256(contents)
	var output bytes.Buffer
	_, err := base32.NewEncoder(base32.StdEncoding, &output).Write(hash[:])

	if err != nil {
		panic(err)
	}

	return output.String()
}

var (
	cachedir string
)

func Run(cmd *cobra.Command, args []string) {
	sourcefile := args[0]

	// get contents
	contents, err := os.ReadFile(sourcefile)
	if err != nil {
		panic(err)
	}

	hash := HashContents(contents)

	// get filenames and directories
	cachedir = strings.ReplaceAll(cachedir, "\\", "/")
	executable := path.Join(cachedir, hash+".exe")

	// create cache dir
	_, err = os.Stat(cachedir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(cachedir, 0666)
		panic(err)
	}

	// check for the executable in cgscript dir
	_, err = os.Stat(executable)
	if os.IsNotExist(err) {
		// compiling go executable

		compile := exec.Command("go", "build", "-o", executable, sourcefile)
		compile.Stderr = os.Stderr
		compile.Stdout = os.Stdout
		compile.Stdin = os.Stdin

		err = compile.Run()
		if err != nil {
			panic(err)
		}
	}

	// run the script
	script := exec.Command(executable)
	script.Stderr = os.Stderr
	script.Stdout = os.Stdout
	script.Stdin = os.Stdin

	err = script.Run()
	if err != nil {
		panic(err)
	}
}

var rootCmd = &cobra.Command{
	Use:   "cgscript [script]",
	Short: "A caching runner for scripts written in go",
	Args:  cobra.ExactArgs(1),
	Run:   Run,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().StringVar(&cachedir, "cache", path.Join(os.TempDir(), "cgscript"), "--cache <CACHE DIR>")
}
