// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package modules

import (
	"fmt"
	"runtime"
	"mugomes/miconvertimage/locale"
)

type Translations map[string]string

var tr Translations

// Detecta o idioma do sistema e retorna apenas o código base (pt, en, es, etc.)
func LoadTranslations() {
	platform := runtime.GOOS
	var lang string
	if platform == "linux" {
		lang = locale.Current.GetSystemLanguage()
	} else {
		lang = locale.Current.GetSystemLanguage()
	}

	if lang == "pt" {
		valor := make(map[string]string)
		valor["About"] = "Sobre"
		valor["Check Update"] = "Verificar Atualização"
		valor["Support MiConvertImage"] = "Apoie MiConvertImage"
		valor["About MiConvertImage"] = "Sobre MiConvertImage"
		valor["Add Image"] = "Add Imagem"
		valor["Remove Selected"] = "Remover Selecionado"
		valor["Remove All"] = "Remover Todos"
		valor["Format"] = "Formato"
		valor["Quality"] = "Qualidade"
		valor["Size"] = "Tamanho"
		valor["Proportion"] = "Proporção"
		valor["Keep"] = "Manter"
		valor["Do not keep"] = "Não Manter"
		valor["Invalid format! Only PNG, JPG, or WEBP files are accepted."] = "Formato inválido! Somente arquivos PNG, JPG ou WEBP são aceitos."
		valor["File"] = "Arquivo"
		valor["Convert Image"] = "Converter Imagem"
		valor["Information"] = "Informação"
		valor["Converting images..."] = "Convertendo imagens..."
		valor["Completed"] = "Concluído"
		valor["error opening input image: %v"] = "erro ao abrir a imagem de entrada: %v"
		valor["Cannot open image!"] = "Não foi possível abrir a imagem!"
		valor["error decoding image: %v"] = "erro ao decodificar a imagem: %v"
		valor["Error decoding the image!"] = "Erro ao decodificar a imagem!"
		valor["error creating output file: %v"] = "Erro ao criar o arquivo de saída: %v"
		valor["Error creating image!"] = "Erro ao criar a imagem!"
		valor["unsupported output format: %s"] = "Formato de saída não suportado: %s"
		valor["Output format not supported!"] = "Formato de saída não suportado!"
		valor["error when converting: %v"] = "Erro ao converter: %v"
		valor["Error converting!"] = "Erro na conversão!"
		valor["Converted File"] = "Arquivo convertido"
		valor["Convert"] = "Converter"
		tr = valor
	}
}

// T retorna o texto traduzido com formatação opcional.
func T(key string, args ...any) string {
	msg, ok := tr[key]
	if !ok {
		msg = key // fallback se não achar
	}
	return fmt.Sprintf(msg, args...)
}
