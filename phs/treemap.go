package phs

import (
	"harry.com/jianzhi_offer/zhan"
	"math"
)

//const RED = 0
//const BLACK = 1
//
//type Type int

type RBNodeMap struct {
	color               uint8
	key                 int
	value               int
	left, right, parent *RBNodeMap
}

func (node *RBNodeMap) GetValue () int {
	if node!=nil{
		return node.value
	}
	return math.MaxInt
}

type RBRootMap struct {
	root *RBNodeMap
}

func (node *RBNodeMap) rb_is_red() bool {
	return node.color == RED
}

func (node *RBNodeMap) rb_is_black() bool {
	return node.color == BLACK
}

func (node *RBNodeMap) rb_set_red() {
	node.color = RED
}

func (node *RBNodeMap) rb_set_black() {
	node.color = BLACK
}

func (root *RBRootMap) left_rotate(x *RBNodeMap) {
	var y *RBNodeMap = x.right

	// ly 和 x 的关系
	x.right = y.left
	if y.left != nil {
		y.left.parent = x
	}

	// px 和 y 的关系（要考虑px为空，即x为根节点的情况）
	y.parent = x.parent
	if x.parent == nil {
		root.root = y
	} else {
		if x.parent.left == x {
			x.parent.left = y
		} else {
			x.parent.right = y
		}
	}

	// y 和 x 的关系
	y.left = x
	x.parent = y

}

func (root *RBRootMap) right_rotate(y *RBNodeMap) {
	var x *RBNodeMap = y.left

	// rx 和 y 的关系
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}

	// py 和 x 的关系（要考虑py为空，即y为根节点的情况）
	x.parent = y.parent
	if y.parent == nil {
		root.root = x
	} else {
		if y.parent.right == y {
			y.parent.right = x
		} else {
			y.parent.left = x
		}
	}

	// y 和 x 的关系
	x.right = y
	y.parent = x
}

// 添加节点：将节点(node)插入到红黑树中
func (root *RBRootMap) Set(key int, value int) {
	node := root.root
	for node != nil && node.key != key {
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	if node == nil {
		node = &RBNodeMap{key: key, value: value}
	} else {
		node.value = value
		return
	}

	var y *RBNodeMap
	var x *RBNodeMap = root.root

	// 找到node应插入位置的父节点
	for x != nil {
		y = x
		if node.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	// 设置node和父节点的关系
	node.parent = y
	if y != nil {
		if node.key < y.key {
			y.left = node
		} else {
			y.right = node
		}
	} else {
		root.root = node
	}
	// 设置节点为红色
	node.color = RED
	// 修正为红黑树
	root.add_fixup(node)

}

// 红黑树插入修正
func (root *RBRootMap) add_fixup(node *RBNodeMap) {
	var parent, gparent *RBNodeMap

	// 若“父节点存在，并且父节点的颜色是红色”
	for parent = node.parent; parent != nil && parent.rb_is_red(); {
		gparent = parent.parent

		//若“父节点”是“祖父节点的左孩子”
		if parent == gparent.left {
			var uncle *RBNodeMap = gparent.right
			// Case 1条件：叔叔节点是红色
			if uncle != nil && uncle.rb_is_red() {
				uncle.rb_set_black()
				parent.rb_set_black()
				gparent.rb_set_red()
				node = gparent
				continue
			}
			// Case 2条件：叔叔是黑色，且当前节点是右孩子
			if parent.right == node {
				root.left_rotate(parent)
				var tmp *RBNodeMap = parent
				parent = node
				node = tmp
			}
			// Case 3条件：叔叔是黑色，且当前节点是左孩子。
			parent.rb_set_black()
			gparent.rb_set_red()
			root.right_rotate(gparent)
		} else {
			//若“父节点”是“祖父节点的右孩子”

			var uncle *RBNodeMap = gparent.left
			// Case 1条件：叔叔节点是红色
			if uncle != nil && uncle.rb_is_red() {
				uncle.rb_set_black()
				parent.rb_set_black()
				gparent.rb_set_red()
				node = gparent
				continue
			}
			// Case 2条件：叔叔是黑色，且当前节点是左孩子
			if parent.left == node {
				root.right_rotate(parent)
				var tmp *RBNodeMap = parent
				parent = node
				node = tmp
			}
			// Case 3条件：叔叔是黑色，且当前节点是右孩子。
			parent.rb_set_black()
			gparent.rb_set_red()
			root.left_rotate(gparent)
		}
	}
	// 将根节点设为黑色
	root.root.rb_set_black()
}

func (root *RBRootMap) Delete(key int) {
	node := root.root
	for node != nil && node.key != key {
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}
	if node == nil {
		return
	}
	var child, parent *RBNodeMap
	var color uint8

	// 被删除节点的"左右孩子都不为空"的情况。
	if node.left != nil && node.right != nil {
		// 获取后继节点
		var replace *RBNodeMap = node.right
		for replace.left != nil {
			replace = replace.left
		}

		// "node节点"不是根节点
		if node.parent != nil {
			if node.parent.left == node {
				node.parent.left = replace
			} else {
				node.parent.right = replace
			}
		} else {
			// "node节点"是根节点，更新根节点。
			root.root = replace
		}
		// child是"取代节点"的右孩子，也是需要"调整的节点"。
		// "取代节点"肯定不存在左孩子！因为它是一个后继节点。
		child = replace.right
		parent = replace.parent
		// 保存"取代节点"的颜色(注意这里删掉的是node节点，但实际删掉的颜色是replace的)
		color = replace.color

		// "被删除节点"是"它的后继节点的父节点"
		if parent == node {
			parent = replace
		} else {
			// child不为空
			if child != nil {
				child.parent = parent
			}
			parent.left = child

			replace.right = node.right
			node.right.parent = replace
		}

		replace.parent = node.parent
		replace.color = node.color
		replace.left = node.left
		node.left.parent = replace

		if color == BLACK {
			root.delete_fixup(child, parent)
		}

		return
	}

	if node.left != nil {
		child = node.left
	} else {
		child = node.right
	}

	parent = node.parent
	color = node.color // 保存"取代节点"的颜色

	if child != nil {
		child.parent = parent
	}

	// "node节点"不是根节点
	if parent != nil {
		if parent.left == node {
			parent.left = child
		} else {
			parent.right = child
		}
	} else {
		root.root = child
	}

	if color == BLACK {
		root.delete_fixup(child, parent)
	}

}

func (root *RBRootMap) delete_fixup(node *RBNodeMap, parent *RBNodeMap) {
	var other *RBNodeMap

	for (node == nil || node.rb_is_black()) && node != root.root {
		// node是父节点的左孩子
		if parent.left == node {
			other = parent.right
			// Case 1: node的兄弟节点是红色的
			if other.rb_is_red() {
				other.rb_set_black()
				parent.rb_set_red()
				root.left_rotate(parent)
				other = parent.right
			}
			// Case 2: node的兄弟w是黑色，且w的俩个孩子也都是黑色的
			if (other.left == nil || other.left.rb_is_black()) && (other.right == nil || other.right.rb_is_black()) {
				other.rb_set_red()
				node = parent
				parent = node.parent
			} else {
				// Case 3: node的兄弟w是黑色的，并且w的左孩子是红色，右孩子为黑色。
				if other.right == nil || other.right.rb_is_black() {
					other.left.rb_set_black()
					other.rb_set_red()
					root.right_rotate(other)
					other = parent.right
				}
				// Case 4: node的兄弟w是黑色的；并且w的右孩子是红色的，左孩子任意颜色。
				other.color = parent.color
				parent.rb_set_black()
				other.right.rb_set_black()
				root.left_rotate(parent)
				node = root.root
				break
			}
		} else {
			other = parent.left
			// Case 1: node的兄弟w是红色的
			if other.rb_is_red() {
				other.rb_set_black()
				parent.rb_set_red()
				root.right_rotate(parent)
				other = parent.left
			}
			// Case 2: node的兄弟w是黑色，且w的俩个孩子也都是黑色的
			if (other.left == nil || other.left.rb_is_black()) && (other.right == nil || other.right.rb_is_black()) {
				other.rb_set_red()
				node = parent
				parent = node.parent
			} else {
				// Case 3: node的兄弟w是黑色的，并且w的左孩子是红色，右孩子为黑色。
				if other.left == nil || other.left.rb_is_black() {
					other.right.rb_set_black()
					other.rb_set_red()
					root.left_rotate(other)
					other = parent.left
				}
				// Case 4: node的兄弟w是黑色的；并且w的右孩子是红色的，左孩子任意颜色。
				other.color = parent.color
				parent.rb_set_black()
				other.left.rb_set_black()
				root.right_rotate(parent)
				node = root.root
				break
			}
		}
	}

	if node != nil {
		node.rb_set_black()
	}
}

type MapIterator struct {
	cur   *RBNodeMap
	stack *zhan.Stack
}

func constructor(root *RBNodeMap) MapIterator {
	return MapIterator{
		cur:   root,
		stack: zhan.NewStack(),
	}
}
func (bs *MapIterator) hasNext() bool {
	if bs.cur != nil || !bs.stack.IsEmpty() {
		return true
	} else {
		return false
	}
}
func (bs *MapIterator) next() *RBNodeMap {
	for bs.cur != nil {
		bs.stack.Push(bs.cur)
		bs.cur = bs.cur.left
	}
	node, _ := bs.stack.Pop().(*RBNodeMap)
	bs.cur = node.right
	return node
}

type MapIteratorReversed struct {
	cur   *RBNodeMap
	stack *zhan.Stack
}

func constructorMapR(root *RBNodeMap) (br MapIteratorReversed) {
	br.cur = root
	br.stack = zhan.NewStack()
	return br
}
func (br *MapIteratorReversed) hasPrev() bool {
	if br.cur != nil || !br.stack.IsEmpty() {
		return true
	} else {
		return false
	}
}
func (br *MapIteratorReversed) prev() *RBNodeMap {
	for br.cur != nil {
		br.stack.Push(br.cur)
		br.cur = br.cur.right
	}
	node, _ := br.stack.Pop().(*RBNodeMap)
	br.cur = node.left
	return node
}
func (root *RBRootMap) Floor(key int) *RBNodeMap {
	if root.root == nil {
		return nil
	}
	R := constructorMapR(root.root)
	for R.hasPrev() {
		prev := R.prev()
		if prev.key <= key {
			return prev
		}
	}
	return nil
}

func (root *RBRootMap) Ceil(key int) *RBNodeMap {
	if root.root == nil {
		return nil
	}
	R := constructor(root.root)
	for R.hasNext() {
		next := R.next()
		if next.key >= key {
			return next
		}
	}
	return nil
}