package cli

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func SetupWatcher(onModified func(filename string)) *fsnotify.Watcher {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					onModified(filepath.Base(event.Name))
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				Logger(err.Error()).Error()
			}
		}
	}()

	return watcher
}
