/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
    iter *Iterator
    current int
    hasPeeked bool
}

func Constructor(iter *Iterator) *PeekingIterator {
    return &PeekingIterator{iter: iter, current: 0, hasPeeked: false}
}

func (this *PeekingIterator) hasNext() bool {
    if (this.hasPeeked) {
        return true
    }
    return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
    if (this.hasPeeked) {
        this.hasPeeked = false
        return this.current
    }
    return this.iter.next()
}

func (this *PeekingIterator) peek() int {
    if (this.hasPeeked) {
        return this.current
    }
    next := this.next()
    this.current = next
    this.hasPeeked = true
    return next
}
