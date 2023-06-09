package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func splitFileIntoChunks(filePath string, chunkSize int64) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	chunkNumber := 0
	buffer := make([]byte, chunkSize)

	for {
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			log.Fatal("Error reading file:", err)
		}
		if n == 0 {
			break
		}

		outputPath := filepath.Join("temp", "chunk_"+strconv.Itoa(chunkNumber)+".mp4")
		chunkFile, err := os.Create(outputPath)
		if err != nil {
			log.Fatal("Error creating chunk file:", err)
		}
		defer chunkFile.Close()

		_, err = chunkFile.Write(buffer[:n])
		if err != nil {
			log.Fatal("Error writing chunk:", err)
		}

		chunkNumber++
	}
}

func main() {
	filePath := "'for_hemma' (1080p).mp4"
	chunkSize := 1024 * 1024 * 64 // Chunk size in bytes

	splitFileIntoChunks(filePath, int64(chunkSize))
}
