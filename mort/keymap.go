package mort

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type KeyMapper interface {
	LoadMap(file_name string, split int)
	GetDescription() string
	Size() int
	Print()
}

type KeyMap struct {
	title  string
	keymap map[string]string
	split  int
}

func NewKeyMap(filename string, split int) *KeyMap {
	m := new(KeyMap)
	m.LoadMap(filename, split)
	return m
}

func (m *KeyMap) LoadMap(filename string, split int) {
	fmt.Println("LoadMap() file_name=", filename, "split=", split)
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m.keymap = make(map[string]string)
	m.title = filepath.Base(filename)
	m.split = split
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		m.keymap[strings.Trim(line[:split-1], " ")] = strings.Trim(line[split:], " ")
		fmt.Print(("."))
	}
	fmt.Println()
	fmt.Println(len(m.keymap), "lines read.")
}

func (m KeyMap) GetDescription() string {
	return m.title
}

func (m KeyMap) Size() int {
	return len(m.keymap)
}

func (m KeyMap) Print() {
	fmt.Println(m.title)
	fmt.Println("size:", m.Size())
	keys := []string{}
	for k := range m.keymap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmtstr := fmt.Sprintf("%%-%vv %%v\n", m.split)
	fmt.Println("fmtstr:", fmtstr)
	for _, k := range keys {
		fmt.Printf(fmtstr, k, m.keymap[k])
	}
}
