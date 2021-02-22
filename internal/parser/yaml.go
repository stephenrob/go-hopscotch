// Package parser provides functions for parsing workflows defined in yaml and json files
package parser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type InvalidYamlFile struct {
	err      error
	filepath string
	message  string
}

func (e *InvalidYamlFile) Error() string {
	return fmt.Sprintf("Invalid Yaml File: %s", e.message)
}

// ParseYamlWorkflow reads the file at the given filepath and unmarshals the data as yaml into a workflow.
// Returns an error if the filepath is invalid or there is an error unmarshalling the data.
func ParseYamlWorkflow(filepath string) (*Workflow, error) {

	if filepath == "" {
		err := &InvalidYamlFile{
			err:      fmt.Errorf("provided filepath was empty, please provide path to yaml file"),
			filepath: filepath,
			message:  "provided filepath was empty",
		}
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(filepath)

	if err != nil {
		err := &InvalidYamlFile{
			err:      err,
			filepath: filepath,
			message:  fmt.Sprintf("unable to open file %s", filepath),
		}
		return nil, err
	}

	var workflow Workflow
	err = yaml.Unmarshal(yamlFile, &workflow)

	if err != nil {
		err := &InvalidYamlFile{
			err:      err,
			filepath: filepath,
			message:  fmt.Sprintf("error parsing yaml file"),
		}
		return nil, err
	}

	return &workflow, nil

}
