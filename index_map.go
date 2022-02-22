package delta

const (
	// MatchLimit specifies the maximum number of positions tracked
	// for each unique key in the map of source data. See makeMap().
	MatchLimit = 50

	// MatchSize specifies the size of unique
	// chunks being searched for, in bytes.
	MatchSize = 9
)

type chunk [MatchSize]byte

// indexMap _ _
type indexMap struct {
	m map[chunk][]int
} //                                                                    indexMap

// makeMap creates a map of unique chunks in 'data'.
// The key specifies the unique chunk of bytes, while the
// values array returns the positions of the chunk in 'data'.
func makeMap(data []byte) indexMap {
	lenData := len(data)
	if lenData < MatchSize {
		return indexMap{m: map[chunk][]int{}}
	}
	ret := indexMap{m: make(map[chunk][]int, lenData/4)}
	var key chunk
	lenData -= MatchSize
	for i := 0; i < lenData; {
		copy(key[:], data[i:])
		ar, found := ret.m[key]
		if !found {
			ret.m[key] = []int{i}
			i++
			continue
		}
		if len(ar) >= MatchLimit {
			i++
			continue
		}
		ret.m[key] = append(ret.m[key], i)
		i += MatchSize
	}
	return ret
} //                                                                     makeMap

// get _ _
func (ob *indexMap) get(key chunk) (locs []int, found bool) {
	locs, found = ob.m[key]
	return
} //                                                                         get

// end
