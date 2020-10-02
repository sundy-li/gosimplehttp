package util

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func ConvertToPdf(srcFile string, writer io.Writer) (err error) {
	args := []string{"--headless", "--invisible", "--convert-to", "pdf", srcFile, "--outdir", os.TempDir()}
	cmd := exec.Command("soffice", args...)

	expire := 100
	time.AfterFunc(time.Duration(expire)*time.Second, func() {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
	})

	err = cmd.Run()
	if err != nil {
		return
	}

	fs, err := os.Open(srcFile)
	if err != nil {
		return
	}
	stat, _ := fs.Stat()

	lastIdx := strings.LastIndex(stat.Name(), ".")

	outFile := filepath.Join(os.TempDir(), stat.Name()[:lastIdx]+".pdf")
	fsout, err := os.Open(outFile)
	if err != nil {
		return
	}

	_, err = io.Copy(writer, fsout)
	return
}
