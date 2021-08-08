package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestBinaryToInt(t *testing.T) {
	tests := []struct {
		param  string
		expect int64
	}{
		{
			param:  "0100101001010001",
			expect: 19025,
		},
		{
			param:  "1100101001010001",
			expect: -19025,
		},
	}

	for _, test := range tests {
		num, err := binaryToInt(test.param)
		if err != nil || num != test.expect {
			t.Errorf("Error: err='%s', %d =? %d", err, num, test.expect)
		}
	}
}

func TestHexBinaryConversion(t *testing.T) {
	tests := []struct {
		param  int
		expect string
	}{
		{
			param:  19025,
			expect: "0x4A51",
		},
		{
			param:  -19025,
			expect: "-0x4A51",
		},
	}

	for _, test := range tests {
		got := binaryToHexadecimal(test.param)
		if got != test.expect {
			t.Errorf("Error: '%s'!='%s'", got, test.expect)
		}
	}
}

func TestDecimalBinaryConversion(t *testing.T) {
	tests := []struct {
		param  int
		expect string
	}{
		{
			param:  19025,
			expect: "19025",
		},
		{
			param:  -19025,
			expect: "-19025",
		},
	}

	for _, test := range tests {
		got := binaryToDecimal(test.param)
		if got != test.expect {
			t.Errorf("Error: '%s'!='%s'", got, test.expect)
		}
	}
}

func TestScriptError(t *testing.T) {
	var in bytes.Buffer
	var out bytes.Buffer

	args := []string{"./format_circuit_output"}

	program(args, &in, &out)

	stdout := out.String()
	if !strings.Contains(stdout, "error") {
		t.Errorf("Expected error: %s", stdout)
	}
}

func TestScript(t *testing.T) {
	tests := []struct {
		format string
		input  string
		expect string
	}{
		{
			format: "b",
			input:  "0101 1010 1111 0001      1111",
			expect: "0101101011110001",
		},
		{
			format: "h",
			input:  "0101 1010 1111 0001      1111",
			expect: "0x5AF1",
		},
		{
			format: "d",
			input:  "0101 1010 1111 0001      1111",
			expect: "23281",
		},
	}

	for _, test := range tests {
		var in bytes.Buffer
		var out bytes.Buffer

		in.WriteString(test.input)

		args := []string{"./format_circuit_output", test.format}

		program(args, &in, &out)

		stdout := out.String()
		if stdout != test.expect+"\n" {
			t.Errorf("Error: '%s' != '%s'", stdout, test.expect)
		}
	}
}
