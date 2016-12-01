package WorkerPool

import (
	"errors"
	"time"
)

const (
	ERR_WORKERPOOLSHUTDOWN = "The worker pool is shutting down and wont take new assignments"
	ERR_TIMEDOUTREQUSTSLOT = "The request for a free execution slot has been timed out"
)

// Represents simple assignment
type WorkItem func()

// Represents simple workers pool operating on Projects and WorkITems
type WorkerPool interface {

	// Synchronously requests worker slot and end exectes/does WorkItem in parallel routine as soon as slot is obtained
	// If no slot aquired upon timeOut exceeds - returns ERR_TIMEDOUTREQUSTSLOT
	Do(wi WorkItem, timeOut time.Duration) error

	// Closes the worker pool
	Close()
}

// Private custom  implementation of worker pool
type workerPool struct {
	isShuttingDown   bool
	workersChan      chan struct{}
	cancellationChan chan struct{}
}

// A new Deapertment Factory
func NewWorkerPool(initWorkerNumber int) WorkerPool {
	// instantiate  pool
	workersChan := make(chan struct{}, initWorkerNumber)

	// fill up pool
	// for each initially empty slot we shoul put one value
	for i := 0; i < initWorkerNumber; i++ {
		workersChan <- struct{}{}
	}

	// chan to notify processes on closing pool
	cancellationChan := make(chan struct{})

	return &workerPool{
		isShuttingDown:   false,
		workersChan:      workersChan,
		cancellationChan: cancellationChan,
	}

}

// Implements WorkerPool.Do(wi WorkItem) method
func (this *workerPool) Do(wi WorkItem, timeOut time.Duration) error {

	if this.isShuttingDown {
		return errors.New(ERR_WORKERPOOLSHUTDOWN)
	}

	t := time.NewTimer(timeOut)

	select {
	case _ = <-this.workersChan:
		if !t.Stop() {
			<-t.C
		}
	// The channel will be only closed
	case _ = <-this.cancellationChan:
		if !t.Stop() {
			<-t.C
		}
		return errors.New(ERR_WORKERPOOLSHUTDOWN)
	case _ = <-t.C:
		return errors.New(ERR_TIMEDOUTREQUSTSLOT)
	}

	go func() {
		defer this.releaseSlot()
		wi()
	}()

	return nil
}

func (this *workerPool) Close() {
	this.isShuttingDown = true

	// notify those who are actually waiting about closing
	close(this.cancellationChan)

	// wait while all left assignments are done
	for i := 0; i < cap(this.workersChan); i++ {
		<-this.workersChan
	}

	close(this.workersChan)

}

func (this *workerPool) releaseSlot() {
	this.workersChan <- struct{}{}
}
