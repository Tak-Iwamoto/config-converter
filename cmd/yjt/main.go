package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Tak-Iwamoto/yjt"
)

func main() {
	os.Exit(run())
}

func run() int {
	if len(os.Args) < 2 {
		fmt.Printf("Please input filename")
		return 1
	}
	fileName := os.Args[1]
	file, err := os.ReadFile(fileName)

	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	var output = fs.String("o", "", "output file format")
	fs.Parse(os.Args[2:])

	if err != nil {
		fmt.Println(err.Error())
		return 1
	}

	ext := fetchExt(fileName)

	switch ext {
	case "json":
		j, err := convertJson(file, *output)
		if err != nil {
			writeToStdOut(err.Error())
		}
		writeToStdOut(string(j))
	case "yml", "yaml":
		y, err := convertYaml(file, *output)
		if err != nil {
			writeToStdOut(err.Error())
		}
		writeToStdOut(string(y))
	case "toml":
		t, err := convertToml(file, *output)
		if err != nil {
			writeToStdOut(err.Error())
		}
		writeToStdOut(string(t))
	default:
		return 1
	}
	return 0
}

func writeToStdOut(s string) error {
	_, err := io.WriteString(os.Stdout, s)
	if err != nil {
		return err
	}
	return nil
}

func fetchExt(fileName string) string {
	return filepath.Ext(fileName)[1:]
}

func convertJson(j []byte, output string) ([]byte, error) {
	switch output {
	case "yml", "yaml":
		fmt.Println("convert to yaml")
		return yjt.JsonToYaml(j)
	case "toml":
		return yjt.JsonToToml(j)
	case "json":
		return j, nil
	default:
		fmt.Println(output)
		return nil, fmt.Errorf("failed to convert json to %v", output)
	}
}

func convertYaml(y []byte, output string) ([]byte, error) {
	switch output {
	case "json":
		return yjt.YamlToJson(y)
	case "toml":
		return yjt.YamlToToml(y)
	case "yml", "yaml":
		return y, nil
	default:
		return nil, fmt.Errorf("failed to convert yaml to %v", output)
	}
}

func convertToml(t []byte, output string) ([]byte, error) {
	switch output {
	case "json":
		return yjt.TomlToJson(t)
	case "yml", "yaml":
		return yjt.TomlToYaml(t)
	case "toml":
		return t, nil
	default:
		return nil, fmt.Errorf("failed to convert toml to %v", output)
	}
}
