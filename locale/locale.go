// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package locale

type Provider interface {
    GetSystemLanguage() string
}

// Usado pelo resto do app:
var Current Provider