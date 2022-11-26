package main
 
import (
	"fmt"
	"os"
	"strconv"
	"time"
 
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
 
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)
 
var mp3name = "coin01.mp3"
 
func doBeep(path string) error {
	if path == "" {
		path = mp3name
	}
	f, _ := os.Open(path)
	s, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	done := make(chan struct{})
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return err
	}
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(done)
	})))
	<-done
 
	return nil
}
 
//Dummy
func makeMenu() *fyne.MainMenu {
	var menuMain = fyne.NewMenu("File", fyne.NewMenuItem("New", func() { fmt.Println("Menu New") }))
	var menuSub1 = fyne.NewMenu("Edit",
		fyne.NewMenuItem("Cut", func() { fmt.Println("Menu Cut") }),
		fyne.NewMenuItem("Copy", func() { fmt.Println("Menu Copy") }),
		fyne.NewMenuItem("Paste", func() { fmt.Println("Menu Paste") }),
	)
	var menu = fyne.NewMainMenu(menuMain, menuSub1)
	return menu
}
 
//Timer
func timer(sec int64, lbl *widget.Label, prg *widget.ProgressBar, stopCh chan struct{}) {
	prg.SetValue(0)
	var s int64
	st := time.NewTicker(time.Duration(1) * time.Second)
	defer st.Stop()
	for {
		select {
		case <-stopCh:
			return
		case <-st.C:
			s++
			lbl.SetText(fmt.Sprintf("%d : %02d : %02d", s/3600, (s%3600)/60, s%60))
			prg.SetValue(float64(s) / float64(sec))
			if s > sec {
				doBeep("")
			}
		}
	}
}
 
//Screen
func makeWelcome() *widget.Box {
	var stopCh chan struct{}
 
	//Labels and Progress
	lbl1 := widget.NewLabelWithStyle("0 : 00 : 00", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	lbl2 := widget.NewLabelWithStyle("0 : 03 : 00", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	prg := widget.NewProgressBar()
 
	//Set Timer Area
	var sec int64 = 180
	e1, e2, e3 := widget.NewEntry(), widget.NewEntry(), widget.NewEntry()
	e1.SetText("0")
	e2.SetText("03")
	e3.SetText("00")
	eb := widget.NewButton("Set", func() {
		t1, _ := strconv.ParseInt(e1.Text, 10, 64)
		t2, _ := strconv.ParseInt(e2.Text, 10, 64)
		t3, _ := strconv.ParseInt(e3.Text, 10, 64)
		sec = t1*3600 + t2*60 + t3
		lbl2.SetText(fmt.Sprintf("%d : %02d : %02d", t1, t2, t3))
	})
	timerSetArea := fyne.NewContainerWithLayout(layout.NewCenterLayout(), widget.NewHBox(e1, e2, e3, eb))
 
	//Buttons
	btnStart := widget.NewButton("Start", func() {
		if stopCh != nil {
			close(stopCh)
		}
		stopCh = make(chan struct{})
		go timer(sec, lbl1, prg, stopCh)
	})
	btnStop := widget.NewButton("Stop", func() {
		if stopCh != nil {
			close(stopCh)
		}
		stopCh = make(chan struct{})
	})
	btn := fyne.NewContainerWithLayout(layout.NewGridLayout(2), btnStart, btnStop)
	return widget.NewVBox(timerSetArea, btn, lbl1, lbl2, prg)
}
 
func main() {
	app := app.New()
 
	w := app.NewWindow("Timer") //w.SetTitle("TitleSet")
	w.SetMainMenu(makeMenu())
 
	tabs := widget.NewTabContainer(
		widget.NewTabItemWithIcon("Welcome", theme.HomeIcon(), makeWelcome()),
		widget.NewTabItemWithIcon("Setting", theme.SettingsIcon(), widget.NewVBox())) //Dummy
	//tabs.SetTabLocation(widget.TabLocationLeading) //Vertical tab menu
	w.SetContent(tabs)
 
	w.ShowAndRun()
}
