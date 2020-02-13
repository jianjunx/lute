// Lute - 一款对中文语境优化的 Markdown 引擎，支持 Go 和 JavaScript
// Copyright (c) 2019-present, b3log.org
//
// Lute is licensed under the Mulan PSL v1.
// You can use this software according to the terms and conditions of the Mulan PSL v1.
// You may obtain a copy of Mulan PSL v1 at:
//     http://license.coscl.org.cn/MulanPSL
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR
// PURPOSE.
// See the Mulan PSL v1 for more details.

package lute

import (
	"strings"
)

// EChartsJSONRenderer 描述了 JSON 渲染器。
type EChartsJSONRenderer struct {
	*BaseRenderer
}

// newEChartsJSONRenderer 创建一个 ECharts JSON 渲染器。
func (lute *Lute) newEChartsJSONRenderer(tree *Tree) Renderer {
	ret := &EChartsJSONRenderer{lute.newBaseRenderer(tree)}
	ret.rendererFuncs[NodeDocument] = ret.renderDocument
	ret.rendererFuncs[NodeParagraph] = ret.renderParagraph
	ret.rendererFuncs[NodeText] = ret.renderText
	ret.rendererFuncs[NodeCodeSpan] = ret.renderCodeSpan
	ret.rendererFuncs[NodeCodeBlock] = ret.renderCodeBlock
	ret.rendererFuncs[NodeMathBlock] = ret.renderMathBlock
	ret.rendererFuncs[NodeInlineMath] = ret.renderInlineMath
	ret.rendererFuncs[NodeEmphasis] = ret.renderEmphasis
	ret.rendererFuncs[NodeStrong] = ret.renderStrong
	ret.rendererFuncs[NodeBlockquote] = ret.renderBlockquote
	ret.rendererFuncs[NodeHeading] = ret.renderHeading
	ret.rendererFuncs[NodeList] = ret.renderList
	ret.rendererFuncs[NodeListItem] = ret.renderListItem
	ret.rendererFuncs[NodeThematicBreak] = ret.renderThematicBreak
	ret.rendererFuncs[NodeHardBreak] = ret.renderHardBreak
	ret.rendererFuncs[NodeSoftBreak] = ret.renderSoftBreak
	ret.rendererFuncs[NodeHTMLBlock] = ret.renderHTML
	ret.rendererFuncs[NodeInlineHTML] = ret.renderInlineHTML
	ret.rendererFuncs[NodeLink] = ret.renderLink
	ret.rendererFuncs[NodeImage] = ret.renderImage
	ret.rendererFuncs[NodeStrikethrough] = ret.renderStrikethrough
	ret.rendererFuncs[NodeTaskListItemMarker] = ret.renderTaskListItemMarker
	ret.rendererFuncs[NodeTable] = ret.renderTable
	ret.rendererFuncs[NodeTableHead] = ret.renderTableHead
	ret.rendererFuncs[NodeTableRow] = ret.renderTableRow
	ret.rendererFuncs[NodeTableCell] = ret.renderTableCell
	ret.rendererFuncs[NodeEmoji] = ret.renderEmoji
	ret.rendererFuncs[NodeEmojiUnicode] = ret.renderEmojiUnicode
	ret.rendererFuncs[NodeEmojiImg] = ret.renderEmojiImg
	ret.rendererFuncs[NodeEmojiAlias] = ret.renderEmojiAlias

	ret.defaultRendererFunc = ret.renderDefault
	return ret
}

func (r *EChartsJSONRenderer) renderDefault(n *Node, entering bool) (WalkStatus, error) {
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderInlineMath(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.leaf("Inline Math\nspan", node)
	}
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderMathBlock(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.leaf("Math Block\ndiv", node)
	}
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderEmojiImg(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("Emoji Img\n", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderEmojiUnicode(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("Emoji Unicode\n", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderEmojiAlias(node *Node, entering bool) (WalkStatus, error) {
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderEmoji(node *Node, entering bool) (WalkStatus, error) {
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderTableCell(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("Table Cell\ntd", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderTableRow(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("Table Row\ntr", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderTableHead(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("Table Head\nthead", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderTable(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("Table\ntable", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderStrikethrough(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("Strikethrough\ndel", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderImage(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		r.val("Image\nimg", node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderLink(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		r.val("Link\na", node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderHTML(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("HTML Block\n", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderInlineHTML(node *Node, entering bool) (WalkStatus, error) {
	r.leaf("Inline HTML\n", node)
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderDocument(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.writeByte(itemOpenBracket)
		r.openObj()
		r.val("Document", node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
		r.writeByte(itemCloseBracket)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderParagraph(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		r.val("Paragraph\np", node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderText(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		text := bytesToStr(node.tokens)
		var i int
		summary := ""
		for _, r := range text {
			i++
			summary += string(r)
			if 4 < i {
				summary += "..."
				break
			}
		}
		r.openObj()
		r.val("Text\n"+summary, node)
		r.closeObj(node)
	}
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderCodeSpan(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.leaf("Code Span\ncode", node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderEmphasis(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		r.val("Emphasis\nem", node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderStrong(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		r.val("Strong\nstrong", node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderBlockquote(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		r.val("Blockquote\nblockquote", node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderHeading(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		h := "h" + " 123456"[node.headingLevel:node.headingLevel+1]
		r.val("Heading\n"+h, node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderList(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		list := "ul"
		if 1 == node.listData.typ || (3 == node.listData.typ && 0 == node.listData.bulletChar) {
			list = "ol"
		}
		r.val("List\n"+list, node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderListItem(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		r.val("List Item\nli "+bytesToStr(node.listData.marker), node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderTaskListItemMarker(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.openObj()
		check := " "
		if node.taskListItemChecked {
			check = "X"
		}
		r.val("Task List Item Marker\n["+check+"]", node)
		r.openChildren(node)
	} else {
		r.closeChildren(node)
		r.closeObj(node)
	}
	return WalkContinue, nil
}

func (r *EChartsJSONRenderer) renderThematicBreak(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.leaf("Thematic Break\nhr", node)
	}
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderHardBreak(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.leaf("Hard Break\nbr", node)
	}
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderSoftBreak(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.leaf("Soft Break\n", node)
	}
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) renderCodeBlock(node *Node, entering bool) (WalkStatus, error) {
	if entering {
		r.leaf("Code Block\npre.code", node)
	}
	return WalkStop, nil
}

func (r *EChartsJSONRenderer) leaf(val string, node *Node) {
	r.openObj()
	r.val(val, node)
	r.closeObj(node)
}

func (r *EChartsJSONRenderer) val(val string, node *Node) {
	val = strings.ReplaceAll(val, "\\", "\\\\")
	val = strings.ReplaceAll(val, "\n", "\\n")
	val = strings.ReplaceAll(val, "\"", "")
	val = strings.ReplaceAll(val, "'", "")
	r.writeString("\"name\":\"" + val + "\"")
}

func (r *EChartsJSONRenderer) openObj() {
	r.writeByte('{')
}

func (r *EChartsJSONRenderer) closeObj(node *Node) {
	r.writeByte('}')
	if !r.ignore(node.next) {
		r.comma()
	}
}

func (r *EChartsJSONRenderer) openChildren(node *Node) {
	if nil != node.firstChild {
		r.writeString(",\"children\":[")
	}
}

func (r *EChartsJSONRenderer) closeChildren(node *Node) {
	if nil != node.firstChild {
		r.writeByte(']')
	}
}

func (r *EChartsJSONRenderer) comma() {
	r.writeString(",")
}

func (r *EChartsJSONRenderer) ignore(node *Node) bool {
	return nil == node ||
		// 以下类型的节点不进行渲染，否则图画出来节点太多
		NodeBlockquoteMarker == node.Typ ||
		NodeEmA6kOpenMarker == node.Typ || NodeEmA6kCloseMarker == node.Typ ||
		NodeEmU8eOpenMarker == node.Typ || NodeEmU8eCloseMarker == node.Typ ||
		NodeStrongA6kOpenMarker == node.Typ || NodeStrongA6kCloseMarker == node.Typ ||
		NodeStrongU8eOpenMarker == node.Typ || NodeStrongU8eCloseMarker == node.Typ ||
		NodeStrikethrough1OpenMarker == node.Typ || NodeStrikethrough1CloseMarker == node.Typ ||
		NodeStrikethrough2OpenMarker == node.Typ || NodeStrikethrough2CloseMarker == node.Typ ||
		NodeMathBlockOpenMarker == node.Typ || NodeMathBlockContent == node.Typ || NodeMathBlockCloseMarker == node.Typ ||
		NodeInlineMathOpenMarker == node.Typ || NodeInlineMathContent == node.Typ || NodeInlineMathCloseMarker == node.Typ
}
