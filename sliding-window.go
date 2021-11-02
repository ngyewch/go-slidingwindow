package slidingwindow

import (
	"sync"
	"time"
)

type SlidingWindow struct {
	mutex          sync.Mutex
	windowDuration time.Duration
	entries        []slidingWindowEntry
}

type slidingWindowEntry struct {
	t      time.Time
	record interface{}
}

func NewSlidingWindow(windowDuration time.Duration) *SlidingWindow {
	return &SlidingWindow{
		windowDuration: windowDuration,
		entries:        make([]slidingWindowEntry, 0),
	}
}

func (w *SlidingWindow) doPurgeExpired() {
	t := time.Now()
	index := 0
	for ; index < len(w.entries); index++ {
		entry := w.entries[index]
		if t.Sub(entry.t) < w.windowDuration {
			break
		}
	}
	if index > 0 {
		arr := make([]slidingWindowEntry, len(w.entries) - index)
		for i := 0; i < len(arr); i++ {
			arr[i] = w.entries[i + index]
		}
		w.entries = arr
	}
}

func (w *SlidingWindow) doGet() []interface{} {
	records := make([]interface{}, len(w.entries))
	for i := 0; i < len(w.entries); i++ {
		records[i] = w.entries[i].record
	}
	return records
}

func (w *SlidingWindow) Get() []interface{} {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.doPurgeExpired()
	return w.doGet()
}

func (w *SlidingWindow) Append(record interface{}) {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	w.entries = append(w.entries, slidingWindowEntry{t: time.Now(), record: record})
	w.doPurgeExpired()
}
