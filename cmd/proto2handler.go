package main

import (
	"flag"
	"fmt"
	"os"

	"proto-handler-generator/generator"
)

func main() {
	protoPath := flag.String("proto-path", "./protos", "Path to .proto files")
	outputPath := flag.String("output-path", "./generated", "Path to output generated code")
	flag.Parse()

	err := generator.GenerateCode(*protoPath, *outputPath)
	if err != nil {
		fmt.Println("Error during code generation:", err)
		os.Exit(1)
	}

	fmt.Println("Code generation completed successfully.")
}
