package fabslang

import (
	"bufio"
	"io"
	"strconv"
)

type FabsLang struct {
	code         string
	compiledCode string
	codePointer  int
	oldPoints    []int

	pointer int
	memory  map[int]byte

	input  io.Reader
	output io.Writer
}

func New(code string, input io.Reader, output io.Writer) *FabsLang {
	return &FabsLang{code: code, memory: make(map[int]byte), oldPoints: make([]int, 0), input: input, output: output}
}

func (b *FabsLang) Run() {
	b.compile()
	for b.codePointer < len(b.compiledCode) {
		b.interpret()
		b.codePointer++
	}
}

func (b *FabsLang) applyCount(count, character int) {
	if count == 0 {
		panic("syntax error at character " + strconv.Itoa(character))
	}
	b.compiledCode += strconv.Itoa(count)
}

func (b *FabsLang) compile() {
	count := 0
	for i := 0; i < len(b.code); i++ {
		switch {
		case count == 0 && b.code[i] == 'f':
			count++
		case count == 1 && b.code[i] == 'a':
			count++
		case count == 2 && b.code[i] == 'b':
			count++
		case count == 3 && b.code[i] == 's':
			count++
		case count == 4 && b.code[i] == 'l':
			count++
		case count == 5 && b.code[i] == 'a':
			count++
		case count == 6 && b.code[i] == 'n':
			count++
		case count == 7 && b.code[i] == 'g':
			count++
		default:
			b.applyCount(count, i)
			count = 0
			i--
		}
	}
	if count > 0 {
		b.applyCount(count, 0)
	}
}

func (b *FabsLang) interpret() {
	switch b.getCommand() {
	case '1':
		byteList := []byte{b.getCurrentByte()}
		b.output.Write(byteList)
	case '2':
		b.setCurrentByte(b.getCurrentByte() + 1)
	case '3':
		b.setCurrentByte(b.getCurrentByte() - 1)
	case '4':
		b.pointer++
	case '5':
		b.pointer--
	case '6':
		reader := bufio.NewReader(b.input)
		input, _ := reader.ReadByte()
		b.setCurrentByte(input)
	case '7':
		for b.getCurrentByte() != 0 {
			codePointer := b.codePointer
			for b.codePointer++; b.getCommand() != '8'; b.codePointer++ {
				b.interpret()
			}
			b.codePointer = codePointer
		}

		numberOfJumps := 1
		for numberOfJumps > 0 {
			b.codePointer++
			switch b.getCommand() {
			case '7':
				numberOfJumps++
			case '8':
				numberOfJumps--
			}
		}
	}
}

func (b *FabsLang) getCommand() rune {
	return rune(b.compiledCode[b.codePointer])
}

func (b *FabsLang) getCurrentByte() byte {
	currentByte, _ := b.memory[b.pointer]
	return currentByte
}

func (b *FabsLang) setCurrentByte(val byte) {
	b.memory[b.pointer] = val
}
