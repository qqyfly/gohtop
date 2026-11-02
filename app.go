package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/shirou/gopsutil/v4/process"
)

func main() {
	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(false).
		SetSelectable(true, false)

	headers := []string{"PID", "Name", "CPU %", "Memory %"}

	for col, h := range headers {
		table.SetCell(0, col, tview.NewTableCell(h))
	}

	processes, _ := process.Processes()

	for i, p := range processes {
		name, _ := p.Name()
		cpu, _ := p.CPUPercent()
		mem, _ := p.MemoryPercent()

		values := []string{
			fmt.Sprint(p.Pid),
			name,
			fmt.Sprintf("%.1f", cpu),
			fmt.Sprintf("%.1f", mem),
		}

		for col, v := range values {
			table.SetCell(i+1, col,
				tview.NewTableCell(v))
		}
	}

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyRune && event.Rune() == 'q' {
			app.Stop()
			return nil
		}
		return event
	})

	if err := app.SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
