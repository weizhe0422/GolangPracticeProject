package Parser

import (
	"fmt"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocs/Crawler/engine"
	"regexp"
)

var boardListRe = ` <a class="board" href="/bbs/[0-9a-z]+/index.html">`
func ParseBoardList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(boardListRe)
	matches:= re.FindAllSubmatch(contents, -1)
	for _, m := range matches{
		fmt.Print(m)
	}

	fmt.Printf("Board Count %d", len(matches))

}