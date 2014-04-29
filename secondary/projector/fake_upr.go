package projector

import (
	c "github.com/couchbase/indexing/secondary/common"
)

// FakeBucket fot unit testing.
type FakeBucket struct {
	bucket  string
	vbmap   map[string][]uint16
	flogs   map[uint16][][2]uint64
	C       chan *MutationEvent
	streams map[uint16]*FakeStream
}

// FakeStream fot unit testing.
type FakeStream struct {
	seqno  uint64
	vbuuid uint64
	killch chan bool
}

// NewFakeBuckets returns a reference to new FakeBucket.
func NewFakeBuckets(buckets []string) map[string]*FakeBucket {
	fakebuckets := make(map[string]*FakeBucket)
	for _, bucket := range buckets {
		fakebuckets[bucket] = &FakeBucket{
			bucket:  bucket,
			vbmap:   make(map[string][]uint16),
			flogs:   make(map[uint16][][2]uint64),
			C:       make(chan *MutationEvent, c.MutationChannelSize),
			streams: make(map[uint16]*FakeStream),
		}
	}
	return fakebuckets
}

// BucketAccess interface

// GetVBmap is method receiver for BucketAccess interface
func (b *FakeBucket) GetVBmap(kvaddrs []string) (map[string][]uint16, error) {
	m := make(map[string][]uint16)
	for kvaddr, vbnos := range b.vbmap {
		m[kvaddr] = vbnos
	}
	return m, nil
}

// GetFailoverLog is method receiver for BucketAccess interface
func (b *FakeBucket) GetFailoverLog(vbno uint16) (flog [][2]uint64, err error) {
	return b.flogs[vbno], nil
}

// OpenKVFeed is method receiver for BucketAccess interface
func (b *FakeBucket) OpenKVFeed(kvaddr string) (KVFeeder, error) {
	return b, nil
}

// Close is method receiver for BucketAccess interface
func (b *FakeBucket) Close(kvaddr string) {
	close(b.C)
}

// SetVbmap fake initialization method.
func (b *FakeBucket) SetVbmap(kvaddr string, vbnos []uint16) {
	b.vbmap[kvaddr] = vbnos
}

// SetFailoverLog fake initialization method.
func (b *FakeBucket) SetFailoverLog(vbno uint16, flog [][2]uint64) {
	b.flogs[vbno] = flog
}

// KVFeeder interface

// GetChannel is method receiver for KVFeeder interface
func (b *FakeBucket) GetChannel() <-chan *MutationEvent {
	return b.C
}

// StartVbStreams is method receiver for KVFeeder interface
func (b *FakeBucket) StartVbStreams(restartTs *c.Timestamp) (failoverTs, kvTs *c.Timestamp, err error) {
	for i, vbno := range restartTs.Vbnos {
		if stream, ok := b.streams[vbno]; ok {
			close(stream.killch)
		}
		stream := &FakeStream{
			seqno:  restartTs.Seqnos[i],
			vbuuid: restartTs.Vbuuids[i],
			killch: make(chan bool),
		}
		b.streams[vbno] = stream
		go stream.run(b.C)
	}
	return restartTs, restartTs, nil
}

// EndVbStreams is method receiver for KVFeeder interface
func (b *FakeBucket) EndVbStreams(endTs *c.Timestamp) (err error) {
	for _, vbno := range endTs.Vbnos {
		if stream, ok := b.streams[vbno]; ok {
			close(stream.killch)
			delete(b.streams, vbno)
		}
	}
	return
}

// CloseKVFeed is method receiver for KVFeeder interface
func (b *FakeBucket) CloseKVFeed() (err error) {
	for _, stream := range b.streams {
		close(stream.killch)
	}
	return
}

func (s *FakeStream) run(mutch chan<- *MutationEvent) {
	// TODO: generate mutation events
}