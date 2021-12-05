package main

import (
	"fmt"
	"harry.com/jianzhi_offer/qian"
)

/**
替换单词
 */
/**
题目：英语中有一个概念叫词根。在词根后面加上若干字符就能拼出更长的单词。例如，"an"是一个词根，在它后面加上"other"就能得到另一个单词"another"。现在给定一个由词根组成的字典和一个英语句子，如果句子中的单词在字典中有它的词根，则用它的词根替换该单词；如果单词没有词根，则保留该单词。请输出替换后的句子。例如，如果词根字典包含字符串["cat"，"bat"，"rat"]，英语句子为"the cattlewas rattled by the battery"，则替换之后的句子是"the cat was rat by the bat"。
 */

func main() {
	result:=qian.ReplaceWords([]string{"cat","bat","rat"},"the cattlewas rattled by the battery")
	fmt.Println(result)
}


