package keymap

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

type MapHolder struct {
	title  string
	keymap map[string]string
	split  int
}

func NewMap(filename string, split int) *MapHolder {
	m := new(MapHolder)
	m.LoadMap(filename, split)
	return m
}

func (m *MapHolder) LoadMap(filename string, split int) {
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

func (m MapHolder) GetDescription() string {
	return m.title
}

func (m MapHolder) Size() int {
	return len(m.keymap)
}

func (m MapHolder) Print() {
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
