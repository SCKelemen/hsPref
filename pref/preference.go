package preference

import (
  "time"
)

type Preferences map[string]interface{}

func (p Preferences) Get(key string)interface{} {
  if p == nil {
    return nil
  }
  return p[key]
}

func (p Preferences) Set(key string, val interface{}) {
  p[key] = val
}

func (p Preferences) Rem(key string) {
  delete(p, key)
}
