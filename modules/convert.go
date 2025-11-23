// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package modules

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"

	"os"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func ConvertImage(inputPath, outputPath string, format string, width, height uint, resizeOnlyOne bool) (string, string) {
	if format == "" {
		format = "webp"
		outputPath = fmt.Sprintf("%s%s", outputPath, format)
	}
	// Abre arquivo de entrada
	inFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Print(T("error opening input image: %v", err))
		return "", T("Cannot open image!")
	}
	defer inFile.Close()

	// Decodifica qualquer formato suportado
	img, _, err := image.Decode(inFile)
	if err != nil {
		fmt.Print(T("error decoding image: %v", err))
		return "", T("Error decoding the image!")
	}

	// Se resizeOnlyOne for verdadeiro, ajusta apenas a largura ou altura
	if resizeOnlyOne {
		// Redimensiona apenas a largura, mantendo a altura original
		if width > 0 {
			img = resize.Resize(width, 0, img, resize.Lanczos3)
		}
		// Redimensiona apenas a altura, mantendo a largura original
		if height > 0 {
			img = resize.Resize(0, height, img, resize.Lanczos3)
		}
	} else {
		// Redimensiona proporcionalmente, caso largura e altura sejam ambos fornecidos
		if width > 0 || height > 0 {
			img = resize.Resize(width, height, img, resize.Lanczos3)
		}
	}

	// Cria arquivo de saída
	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Print(T("error creating output file: %v", err))
		return "", T("Error creating image!")
	}
	defer outFile.Close()

	// Detecta formato destino pela extensão
	switch format {
	case "jpg", "jpeg":
		// Salva como JPG
		options := &jpeg.Options{Quality: 90}
		err = jpeg.Encode(outFile, img, options)

	case "png":
		// Salva como PNG
		err = png.Encode(outFile, img)

	case "webp":
		// Salva como WEBP
		options := &webp.Options{Quality: 90}
		err = webp.Encode(outFile, img, options)

	default:
		fmt.Print(T("unsupported output format: %s", format))
		return "", T("Output format not supported!")
	}

	if err != nil {
		fmt.Print(T("error when converting: %v", err))
		return "", T("Error converting")
	}

	return T("Converted File"), ""
}
