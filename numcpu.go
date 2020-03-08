// Package runtime counts the number of threads currently available
// to your process, without caching.
//
// Handy to dynamically adjust load in response to changing circumstances,
// such as hot-plugging of cpus or a changed cpu affinity.
package runtime // import "blitznote.com/src/go.runtime"

// NumCPU queries the operating system for the count of threads available
// for use to this process, without caching.
//
// Please note that you will still need to examine the underlying cpu topology,
// i. e. to learn what datastructure will perform the best.
// E. g., four available threads on a four-socket 32-core machine might span
// caches that are not shared.
//
// Returns 0 on errors (extremely unlikely); try |runtime.NumCPU| in that case.
func NumCPU() int {
	return int(numcpu())
}
