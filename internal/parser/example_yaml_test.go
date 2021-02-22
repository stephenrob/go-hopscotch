package parser

import "fmt"

func ExampleParseYamlWorkflow() {

	inputFile := "testdata/complete_workflow.yaml"

	wf, _ := ParseYamlWorkflow(inputFile)

	fmt.Printf("%+v\n", wf)

	// Output:
	// &{ID:lulibrary.workflows.reloadAllUsers Name:Reload All Users}
}

func ExampleParseYamlWorkflow_noFile() {

	inputFile := "testdata/no_file.yaml"

	_, err := ParseYamlWorkflow(inputFile)

	if err != nil {
		fmt.Printf("Invalid Yaml File: unable to open file testdata/no_file.yaml")
	}

	// Output:
	// Invalid Yaml File: unable to open file testdata/no_file.yaml
}

func ExampleParseYamlWorkflow_invalidFile() {

	inputFile := "testdata/workflow.csv"

	_, err := ParseYamlWorkflow(inputFile)

	if err != nil {
		fmt.Printf("Invalid Yaml File: error parsing yaml file")
	}

	// Output:
	// Invalid Yaml File: error parsing yaml file
}
