package repl

import (
	"bufio"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
	"strings"

	"github.com/fatih/color"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		color.New(color.FgGreen).Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}
		evaluator.DefineMacros(program, env)
		expanded := evaluator.ExpandMacros(program, env)
		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			if strings.HasPrefix(evaluated.Inspect(), "ERROR") {
				color.New(color.FgRed).Fprintln(out, evaluated.Inspect())
			} else {
				color.New(color.FgBlue).Fprintln(out, evaluated.Inspect())
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "parser error happen!!!: \n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
