package service

import "github.com/jeonyunjae/fiber-api/service/userlocation/list"

func ServiceInit() {
	//각각의 방식 정의
	// 1.list
	list.UserlocationList.UserlocationsInit()
	// 2.decimaltree

	// 3.dictionary

	// 4.kdtree

	// 5.orm

	// 6.query

}
