// Copyright (C) 2024-2026 Murilo Gomes Julio
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
	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgcolumnview"
	"github.com/mugomes/mgdialogbox"
	"github.com/mugomes/mgnumericentry"
	"github.com/mugomes/mgsmartflow"
)

const VERSION_APP string = "2.1.0"

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
	imagens       [][]string
	format        string
	qualidade     int
	tamanhoWidth  int
	tamanhoHeight int
	proporcao     bool
}

func main() {
	m.LoadTranslations()

	sIcon := fyne.NewStaticResource("miconvertimage.png", resourceAppIconData)
	a := app.NewWithID("br.com.mugomes.miconvertimage")
	a.SetIcon(sIcon)
	w := a.NewWindow("MiConvertImage")
	w.Resize(fyne.NewSize(800, 559))
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

	flow := mgsmartflow.New()

	headers := []string{m.T("File")}
	sWidths := []float32{30, w.Canvas().Size().Width - 7}
	data := [][]string{}

	cv := mgcolumnview.NewColumnView(headers, sWidths, true)

	btnAddFile := widget.NewButton(m.T("Add Image"), func() {
		mgdialogbox.NewOpenFile(a, m.T("Open File"), []string{".webp", ".jpg", ".jpeg", ".png"}, true, func(filenames []string) {
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
					mgdialogbox.NewAlert(a, "MiConvertImage", m.T("Invalid format! Only PNG, JPG, or WEBP files are accepted."), true, "Ok")
				}

			}
		})
	})

	btnRemoveFile := widget.NewButton(m.T("Remove Selected"), func() {
		cv.RemoveSelected()
	})

	btnRemoveFiles := widget.NewButton(m.T("Remove All"), func() {
		cv.RemoveAll()
	})
	btnRemoveFiles.Resize(fyne.NewSize(159, 30))

	flow.AddColumn(btnAddFile, btnRemoveFile, btnRemoveFiles)

	flow.AddRow(cv)

	flow.SetResize(cv, fyne.NewSize(w.Canvas().Size().Width - 7, 300))

	lblFormat := widget.NewLabel(m.T("Format"))
	lblFormat.TextStyle = fyne.TextStyle{Bold: true}
	sFormats := []string{"webp", "jpg", "png"}
	cboFormat := widget.NewSelectEntry(sFormats)
	cboFormat.SetText("webp")
	cboFormat.Entry.Disable()
	ctnFormat := container.NewVBox(lblFormat, cboFormat)

	lblQualidade := widget.NewLabel(m.T("Quality"))
	lblQualidade.TextStyle = fyne.TextStyle{Bold: true}
	txtQualidade, vQualidade := mgnumericentry.NewMGNumericEntryWithButtons(0, 100, 90)
	cmpQualidade := container.NewHBox(widget.NewLabel(""), txtQualidade)
	ctnQualidade := container.NewVBox(lblQualidade, cmpQualidade)

	lblTamanho := widget.NewLabel(m.T("Size"))
	lblTamanho.TextStyle = fyne.TextStyle{Bold: true}
	txtTamanhoWidth := widget.NewEntry()
	txtTamanhoWidth.SetText("0")
	lblX := widget.NewLabel("x")
	txtTamanhoHeight := widget.NewEntry()
	txtTamanhoHeight.SetText("0")
	ctnTamanho := container.NewVBox(
		lblTamanho,
		container.NewHBox(
			txtTamanhoWidth,
			lblX,
			txtTamanhoHeight,
		),
	)
	
	lblProporcao := widget.NewLabel(m.T("Proportion"))
	sProporcao := []string{m.T("Keep"), m.T("Do not keep")}
	cboProporcao := widget.NewSelectEntry(sProporcao)
	cboProporcao.SetText(m.T("Keep"))
	cboProporcao.Entry.Disable()
	ctnProporcao := container.NewVBox(lblProporcao, cboProporcao)

	flow.AddColumn(ctnFormat, ctnQualidade, ctnTamanho, ctnProporcao)

	flow.SetResize(ctnQualidade, fyne.NewSize(100, 38))
	flow.SetResize(ctnTamanho, fyne.NewSize(117, 38))
	btnConvert := widget.NewButton(m.T("Convert"), func() {
		s := &sDados{}
		s.imagens = cv.ListAll()
		s.format = cboFormat.Text

		s.qualidade = vQualidade.GetValue()

		sTamanhoWidth, _ := strconv.Atoi(txtTamanhoWidth.Text)
		s.tamanhoWidth = sTamanhoWidth

		sTamanhoHeight, _ := strconv.Atoi(txtTamanhoHeight.Text)
		s.tamanhoHeight = sTamanhoHeight

		if cboProporcao.Text == m.T("Keep") {
			s.proporcao = true
		} else {
			s.proporcao = false
		}
		s.showConvert(a)
	})

	flow.SetGap(ctnTamanho, fyne.NewPos(0, 37))

	flow.AddColumn(
		layout.NewSpacer(),
		btnConvert,
		layout.NewSpacer(),
	)
	w.SetContent(flow.Container)
	w.ShowAndRun()
}
