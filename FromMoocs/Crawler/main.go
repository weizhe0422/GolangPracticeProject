package main

import (
	"fmt"
	"github.com/weizhe0422/GolangPracticeProject/FromMoocs/Crawler/fetcher"
	"regexp"
)

func main() {

	bytes, err := fetcher.Fetch("https://www.ptt.cc/bbs/hotboards.html")
	if err!= nil{
		fmt.Errorf("fail to link to PTT")
	}

	//fmt.Printf("contents: %s", bytes)

	re := regexp.MustCompile(`<a class="board" href="/bbs/[0-9a-zA-Z]+/index.html"><div class="board-name">Gossiping</div>
                <div class="board-nuser"><span class="hl f6">15081</span></div>
                <div class="board-class">綜合</div>
                <div class="board-title">&#9678;[八卦]政問水桶三個月,請珍惜帳號</div>
            </a>`)
	matches:= re.FindAllStringSubmatch(string(bytes), -1)
	for _, m := range matches{
		fmt.Println(m)
	}

	fmt.Printf("Board Count %d", len(matches))
}
