package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const endSymbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint8

const (
	ruleUnknown pegRule = iota
	ruleroot
	ruletext
	ruleline
	rulecomments
	ruleblanks
	rulelabel
	ruleexpression
	rulesign
	ruleterm_op
	ruleterm
	rulefactor_op
	rulefactor
	ruleident
	rulenumber
	rule__
	rule_
	ruleEOL
	ruleEOT
	rulePegText
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
)

var rul3s = [...]string{
	"Unknown",
	"root",
	"text",
	"line",
	"comments",
	"blanks",
	"label",
	"expression",
	"sign",
	"term_op",
	"term",
	"factor_op",
	"factor",
	"ident",
	"number",
	"__",
	"_",
	"EOL",
	"EOT",
	"PegText",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
}

type token32 struct {
	pegRule
	begin, end uint32
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v", rul3s[t.pegRule], t.begin, t.end)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(pretty bool, buffer string) {
	var print func(node *node32, depth int)
	print = func(node *node32, depth int) {
		for node != nil {
			for c := 0; c < depth; c++ {
				fmt.Printf(" ")
			}
			rule := rul3s[node.pegRule]
			quote := strconv.Quote(string(([]rune(buffer)[node.begin:node.end])))
			if !pretty {
				fmt.Printf("%v %v\n", rule, quote)
			} else {
				fmt.Printf("\x1B[34m%v\x1B[m %v\n", rule, quote)
			}
			if node.up != nil {
				print(node.up, depth+1)
			}
			node = node.next
		}
	}
	print(node, 0)
}

func (node *node32) Print(buffer string) {
	node.print(false, buffer)
}

func (node *node32) PrettyPrint(buffer string) {
	node.print(true, buffer)
}

type tokens32 struct {
	tree []token32
}

func (t *tokens32) Trim(length uint32) {
	t.tree = t.tree[:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) AST() *node32 {
	type element struct {
		node *node32
		down *element
	}
	tokens := t.Tokens()
	var stack *element
	for _, token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	if stack != nil {
		return stack.node
	}
	return nil
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	t.AST().Print(buffer)
}

func (t *tokens32) PrettyPrintSyntaxTree(buffer string) {
	t.AST().PrettyPrint(buffer)
}

func (t *tokens32) Add(rule pegRule, begin, end, index uint32) {
	if tree := t.tree; int(index) >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	t.tree[index] = token32{
		pegRule: rule,
		begin:   begin,
		end:     end,
	}
}

func (t *tokens32) Tokens() []token32 {
	return t.tree
}

type Parser struct {
	s Scan

	Buffer string
	buffer []rune
	rules  [26]func() bool
	parse  func(rule ...int) error
	reset  func()
	Pretty bool
	tokens32
}

func (p *Parser) Parse(rule ...int) error {
	return p.parse(rule...)
}

func (p *Parser) Reset() {
	p.reset()
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer []rune, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range buffer {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p   *Parser
	max token32
}

func (e *parseError) Error() string {
	tokens, error := []token32{e.max}, "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.buffer, positions)
	format := "parse error near %v (line %v symbol %v - line %v symbol %v):\n%v\n"
	if e.p.Pretty {
		format = "parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n"
	}
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf(format,
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			strconv.Quote(string(e.p.buffer[begin:end])))
	}

	return error
}

func (p *Parser) PrintSyntaxTree() {
	if p.Pretty {
		p.tokens32.PrettyPrintSyntaxTree(p.Buffer)
	} else {
		p.tokens32.PrintSyntaxTree(p.Buffer)
	}
}

func (p *Parser) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for _, token := range p.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:
			p.s.Err(begin)
		case ruleAction1:
			p.s.Err(begin)
		case ruleAction2:

			p.s.line++
			p.s.lineHead = begin + 1

		case ruleAction3:

			p.s.line++
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d COMMENT\t%s\n", p.s.line, text)

		case ruleAction4:

			p.s.line++
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d BLANK  \t%s\n", p.s.line, text)

		case ruleAction5:

			p.s.line++
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d LABEL  \t%s\n", p.s.line, text)

		}
	}
	_, _, _, _, _ = buffer, _buffer, text, begin, end
}

func (p *Parser) Init() {
	var (
		max                  token32
		position, tokenIndex uint32
		buffer               []rune
	)
	p.reset = func() {
		max = token32{}
		position, tokenIndex = 0, 0

		p.buffer = []rune(p.Buffer)
		if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != endSymbol {
			p.buffer = append(p.buffer, endSymbol)
		}
		buffer = p.buffer
	}
	p.reset()

	_rules := p.rules
	tree := tokens32{tree: make([]token32, math.MaxInt16)}
	p.parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokens32 = tree
		if matches {
			p.Trim(tokenIndex)
			return nil
		}
		return &parseError{p, max}
	}

	add := func(rule pegRule, begin uint32) {
		tree.Add(rule, begin, position, tokenIndex)
		tokenIndex++
		if begin != position && position > max.end {
			max = token32{rule, begin, position}
		}
	}

	matchDot := func() bool {
		if buffer[position] != endSymbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 root <- <((text EOT) / (text <.+> Action0 EOT) / (<.+> Action1 EOT))> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position2, tokenIndex2 := position, tokenIndex
					if !_rules[ruletext]() {
						goto l3
					}
					if !_rules[ruleEOT]() {
						goto l3
					}
					goto l2
				l3:
					position, tokenIndex = position2, tokenIndex2
					if !_rules[ruletext]() {
						goto l4
					}
					{
						position5 := position
						if !matchDot() {
							goto l4
						}
					l6:
						{
							position7, tokenIndex7 := position, tokenIndex
							if !matchDot() {
								goto l7
							}
							goto l6
						l7:
							position, tokenIndex = position7, tokenIndex7
						}
						add(rulePegText, position5)
					}
					if !_rules[ruleAction0]() {
						goto l4
					}
					if !_rules[ruleEOT]() {
						goto l4
					}
					goto l2
				l4:
					position, tokenIndex = position2, tokenIndex2
					{
						position8 := position
						if !matchDot() {
							goto l0
						}
					l9:
						{
							position10, tokenIndex10 := position, tokenIndex
							if !matchDot() {
								goto l10
							}
							goto l9
						l10:
							position, tokenIndex = position10, tokenIndex10
						}
						add(rulePegText, position8)
					}
					if !_rules[ruleAction1]() {
						goto l0
					}
					if !_rules[ruleEOT]() {
						goto l0
					}
				}
			l2:
				add(ruleroot, position1)
			}
			return true
		l0:
			position, tokenIndex = position0, tokenIndex0
			return false
		},
		/* 1 text <- <(line line*)> */
		func() bool {
			position11, tokenIndex11 := position, tokenIndex
			{
				position12 := position
				if !_rules[ruleline]() {
					goto l11
				}
			l13:
				{
					position14, tokenIndex14 := position, tokenIndex
					if !_rules[ruleline]() {
						goto l14
					}
					goto l13
				l14:
					position, tokenIndex = position14, tokenIndex14
				}
				add(ruletext, position12)
			}
			return true
		l11:
			position, tokenIndex = position11, tokenIndex11
			return false
		},
		/* 2 line <- <((comments EOL) / (blanks EOL) / (label EOL) / (<((!'\n' .)+ _ EOL)> Action2))> */
		func() bool {
			position15, tokenIndex15 := position, tokenIndex
			{
				position16 := position
				{
					position17, tokenIndex17 := position, tokenIndex
					if !_rules[rulecomments]() {
						goto l18
					}
					if !_rules[ruleEOL]() {
						goto l18
					}
					goto l17
				l18:
					position, tokenIndex = position17, tokenIndex17
					if !_rules[ruleblanks]() {
						goto l19
					}
					if !_rules[ruleEOL]() {
						goto l19
					}
					goto l17
				l19:
					position, tokenIndex = position17, tokenIndex17
					if !_rules[rulelabel]() {
						goto l20
					}
					if !_rules[ruleEOL]() {
						goto l20
					}
					goto l17
				l20:
					position, tokenIndex = position17, tokenIndex17
					{
						position21 := position
						{
							position24, tokenIndex24 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l24
							}
							position++
							goto l15
						l24:
							position, tokenIndex = position24, tokenIndex24
						}
						if !matchDot() {
							goto l15
						}
					l22:
						{
							position23, tokenIndex23 := position, tokenIndex
							{
								position25, tokenIndex25 := position, tokenIndex
								if buffer[position] != rune('\n') {
									goto l25
								}
								position++
								goto l23
							l25:
								position, tokenIndex = position25, tokenIndex25
							}
							if !matchDot() {
								goto l23
							}
							goto l22
						l23:
							position, tokenIndex = position23, tokenIndex23
						}
						if !_rules[rule_]() {
							goto l15
						}
						if !_rules[ruleEOL]() {
							goto l15
						}
						add(rulePegText, position21)
					}
					if !_rules[ruleAction2]() {
						goto l15
					}
				}
			l17:
				add(ruleline, position16)
			}
			return true
		l15:
			position, tokenIndex = position15, tokenIndex15
			return false
		},
		/* 3 comments <- <(<(_ ';' (!'\n' .)* _)> Action3)> */
		func() bool {
			position26, tokenIndex26 := position, tokenIndex
			{
				position27 := position
				{
					position28 := position
					if !_rules[rule_]() {
						goto l26
					}
					if buffer[position] != rune(';') {
						goto l26
					}
					position++
				l29:
					{
						position30, tokenIndex30 := position, tokenIndex
						{
							position31, tokenIndex31 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l31
							}
							position++
							goto l30
						l31:
							position, tokenIndex = position31, tokenIndex31
						}
						if !matchDot() {
							goto l30
						}
						goto l29
					l30:
						position, tokenIndex = position30, tokenIndex30
					}
					if !_rules[rule_]() {
						goto l26
					}
					add(rulePegText, position28)
				}
				if !_rules[ruleAction3]() {
					goto l26
				}
				add(rulecomments, position27)
			}
			return true
		l26:
			position, tokenIndex = position26, tokenIndex26
			return false
		},
		/* 4 blanks <- <(<_> Action4)> */
		func() bool {
			position32, tokenIndex32 := position, tokenIndex
			{
				position33 := position
				{
					position34 := position
					if !_rules[rule_]() {
						goto l32
					}
					add(rulePegText, position34)
				}
				if !_rules[ruleAction4]() {
					goto l32
				}
				add(ruleblanks, position33)
			}
			return true
		l32:
			position, tokenIndex = position32, tokenIndex32
			return false
		},
		/* 5 label <- <(<(ident ':' _)> Action5)> */
		func() bool {
			position35, tokenIndex35 := position, tokenIndex
			{
				position36 := position
				{
					position37 := position
					if !_rules[ruleident]() {
						goto l35
					}
					if buffer[position] != rune(':') {
						goto l35
					}
					position++
					if !_rules[rule_]() {
						goto l35
					}
					add(rulePegText, position37)
				}
				if !_rules[ruleAction5]() {
					goto l35
				}
				add(rulelabel, position36)
			}
			return true
		l35:
			position, tokenIndex = position35, tokenIndex35
			return false
		},
		/* 6 expression <- <(sign term (term_op term)*)> */
		nil,
		/* 7 sign <- <(('-' / '+')? _)> */
		nil,
		/* 8 term_op <- <(('-' / '+') _)> */
		nil,
		/* 9 term <- <(factor (factor_op factor)*)> */
		nil,
		/* 10 factor_op <- <(('*' / '/') _)> */
		nil,
		/* 11 factor <- <(ident / number / ('(' _ expression ')' _))> */
		nil,
		/* 12 ident <- <(([a-z] / [A-Z]) ([a-z] / [A-Z] / [0-9])* _)> */
		func() bool {
			position44, tokenIndex44 := position, tokenIndex
			{
				position45 := position
				{
					position46, tokenIndex46 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l47
					}
					position++
					goto l46
				l47:
					position, tokenIndex = position46, tokenIndex46
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l44
					}
					position++
				}
			l46:
			l48:
				{
					position49, tokenIndex49 := position, tokenIndex
					{
						position50, tokenIndex50 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l51
						}
						position++
						goto l50
					l51:
						position, tokenIndex = position50, tokenIndex50
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l52
						}
						position++
						goto l50
					l52:
						position, tokenIndex = position50, tokenIndex50
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l49
						}
						position++
					}
				l50:
					goto l48
				l49:
					position, tokenIndex = position49, tokenIndex49
				}
				if !_rules[rule_]() {
					goto l44
				}
				add(ruleident, position45)
			}
			return true
		l44:
			position, tokenIndex = position44, tokenIndex44
			return false
		},
		/* 13 number <- <([0-9]+ _)> */
		nil,
		/* 14 __ <- <(!([a-z] / [A-Z] / [0-9] / '_') _)> */
		nil,
		/* 15 _ <- <(' ' / '\t')*> */
		func() bool {
			{
				position56 := position
			l57:
				{
					position58, tokenIndex58 := position, tokenIndex
					{
						position59, tokenIndex59 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l60
						}
						position++
						goto l59
					l60:
						position, tokenIndex = position59, tokenIndex59
						if buffer[position] != rune('\t') {
							goto l58
						}
						position++
					}
				l59:
					goto l57
				l58:
					position, tokenIndex = position58, tokenIndex58
				}
				add(rule_, position56)
			}
			return true
		},
		/* 16 EOL <- <(('\r' '\n') / '\n')> */
		func() bool {
			position61, tokenIndex61 := position, tokenIndex
			{
				position62 := position
				{
					position63, tokenIndex63 := position, tokenIndex
					if buffer[position] != rune('\r') {
						goto l64
					}
					position++
					if buffer[position] != rune('\n') {
						goto l64
					}
					position++
					goto l63
				l64:
					position, tokenIndex = position63, tokenIndex63
					if buffer[position] != rune('\n') {
						goto l61
					}
					position++
				}
			l63:
				add(ruleEOL, position62)
			}
			return true
		l61:
			position, tokenIndex = position61, tokenIndex61
			return false
		},
		/* 17 EOT <- <!.> */
		func() bool {
			position65, tokenIndex65 := position, tokenIndex
			{
				position66 := position
				{
					position67, tokenIndex67 := position, tokenIndex
					if !matchDot() {
						goto l67
					}
					goto l65
				l67:
					position, tokenIndex = position67, tokenIndex67
				}
				add(ruleEOT, position66)
			}
			return true
		l65:
			position, tokenIndex = position65, tokenIndex65
			return false
		},
		nil,
		/* 20 Action0 <- <{p.s.Err(begin)}> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
		/* 21 Action1 <- <{p.s.Err(begin)}> */
		func() bool {
			{
				add(ruleAction1, position)
			}
			return true
		},
		/* 22 Action2 <- <{
			p.s.line++;
			p.s.lineHead = begin + 1
		}> */
		func() bool {
			{
				add(ruleAction2, position)
			}
			return true
		},
		/* 23 Action3 <- <{
			p.s.line++;
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d COMMENT\t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction3, position)
			}
			return true
		},
		/* 24 Action4 <- <{
			p.s.line++;
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d BLANK  \t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction4, position)
			}
			return true
		},
		/* 25 Action5 <- <{
			p.s.line++;
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d LABEL  \t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction5, position)
			}
			return true
		},
	}
	p.rules = _rules
}
