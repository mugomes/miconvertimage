// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"image/color"
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func showAbout(a fyne.App) {
	w := a.NewWindow("Sobre")
	w.Resize(fyne.NewSize(597, 470))
	w.CenterOnScreen()
	w.SetFixedSize(true)

	lblSoftware := canvas.NewText("MiConvertImage - Version:"+VERSION_APP, color.Opaque)
	lblSoftware.TextSize = 18
	lblSoftware.TextStyle.Bold = true
	lblSoftware.Move(fyne.NewPos(9, 7))

	lblDesenvolvedor1 := widget.NewLabel("Desenvolvido por:")
	lblDesenvolvedor1.TextStyle = fyne.TextStyle{Bold: true}
	lblDesenvolvedor1.Move(fyne.NewPos(0, lblSoftware.MinSize().Height+10))

	lblDesenvolvedor2 := widget.NewLabel("Murilo Gomes Julio")
	lblDesenvolvedor2.Move(fyne.NewPos(lblDesenvolvedor1.MinSize().Width-10, lblDesenvolvedor1.Position().Y))

	lblSite1 := widget.NewLabel("Site:")
	lblSite1.TextStyle = fyne.TextStyle{Bold: true}
	lblSite1.Move(fyne.NewPos(0, lblDesenvolvedor1.Position().Y+37))

	sURL, _ := url.Parse("https://mugomes.github.io")
	lblSite2 := widget.NewHyperlink("https://mugomes.github.io", sURL)
	lblSite2.Move(fyne.NewPos(lblSite1.MinSize().Width-10, lblDesenvolvedor2.Position().Y+37))

	lblCopyright1 := widget.NewLabel("Copyright (C) 2024-2025 Murilo Gomes Julio")
	lblCopyright1.TextStyle = fyne.TextStyle{Bold: true}
	lblCopyright1.Move(fyne.NewPos(0, lblSite1.Position().Y+37))

	lblLicense1 := widget.NewLabel("License:")
	lblLicense1.TextStyle = fyne.TextStyle{Bold: true}
	lblLicense1.Move(fyne.NewPos(0, lblCopyright1.Position().Y+37))

	lblLicense2 := widget.NewLabel("GPL-2.0-only")
	lblLicense2.Move(fyne.NewPos(lblLicense1.MinSize().Width-10, lblCopyright1.Position().Y+37))

	txtLicense := widget.NewRichTextFromMarkdown(`
	MiConvertImage is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, only version 2 of the License.
	
	MiConvertImage is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
	`)
	txtLicense.Wrapping = fyne.TextWrapWord
	
	vBoxLicense := container.NewVScroll(txtLicense)
	vBoxLicense.Move(fyne.NewPos(0, lblLicense1.Position().Y+37))
	vBoxLicense.Resize(fyne.NewSize(597, 257))

	layout := container.NewWithoutLayout(
		lblSoftware,
		lblDesenvolvedor1,
		lblDesenvolvedor2,
		lblSite1,
		lblSite2,
		lblCopyright1,
		lblLicense1,
		lblLicense2,
		vBoxLicense)

	w.SetContent(layout)
	w.Show()
}
