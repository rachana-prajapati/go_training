package main

import (
	"io"
	"os"
)

type AlphaReader string

func (alphaReader AlphaReader) Read(p []byte) (int, error) {
	count := 0
	for idx := 0; idx < len(alphaReader); idx++ {
		if (alphaReader[idx] >= 'A' && alphaReader[idx] <= 'Z') || (alphaReader[idx] >= 'a' && alphaReader[idx] <= 'z') {
			p[count] = alphaReader[idx]
			count++
		}
	}
	return count, io.EOF
}

func Copy(writer io.Writer, reader AlphaReader) {
	buffer := make([]byte, 2048)
	count, err := reader.Read(buffer)
	totalBytesRead := 0
	if count > 0 {
		totalBytesRead += count
		writer.Write(buffer)
	}
	if err == io.EOF {
		return
	}
	if err != nil {
		panic(err)
	}
}

func main() {
	alphaReader := AlphaReader("378535 Consequat cillum^(*&(&* adipisicing ea nisi tempor consequat velit commodo nisi dolor culpa magna nulla. Aute in elit ea occaecat elit. Ut laborum dolor occaecat laborum pariatur ipsum adipisicing cupidatat irure cupidatat reprehenderit esse. Amet ut irure Lorem in mollit proident exercitation consectetur magna et cillum aliqua. Eu sit aliqua elit commodo quis amet enim. Id sint et proident nisi aliquip dolore consectetur ullamco ad.")
	Copy(os.Stdout, alphaReader)
}
