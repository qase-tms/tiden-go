package gotest

import "fmt"

// boundedBuffer keeps the head and tail of an unbounded write stream within
// a fixed byte budget. Failures usually matter at the edges: the first error
// and the final panic both survive truncation.
type boundedBuffer struct {
	max     int
	head    []byte
	tail    []byte // ring once head is full
	tailOff int
	dropped int
}

func newBoundedBuffer(max int) *boundedBuffer {
	if max < 64 {
		max = 64
	}
	return &boundedBuffer{max: max}
}

func (b *boundedBuffer) WriteString(s string) {
	headCap := b.max / 2
	data := []byte(s)
	if len(b.head) < headCap {
		take := min(headCap-len(b.head), len(data))
		b.head = append(b.head, data[:take]...)
		data = data[take:]
	}
	if len(data) == 0 {
		return
	}
	tailCap := b.max - b.max/2
	if b.tail == nil {
		b.tail = make([]byte, 0, tailCap)
	}
	for _, c := range data {
		if len(b.tail) < tailCap {
			b.tail = append(b.tail, c)
		} else {
			b.dropped++
			b.tail[b.tailOff] = c
			b.tailOff = (b.tailOff + 1) % tailCap
		}
	}
}

// String assembles the buffered content and reports whether bytes were lost.
func (b *boundedBuffer) String() (string, bool) {
	if b.dropped == 0 {
		return string(b.head) + string(b.tail), false
	}
	marker := fmt.Sprintf("\n... [%d bytes truncated] ...\n", b.dropped)
	return string(b.head) + marker + string(b.tail[b.tailOff:]) + string(b.tail[:b.tailOff]), true
}
