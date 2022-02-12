package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"xml_ms_word_helper/base"
)

func main() {
	doc := base.InitDoc()
	body := doc.Body

	paragraph := base.Paragraph{}
	fonts := base.Fonts{}
	fonts.Content = "I had written some words to mydoc.doc, can I read it？"
	fonts.SetUnderLine(base.LineTypeSingle)
	fonts.SetBold()
	fonts.SetItalic()
	fonts.SetStrike()
	fonts.Content = "I had written some words to mydoc.doc, can I read it？"
	fonts.SetUnderLine(base.LineTypeDotDotDash)
	paragraph.Fonts = append(paragraph.Fonts, fonts)
	body.Paragraphs = append(body.Paragraphs, paragraph)

	paragraph1 := base.Paragraph{}
	pic := base.Picture{}
	const sourcePic = "./11111.png"
	file, err := ioutil.ReadFile(sourcePic)
	if err != nil {
		_ = fmt.Sprintf("Something wrong when read the file from disk:%s", sourcePic)
	}
	fileStr := base64.StdEncoding.EncodeToString(file)
	pic.BinDataContent = fileStr
	paragraph1.Pictures = append(paragraph1.Pictures, pic)
	body.Paragraphs = append(body.Paragraphs, paragraph1)

	doc.Body = body
	doc = doc.Build()
	fmt.Print(doc)
	const targetFile = "./mydoc.doc"
	err = ioutil.WriteFile(targetFile, []byte(doc.Content), 0666)
	if err != nil {
		_ = fmt.Sprintf("File write fail:%s", targetFile)
	}
}
