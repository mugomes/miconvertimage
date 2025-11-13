// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package modules

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Translations map[string]string

var tr Translations

// Detecta o idioma do sistema e retorna apenas o código base (pt, en, es, etc.)
func GetSystemLanguage() string {
    lang := os.Getenv("LANG")
    if lang == "" {
        lang = os.Getenv("LC_ALL")
    }
    if lang == "" {
        lang = os.Getenv("LC_MESSAGES")
    }
    if lang == "" {
        return "en" // fallback padrão
    }

    // Exemplo: "pt_BR.UTF-8" → "pt"
    parts := strings.Split(lang, ".")
    base := parts[0]
    baseParts := strings.Split(base, "_")
    return strings.ToLower(baseParts[0])
}

func LoadTranslations() error {
    file, err := os.Open("langs/" + GetSystemLanguage() + ".json")
    if err != nil {
        return err
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    return decoder.Decode(&tr)
}

// T retorna o texto traduzido com formatação opcional.
func T(key string, args ...interface{}) string {
    msg, ok := tr[key]
    if !ok {
        msg = key // fallback se não achar
    }
    return fmt.Sprintf(msg, args...)
}
