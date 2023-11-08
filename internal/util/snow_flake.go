package util

import (
	"fmt"
	"sync"
	"time"
)

const (
	workerBits   uint8 = 10
	sequenceBits uint8 = 12
	workerMax    int64 = -1 ^ (-1 << workerBits)
	sequenceMask int64 = -1 ^ (-1 << sequenceBits)
)

type IdWorker struct {
	mu            sync.Mutex
	workerId      int64
	lastTimestamp int64
	sequence      int64
}

func NewIdWorker(workerId int64) *IdWorker {
	if workerId < 0 || workerId > workerMax {
		panic(fmt.Sprintf("workerId can't be greater than %d or less than 0", workerMax))
	}
	return &IdWorker{
		workerId:      workerId,
		lastTimestamp: -1,
		sequence:      0,
	}
}

func (w *IdWorker) NextId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	timestamp := time.Now().UnixNano() / 1e6
	if timestamp < w.lastTimestamp {
		panic(fmt.Sprintf("Clock moved backwards. Refused to generate id for %d milliseconds", w.lastTimestamp-timestamp))
	}
	if timestamp == w.lastTimestamp {
		w.sequence = (w.sequence + 1) & sequenceMask
		if w.sequence == 0 {
			for timestamp <= w.lastTimestamp {
				timestamp = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.sequence = 0
	}
	w.lastTimestamp = timestamp
	id := (timestamp << (workerBits + sequenceBits)) | (w.workerId << sequenceBits) | w.sequence
	return id
}
