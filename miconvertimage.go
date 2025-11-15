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

	// "fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgcolumnview"
	"github.com/mugomes/mgdialogopenfile"
	"github.com/mugomes/mgnumericentry"
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
	sWidths := []float32{30, w.Canvas().Size().Width - 7}
	data := [][]string{}

	cv := mgcolumnview.NewColumnView(headers, sWidths, true)

	btnAddFile := widget.NewButton("Add Imagem", func() {
		fs := mgdialogopenfile.New(a, "Abrir Arquivo", []string{".webp", ".jpg", ".png"}, true, func(filenames []string) {
			for _, filename := range filenames {
				ext := filepath.Ext(filename)
				if ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".webp" {
					data = append(data, []string{
						fmt.Sprintf("%d", len(data)+1),
						filename,
					})
					cv.AddRow([]string{
						filename,
					})
				} else {
					dialog.ShowInformation("MiCheckHash", "Formato Inválido! Somente arquivos PNG, JPG ou WEBP são aceitos.", w)
				}

			}
		})
		fs.Show()
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

	lblQualidade := widget.NewLabel("Qualidade")
	lblQualidade.TextStyle = fyne.TextStyle{Bold: true}
	lblQualidade.Move(fyne.NewPos(cboFormat.Size().Width+17, cv.Position().Y+300))
	txtQualidade := mgnumericentry.NewMGNumericEntryWithButtons(0, 100, 90)
	txtQualidade.Resize(fyne.NewSize(100, 38))
	txtQualidade.Move(fyne.NewPos(cboFormat.Size().Width+17, lblQualidade.Position().Y+37))
	
	// lblTamanho := widget.NewLabel("Tamanho")
	// lblTamanho.TextStyle = fyne.TextStyle{Bold: true}
	// lblTamanho.Resize(fyne.NewSize(100, 30))
	// lblTamanho.Move(fyne.NewPos(cboFormat.Size().Width+17, cv.Position().Y+300))
	// txtTamanhoWidth := widget.NewEntry()
	// txtTamanhoWidth.Resize(fyne.NewSize(50, 38))
	// txtTamanhoWidth.Move(fyne.NewPos(cboFormat.Size().Width+24, lblTamanho.Position().Y+37))
	// lblX := widget.NewLabel("x")
	// lblX.Resize(fyne.NewSize(50, 38))
	// lblX.Move(fyne.NewPos(txtTamanhoWidth.Position().X+49, lblTamanho.Position().Y+37))
	// txtTamanhoHeight := widget.NewEntry()
	// txtTamanhoHeight.Resize(fyne.NewSize(50,38))
	// txtTamanhoHeight.Move(fyne.NewPos(lblX.Position().X+24, lblTamanho.Position().Y+37))

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
		cboFormat,
		lblQualidade,
		txtQualidade,
		// lblTamanho,
		// txtTamanhoWidth,
		// lblX,
		// txtTamanhoHeight,
		btnConvert,
	)
	w.SetContent(layout)
	w.ShowAndRun()
}
