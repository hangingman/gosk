package eval

import (
	"fmt"
	"log"
)

type LabelManagement struct {
}

func (l *LabelManagement) AddLabelCallback(ident string) {
	log.Println(fmt.Sprintf("info: add label %s !!", ident))
}

func (l *LabelManagement) RemoveLabelCallback(ident string) {
	log.Println(fmt.Sprintf("info: remove label %s !!", ident))
}

func (l *LabelManagement) Emit(ident string) {
	log.Println(fmt.Sprintf("info: emit label %s !!", ident))
}
