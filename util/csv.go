package util

import (
	"bytes"
	"database/sql/driver"
	"encoding/csv"
	"io"
	"os"

	"github.com/unknwon/com"
)

//ReadCsv Read all rows from csv
func ReadCsv(path string) ([][]string, error) {
	fs, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//CsvFormat build csv format
func CsvFormat(header []string, rows [][]interface{}) (*bytes.Buffer, error){
	b := &bytes.Buffer{}
	writer := csv.NewWriter(b)
	err := writer.Write(header)
	if err != nil {
		return nil, err
	}
	for _, row := range rows {
		var tmp []interface{}
		for _, r := range row {
			if rr, ok := r.(driver.Valuer); ok {
				vv, _ := rr.Value()
				tmp = append(tmp, vv)
			}else {
				tmp = append(tmp, r)
			}
		}
		err := writer.Write(ToStrSlice(tmp))
		if err != nil {
			return nil, err
		}
	}
	writer.Flush()
	return b, nil
}

func ToStrSlice(src []interface{})[]string {
	var res = make([]string,len(src))
	for i, v := range src {
		res[i] = com.ToStr(v)
	}
	return res
}


func WriteFile(filename string, data []byte) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
