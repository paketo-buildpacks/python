package integration_test

import (
	"bytes"
	"path/filepath"
	"testing"
	"time"

	"github.com/paketo-buildpacks/packit/v2/pexec"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var pythonBuildpack string

func TestIntegration(t *testing.T) {
	Expect := NewWithT(t).Expect

	format.MaxLength = 0

	bash := pexec.NewExecutable("bash")
	buffer := bytes.NewBuffer(nil)
	err := bash.Execute(pexec.Execution{
		Args:   []string{"-c", "../scripts/package.sh --version 1.2.3"},
		Stdout: buffer,
		Stderr: buffer,
	})
	Expect(err).NotTo(HaveOccurred(), buffer.String)

	pythonBuildpack, err = filepath.Abs("../build/buildpackage.cnb")
	Expect(err).NotTo(HaveOccurred())

	SetDefaultEventuallyTimeout(10 * time.Second)

	suite := spec.New("Integration", spec.Report(report.Terminal{}))
	suite("Conda", testConda)
	suite("Pip", testPip)
	suite("Pipenv", testPipenv)
	suite("PoetryDepOnly", testPoetryDepOnly)
	suite("PoetryRun", testPoetryRun)
	suite("NoPackageManager", testNoPackageManager)
	suite.Run(t)
}
