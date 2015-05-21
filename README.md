# go-yaml2json

[![API Documentation](http://img.shields.io/badge/api-Godoc-blue.svg?style=flat-square)](https://godoc.org/github.com/peter-edge/go-yaml2json)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/peter-edge/go-yaml2json/blob/master/LICENSE)

yaml2json is a super small library to transform YAML input to JSON output.

```go
import (
	"io/ioutil"
	"os"

	"github.com/peter-edge/go-yaml2json"
)

func do() error {
	yamlData, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	jsonData, err := yaml2json.Transform(yamlData, yaml2json.TransformOptions{Pretty: true})
	if err != nil {
		return err
	}
	if _, err := os.Stdout.Write(jsonData); err != nil {
		return err
	}
}
```

There is also a command-line tool at [cmd/yaml2json](cmd/yaml2json) that reads a YAML file from stdin
and outputs a JSON file to stdout. Run `make install` or `go get -v github.com/peter-edge/go-yaml2json/cmd/yaml2json`
to install.

Originally inspired by https://github.com/bronze1man/yaml2json.
