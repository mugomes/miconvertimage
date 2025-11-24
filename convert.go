// Copyright (C) 2024-2025 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"fmt"
	m "mugomes/miconvertimage/modules"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgcolumnview"
)

func (s *sDados) showConvert(a fyne.App) {
	m.LoadTranslations()

	w := a.NewWindow(m.T("Convert Image"))
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(800, 600))
	w.SetFixedSize(true)

	lblInfo := widget.NewLabel(m.T("Information"))
	lblInfo.TextStyle = fyne.TextStyle{Bold: true}
	lblInfo.Resize(fyne.NewSize(w.Canvas().Size().Width-7, 30))
	lblInfo.Move(fyne.NewPos(7, 7))

	lblStatus := widget.NewLabel(m.T("Converting images..."))
	lblStatus.Resize(fyne.NewSize(w.Canvas().Size().Width-7, 30))
	lblStatus.Move(fyne.NewPos(7, lblInfo.Position().Y+34))

	sHeader := []string{m.T("File"), m.T("Information")}
	sWidths := []float32{500, 200}
	lstFiles := mgcolumnview.NewColumnView(sHeader, sWidths, false)
	lstFiles.Resize(fyne.NewSize(w.Canvas().Size().Width-7, w.Canvas().Size().Height-7))
	lstFiles.Move(fyne.NewPos(7, lblStatus.Position().Y+38))
	sContainer := container.NewWithoutLayout(
		lblInfo,
		lblStatus,
		lstFiles,
	)

	w.SetContent(sContainer)
	w.Show()

	go func() {
		for _, row := range s.imagens {
			//sFile := filepath.Base(row[0])
			sExtension := filepath.Ext(row[0])

			sInfo, msgError := m.ConvertImage(row[0], strings.Replace(row[0], sExtension, fmt.Sprintf(".%s", s.format), 1), s.format, s.qualidade, uint(s.tamanhoWidth), uint(s.tamanhoHeight), s.proporcao)

			fyne.Do(func() {
				if sInfo == "" {
					i := []string{row[0], msgError}
					lstFiles.AddRow(i)
				} else {
					i := []string{row[0], sInfo}
					lstFiles.AddRow(i)
				}
			})
		}

		fyne.Do(func() {
			lblStatus.Text = m.T("Completed")
			lblStatus.Refresh()
		})
	}()
}
