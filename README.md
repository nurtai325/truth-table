Truth Table Generator

A command-line tool to generate truth tables for logical expressions. This tool evaluates logical expressions with standard logical operators and displays the corresponding truth table.
Features

    Logical Operators Supported:
        ! (NOT)
        and (AND)
        or (OR)
        -> (Implication)
        <=> (If and only if)
    Command-line Interface: Input logical expressions and receive their truth tables.
    REPL Mode: Enter an expression to evaluate it immediately.

Installation
Prerequisites

    Go (version 1.16+)
    Make

Steps

    Clone the repository:

git clone https://github.com/yourusername/truth-table-generator.git
cd truth-table-generator

Make sure you have Go and Make installed. You can check by running:

go version
make --version

Build the project:

make

Run the program:

    ./truth-table

Usage

Once the program is running, you can enter expressions directly via the REPL. Here are some key commands:

    -h : Show help information (displays usage and available commands).

    -e <expression> : Specify a logical expression to generate its truth table. For example:

    ./truth-table -e "(A and B) or !C"

    This command will output the truth table for the provided logical expression.

Example

$ ./truth-table -e "(A and B) or !C"
A | B | C | (A and B) or !C
---------------------------
T | T | T | T
T | T | F | T
T | F | T | F
T | F | F | T
F | T | T | T
F | T | F | F
F | F | T | F
F | F | F | F

Supported Operators

    NOT (!): Negates the truth value of the operand.
    AND (and): True if both operands are true.
    OR (or): True if at least one operand is true.
    Implication (->): True unless the first operand is true and the second is false.
    If and Only If (<=>): True if both operands are the same (either both true or both false).

Contributing

Contributions are welcome! Fork the repository, make improvements or fix bugs, and submit a pull request. For any issues or feature requests, please open an issue.
License

This project is licensed under the MIT License - see the LICENSE file for details.
