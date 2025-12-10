package integration_test

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/paketo-buildpacks/occam"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
	. "github.com/paketo-buildpacks/occam/matchers"
)

func testPoetryRun(t *testing.T, context spec.G, it spec.S) {
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

	context("when building a poetry app with a poetry run script", func() {
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
			source, err = occam.Source(filepath.Join("testdata", "poetry-run"))
			Expect(err).NotTo(HaveOccurred())

			var logs fmt.Stringer
			image, logs, err = pack.WithNoColor().Build.
				WithBuildpacks(pythonBuildpack).
				WithPullPolicy("never").
				WithEnv(map[string]string{
					"BPE_SOME_VARIABLE":      "some-value",
					"BP_IMAGE_LABELS":        "some-label=some-value",
					"BP_LIVE_RELOAD_ENABLED": "true",
				}).
				Execute(name, source)
			Expect(err).NotTo(HaveOccurred(), logs.String())

			container, err = docker.Container.Run.
				WithEnv(map[string]string{"PORT": "8080"}).
				WithPublish("8080").
				WithPublishAll().
				Execute(image.ID)
			Expect(err).NotTo(HaveOccurred())

			Eventually(container).Should(BeAvailable())

			response, err := http.Get(fmt.Sprintf("http://localhost:%s", container.HostPort("8080")))
			Expect(err).NotTo(HaveOccurred())
			defer func() { Expect(response.Body.Close()).ToNot(HaveOccurred()) }()

			Expect(response.StatusCode).To(Equal(http.StatusOK))

			content, err := io.ReadAll(response.Body)
			Expect(err).NotTo(HaveOccurred())
			Expect(string(content)).To(ContainSubstring("Hello, World!"))

			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for CA Certificates")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for CPython")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Pip")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Poetry")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Poetry Install")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Poetry Run")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Environment Variables")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Image Labels")))
			Expect(logs).To(ContainLines(ContainSubstring("Buildpack for Watchexec")))

			Expect(image.Buildpacks[7].Key).To(Equal("paketo-buildpacks/environment-variables"))
			Expect(image.Buildpacks[7].Layers["environment-variables"].Metadata["variables"]).To(Equal(map[string]interface{}{"SOME_VARIABLE": "some-value"}))
			Expect(image.Labels["some-label"]).To(Equal("some-value"))
		})
	})
}
