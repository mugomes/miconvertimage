// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package modules

import "github.com/mugomes/mglang"

func LoadTranslations() {
	lang := mglang.GetLang()

	if lang == "pt" {
		mglang.Set("About", "Sobre")
		mglang.Set("Check Update", "Verificar Atualização")
		mglang.Set("Support MiConvertImage", "Apoie MiConvertImage")
		mglang.Set("About MiConvertImage", "Sobre MiConvertImage")
		mglang.Set("Add Image", "Add Imagem")
		mglang.Set("Remove Selected", "Remover Selecionado")
		mglang.Set("Remove All", "Remover Todos")
		mglang.Set("Format", "Formato")
		mglang.Set("Quality", "Qualidade")
		mglang.Set("Size", "Tamanho")
		mglang.Set("Proportion", "Proporção")
		mglang.Set("Keep", "Manter")
		mglang.Set("Do not keep", "Não Manter")
		mglang.Set("Invalid format! Only PNG, JPG, or WEBP files are accepted.", "Formato inválido! Somente arquivos PNG, JPG ou WEBP são aceitos.")
		mglang.Set("File", "Arquivo")
		mglang.Set("Convert Image", "Converter Imagem")
		mglang.Set("Information", "Informação")
		mglang.Set("Converting images...", "Convertendo imagens...")
		mglang.Set("Completed", "Concluído")
		mglang.Set("error opening input image: %v", "erro ao abrir a imagem de entrada: %v")
		mglang.Set("Cannot open image!", "Não foi possível abrir a imagem!")
		mglang.Set("error decoding image: %v", "erro ao decodificar a imagem: %v")
		mglang.Set("Error decoding the image!", "Erro ao decodificar a imagem!")
		mglang.Set("error creating output file: %v", "Erro ao criar o arquivo de saída: %v")
		mglang.Set("Error creating image!", "Erro ao criar a imagem!")
		mglang.Set("unsupported output format: %s", "Formato de saída não suportado: %s")
		mglang.Set("Output format not supported!", "Formato de saída não suportado!")
		mglang.Set("error when converting: %v", "Erro ao converter: %v")
		mglang.Set("Error converting!", "Erro na conversão!")
		mglang.Set("Converted File", "Arquivo convertido")
		mglang.Set("Convert", "Converter")
	}
}

func T(key string, args ...any) string {
	return mglang.T(key, args...)
}
