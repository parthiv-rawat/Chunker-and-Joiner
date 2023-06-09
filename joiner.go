package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

func reassembleVideoChunks(chunksDirectory, outputFile string) {
	chunkFiles := []string{}
	chunkNumber := 0

	for {
		chunkPath := filepath.Join(chunksDirectory, "chunk_"+strconv.Itoa(chunkNumber)+".mp4")
		if _, err := os.Stat(chunkPath); os.IsNotExist(err) {
			break
		}

		chunkFiles = append(chunkFiles, chunkPath)
		chunkNumber++
	}

	// Step 1: Concatenate video chunks into a single file without re-encoding
	tempOutputFile := filepath.Join(chunksDirectory, "temp_output.mp4")
	output, err := os.Create(tempOutputFile)
	if err != nil {
		log.Fatal("Error creating temporary output file:", err)
	}
	defer output.Close()

	for _, chunkFile := range chunkFiles {
		chunk, err := os.Open(chunkFile)
		if err != nil {
			log.Fatal("Error opening chunk file:", err)
		}
		defer chunk.Close()

		_, err = io.Copy(output, chunk)
		if err != nil {
			log.Fatal("Error writing chunk to temporary output file:", err)
		}
	}

	// Step 2: Optimize the video file using ffmpeg
	cmd := exec.Command("ffmpeg", "-i", tempOutputFile, "-c", "copy", outputFile)
	err = cmd.Run()
	if err != nil {
		log.Fatal("Error running ffmpeg command:", err)
	}

	// Cleanup temporary files
	os.Remove(tempOutputFile)
}

func main() {
	chunksDirectory := "temp"
	outputFile := filepath.Join(chunksDirectory, "output_video.mp4")

	reassembleVideoChunks(chunksDirectory, outputFile)
}
