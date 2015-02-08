package ghcal

import (
	"bytes"
	"fmt"
	"time"
)

func time2str(t time.Time) string {
	return t.Format("20060102T150405Z")
}

type Event struct {
	UID         string
	DTStamp     time.Time
	DTStart     time.Time
	Summary     *string
	Description *string
}

func (e *Event) ICalString() string {
	buf := bytes.Buffer{}
	buf.WriteString("BEGIN:VEVENT\r\n")
	buf.WriteString(fmt.Sprintf("DTSTAMP:%s\r\n", time2str(e.DTStamp)))
	buf.WriteString(fmt.Sprintf("UID:%s\r\n", e.UID))
	buf.WriteString(fmt.Sprintf("DTSTART:%s\r\n", time2str(e.DTStart)))
	if e.Summary != nil {
		buf.WriteString(fmt.Sprintf("SUMMARY:%s\r\n", *e.Summary))
	}
	if e.Description != nil {
		buf.WriteString(fmt.Sprintf("DESCRIPTION:%s\r\n", *e.Description))
	}
	buf.WriteString("END:VEVENT\r\n")
	return buf.String()
}

type Calendar struct {
	Prodid string
	Events []Event
}

func (c *Calendar) ICalString() string {
	buf := bytes.Buffer{}
	buf.WriteString("BEGIN:VCALENDAR\r\n")
	buf.WriteString("VERSION:2.0\r\n")
	buf.WriteString(fmt.Sprintf("PRODID:%s\r\n", c.Prodid))
	for _, ev := range c.Events {
		buf.WriteString(ev.ICalString())
	}
	buf.WriteString("END:VCALENDAR\r\n")
	return buf.String()
}
