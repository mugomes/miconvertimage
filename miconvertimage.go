// Copyright (C) 2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"fmt"
	"net/url"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgcolumnview"
)

const VERSION_APP string = "2.0.0"

type myTheme struct {
	fyne.Theme
}

func (m myTheme) Size(name fyne.ThemeSizeName) float32 {
	if name == theme.SizeNameText {
		return 16 // Set your desired font size here
	}
	return m.Theme.Size(name)
}

type sDados struct {
	imagens [][]string
	format  string
}

func main() {
	sIcon, err := fyne.LoadResourceFromPath("icon/miconvertimage.png")
	if err != nil {
		panic(err)
	}
	a := app.NewWithID("br.com.mugomes.miconvertimage")
	a.SetIcon(sIcon)
	w := a.NewWindow("MiConvertImage")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	a.Settings().SetTheme(&myTheme{theme.DarkTheme()})

	mnuAbout := fyne.NewMenu("Sobre",
		fyne.NewMenuItem("Check Update", func() {
			url, _ := url.Parse("https://github.com/mugomes/miconvertimage/releases")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("Support MiConvertImage", func() {
			url, _ := url.Parse("https://mugomes.github.io/apoie.html")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem("About MiConvertImage", func() {
			showAbout(a)
		}),
	)

	w.SetMainMenu(fyne.NewMainMenu(mnuAbout))

	headers := []string{"Arquivo"}
	sWidths := []float32{30,400}
	data := [][]string{}

	cv := mgcolumnview.NewColumnView(headers, sWidths, true)

	btnAddFile := widget.NewButton("Add Imagem", func() {
		fd := dialog.NewFileOpen(func(r fyne.URIReadCloser, err error) {
			if err != nil || r == nil {
				return
			}

			ext := filepath.Ext(r.URI().Path())
			if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
				data = append(data, []string{
					fmt.Sprintf("%d", len(data)+1),
					r.URI().Path(),
				})
				cv.AddRow([]string{
					r.URI().Path(),
				})
			} else {
				dialog.ShowInformation("MiCheckHash", "Formato Inválido! Somente arquivos PNG, JPG ou WEBP são aceitos.", w)
			}
		}, w)

		fd.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg", ".webp"}))
		fd.Show()
	})
	btnAddFile.Resize(fyne.NewSize(137, 30))
	btnAddFile.Move(fyne.NewPos(7, 7))

	btnRemoveFile := widget.NewButton("Remover Selecionados", func() {
		cv.RemoveSelected()
	})
	btnRemoveFile.Resize(fyne.NewSize(197, 30))
	btnRemoveFile.Move(fyne.NewPos(btnAddFile.Size().Width+19, 7))

	btnRemoveFiles := widget.NewButton("Remover Todos", func() {
		cv.RemoveAll()
	})
	btnRemoveFiles.Resize(fyne.NewSize(159, 30))
	btnRemoveFiles.Move(fyne.NewPos(btnRemoveFile.Size().Width+167, 7))

	cv.Resize(fyne.NewSize(w.Canvas().Size().Width-7, 300))
	cv.Move(fyne.NewPos(0, btnAddFile.Position().Y+37))
	
	lblFormat := widget.NewLabel("Formatos")
	lblFormat.TextStyle = fyne.TextStyle{Bold: true}
	lblFormat.Resize(fyne.NewSize(137, 30))
	lblFormat.Move(fyne.NewPos(0, cv.Position().Y+300))
	sFormats := []string{"webp", "jpg", "png"}
	cboFormat := widget.NewSelectEntry(sFormats)
	cboFormat.Resize(fyne.NewSize(137, 38))
	cboFormat.Move(fyne.NewPos(7, lblFormat.Position().Y+37))

	btnConvert := widget.NewButton("Gerar", func() {
		s := &sDados{}
		s.imagens = cv.ListAll()
		s.format = cboFormat.Text
		s.showConvert(a)
	})

	btnConvert.Resize(fyne.NewSize(157, 59))
	btnConvert.Move(fyne.NewPos(7, cboFormat.Position().Y+67))

	layout := container.NewWithoutLayout(
		btnAddFile,
		btnRemoveFile,
		btnRemoveFiles,
		cv,
		lblFormat,
		cboFormat,
		btnConvert,
	)
	w.SetContent(layout)
	w.ShowAndRun()
}
