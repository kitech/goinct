package runtime

// return os thread no, 0+, not really os thread id
func Gettno() int32 {
	// struct runtime.p means process
	return getg().m.p.ptr().id
}

// this is really pthread_self() value
func Gettid() uint64 {
	// struct runtime.p means process
	return getg().m.procid
}
