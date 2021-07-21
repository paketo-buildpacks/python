package integration_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

func testNoPackageManager(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect     = NewWithT(t).Expect
		Eventually = NewWithT(t).Eventually

		pack   occam.Pack
		docker occam.Docker
	)

	it.Before(func() {
		pack = occam.NewPack()
		docker = occam.NewDocker()
	})

	context("when building an app with no package manager", func() {
		var (
			image     occam.Image
			container occam.Container

			name   string
			source string
		)

		it.Before(func() {
			var err error
			name, err = occam.RandomName()
			Expect(err).NotTo(HaveOccurred())
		})

		it.After(func() {
			Expect(docker.Container.Remove.Execute(container.ID)).To(Succeed())
			Expect(docker.Image.Remove.Execute(image.ID)).To(Succeed())
			Expect(docker.Volume.Remove.Execute(occam.CacheVolumeNames(name))).To(Succeed())
			Expect(os.RemoveAll(source)).To(Succeed())
		})

		it("creates a working OCI image with a start command", func() {
			var err error
			source, err = occam.Source(filepath.Join("testdata", "no_package_manager"))
			Expect(err).NotTo(HaveOccurred())

			var logs fmt.Stringer
			image, logs, err = pack.WithNoColor().Build.
				WithBuildpacks(pythonBuildpack).
				WithPullPolicy("never").
				Execute(name, source)
			Expect(err).NotTo(HaveOccurred(), logs.String())

			Expect(logs).To(ContainLines(ContainSubstring("CPython Buildpack")))
			Expect(logs).To(ContainLines(ContainSubstring("Python Start Buildpack")))

			container, err = docker.Container.Run.
				WithCommand("hello.py").
				Execute(image.ID)
			Expect(err).NotTo(HaveOccurred())

			Eventually(func() string {
				cLogs, err := docker.Container.Logs.Execute(container.ID)
				Expect(err).NotTo(HaveOccurred())
				return cLogs.String()
			}).Should(ContainSubstring("Hello"))
		})

		context("when using Procfile to set the start command", func() {
			it("creates a working OCI image that starts the right process", func() {
				var err error
				source, err = occam.Source(filepath.Join("testdata", "no_package_manager"))
				Expect(err).NotTo(HaveOccurred())

				Expect(ioutil.WriteFile(filepath.Join(source, "Procfile"),
					[]byte("web: python hello.py"), os.ModePerm)).
					To(Succeed())

				var logs fmt.Stringer
				image, logs, err = pack.WithNoColor().Build.
					WithBuildpacks(pythonBuildpack).
					WithPullPolicy("never").
					WithEnv(map[string]string{
						"BPE_SOME_VARIABLE": "some-value",
					}).
					Execute(name, source)
				Expect(err).NotTo(HaveOccurred(), logs.String())

				Expect(logs).To(ContainLines(ContainSubstring("CPython Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Python Start Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Procfile Buildpack")))
				Expect(logs).To(ContainLines(ContainSubstring("Environment Variables Buildpack")))

				Expect(image.Buildpacks[3].Key).To(Equal("paketo-buildpacks/environment-variables"))
				Expect(image.Buildpacks[3].Layers["environment-variables"].Metadata["variables"]).To(Equal(map[string]interface{}{"SOME_VARIABLE": "some-value"}))

				container, err = docker.Container.Run.Execute(image.ID)
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() string {
					cLogs, err := docker.Container.Logs.Execute(container.ID)
					Expect(err).NotTo(HaveOccurred())
					return cLogs.String()
				}).Should(ContainSubstring("Hello"))
			})
		})
	})
}
