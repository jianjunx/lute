// Lute - A structured markdown engine.
// Copyright (C) 2019-present, b3log.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lute

type ListType int

type List struct {
	*BaseNode

	Bullet bool
	Start  int
	Tight  bool

	IndentSpaces int
	Marker       string
	WNSpaces     int
}

func newList(indentSpaces int, marker string, wnSpaces int, t *Tree) (ret Node) {
	baseNode := &BaseNode{typ: NodeList}
	ret = &List{
		baseNode,
		true,
		1,
		false,
		indentSpaces,
		marker,
		wnSpaces,
	}
	t.context.CurNode = ret

	return
}

func (t *Tree) parseList(line items) (ret Node) {
	spaces, tabs, tokens, firstNonWhitespace := t.nonWhitespace(line)
	marker := firstNonWhitespace
	indentSpaces := spaces + tabs*4
	line = line[len(tokens):]
	spaces, tabs, _, firstNonWhitespace = t.nonWhitespace(line)
	w := len(marker.val)
	n := spaces + tabs*4
	wnSpaces := w + n
	oldContextIndentSpaces := t.context.IndentSpaces
	t.context.IndentSpaces = indentSpaces + wnSpaces
	if 4 <= n {
		line = indentOffset(line, w+1, t)
		t.context.IndentSpaces = 2
	} else {
		line = indentOffset(line, indentSpaces+wnSpaces, t)
	}
	ret = newList(indentSpaces, marker.val, wnSpaces, t)
	tight := false

	defer func() { t.context.IndentSpaces = oldContextIndentSpaces }()

	for {
		n := t.parseListItem(line)
		if nil == n {
			break
		}
		ret.AppendChild(ret, n)

		if n.(*ListItem).Tight {
			tight = true
		}

		line = t.nextLine()
		if line.isEOF() {
			break
		}

		if t.isThematicBreak(line) {
			t.backupLine(line)
			break
		}

		if marker.val != line[0].val {
			// TODO: 考虑有序列表序号递增
			t.backupLine(line)

			break
		} else {
			line = line[len(marker.val):]
		}
	}

	ret.(*List).Tight = tight
	//for child := ret.FirstChild();nil != child;child = child.Next() {
	//	child.(*ListItem).Tight = tight
	//}

	return
}

func (t *Tree) isList(line items) bool {
	if 2 > len(line) { // at least marker and newline
		return false
	}

	line = line.trimLeft()
	firstNonWhitespace := line[0]
	if itemAsterisk == firstNonWhitespace.typ || itemHyphen == firstNonWhitespace.typ || itemPlus == firstNonWhitespace.typ {
		return line[1].isWhitespace()
	} else if firstNonWhitespace.isNumInt() && 9 >= len(firstNonWhitespace.val) {
		return (itemDot == line[1].typ || itemCloseParen == line[1].typ) && line[2].isWhitespace()
	}

	return false
}
