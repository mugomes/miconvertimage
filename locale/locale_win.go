// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

//go:build windows
package locale

import (
    "golang.org/x/sys/windows/registry"
    "strings"
)

type winProvider struct{}

func init() {
    Current = winProvider{}
}

func (winProvider) GetSystemLanguage() string {
    k, _, err := registry.CreateKey(
        registry.CURRENT_USER,
        `Control Panel\International`,
        registry.QUERY_VALUE,
    )
    if err != nil {
        return "en"
    }
    defer k.Close()

    s, _, err := k.GetStringValue("LocaleName")
    if err != nil || len(s) < 2 {
        return "en"
    }

    return strings.ToLower(s[:2])
}
