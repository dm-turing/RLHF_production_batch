package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
)

type Composite interface {
	Add(Composite)
	Children() []Composite
	String() string
}

type Leaf struct {
	XMLName xml.Name `xml:"leaf"`
	Value   string   `xml:"value"`
}

func (l Leaf) String() string {
	return l.Value
}

type Branch struct {
	XMLName  xml.Name    `xml:"branch"`
	Children []Composite `xml:"any"`
}

func (b Branch) Add(c Composite) {
	b.Children = append(b.Children, c)
}

func (b Branch) Children() []Composite {
	return b.Children
}

func (b Branch) String() string {
	var result []string
	for _, child := range b.Children {
		result = append(result, child.String())
	}
	return "Branch{" + fmt.Sprint(result) + "}"
}

func parseXML(r io.Reader) (Composite, error) {
	var root Composite
	d := xml.NewDecoder(r)

	for {
		t, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		switch se := t.(type) {
		case xml.StartElement:
			switch se.Name.Local {
			case "leaf":
				var l Leaf
				if err := d.Decode(&l); err != nil {
					return nil, err
				}
				root = l
			case "branch":
				var b Branch
				if err := d.Decode(&b); err != nil {
					return nil, err
				}
				root = b
			}
		}
	}

	return root, nil
}

func main() {
	xmlData := []byte(`
    <branch>
        <leaf value="a"/>
        <branch>
            <leaf value="b"/>
            <leaf value="c"/>
        </branch>
        <leaf value="d"/>
    </branch>
    `)

	root, err := parseXML(ioutil.NopCloser(bytes.NewReader(xmlData)))
	if err != nil {
		panic(err)
	}
	fmt.Println(root)
}
