package ttime

import "time"

var (
  currentTime time.Time
  frozen bool
)

func Freeze(t time.Time) {
  currentTime = t
  frozen = true
}

func Unfreeze() {
  frozen = false
}

func IsFrozen() bool {
  return frozen
}

func Now() time.Time {
  if frozen {
    return currentTime
  } else {
    return time.Now()
  }
}

func After(d time.Duration) <-chan time.Time {
  if frozen {
    currentTime = currentTime.Add(d)
    c := make(chan time.Time, 1)
    c <- currentTime
    return c
  } else {
    return time.After(d)
  }
}

func Tick(d time.Duration) <-chan time.Time {
  if frozen {
    currentTime = currentTime.Add(d)
    c := make(chan time.Time, 1)
    go func() {
      for {
        c <- currentTime
      }
    }()
    return c
  } else {
    return time.Tick(d)
  }
}

func Sleep(d time.Duration) {
  if frozen && d > 0 {
    currentTime = currentTime.Add(d)
  } else {
    time.Sleep(d)
  }
}

