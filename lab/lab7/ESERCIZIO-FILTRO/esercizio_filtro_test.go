package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFiltro(t *testing.T) {
	RunTest(t, []string{"344"}, "Output: 34")
	RunTest(t, []string{"35565"}, "Output: 3565")
	RunTest(t, []string{"5555"}, "Output: 5")
}

func RunTest(t *testing.T, input []string, result string) {
	tempFile, err := ioutil.TempFile("", "out")
	if err != nil {
		t.Error("Error while creating temporary file")
	}
	defer tempFile.Close()

	tempFile, os.Stdout = os.Stdout, tempFile

	inFile, err := ioutil.TempFile("", "in")
	if err != nil {
		t.Error("Error while creating temporary file")
	}
	defer inFile.Close()

	for i, s := range input {
		if i == 0 {
			fmt.Fprint(inFile, s)
		} else {
			fmt.Fprint(inFile, " ", s)
		}
	}

	inFile, os.Stdin = os.Stdin, inFile

	os.Stdin.Seek(0, 0)

	main()

	_, err = os.Stdout.Seek(0, 0)
	if err != nil {
		t.Error("Error while resetting file pointer temporary file")
	}
	output, err := ioutil.ReadAll(os.Stdout)

	if err != nil {
		t.Error("Error while reading output temporary file")
	}
	if tOut, tRes := strings.Trim(string(output), "\n"), strings.Trim(result, "\n"); tOut != tRes {
		t.Error(string(output))
		t.Error(fmt.Sprintf("\nInput: %s\nExpected result: <%s>\nGiven result: <%s>\n", strings.Join(input, " "), tRes, tOut))
	}

	tempFile, os.Stdout = os.Stdout, tempFile
	inFile, os.Stdin = os.Stdin, inFile
}
