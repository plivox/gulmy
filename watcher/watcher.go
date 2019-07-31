package watcher

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/plivox/glumy/shell"

	"github.com/fsnotify/fsnotify"
)

type watcher struct {
	filters []fsnotify.Op
	paths   []string
}

func New() *watcher {
	return &watcher{filters: []fsnotify.Op{
		fsnotify.Chmod,
		fsnotify.Rename,
	}}
}

func (w *watcher) Add(path string) error {
	fullPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	if !shell.IsExist(fullPath) {
		return fmt.Errorf("Path %s not exist", fullPath)
	}

	if shell.IsDir(fullPath) {
		err = filepath.Walk(fullPath, func(walkPath string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fullPath, err := filepath.Abs(walkPath)
			if err != nil {
				return err
			}
			w.paths = append(w.paths, fullPath)
			return nil
		})
	} else {
		w.paths = append(w.paths, fullPath)
	}
	return nil
}

func (w *watcher) Start(callback func(fsnotify.Event)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				var ignore = false
				for _, filter := range w.filters {
					if event.Op&filter == filter {
						ignore = true
						continue
					}
				}

				if !ignore {
					callback(event)
				}
			}
		}
	}()

	fmt.Printf("Watch %d files ...\n", len(w.paths))
	for _, path := range w.paths {
		err = watcher.Add(path)
		if err != nil {
			log.Fatal(err)
		}
	}
	<-done
	return nil
}
