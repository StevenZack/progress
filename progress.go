package progress

import "time"

type Progress struct {
	total int
	index int
	start time.Time
}

func New(total int) *Progress {
	return &Progress{
		total: total,
	}
}

func (p *Progress) Start() {
	p.start = time.Now()
}

func (p *Progress) End() (dur time.Duration, remains time.Duration) {
	dur = time.Now().Sub(p.start)
	remains = time.Duration(int(dur) * (p.total - p.index))
	p.index++
	return
}

