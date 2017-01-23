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
	rulenim_only
	rulenim_params
	rulenim_operand
	rulecomments
	ruleblanks
	rulelabel
	ruledst
	rulesrc
	rulememory
	ruleimmediate
	rulequoted
	rulecomment
	ruleexpression
	rulesign
	ruleterm_op
	ruleterm
	rulefactor_op
	rulefactor
	ruleident
	rulehex
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
	ruleAction6
	ruleAction7
	ruleAction8
)

var rul3s = [...]string{
	"Unknown",
	"root",
	"text",
	"line",
	"nim_only",
	"nim_params",
	"nim_operand",
	"comments",
	"blanks",
	"label",
	"dst",
	"src",
	"memory",
	"immediate",
	"quoted",
	"comment",
	"expression",
	"sign",
	"term_op",
	"term",
	"factor_op",
	"factor",
	"ident",
	"hex",
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
	"Action6",
	"Action7",
	"Action8",
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
	rules  [39]func() bool
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
			fmt.Printf("line %04d OTHERS \t%s", p.s.line, text)

		case ruleAction3:

			p.s.line++
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d NIMONIC\t%s\n", p.s.line, text)

		case ruleAction4:

			p.s.line++
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d NIMONIC\t%s\n", p.s.line, text)

		case ruleAction5:

			p.s.line++
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d NIM & OP\t%s\n", p.s.line, text)

		case ruleAction6:

			p.s.line++
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d COMMENT\t%s\n", p.s.line, text)

		case ruleAction7:

			p.s.line++
			p.s.lineHead = begin + 1
			fmt.Printf("line %04d BLANK  \t%s\n", p.s.line, text)

		case ruleAction8:

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
		/* 2 line <- <((comments EOL) / (blanks EOL) / (label EOL) / (nim_only EOL) / (nim_params EOL) / (nim_operand EOL) / (<((!'\n' .)+ _ EOL)> Action2))> */
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
					if !_rules[rulenim_only]() {
						goto l21
					}
					if !_rules[ruleEOL]() {
						goto l21
					}
					goto l17
				l21:
					position, tokenIndex = position17, tokenIndex17
					if !_rules[rulenim_params]() {
						goto l22
					}
					if !_rules[ruleEOL]() {
						goto l22
					}
					goto l17
				l22:
					position, tokenIndex = position17, tokenIndex17
					if !_rules[rulenim_operand]() {
						goto l23
					}
					if !_rules[ruleEOL]() {
						goto l23
					}
					goto l17
				l23:
					position, tokenIndex = position17, tokenIndex17
					{
						position24 := position
						{
							position27, tokenIndex27 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l27
							}
							position++
							goto l15
						l27:
							position, tokenIndex = position27, tokenIndex27
						}
						if !matchDot() {
							goto l15
						}
					l25:
						{
							position26, tokenIndex26 := position, tokenIndex
							{
								position28, tokenIndex28 := position, tokenIndex
								if buffer[position] != rune('\n') {
									goto l28
								}
								position++
								goto l26
							l28:
								position, tokenIndex = position28, tokenIndex28
							}
							if !matchDot() {
								goto l26
							}
							goto l25
						l26:
							position, tokenIndex = position26, tokenIndex26
						}
						if !_rules[rule_]() {
							goto l15
						}
						if !_rules[ruleEOL]() {
							goto l15
						}
						add(rulePegText, position24)
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
		/* 3 nim_only <- <(<(_ ident _ comment)> Action3)> */
		func() bool {
			position29, tokenIndex29 := position, tokenIndex
			{
				position30 := position
				{
					position31 := position
					if !_rules[rule_]() {
						goto l29
					}
					if !_rules[ruleident]() {
						goto l29
					}
					if !_rules[rule_]() {
						goto l29
					}
					if !_rules[rulecomment]() {
						goto l29
					}
					add(rulePegText, position31)
				}
				if !_rules[ruleAction3]() {
					goto l29
				}
				add(rulenim_only, position30)
			}
			return true
		l29:
			position, tokenIndex = position29, tokenIndex29
			return false
		},
		/* 4 nim_params <- <(<(_ ident _ src _ comment)> Action4)> */
		func() bool {
			position32, tokenIndex32 := position, tokenIndex
			{
				position33 := position
				{
					position34 := position
					if !_rules[rule_]() {
						goto l32
					}
					if !_rules[ruleident]() {
						goto l32
					}
					if !_rules[rule_]() {
						goto l32
					}
					if !_rules[rulesrc]() {
						goto l32
					}
					if !_rules[rule_]() {
						goto l32
					}
					if !_rules[rulecomment]() {
						goto l32
					}
					add(rulePegText, position34)
				}
				if !_rules[ruleAction4]() {
					goto l32
				}
				add(rulenim_params, position33)
			}
			return true
		l32:
			position, tokenIndex = position32, tokenIndex32
			return false
		},
		/* 5 nim_operand <- <(<(_ ident _ dst _ ',' _ src _ comment)> Action5)> */
		func() bool {
			position35, tokenIndex35 := position, tokenIndex
			{
				position36 := position
				{
					position37 := position
					if !_rules[rule_]() {
						goto l35
					}
					if !_rules[ruleident]() {
						goto l35
					}
					if !_rules[rule_]() {
						goto l35
					}
					if !_rules[ruledst]() {
						goto l35
					}
					if !_rules[rule_]() {
						goto l35
					}
					if buffer[position] != rune(',') {
						goto l35
					}
					position++
					if !_rules[rule_]() {
						goto l35
					}
					if !_rules[rulesrc]() {
						goto l35
					}
					if !_rules[rule_]() {
						goto l35
					}
					if !_rules[rulecomment]() {
						goto l35
					}
					add(rulePegText, position37)
				}
				if !_rules[ruleAction5]() {
					goto l35
				}
				add(rulenim_operand, position36)
			}
			return true
		l35:
			position, tokenIndex = position35, tokenIndex35
			return false
		},
		/* 6 comments <- <(<(_ ';' (!'\n' .)* _)> Action6)> */
		func() bool {
			position38, tokenIndex38 := position, tokenIndex
			{
				position39 := position
				{
					position40 := position
					if !_rules[rule_]() {
						goto l38
					}
					if buffer[position] != rune(';') {
						goto l38
					}
					position++
				l41:
					{
						position42, tokenIndex42 := position, tokenIndex
						{
							position43, tokenIndex43 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l43
							}
							position++
							goto l42
						l43:
							position, tokenIndex = position43, tokenIndex43
						}
						if !matchDot() {
							goto l42
						}
						goto l41
					l42:
						position, tokenIndex = position42, tokenIndex42
					}
					if !_rules[rule_]() {
						goto l38
					}
					add(rulePegText, position40)
				}
				if !_rules[ruleAction6]() {
					goto l38
				}
				add(rulecomments, position39)
			}
			return true
		l38:
			position, tokenIndex = position38, tokenIndex38
			return false
		},
		/* 7 blanks <- <(<_> Action7)> */
		func() bool {
			position44, tokenIndex44 := position, tokenIndex
			{
				position45 := position
				{
					position46 := position
					if !_rules[rule_]() {
						goto l44
					}
					add(rulePegText, position46)
				}
				if !_rules[ruleAction7]() {
					goto l44
				}
				add(ruleblanks, position45)
			}
			return true
		l44:
			position, tokenIndex = position44, tokenIndex44
			return false
		},
		/* 8 label <- <(<(ident ':' _)> Action8)> */
		func() bool {
			position47, tokenIndex47 := position, tokenIndex
			{
				position48 := position
				{
					position49 := position
					if !_rules[ruleident]() {
						goto l47
					}
					if buffer[position] != rune(':') {
						goto l47
					}
					position++
					if !_rules[rule_]() {
						goto l47
					}
					add(rulePegText, position49)
				}
				if !_rules[ruleAction8]() {
					goto l47
				}
				add(rulelabel, position48)
			}
			return true
		l47:
			position, tokenIndex = position47, tokenIndex47
			return false
		},
		/* 9 dst <- <(ident / immediate / memory / quoted)> */
		func() bool {
			position50, tokenIndex50 := position, tokenIndex
			{
				position51 := position
				{
					position52, tokenIndex52 := position, tokenIndex
					if !_rules[ruleident]() {
						goto l53
					}
					goto l52
				l53:
					position, tokenIndex = position52, tokenIndex52
					if !_rules[ruleimmediate]() {
						goto l54
					}
					goto l52
				l54:
					position, tokenIndex = position52, tokenIndex52
					if !_rules[rulememory]() {
						goto l55
					}
					goto l52
				l55:
					position, tokenIndex = position52, tokenIndex52
					if !_rules[rulequoted]() {
						goto l50
					}
				}
			l52:
				add(ruledst, position51)
			}
			return true
		l50:
			position, tokenIndex = position50, tokenIndex50
			return false
		},
		/* 10 src <- <(ident / immediate / memory / quoted)> */
		func() bool {
			position56, tokenIndex56 := position, tokenIndex
			{
				position57 := position
				{
					position58, tokenIndex58 := position, tokenIndex
					if !_rules[ruleident]() {
						goto l59
					}
					goto l58
				l59:
					position, tokenIndex = position58, tokenIndex58
					if !_rules[ruleimmediate]() {
						goto l60
					}
					goto l58
				l60:
					position, tokenIndex = position58, tokenIndex58
					if !_rules[rulememory]() {
						goto l61
					}
					goto l58
				l61:
					position, tokenIndex = position58, tokenIndex58
					if !_rules[rulequoted]() {
						goto l56
					}
				}
			l58:
				add(rulesrc, position57)
			}
			return true
		l56:
			position, tokenIndex = position56, tokenIndex56
			return false
		},
		/* 11 memory <- <('[' (ident / hex) ']')> */
		func() bool {
			position62, tokenIndex62 := position, tokenIndex
			{
				position63 := position
				if buffer[position] != rune('[') {
					goto l62
				}
				position++
				{
					position64, tokenIndex64 := position, tokenIndex
					if !_rules[ruleident]() {
						goto l65
					}
					goto l64
				l65:
					position, tokenIndex = position64, tokenIndex64
					if !_rules[rulehex]() {
						goto l62
					}
				}
			l64:
				if buffer[position] != rune(']') {
					goto l62
				}
				position++
				add(rulememory, position63)
			}
			return true
		l62:
			position, tokenIndex = position62, tokenIndex62
			return false
		},
		/* 12 immediate <- <(hex / number / expression)> */
		func() bool {
			position66, tokenIndex66 := position, tokenIndex
			{
				position67 := position
				{
					position68, tokenIndex68 := position, tokenIndex
					if !_rules[rulehex]() {
						goto l69
					}
					goto l68
				l69:
					position, tokenIndex = position68, tokenIndex68
					if !_rules[rulenumber]() {
						goto l70
					}
					goto l68
				l70:
					position, tokenIndex = position68, tokenIndex68
					if !_rules[ruleexpression]() {
						goto l66
					}
				}
			l68:
				add(ruleimmediate, position67)
			}
			return true
		l66:
			position, tokenIndex = position66, tokenIndex66
			return false
		},
		/* 13 quoted <- <('"' (!'"' .)+ '"')> */
		func() bool {
			position71, tokenIndex71 := position, tokenIndex
			{
				position72 := position
				if buffer[position] != rune('"') {
					goto l71
				}
				position++
				{
					position75, tokenIndex75 := position, tokenIndex
					if buffer[position] != rune('"') {
						goto l75
					}
					position++
					goto l71
				l75:
					position, tokenIndex = position75, tokenIndex75
				}
				if !matchDot() {
					goto l71
				}
			l73:
				{
					position74, tokenIndex74 := position, tokenIndex
					{
						position76, tokenIndex76 := position, tokenIndex
						if buffer[position] != rune('"') {
							goto l76
						}
						position++
						goto l74
					l76:
						position, tokenIndex = position76, tokenIndex76
					}
					if !matchDot() {
						goto l74
					}
					goto l73
				l74:
					position, tokenIndex = position74, tokenIndex74
				}
				if buffer[position] != rune('"') {
					goto l71
				}
				position++
				add(rulequoted, position72)
			}
			return true
		l71:
			position, tokenIndex = position71, tokenIndex71
			return false
		},
		/* 14 comment <- <(';' (!'\n' .)*)?> */
		func() bool {
			{
				position78 := position
				{
					position79, tokenIndex79 := position, tokenIndex
					if buffer[position] != rune(';') {
						goto l79
					}
					position++
				l81:
					{
						position82, tokenIndex82 := position, tokenIndex
						{
							position83, tokenIndex83 := position, tokenIndex
							if buffer[position] != rune('\n') {
								goto l83
							}
							position++
							goto l82
						l83:
							position, tokenIndex = position83, tokenIndex83
						}
						if !matchDot() {
							goto l82
						}
						goto l81
					l82:
						position, tokenIndex = position82, tokenIndex82
					}
					goto l80
				l79:
					position, tokenIndex = position79, tokenIndex79
				}
			l80:
				add(rulecomment, position78)
			}
			return true
		},
		/* 15 expression <- <(sign term (term_op term)*)> */
		func() bool {
			position84, tokenIndex84 := position, tokenIndex
			{
				position85 := position
				if !_rules[rulesign]() {
					goto l84
				}
				if !_rules[ruleterm]() {
					goto l84
				}
			l86:
				{
					position87, tokenIndex87 := position, tokenIndex
					if !_rules[ruleterm_op]() {
						goto l87
					}
					if !_rules[ruleterm]() {
						goto l87
					}
					goto l86
				l87:
					position, tokenIndex = position87, tokenIndex87
				}
				add(ruleexpression, position85)
			}
			return true
		l84:
			position, tokenIndex = position84, tokenIndex84
			return false
		},
		/* 16 sign <- <(('-' / '+')? _)> */
		func() bool {
			position88, tokenIndex88 := position, tokenIndex
			{
				position89 := position
				{
					position90, tokenIndex90 := position, tokenIndex
					{
						position92, tokenIndex92 := position, tokenIndex
						if buffer[position] != rune('-') {
							goto l93
						}
						position++
						goto l92
					l93:
						position, tokenIndex = position92, tokenIndex92
						if buffer[position] != rune('+') {
							goto l90
						}
						position++
					}
				l92:
					goto l91
				l90:
					position, tokenIndex = position90, tokenIndex90
				}
			l91:
				if !_rules[rule_]() {
					goto l88
				}
				add(rulesign, position89)
			}
			return true
		l88:
			position, tokenIndex = position88, tokenIndex88
			return false
		},
		/* 17 term_op <- <(('-' / '+') _)> */
		func() bool {
			position94, tokenIndex94 := position, tokenIndex
			{
				position95 := position
				{
					position96, tokenIndex96 := position, tokenIndex
					if buffer[position] != rune('-') {
						goto l97
					}
					position++
					goto l96
				l97:
					position, tokenIndex = position96, tokenIndex96
					if buffer[position] != rune('+') {
						goto l94
					}
					position++
				}
			l96:
				if !_rules[rule_]() {
					goto l94
				}
				add(ruleterm_op, position95)
			}
			return true
		l94:
			position, tokenIndex = position94, tokenIndex94
			return false
		},
		/* 18 term <- <(factor (factor_op factor)*)> */
		func() bool {
			position98, tokenIndex98 := position, tokenIndex
			{
				position99 := position
				if !_rules[rulefactor]() {
					goto l98
				}
			l100:
				{
					position101, tokenIndex101 := position, tokenIndex
					if !_rules[rulefactor_op]() {
						goto l101
					}
					if !_rules[rulefactor]() {
						goto l101
					}
					goto l100
				l101:
					position, tokenIndex = position101, tokenIndex101
				}
				add(ruleterm, position99)
			}
			return true
		l98:
			position, tokenIndex = position98, tokenIndex98
			return false
		},
		/* 19 factor_op <- <(('*' / '/') _)> */
		func() bool {
			position102, tokenIndex102 := position, tokenIndex
			{
				position103 := position
				{
					position104, tokenIndex104 := position, tokenIndex
					if buffer[position] != rune('*') {
						goto l105
					}
					position++
					goto l104
				l105:
					position, tokenIndex = position104, tokenIndex104
					if buffer[position] != rune('/') {
						goto l102
					}
					position++
				}
			l104:
				if !_rules[rule_]() {
					goto l102
				}
				add(rulefactor_op, position103)
			}
			return true
		l102:
			position, tokenIndex = position102, tokenIndex102
			return false
		},
		/* 20 factor <- <(ident / number / ('(' _ expression ')' _))> */
		func() bool {
			position106, tokenIndex106 := position, tokenIndex
			{
				position107 := position
				{
					position108, tokenIndex108 := position, tokenIndex
					if !_rules[ruleident]() {
						goto l109
					}
					goto l108
				l109:
					position, tokenIndex = position108, tokenIndex108
					if !_rules[rulenumber]() {
						goto l110
					}
					goto l108
				l110:
					position, tokenIndex = position108, tokenIndex108
					if buffer[position] != rune('(') {
						goto l106
					}
					position++
					if !_rules[rule_]() {
						goto l106
					}
					if !_rules[ruleexpression]() {
						goto l106
					}
					if buffer[position] != rune(')') {
						goto l106
					}
					position++
					if !_rules[rule_]() {
						goto l106
					}
				}
			l108:
				add(rulefactor, position107)
			}
			return true
		l106:
			position, tokenIndex = position106, tokenIndex106
			return false
		},
		/* 21 ident <- <(([a-z] / [A-Z]) ([a-z] / [A-Z] / [0-9])* _)> */
		func() bool {
			position111, tokenIndex111 := position, tokenIndex
			{
				position112 := position
				{
					position113, tokenIndex113 := position, tokenIndex
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l114
					}
					position++
					goto l113
				l114:
					position, tokenIndex = position113, tokenIndex113
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l111
					}
					position++
				}
			l113:
			l115:
				{
					position116, tokenIndex116 := position, tokenIndex
					{
						position117, tokenIndex117 := position, tokenIndex
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l118
						}
						position++
						goto l117
					l118:
						position, tokenIndex = position117, tokenIndex117
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l119
						}
						position++
						goto l117
					l119:
						position, tokenIndex = position117, tokenIndex117
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l116
						}
						position++
					}
				l117:
					goto l115
				l116:
					position, tokenIndex = position116, tokenIndex116
				}
				if !_rules[rule_]() {
					goto l111
				}
				add(ruleident, position112)
			}
			return true
		l111:
			position, tokenIndex = position111, tokenIndex111
			return false
		},
		/* 22 hex <- <('0' ('x' / 'X') ([0-9] / [a-z] / [A-Z])+ _)> */
		func() bool {
			position120, tokenIndex120 := position, tokenIndex
			{
				position121 := position
				if buffer[position] != rune('0') {
					goto l120
				}
				position++
				{
					position122, tokenIndex122 := position, tokenIndex
					if buffer[position] != rune('x') {
						goto l123
					}
					position++
					goto l122
				l123:
					position, tokenIndex = position122, tokenIndex122
					if buffer[position] != rune('X') {
						goto l120
					}
					position++
				}
			l122:
				{
					position126, tokenIndex126 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l127
					}
					position++
					goto l126
				l127:
					position, tokenIndex = position126, tokenIndex126
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l128
					}
					position++
					goto l126
				l128:
					position, tokenIndex = position126, tokenIndex126
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l120
					}
					position++
				}
			l126:
			l124:
				{
					position125, tokenIndex125 := position, tokenIndex
					{
						position129, tokenIndex129 := position, tokenIndex
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l130
						}
						position++
						goto l129
					l130:
						position, tokenIndex = position129, tokenIndex129
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l131
						}
						position++
						goto l129
					l131:
						position, tokenIndex = position129, tokenIndex129
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l125
						}
						position++
					}
				l129:
					goto l124
				l125:
					position, tokenIndex = position125, tokenIndex125
				}
				if !_rules[rule_]() {
					goto l120
				}
				add(rulehex, position121)
			}
			return true
		l120:
			position, tokenIndex = position120, tokenIndex120
			return false
		},
		/* 23 number <- <([0-9]+ _)> */
		func() bool {
			position132, tokenIndex132 := position, tokenIndex
			{
				position133 := position
				if c := buffer[position]; c < rune('0') || c > rune('9') {
					goto l132
				}
				position++
			l134:
				{
					position135, tokenIndex135 := position, tokenIndex
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l135
					}
					position++
					goto l134
				l135:
					position, tokenIndex = position135, tokenIndex135
				}
				if !_rules[rule_]() {
					goto l132
				}
				add(rulenumber, position133)
			}
			return true
		l132:
			position, tokenIndex = position132, tokenIndex132
			return false
		},
		/* 24 __ <- <(!([a-z] / [A-Z] / [0-9] / '_') _)> */
		nil,
		/* 25 _ <- <(' ' / '\t')*> */
		func() bool {
			{
				position138 := position
			l139:
				{
					position140, tokenIndex140 := position, tokenIndex
					{
						position141, tokenIndex141 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l142
						}
						position++
						goto l141
					l142:
						position, tokenIndex = position141, tokenIndex141
						if buffer[position] != rune('\t') {
							goto l140
						}
						position++
					}
				l141:
					goto l139
				l140:
					position, tokenIndex = position140, tokenIndex140
				}
				add(rule_, position138)
			}
			return true
		},
		/* 26 EOL <- <(('\r' '\n') / '\n')> */
		func() bool {
			position143, tokenIndex143 := position, tokenIndex
			{
				position144 := position
				{
					position145, tokenIndex145 := position, tokenIndex
					if buffer[position] != rune('\r') {
						goto l146
					}
					position++
					if buffer[position] != rune('\n') {
						goto l146
					}
					position++
					goto l145
				l146:
					position, tokenIndex = position145, tokenIndex145
					if buffer[position] != rune('\n') {
						goto l143
					}
					position++
				}
			l145:
				add(ruleEOL, position144)
			}
			return true
		l143:
			position, tokenIndex = position143, tokenIndex143
			return false
		},
		/* 27 EOT <- <!.> */
		func() bool {
			position147, tokenIndex147 := position, tokenIndex
			{
				position148 := position
				{
					position149, tokenIndex149 := position, tokenIndex
					if !matchDot() {
						goto l149
					}
					goto l147
				l149:
					position, tokenIndex = position149, tokenIndex149
				}
				add(ruleEOT, position148)
			}
			return true
		l147:
			position, tokenIndex = position147, tokenIndex147
			return false
		},
		nil,
		/* 30 Action0 <- <{p.s.Err(begin)}> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
		/* 31 Action1 <- <{p.s.Err(begin)}> */
		func() bool {
			{
				add(ruleAction1, position)
			}
			return true
		},
		/* 32 Action2 <- <{
		 p.s.line++;
		 p.s.lineHead = begin + 1
		 fmt.Printf("line %04d OTHERS \t%s", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction2, position)
			}
			return true
		},
		/* 33 Action3 <- <{
		 p.s.line++;
		 p.s.lineHead = begin + 1
		 fmt.Printf("line %04d NIMONIC\t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction3, position)
			}
			return true
		},
		/* 34 Action4 <- <{
		 p.s.line++;
		 p.s.lineHead = begin + 1
		 fmt.Printf("line %04d NIMONIC\t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction4, position)
			}
			return true
		},
		/* 35 Action5 <- <{
		 p.s.line++;
		 p.s.lineHead = begin + 1
		 fmt.Printf("line %04d NIM & OP\t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction5, position)
			}
			return true
		},
		/* 36 Action6 <- <{
		 p.s.line++;
		 p.s.lineHead = begin + 1
		 fmt.Printf("line %04d COMMENT\t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction6, position)
			}
			return true
		},
		/* 37 Action7 <- <{
		 p.s.line++;
		 p.s.lineHead = begin + 1
		 fmt.Printf("line %04d BLANK  \t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction7, position)
			}
			return true
		},
		/* 38 Action8 <- <{
		 p.s.line++;
		 p.s.lineHead = begin + 1
		 fmt.Printf("line %04d LABEL  \t%s\n", p.s.line, text)
		}> */
		func() bool {
			{
				add(ruleAction8, position)
			}
			return true
		},
	}
	p.rules = _rules
}
