package ui

import (
	"image/color"
	"math"
	"sync"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type CircularTimer struct {
	widget.BaseWidget
	diameter       float32
	strokeWidth    float32
	remainingTime  time.Duration
	totalDuration  time.Duration
	dotColor       color.Color
	circleColor    color.Color
	ticker         *time.Ticker
	tickerStopChan chan bool
	mutex          sync.Mutex
}

func NewCircularTimer(duration time.Duration, diameter, strokeWidth float32, circleColor, dotColor color.Color) *CircularTimer {
	t := &CircularTimer{
		diameter:      diameter,
		strokeWidth:   strokeWidth,
		totalDuration: duration,
		remainingTime: duration,
		dotColor:      dotColor,
		circleColor:   circleColor,
	}
	t.ExtendBaseWidget(t)
	return t
}

func (t *CircularTimer) Start() {
	if t.ticker != nil {
		return // Timer is already running
	}

	t.remainingTime = t.totalDuration
	t.ticker = time.NewTicker(1 * time.Second)
	t.tickerStopChan = make(chan bool)

	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.mutex.Lock()
				if t.remainingTime > 0 {
					t.remainingTime -= time.Second
					// Directly call Refresh; Fyne handles thread safety.
					t.Refresh() // Marks the widget as needing to be redrawn
				} else {
					t.Stop()
				}
				t.mutex.Unlock()
			case <-t.tickerStopChan:
				return
			}
		}
	}()
}

func (t *CircularTimer) Stop() {
	if t.ticker != nil {
		t.ticker.Stop()
		close(t.tickerStopChan)
		t.ticker = nil
		t.Refresh() // Also refresh on stop to ensure UI is updated
	}
}

func (t *CircularTimer) CreateRenderer() fyne.WidgetRenderer {
	t.ExtendBaseWidget(t)

	background := canvas.NewCircle(color.Transparent)
	background.StrokeColor = t.circleColor
	background.StrokeWidth = t.strokeWidth

	dot := canvas.NewCircle(t.dotColor)
	dot.Resize(fyne.NewSize(t.strokeWidth, t.strokeWidth))

	return &circularTimerRenderer{
		background: background,
		dot:        dot,
		timer:      t,
	}
}

type circularTimerRenderer struct {
	background *canvas.Circle
	dot        *canvas.Circle
	timer      *CircularTimer
}

func (r *circularTimerRenderer) Layout(size fyne.Size) {
	r.background.Resize(size)
	r.updateDotPosition()
}

func (r *circularTimerRenderer) MinSize() fyne.Size {
	return fyne.NewSize(r.timer.diameter, r.timer.diameter)
}

func (r *circularTimerRenderer) Refresh() {
	canvas.Refresh(r.background)
	r.updateDotPosition()
	canvas.Refresh(r.dot)
}

func (r *circularTimerRenderer) updateDotPosition() {
	progress := float64(r.timer.totalDuration-r.timer.remainingTime) / float64(r.timer.totalDuration)
	angle := progress * 2 * math.Pi
	x := float64(r.timer.diameter/2) + (float64(r.timer.diameter/2)-float64(r.timer.strokeWidth))*math.Cos(angle-math.Pi/2)
	y := float64(r.timer.diameter/2) + (float64(r.timer.diameter/2)-float64(r.timer.strokeWidth))*math.Sin(angle-math.Pi/2)
	r.dot.Move(fyne.NewPos(float32(x)-r.dot.Size().Width/2, float32(y)-r.dot.Size().Height/2))
}

func (r *circularTimerRenderer) BackgroundColor() color.Color {
	return color.Transparent
}

func (r *circularTimerRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.dot}
}

func (r *circularTimerRenderer) Destroy() {}

func (t *CircularTimer) Reset() {
	t.Stop()
	t.remainingTime = t.totalDuration
	t.Refresh()
}

func (t *CircularTimer) UpdateDuration(duration time.Duration) {
	t.Stop()
	t.totalDuration = duration
	t.remainingTime = duration
	t.Refresh()
}
