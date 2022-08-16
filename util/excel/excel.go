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

func FileWrite(rows [][]string) {
	// 파일 생성
	file, err := os.Create("./output.csv")
	if err != nil {
		panic(err)
	}

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	for _, row := range rows {
		wr.Write([]string{
			row[0],
			row[1],
			row[2],
			row[3],
			row[4],
			row[5],
			row[6],
			row[7],
			row[8]})
	}

	wr.Flush()
}
