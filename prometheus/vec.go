func (m *metricMap) delete(hash uint64, labels []string) bool {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	variants := m.children[hash]
	for i, v := range variants {
		if slices.Equal(v.labelValues, labels) {
			// Zero out the reference to allow GC
			copy(variants[i:], variants[i+1:])
			variants[len(variants)-1] = metricWithLabelValues{}
			m.children[hash] = variants[:len(variants)-1]
			return true
		}
	}
	return false
}