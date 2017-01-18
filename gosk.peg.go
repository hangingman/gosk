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
	ruleEOT
	ruleexpression
	ruleliteral
	rulePegText
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	ruleAction7
)

var rul3s = [...]string{
	"Unknown",
	"root",
	"EOT",
	"expression",
	"literal",
	"PegText",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"Action7",
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
	// parserが自動生成するフィールド変数と区別するために
	// 敢えて埋め込みを行っていない。

	Buffer string
	buffer []rune
	rules  [14]func() bool
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

			// p,begin,end,text を使用する場合はruleを<…>で囲む。
			// '0'で始まる数字列
			fmt.Printf("line %d(%d) KIND:ZeroNUMBER \"%s\"\n", p.s.line, begin-p.s.lineHead, text)

		case ruleAction3:
			// 数字列
			fmt.Printf("line %d(%d) KIND:NUMBER \"%s\"\n", p.s.line, begin-p.s.lineHead, text)

		case ruleAction4:
			// 大小英字列
			fmt.Printf("line %d(%d) KIND:IDENT \"%s\"\n", p.s.line, begin-p.s.lineHead, text)

		case ruleAction5:

		case ruleAction6:
			// 改行時処理
			p.s.line++
			p.s.lineHead = begin + 1

		case ruleAction7:
			// その他文字
			fmt.Printf("line %d(%d) KIND:OTHER \"%s\"\n", p.s.line, begin-p.s.lineHead, text)

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
		/* 0 root <- <((expression EOT) / (expression <.+> Action0 EOT) / (<.+> Action1 EOT))> */
		func() bool {
			position0, tokenIndex0 := position, tokenIndex
			{
				position1 := position
				{
					position2, tokenIndex2 := position, tokenIndex
					if !_rules[ruleexpression]() {
						goto l3
					}
					if !_rules[ruleEOT]() {
						goto l3
					}
					goto l2
				l3:
					position, tokenIndex = position2, tokenIndex2
					if !_rules[ruleexpression]() {
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
		/* 1 EOT <- <!.> */
		func() bool {
			position11, tokenIndex11 := position, tokenIndex
			{
				position12 := position
				{
					position13, tokenIndex13 := position, tokenIndex
					if !matchDot() {
						goto l13
					}
					goto l11
				l13:
					position, tokenIndex = position13, tokenIndex13
				}
				add(ruleEOT, position12)
			}
			return true
		l11:
			position, tokenIndex = position11, tokenIndex11
			return false
		},
		/* 2 expression <- <(literal literal*)> */
		func() bool {
			position14, tokenIndex14 := position, tokenIndex
			{
				position15 := position
				if !_rules[ruleliteral]() {
					goto l14
				}
			l16:
				{
					position17, tokenIndex17 := position, tokenIndex
					if !_rules[ruleliteral]() {
						goto l17
					}
					goto l16
				l17:
					position, tokenIndex = position17, tokenIndex17
				}
				add(ruleexpression, position15)
			}
			return true
		l14:
			position, tokenIndex = position14, tokenIndex14
			return false
		},
		/* 3 literal <- <((<(&'0' [0-9]+)> Action2) / (<[0-9]+> Action3) / (<([a-z] / [A-Z])+> Action4) / (' '+ Action5) / (<'\n'> Action6) / (<(!([0-9] / [a-z] / [A-Z] / ' ' / '\n') .)+> Action7))> */
		func() bool {
			position18, tokenIndex18 := position, tokenIndex
			{
				position19 := position
				{
					position20, tokenIndex20 := position, tokenIndex
					{
						position22 := position
						{
							position23, tokenIndex23 := position, tokenIndex
							if buffer[position] != rune('0') {
								goto l21
							}
							position++
							position, tokenIndex = position23, tokenIndex23
						}
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l21
						}
						position++
					l24:
						{
							position25, tokenIndex25 := position, tokenIndex
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l25
							}
							position++
							goto l24
						l25:
							position, tokenIndex = position25, tokenIndex25
						}
						add(rulePegText, position22)
					}
					if !_rules[ruleAction2]() {
						goto l21
					}
					goto l20
				l21:
					position, tokenIndex = position20, tokenIndex20
					{
						position27 := position
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l26
						}
						position++
					l28:
						{
							position29, tokenIndex29 := position, tokenIndex
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l29
							}
							position++
							goto l28
						l29:
							position, tokenIndex = position29, tokenIndex29
						}
						add(rulePegText, position27)
					}
					if !_rules[ruleAction3]() {
						goto l26
					}
					goto l20
				l26:
					position, tokenIndex = position20, tokenIndex20
					{
						position31 := position
						{
							position34, tokenIndex34 := position, tokenIndex
							if c := buffer[position]; c < rune('a') || c > rune('z') {
								goto l35
							}
							position++
							goto l34
						l35:
							position, tokenIndex = position34, tokenIndex34
							if c := buffer[position]; c < rune('A') || c > rune('Z') {
								goto l30
							}
							position++
						}
					l34:
					l32:
						{
							position33, tokenIndex33 := position, tokenIndex
							{
								position36, tokenIndex36 := position, tokenIndex
								if c := buffer[position]; c < rune('a') || c > rune('z') {
									goto l37
								}
								position++
								goto l36
							l37:
								position, tokenIndex = position36, tokenIndex36
								if c := buffer[position]; c < rune('A') || c > rune('Z') {
									goto l33
								}
								position++
							}
						l36:
							goto l32
						l33:
							position, tokenIndex = position33, tokenIndex33
						}
						add(rulePegText, position31)
					}
					if !_rules[ruleAction4]() {
						goto l30
					}
					goto l20
				l30:
					position, tokenIndex = position20, tokenIndex20
					if buffer[position] != rune(' ') {
						goto l38
					}
					position++
				l39:
					{
						position40, tokenIndex40 := position, tokenIndex
						if buffer[position] != rune(' ') {
							goto l40
						}
						position++
						goto l39
					l40:
						position, tokenIndex = position40, tokenIndex40
					}
					if !_rules[ruleAction5]() {
						goto l38
					}
					goto l20
				l38:
					position, tokenIndex = position20, tokenIndex20
					{
						position42 := position
						if buffer[position] != rune('\n') {
							goto l41
						}
						position++
						add(rulePegText, position42)
					}
					if !_rules[ruleAction6]() {
						goto l41
					}
					goto l20
				l41:
					position, tokenIndex = position20, tokenIndex20
					{
						position43 := position
						{
							position46, tokenIndex46 := position, tokenIndex
							{
								position47, tokenIndex47 := position, tokenIndex
								if c := buffer[position]; c < rune('0') || c > rune('9') {
									goto l48
								}
								position++
								goto l47
							l48:
								position, tokenIndex = position47, tokenIndex47
								if c := buffer[position]; c < rune('a') || c > rune('z') {
									goto l49
								}
								position++
								goto l47
							l49:
								position, tokenIndex = position47, tokenIndex47
								if c := buffer[position]; c < rune('A') || c > rune('Z') {
									goto l50
								}
								position++
								goto l47
							l50:
								position, tokenIndex = position47, tokenIndex47
								if buffer[position] != rune(' ') {
									goto l51
								}
								position++
								goto l47
							l51:
								position, tokenIndex = position47, tokenIndex47
								if buffer[position] != rune('\n') {
									goto l46
								}
								position++
							}
						l47:
							goto l18
						l46:
							position, tokenIndex = position46, tokenIndex46
						}
						if !matchDot() {
							goto l18
						}
					l44:
						{
							position45, tokenIndex45 := position, tokenIndex
							{
								position52, tokenIndex52 := position, tokenIndex
								{
									position53, tokenIndex53 := position, tokenIndex
									if c := buffer[position]; c < rune('0') || c > rune('9') {
										goto l54
									}
									position++
									goto l53
								l54:
									position, tokenIndex = position53, tokenIndex53
									if c := buffer[position]; c < rune('a') || c > rune('z') {
										goto l55
									}
									position++
									goto l53
								l55:
									position, tokenIndex = position53, tokenIndex53
									if c := buffer[position]; c < rune('A') || c > rune('Z') {
										goto l56
									}
									position++
									goto l53
								l56:
									position, tokenIndex = position53, tokenIndex53
									if buffer[position] != rune(' ') {
										goto l57
									}
									position++
									goto l53
								l57:
									position, tokenIndex = position53, tokenIndex53
									if buffer[position] != rune('\n') {
										goto l52
									}
									position++
								}
							l53:
								goto l45
							l52:
								position, tokenIndex = position52, tokenIndex52
							}
							if !matchDot() {
								goto l45
							}
							goto l44
						l45:
							position, tokenIndex = position45, tokenIndex45
						}
						add(rulePegText, position43)
					}
					if !_rules[ruleAction7]() {
						goto l18
					}
				}
			l20:
				add(ruleliteral, position19)
			}
			return true
		l18:
			position, tokenIndex = position18, tokenIndex18
			return false
		},
		nil,
		/* 6 Action0 <- <{p.s.Err(begin)}> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
		/* 7 Action1 <- <{p.s.Err(begin)}> */
		func() bool {
			{
				add(ruleAction1, position)
			}
			return true
		},
		/* 8 Action2 <- <{
		// p,begin,end,text を使用する場合はruleを<…>で囲む。
		// '0'で始まる数字列
		fmt.Printf("line %d(%d) KIND:ZeroNUMBER \"%s\"\n", p.s.line, begin - p.s.lineHead, text)
		}> */
		func() bool {
			{
				add(ruleAction2, position)
			}
			return true
		},
		/* 9 Action3 <- <{ // 数字列
			fmt.Printf("line %d(%d) KIND:NUMBER \"%s\"\n", p.s.line, begin - p.s.lineHead, text)
		    }> */
		func() bool {
			{
				add(ruleAction3, position)
			}
			return true
		},
		/* 10 Action4 <- <{ // 大小英字列
		fmt.Printf("line %d(%d) KIND:IDENT \"%s\"\n", p.s.line, begin - p.s.lineHead, text)
		}> */
		func() bool {
			{
				add(ruleAction4, position)
			}
			return true
		},
		/* 11 Action5 <- <{  }> */
		func() bool {
			{
				add(ruleAction5, position)
			}
			return true
		},
		/* 12 Action6 <- <{ // 改行時処理
		p.s.line++;
		p.s.lineHead = begin + 1
		}> */
		func() bool {
			{
				add(ruleAction6, position)
			}
			return true
		},
		/* 13 Action7 <- <{ // その他文字
			fmt.Printf("line %d(%d) KIND:OTHER \"%s\"\n", p.s.line, begin - p.s.lineHead, text)
		}> */
		func() bool {
			{
				add(ruleAction7, position)
			}
			return true
		},
	}
	p.rules = _rules
}
