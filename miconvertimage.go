// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"fmt"
	"image/color"
	m "mugomes/miconvertimage/modules"
	"net/url"
	"path/filepath"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgcolumnview"
	"github.com/mugomes/mgdialogbox"
	"github.com/mugomes/mgnumericentry"
)

const VERSION_APP string = "2.0.0"

type myDarkTheme struct{}

func (m myDarkTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	// A lógica para forçar o modo escuro é retornar cores escuras.
	// O Fyne usa estas constantes internamente:
	switch name {
	case theme.ColorNameBackground:
		return color.RGBA{28, 28, 28, 255} // Fundo preto
	case theme.ColorNameForeground:
		return color.White // Texto branco
	// Adicione outros casos conforme a necessidade (InputBackground, Primary, etc.)
	default:
		// Retorna o tema escuro padrão para as outras cores (se existirem)
		// Aqui estamos apenas definindo as cores principais para garantir o Dark Mode
		return theme.DefaultTheme().Color(name, theme.VariantDark)
	}
}

// 3. Implemente os outros métodos necessários da interface fyne.Theme (usando o tema padrão)
func (m myDarkTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}

func (m myDarkTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (m myDarkTheme) Size(n fyne.ThemeSizeName) float32 {
	if n == theme.SizeNameText {
		return 16
	}
	return theme.DefaultTheme().Size(n)
}

type sDados struct {
	imagens [][]string
	format  string
	qualidade int
	tamanhoWidth int
	tamanhoHeight int
	proporcao bool
}

func main() {
	m.LoadTranslations()

	sIcon := fyne.NewStaticResource("miconvertimage.png", resourceMiConvertImagePngData)
	a := app.NewWithID("br.com.mugomes.miconvertimage")
	a.SetIcon(sIcon)
	w := a.NewWindow("MiConvertImage")
	w.Resize(fyne.NewSize(800, 600))
	w.CenterOnScreen()
	w.SetFixedSize(true)
	a.Settings().SetTheme(&myDarkTheme{})

	mnuAbout := fyne.NewMenu(m.T("About"),
		fyne.NewMenuItem(m.T("Check Update"), func() {
			url, _ := url.Parse("https://github.com/mugomes/miconvertimage/releases")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(m.T("Support MiConvertImage"), func() {
			url, _ := url.Parse("https://mugomes.github.io/apoie.html")
			a.OpenURL(url)
		}),
		fyne.NewMenuItemSeparator(),
		fyne.NewMenuItem(m.T("About MiConvertImage"), func() {
			showAbout(a)
		}),
	)

	w.SetMainMenu(fyne.NewMainMenu(mnuAbout))

	headers := []string{m.T("File")}
	sWidths := []float32{30, w.Canvas().Size().Width - 7}
	data := [][]string{}

	cv := mgcolumnview.NewColumnView(headers, sWidths, true)

	btnAddFile := widget.NewButton(m.T("Add Image"), func() {
		mgdialogbox.NewOpenFile(a, m.T("Open File"), []string{".webp", ".jpg", ".png"}, true, func(filenames []string) {
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
					dialog.ShowInformation("MiCheckHash", m.T("Invalid format! Only PNG, JPG, or WEBP files are accepted."), w)
				}

			}
		})
	})
	btnAddFile.Resize(fyne.NewSize(137, 30))
	btnAddFile.Move(fyne.NewPos(7, 7))

	btnRemoveFile := widget.NewButton(m.T("Remove Selected"), func() {
		cv.RemoveSelected()
	})
	btnRemoveFile.Resize(fyne.NewSize(197, 30))
	btnRemoveFile.Move(fyne.NewPos(btnAddFile.Size().Width+19, 7))

	btnRemoveFiles := widget.NewButton(m.T("Remove All"), func() {
		cv.RemoveAll()
	})
	btnRemoveFiles.Resize(fyne.NewSize(159, 30))
	btnRemoveFiles.Move(fyne.NewPos(btnRemoveFile.Size().Width+167, 7))

	cv.Resize(fyne.NewSize(w.Canvas().Size().Width-7, 300))
	cv.Move(fyne.NewPos(0, btnAddFile.Position().Y+37))

	lblFormat := widget.NewLabel(m.T("Formats"))
	lblFormat.TextStyle = fyne.TextStyle{Bold: true}
	lblFormat.Resize(fyne.NewSize(137, 30))
	lblFormat.Move(fyne.NewPos(0, cv.Position().Y+300))
	sFormats := []string{"webp", "jpg", "png"}
	cboFormat := widget.NewSelectEntry(sFormats)
	cboFormat.Resize(fyne.NewSize(137, 38))
	cboFormat.Move(fyne.NewPos(7, lblFormat.Position().Y+37))

	lblQualidade := widget.NewLabel(m.T("Quality"))
	lblQualidade.TextStyle = fyne.TextStyle{Bold: true}
	lblQualidade.Move(fyne.NewPos(cboFormat.Size().Width+12, cv.Position().Y+300))
	txtQualidade, vQualidade := mgnumericentry.NewMGNumericEntryWithButtons(0, 100, 90)
	txtQualidade.Resize(fyne.NewSize(100, 38))
	txtQualidade.Move(fyne.NewPos(cboFormat.Size().Width+17, lblQualidade.Position().Y+37))

	lblTamanho := widget.NewLabel(m.T("Size"))
	lblTamanho.TextStyle = fyne.TextStyle{Bold: true}
	lblTamanho.Resize(fyne.NewSize(100, 30))
	lblTamanho.Move(fyne.NewPos(cboFormat.Size().Width+82, cv.Position().Y+300))
	txtTamanhoWidth := widget.NewEntry()
	txtTamanhoWidth.SetText("0")
	txtTamanhoWidth.Resize(fyne.NewSize(50, 38))
	txtTamanhoWidth.Move(fyne.NewPos(cboFormat.Size().Width+89, lblTamanho.Position().Y+37))
	lblX := widget.NewLabel("x")
	lblX.Resize(fyne.NewSize(50, 38))
	lblX.Move(fyne.NewPos(txtTamanhoWidth.Position().X+49, lblTamanho.Position().Y+37))
	txtTamanhoHeight := widget.NewEntry()
	txtTamanhoHeight.SetText("0")
	txtTamanhoHeight.Resize(fyne.NewSize(50,38))
	txtTamanhoHeight.Move(fyne.NewPos(lblX.Position().X+24, lblTamanho.Position().Y+37))

	lblProporcao := widget.NewLabel("Proporção")
	lblProporcao.Move(fyne.NewPos(txtTamanhoHeight.Position().X+52, cv.Position().Y+300))
	sProporcao := []string{"Manter", "Não Manter"}
	cboProporcao := widget.NewSelectEntry(sProporcao)
	cboProporcao.SetText("Manter")
	cboProporcao.Resize(fyne.NewSize(137, 38))
	cboProporcao.Move(fyne.NewPos(txtTamanhoHeight.Position().X+59, lblProporcao.Position().Y+37))

	btnConvert := widget.NewButton(m.T("Generate"), func() {
		s := &sDados{}
		s.imagens = cv.ListAll()
		s.format = cboFormat.Text

		s.qualidade = vQualidade.GetValue()

		sTamanhoWidth, _ := strconv.Atoi(txtTamanhoWidth.Text)
		s.tamanhoWidth = sTamanhoWidth

		sTamanhoHeight, _ := strconv.Atoi(txtTamanhoHeight.Text)
		s.tamanhoHeight = sTamanhoHeight

		if cboProporcao.Text == "Manter" {
			s.proporcao = true
		} else {
			s.proporcao = false
		}
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
		lblTamanho,
		txtTamanhoWidth,
		lblX,
		txtTamanhoHeight,
		lblProporcao,
		cboProporcao,
		btnConvert,
	)
	w.SetContent(layout)
	w.ShowAndRun()
}
