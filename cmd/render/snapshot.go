package main

import (
	"github.com/harrydb/go/img/grayscale"
	"github.com/nfnt/resize"

	"fmt"
	"image"
	"net/url"
	"runtime"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sourcegraph/go-webkit2/webkit2"
)

func Snapshot(uri url.URL) (*image.Gray, error) {
	runtime.LockOSThread()
	gtk.Init(nil)
	win, err := gtk.OffscreenWindowNew()
	if err != nil {
		return nil, err
	}
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})
	win.SetDefaultSize(512, 0)

	webView := webkit2.NewWebView()
	defer webView.Destroy()
	win.Add(webView)
	win.ShowAll()

	webView.Connect("load-failed", func() {
		fmt.Println("Load failed.")
	})

	var gray *image.Gray

	webView.Connect("load-changed", func(_ *glib.Object, loadEvent webkit2.LoadEvent) {
		switch loadEvent {
		case webkit2.LoadFinished:
			fmt.Printf("Loaded, now waiting\n")
			webView.GetSnapshot(func(result *image.RGBA, err error) {
				resizedSrc := resize.Resize(512, 0, result, resize.Lanczos3)
				gray = grayscale.Convert(resizedSrc, grayscale.ToGrayLuminance)
				gtk.MainQuit()
				return
			})
		}
	})

	glib.IdleAdd(func() bool {
		webView.LoadURI(uri)
		return false
	})

	return gray, nil
}
