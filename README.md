This is an experiment in making time easier to mock in Go tests.

Unfortunately the mock bleeds slightly in to the non-test code by use of the ttime library, instead of time itself,
but I think this is a reasonable trade-off. All other code is identical, and it doesn't require adopting some new Time-like struct.

All methods return actual time.Time structs (if they were supposed to).

Example code coming.