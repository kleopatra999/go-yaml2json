/*
Package main implements the yaml2json command-line tool, which takes
a YAML file as stdin, and outputs a JSON file to stdout.

See the README at https://github.com/peter-edge/go-yaml2json/blob/master/README.md for more details.
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/peter-edge/go-yaml2json"
)

func main() {
	if err := do(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func do() error {
	var pretty bool
	flag.BoolVar(&pretty, "pretty", false, "Make the JSON output pretty.")
	flag.Parse()

	yamlData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	jsonData, err := yaml2json.Transform(yamlData, yaml2json.TransformOptions{Pretty: pretty})
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(jsonData); err != nil {
		return err
	}
	return nil
}
