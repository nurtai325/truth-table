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

    Go
    Make

Steps

    git clone https://github.com/yourusername/truth-table](https://github.com/nurtai325/truth-table
    cd truth-table-generator
    make

Run the program:

    ./truth-table

Usage

Once the program is running, you can enter expressions directly via the REPL. Here are some key commands:

    -h : Show help information (displays usage and available commands).

    -e <expression> : Specify a logical expression to generate its truth table. For example:

    ./truth-table -e 'a||!b->!(c<=>d)'

    This command will output the truth table for the provided logical expression.

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
