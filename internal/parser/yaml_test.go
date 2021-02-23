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
		Context("Workflow Definition", func() {
			It("has correct workflow ID", func() {
				Expect(wf.ID).To(Equal("lulibrary.workflows.reloadAllUsers"))
			})
			It("has correct workflow name", func() {
				Expect(wf.Name).To(Equal("Reload All Users"))
			})
		})
		Context("Workflow Triggers", func() {
			It("has workflow triggers defined", func() {
				Expect(wf.Triggers).ToNot(BeZero())
			})
			It("has correct scheduled trigger", func() {
				Expect(wf.Triggers.Schedule).To(Equal("0/1 * * * *"))
			})
			It("has correct manual trigger", func() {
				Expect(wf.Triggers.Manual).To(BeFalse())
			})
			It("has correct workflows in workflow triggers", func() {
				Expect(wf.Triggers.Workflows).To(ContainElement("lulibrary.workflows.fullReload"))
			})
		})
		Context("Workflow Arguments", func() {
			It("has workflow arguments defined", func() {
				Expect(wf.Args).ToNot(BeZero())
			})
			It("has correct workflow argument", func() {
				Expect(wf.Args).To(HaveKey("a1"))
				Expect(wf.Args["a1"]).To(Equal("Argument1"))
			})
		})
		Context("Workflow Parameters", func() {
			It("has workflow parameters defined", func() {
				Expect(wf.Params).ToNot(BeZero())
			})
			It("has correct parameter definitions", func() {
				p1 := WorkflowParameter{
					Name: "p1",
					Type: "string",
				}
				p2 := WorkflowParameter{
					Name: "p2",
					Type: "boolean",
				}
				Expect(wf.Params).To(ContainElements([]WorkflowParameter{p1, p2}))
			})
		})
		Context("Workflow Tasks", func() {
			It("has workflow tasks defined", func() {
				Expect(wf.Tasks).ToNot(BeZero())
			})
			It("has correct task definition for publish initial message", func() {
				Expect(wf.Tasks).To(HaveKey("publish_initial_message"))
				task := wf.Tasks["publish_initial_message"]
				Expect(task.Name).To(Equal("Publish get all users"))
				Expect(task.Uses).To(Equal("Hopscotch/PublishMessage@v1"))
				Expect(task.Config).ToNot(BeZero())
				Expect(task.Args).To(BeZero())
			})
			It("has correct task config for publish initial message", func() {
				task := wf.Tasks["publish_initial_message"]
				Expect(task.Config).ToNot(BeZero())
				Expect(task.Config).To(HaveKey("type"))
				Expect(task.Config).To(HaveKey("version"))
				Expect(task.Config).To(HaveKey("data_template"))
			})
			It("has correct task definition for get all users", func() {
				Expect(wf.Tasks).To(HaveKey("get_all_users"))
				task := wf.Tasks["get_all_users"]
				Expect(task.Name).To(Equal("Get All Users"))
				Expect(task.Uses).To(Equal("Hopscotch/UseJob@v1"))
				Expect(task.Config).ToNot(BeZero())
				Expect(task.Args).ToNot(BeZero())
			})
			It("has correct task config for get all users", func() {
				task := wf.Tasks["get_all_users"]
				Expect(task.Config).ToNot(BeZero())
				Expect(task.Config).To(HaveKey("job"))
			})
			It("has correct task arguments for get all users", func() {
				task := wf.Tasks["get_all_users"]
				Expect(task.Args).ToNot(BeZero())
				Expect(task.Args).To(HaveKey("hello"))
			})
		})
		Context("Workflow Notifications", func() {
			It("has workflow notifications defined", func() {
				Expect(wf.Notifications).ToNot(BeZero())
			})
			It("has correct workflow notification definition for onSuccess", func() {
				Expect(wf.Notifications).To(HaveKey("onSuccess"))
				Expect(wf.Notifications["onSuccess"]).ToNot(BeZero())
				Expect(wf.Notifications["onSuccess"].User).To(BeTrue())
				Expect(wf.Notifications["onSuccess"].Schedule).To(BeZero())
				Expect(wf.Notifications["onSuccess"].Workflow).To(BeZero())
			})
		})
	})
})
