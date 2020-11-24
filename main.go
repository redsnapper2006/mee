package main

import (
	"fmt"

	"rs.mee/crawl/spider"
	"rs.mee/crawl/utils"
)

func main() {
	fmt.Println("!!!!!BEGIN!!!!!")
	fmt.Println(utils.MeeHost + utils.MeeRootURL)

	spider.FetchDetail("http://permit.mee.gov.cn/permitExt/xkgkAction!xkgk.action?xkgk=getxxgkContent&dataid=11780462637541919d555f667377952e")
	// http://permit.mee.gov.cn/permitExt/xkgkAction!xkgk.action?xkgk=getxxgkContent&dataid=11780462637541919d555f667377952e

	fmt.Println("!!!!!END!!!!!")
}
