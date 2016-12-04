package WorkerPool

import (
	"errors"
	"time"
)

const (
	ERR_WORKERPOOLSHUTDOWN = "The worker pool is shutting down and wont take new assignments"
	ERR_TIMEDOUTREQUSTSLOT = "The request for a free execution slot has been timed out"
)

// Represents simple assignment.
type WorkItem func()

// Represents simple workers pool operating on Projects and WorkItems.
type WorkerPool interface {

	// Synchronously requests worker slot and end exectes/does WorkItem in parallel routine as soon as slot is obtained.
	// If no slot aquired upon timeOut exceeds - returns ERR_TIMEDOUTREQUSTSLOT.
	// If worker pool is already closed or closing while obtaining worker slot - return ERR_WORKERPOOLSHUTDOWN.
	Do(wi WorkItem, timeOut time.Duration) error

	// Closes the worker pool.
	// All new requests will be rejected returning ERR_WORKERPOOLSHUTDOWN.
	// All requests waiting for slots should be notified and canceled returning ERR_WORKERPOOLSHUTDOWN.
	// Processes already obtained their own slot shouldn't be affected and complete normal.
	Close()
}

// Private custom  implementation of worker pool.
type workerPool struct {
	workersChan      chan struct{}
	cancellationChan chan struct{}
}

// A new WorkerPool Factory.
func NewWorkerPool(initWorkerNumber int) WorkerPool {
	// instantiate  pool
	workersChan := make(chan struct{}, initWorkerNumber)

	// fill up pool
	// for each initially empty slot we should put one value
	for i := 0; i < initWorkerNumber; i++ {
		workersChan <- struct{}{}
	}

	// chan to notify processes on closing pool
	cancellationChan := make(chan struct{})

	return &workerPool{
		workersChan:      workersChan,
		cancellationChan: cancellationChan,
	}

}

// Implements WorkerPool.Do(wi WorkItem) method
func (this *workerPool) Do(wi WorkItem, timeOut time.Duration) error {

	select {
	case _ = <-this.workersChan:
	// The channel will be only closed
	case _ = <-this.cancellationChan:
		return errors.New(ERR_WORKERPOOLSHUTDOWN)
	case _ = <-time.After(timeOut):
		return errors.New(ERR_TIMEDOUTREQUSTSLOT)
	}

	go func() {
		defer this.releaseSlot()
		wi()
	}()

	return nil
}

func (this *workerPool) Close() {
	// notify those who are actually waiting about closing
	close(this.cancellationChan)

	// wait while all left assignments are done, and drain slots
	for i := 0; i < cap(this.workersChan); i++ {
		<-this.workersChan
	}

	close(this.workersChan)

}

func (this *workerPool) releaseSlot() {
	// put slot back into pool
	this.workersChan <- struct{}{}
}
