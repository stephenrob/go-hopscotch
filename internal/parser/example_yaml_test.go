package parser

import "fmt"

// Parses a valid yaml workflow definition and returns the workflow.
func ExampleParseYamlWorkflow() {

	inputFile := "testdata/complete_workflow.yaml"

	wf, _ := ParseYamlWorkflow(inputFile)

	fmt.Printf("%#v", wf)

	// Output:
	// &parser.Workflow{ID:"lulibrary.workflows.reloadAllUsers", Name:"Reload All Users", Triggers:parser.WorkflowTriggers{Schedule:"0/1 * * * *", Manual:false, Workflows:[]string{"lulibrary.workflows.fullReload"}}, Args:map[string]string{"a1":"Argument1"}, Params:[]parser.WorkflowParameter{parser.WorkflowParameter{Name:"p1", Type:"string"}, parser.WorkflowParameter{Name:"p2", Type:"boolean"}}, Tasks:map[string]parser.WorkflowTask{"get_all_users":parser.WorkflowTask{Name:"Get All Users", Uses:"Hopscotch/UseJob@v1", Config:map[string]string{"job":"UserLoader::Jobs::GetAllUserCids"}, Args:map[string]string{"hello":"World"}}, "publish_initial_message":parser.WorkflowTask{Name:"Publish get all users", Uses:"Hopscotch/PublishMessage@v1", Config:map[string]string{"data_template":"{\n  \"hello\": \"{{.Args.a1}}\"\n}\n", "type":"userLoader.ldap.getAllUserCids", "version":"0.1.0"}, Args:map[string]string(nil)}, "reload_all_users":parser.WorkflowTask{Name:"Reload All Users", Uses:"Hopscotch/UseJob@v1", Config:map[string]string{"job":"UserLoader::Jobs::ReloadAllUserCids"}, Args:map[string]string(nil)}}, Notifications:map[string]parser.WorkflowNotification{"onSuccess":parser.WorkflowNotification{User:true, Workflow:map[string]bool(nil), Schedule:false}}}
}

// Parses a workflow where there is no file present at the filepath and returns error
func ExampleParseYamlWorkflow_noFile() {

	inputFile := "testdata/no_file.yaml"

	_, err := ParseYamlWorkflow(inputFile)

	if err != nil {
		fmt.Printf("Invalid Yaml File: unable to open file testdata/no_file.yaml")
	}

	// Output:
	// Invalid Yaml File: unable to open file testdata/no_file.yaml
}

// Parses a workflow where there is a non yaml file present at the filepath and returns error
func ExampleParseYamlWorkflow_invalidFile() {

	inputFile := "testdata/workflow.csv"

	_, err := ParseYamlWorkflow(inputFile)

	if err != nil {
		fmt.Printf("Invalid Yaml File: error parsing yaml file")
	}

	// Output:
	// Invalid Yaml File: error parsing yaml file
}
