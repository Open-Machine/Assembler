#!/bin/bash
# This script should be run from this directory.
# The scripts expects two arguments:
# 1. The machine code file
# 2. The number format (check out the format script for all the options)

machineCodeFile="$1"
numFormat="$2"

function run_machine_code() {
    java -jar ../Circuits/logisim-evolution.jar ../Circuits/main.circ -load $machineCodeFile -tty table
}

function format_output() {
    ./go_scripts/format_circuit_output/format_circuit_output $numFormat
}

run_machine_code | format_output
