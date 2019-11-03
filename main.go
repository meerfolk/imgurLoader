package main

import (
	"log"
	"regexp"
	"time"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/martinlindhe/notify"
	"github.com/radovskyb/watcher"

	"github.com/meerfolk/imgurLoader/config"
	"github.com/meerfolk/imgurLoader/imgur"
)

func _checkErr(err error) {
	if err != nil {
		notify.Alert("ImgurLoader", "Imgur Loader", err.Error(), "")
	}
}

func onReady() {
	systray.SetTitle("ImgurLoader")
	systray.SetTooltip("ImgurLoader")
	quitMenu := systray.AddMenuItem("Quit", "Quit app")

	config, err := config.GetOrCreateConfig()
	_checkErr(err)

	w := watcher.New()

	if err := w.Add(config.Path); err != nil {
		log.Fatalln(err)
	}

	go func() {
		for {
			select {
			case event := <-w.Event:
				if !event.IsDir() && regexp.MustCompile(config.File).MatchString(event.Name()) && event.Op == 1 {
					link, err := imgur.Upload(event.Name(), event.Path)
					_checkErr(err)

					notify.Notify("ImgurLoader", "Imgur Loader", "Загрузка завершена", "")
					clipboard.WriteAll(link)
				}
			case error := <-w.Error:
				log.Fatalln(error)
			case <-w.Closed:
				return
			}
		}
	}()

	go func() {
		<-quitMenu.ClickedCh
		w.Close()
		systray.Quit()
	}()

	if err := w.Start(time.Millisecond * 100); err != nil {
		_checkErr(err)
	}
}

func onExit() {

}

func main() {
	systray.Run(onReady, onExit)
}
