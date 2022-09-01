package integration_test

import (
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
)

var pythonBuildpack string

func TestIntegration(t *testing.T) {
	Expect := NewWithT(t).Expect

	format.MaxLength = 0

	output, err := exec.Command("bash", "-c", "../scripts/package.sh --version 1.2.3").CombinedOutput()
	Expect(err).NotTo(HaveOccurred(), string(output))

	pythonBuildpack, err = filepath.Abs("../build/buildpackage.cnb")
	Expect(err).NotTo(HaveOccurred())

	SetDefaultEventuallyTimeout(10 * time.Second)

	suite := spec.New("Integration", spec.Report(report.Terminal{}), spec.Parallel())
	suite("Conda", testConda)
	suite("Pip", testPip)
	suite("Pipenv", testPipenv)
	suite("PoetryDepOnly", testPoetryDepOnly)
	suite("PoetryRun", testPoetryRun)
	suite("NoPackageManager", testNoPackageManager)
	suite.Run(t)
}
