package main

import (
	"context"
	"fmt"

	"github.com/gdamore/tcell"
	"github.com/kyeett/elves-vs-goblin/pkg/client"
	"github.com/kyeett/elves-vs-goblin/pkg/input"
	log "github.com/sirupsen/logrus"

	"github.com/rivo/tview"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	debugView := tview.NewTextView().
		SetDynamicColors(true).
		SetChangedFunc(func() {
			app.Draw()
		})

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Top"), 0, 1, false).
			AddItem(textView, 20, 3, false).
			// AddItem(debugView, 10, 1, false), 0, 2, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Bottom"), 9, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)

	inputCh := make(chan input.Command)
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// textView.Clear()
		// log.Fatal("RUNE:", string(event.Rune()))
		switch event.Key() {
		case tcell.KeyDown:
			inputCh <- input.MoveDown
		case tcell.KeyUp:
			inputCh <- input.MoveUp
		case tcell.KeyLeft:
			inputCh <- input.MoveLeft
		case tcell.KeyRight:
			inputCh <- input.MoveRight
		case tcell.KeyCtrlC:
			log.Info("Received ctrl+C, shutting down client")
			cancel()
			app.Stop()
		}
		return event
	})

	// w := Writer{}
	// c := client.New(&w)
	c := client.New(textView)
	err := c.Connect()
	if err != nil {
		log.Fatal("client failed to connect", err)
	}
	defer c.Close()

	// c.SetOutput(debugView)
	// c.SetLevel(logrus.DebugLevel)
	go c.Run(ctx, inputCh)

	// go func() {

	// 	for index := 0; index < 200; index++ {
	// 		// fmt.Fprintf(textView, "Counter %d", index)
	// 		time.Sleep(100 * time.Millisecond)
	// 		inputCh <- input.Command(rand.Intn(4))
	// 		// textView.Clear()
	// 	}
	// }()

	fmt.Fprintln(debugView, "Ok, seems to work")

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
