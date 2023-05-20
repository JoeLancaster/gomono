# About
Handwritten Go assembly to measure monotonic time via the [VDSO clock_gettime](https://man7.org/linux/man-pages/man7/vdso.7.html) call.

About twice as fast as `time.Now()`

# Requirements
Linux Kernel >= 2.6

64-bit architecture
