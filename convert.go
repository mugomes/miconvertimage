// Copyright (C) 2024-2026 Murilo Gomes Julio
// SPDX-License-Identifier: GPL-2.0-only

// Site: https://mugomes.github.io

package main

import (
	"fmt"
	m "mugomes/miconvertimage/modules"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/mugomes/mgcolumnview"
	"github.com/mugomes/mgsmartflow"
)

func (s *sDados) showConvert(a fyne.App) {
	m.LoadTranslations()

	flow := mgsmartflow.New()

	w := a.NewWindow(m.T("Convert Image"))
	w.CenterOnScreen()
	w.Resize(fyne.NewSize(800, 559))
	w.SetFixedSize(true)

	lblInfo := widget.NewLabel(m.T("Information"))
	lblInfo.TextStyle = fyne.TextStyle{Bold: true}

	flow.AddRow(lblInfo)

	lblStatus := widget.NewLabel(m.T("Converting images..."))

	flow.AddRow(lblStatus)

	sHeader := []string{m.T("File"), m.T("Information")}
	sWidths := []float32{500, 200}
	lstFiles := mgcolumnview.NewColumnView(sHeader, sWidths, false)

	flow.AddRow(lstFiles)

	flow.SetResize(lstFiles, fyne.NewSize(w.Canvas().Size().Width - 7, 467))

	w.SetContent(flow.Container)
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
