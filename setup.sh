#!/bin/bash

echo "Setting up..."

go build
chmod +x assembler

cd go_scripts/format_circuit_output
go build
chmod +x format_circuit_output
cd ../..

echo "Required setps completed!"

add_shell_variables () {
    file="$1"
    cur_pwd=$(pwd)
    echo "Adding shell variables to $file..."
    echo "PATH=\$PATH:$cur_pwd" >> $file
}

may_add_shell_variables () {
    BASHRC="$HOME/.bashrc"
    ZSHRC="$HOME/.zshrc"

    echo "Asking to add shell variables..."

    exists_bash=""
    if [ -f "$BASHRC" ]; then
        exists_bash="true"
    else
        exists_bash="false"
    fi

    exists_zsh=""
    if [ -f "$ZSHRC" ]; then
        exists_zsh="true"
    else
        exists_zsh="false"
    fi

    if [ "$exists_bash" = "true" ] | [ "$exists_zsh" = "true" ]; then
        echo ""
        read -n1 -p "Should I add the assembler command to the shell variables? [Y,n] " doit
        echo ""
        case $doit in  
        n|N) ;;
        y|Y|*)
            if [ "$exists_bash" = "true" ]; then
                add_shell_variables $BASHRC
            fi
            if [ "$exists_zsh" = "true" ]; then
                add_shell_variables $ZSHRC
            fi
            ;;
        esac
        echo ""
    else
        echo "No bashrc nor zshrc found!"
        YELLOW='\033[0;33m'
        NC='\033[0m' # No Color
        printf "${YELLOW}[Warning] Add this directory to your shell variables or execute the ./assemble file directly!${NC}\n"
    fi
}

may_add_shell_variables
echo "All setup up!"
