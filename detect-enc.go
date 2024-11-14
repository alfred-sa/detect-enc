package main

import (
	"fmt"
	"io"
	"math"
	"os"

	"github.com/h2non/filetype"
)

const (
	entropyThreshold = 7.8
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <chemin_du_fichier>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Impossible d'ouvrir le fichier: %s\n", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	kind, err := checkFileType(file)
	if err != nil {
		fmt.Printf("Erreur lors de la vérification du type de fichier: %s\n", err.Error())
		os.Exit(1)
	}
	if kind != "" && kind != "unknown" {
		fmt.Printf("Le fichier n'est pas chiffré\nFile type: %s\n", kind)
		os.Exit(0)
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Printf("Impossible de revenir au début du fichier: %s\n", err.Error())
		os.Exit(1)
	}

	entropy, err := checkFileEntropy(file)
	if err != nil {
		fmt.Printf("Erreur lors de la vérification de l'entropie du fichier: %s\n", err.Error())
		os.Exit(1)
	}

	status := 0
	if entropy > entropyThreshold {
		fmt.Printf("Le fichier est probablement chiffré\n")
		status = 2
	} else {
		fmt.Printf("Le fichier n'est probablement pas chiffré\n")
	}
	fmt.Printf("L'entropie du fichier est: %f bits par octet\n", entropy)
	os.Exit(status)
}

func checkFileEntropy(file *os.File) (float64, error) {
	byteFrequency := make(map[byte]int)
	var totalBytes int

	buffer := make([]byte, 4096)
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, fmt.Errorf("Erreur lors de la lecture du fichier: %w\n", err)
		}
		for _, b := range buffer[:n] {
			byteFrequency[b]++
			totalBytes++
		}
	}

	var entropy float64
	for _, freq := range byteFrequency {
		p := float64(freq) / float64(totalBytes)
		entropy -= p * math.Log2(p)
	}

	return entropy, nil
}

func checkFileType(file *os.File) (string, error) {
	buffer := make([]byte, 262)
	n, err := file.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("Erreur lors de la lecture du fichier: %w\n", err)
	}
	kind, err := filetype.Match(buffer[:n])
	if err != nil {
		return "", fmt.Errorf("Erreur lors de la vérification du type de fichier: %w\n", err)
	}
	return kind.Extension, nil
}
