package events

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
	"unicode"
)

type Event interface {
	Name(Event) string
	Data(Event) string
}

type baseEvent struct{}

func (e *baseEvent) Name(evt Event) string {
	str := strings.TrimPrefix(reflect.TypeOf(evt).String(), "*events.")
	return str[:0] + string(unicode.ToLower(rune(str[0]))) + str[1:]
}

func (e *baseEvent) Data(evt Event) string {
	data, _ := json.Marshal(evt)
	return string(data)
}

type ScanStatus struct {
	baseEvent
	Scanning    bool  `json:"scanning"`
	Count       int64 `json:"count"`
	FolderCount int64 `json:"folderCount"`
}

type KeepAlive struct {
	baseEvent
	TS int64 `json:"ts"`
}

type ServerStart struct {
	baseEvent
	StartTime time.Time `json:"startTime"`
}

const Any = "*"

type RefreshResource struct {
	baseEvent
	resources map[string][]string
}

func (rr *RefreshResource) With(resource string, ids ...string) *RefreshResource {
	if rr.resources == nil {
		rr.resources = make(map[string][]string)
	}
	if len(ids) == 0 {
		rr.resources[resource] = append(rr.resources[resource], Any)
	}
	for i := range ids {
		rr.resources[resource] = append(rr.resources[resource], ids[i])
	}
	return rr
}

func (rr *RefreshResource) Data(evt Event) string {
	if rr.resources == nil {
		return `{"*":"*"}`
	}
	r := evt.(*RefreshResource)
	data, _ := json.Marshal(r.resources)
	return string(data)
}
