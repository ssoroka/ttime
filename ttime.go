package ttime

import "time"

var (
  currentTime time.Time
  timeFrozen  bool
)

type Time struct {
  time.Time
}

func Freeze(t time.Time) {
  currentTime = t
  timeFrozen = true
}

func Unfreeze() {
  timeFrozen = false
}

func IsFrozen() bool {
  return timeFrozen
}

func Now() time.Time {
  if timeFrozen {
    return currentTime
  } else {
    return time.Now()
  }
}

func After(d time.Duration) <-chan time.Time {
  if timeFrozen {
    currentTime = currentTime.Add(d)
    c := make(chan time.Time, 1)
    c <- currentTime
    return c
  } else {
    return time.After(d)
  }
}

func Tick(d time.Duration) <-chan time.Time {
  if timeFrozen {
    c := make(chan time.Time, 1)
    go func() {
      for {
        currentTime = currentTime.Add(d)
        c <- currentTime
      }
    }()
    return c
  } else {
    return time.Tick(d)
  }
}

func Sleep(d time.Duration) {
  if timeFrozen && d > 0 {
    currentTime = currentTime.Add(d)
  } else {
    time.Sleep(d)
  }
}
