package excel

import (
	"bufio"
	"encoding/csv"
	"os"
)

func FileRead(path string) ([][]string, error) {

	// 파일 오픈
	file, error := os.Open(path)

	// csv reader 생성
	rdr := csv.NewReader(bufio.NewReader(file))

	// csv 내용 모두 읽기
	rows, error := rdr.ReadAll()

	return rows, error
}

func FileWrite() {
	// 파일 생성
	file, err := os.Create("./output.csv")
	if err != nil {
		panic(err)
	}

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	wr.Write([]string{"A", "0.25"})
	wr.Write([]string{"B", "55.70"})
	wr.Flush()
}
