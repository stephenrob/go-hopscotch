package parser

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Yaml File Parsing", func() {
	var (
		wf  *Workflow
		err error
	)
	Describe("Invalid Yaml File", func() {
		Context("Empty Filepath", func() {
			BeforeEach(func() {
				wf, err = ParseYamlWorkflow("")
			})
			It("should return a nil workflow", func() {
				Expect(wf).To(BeNil())
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).Should(MatchError("Invalid Yaml File: provided filepath was empty"))
			})
		})
		Context("No file at path", func() {
			BeforeEach(func() {
				wf, err = ParseYamlWorkflow("testdata/invalid_file.yaml")
			})
			It("should return a nil workflow", func() {
				Expect(wf).To(BeNil())
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).Should(MatchError("Invalid Yaml File: unable to open file testdata/invalid_file.yaml"))
			})
		})
		Context("Given a non-yaml file", func() {
			BeforeEach(func() {
				wf, err = ParseYamlWorkflow("testdata/workflow.csv")
			})
			It("should return a nil workflow", func() {
				Expect(wf).To(BeNil())
			})
			It("should error", func() {
				Expect(err).To(HaveOccurred())
				Expect(err).Should(MatchError("Invalid Yaml File: error parsing yaml file"))
			})
		})
	})

	Describe("Valid Yaml File", func() {
		BeforeEach(func() {
			Expect("testdata/complete_workflow.yaml").To(BeAnExistingFile())
			wf, err = ParseYamlWorkflow("testdata/complete_workflow.yaml")
		})
		Context("Unmarshal yaml file to workflow", func() {
			It("should return no error", func() {
				Expect(err).ToNot(HaveOccurred())
			})
			It("should return workflow", func() {
				Expect(wf).ToNot(BeNil())
			})
		})
	})
})
