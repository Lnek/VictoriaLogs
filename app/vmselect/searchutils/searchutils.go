package searchutils

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/VictoriaMetrics/VictoriaLogs/lib/logql"
	"github.com/VictoriaMetrics/VictoriaMetrics/lib/fasttime"
)

var (
	maxExportDuration   = flag.Duration("search.maxExportDuration", time.Hour*24*30, "The maximum duration for /api/v1/export call")
	maxQueryDuration    = flag.Duration("search.maxQueryDuration", time.Second*30, "The maximum duration for query execution; see also -search.storageTimeout")
	denyPartialResponse = flag.Bool("search.denyPartialResponse", false, "Whether to deny partial responses if a part of -storageNode instances fail to perform queries; "+
		"this trades availability over consistency; see also -search.maxQueryDuration and -search.storageTimeout")

	// StorageTimeout limits the duration of query execution on every vmstorage node.
	// See https://github.com/VictoriaMetrics/VictoriaMetrics/issues/711
	StorageTimeout = flag.Duration("search.storageTimeout", 0, "The timeout for per-storage query processing; "+
		"this allows returning partial responses if certain -storageNode instances slowly process the query; "+
		"see also -search.maxQueryDuration and -search.denyPartialResponse command-line flags")
)

func roundToSeconds(ms int64) int64 {
	return ms - ms%1000
}

// GetTime returns time from the given argKey query arg.
//
// If argKey is missing in r, then defaultMs rounded to seconds is returned.
// The rounding is needed in order to align query results in Grafana
// executed at different times. See https://github.com/VictoriaMetrics/VictoriaMetrics/issues/720
func GetTime(r *http.Request, argKey string, defaultMs int64) (int64, error) {
	argValue := r.FormValue(argKey)
	if len(argValue) == 0 {
		return roundToSeconds(defaultMs), nil
	}
	secs, err := strconv.ParseFloat(argValue, 64)
	if err != nil {
		// Try parsing string format
		t, err := time.Parse(time.RFC3339, argValue)
		if err != nil {
			// Handle Prometheus'-provided minTime and maxTime.
			// See https://github.com/prometheus/client_golang/issues/614
			switch argValue {
			case prometheusMinTimeFormatted:
				return minTimeMsecs, nil
			case prometheusMaxTimeFormatted:
				return maxTimeMsecs, nil
			}
			// Try parsing duration relative to the current time
			d, err1 := logql.DurationValue(argValue, 0)
			if err1 != nil {
				return 0, fmt.Errorf("cannot parse %q=%q: %w", argKey, argValue, err)
			}
			if d > 0 {
				d = -d
			}
			t = time.Now().Add(time.Duration(d) * time.Millisecond)
		}
		secs = float64(t.UnixNano()) / 1e9
	}
	for secs > math.MaxInt32 {
		secs = secs / 1e3
	}
	msecs := int64(secs * 1e3)
	if msecs < minTimeMsecs {
		msecs = 0
	}
	if msecs > maxTimeMsecs {
		msecs = maxTimeMsecs
	}
	return msecs, nil
}

var (
	// These constants were obtained from https://github.com/prometheus/prometheus/blob/91d7175eaac18b00e370965f3a8186cc40bf9f55/web/api/v1/api.go#L442
	// See https://github.com/prometheus/client_golang/issues/614 for details.
	prometheusMinTimeFormatted = time.Unix(math.MinInt64/1000+62135596801, 0).UTC().Format(time.RFC3339Nano)
	prometheusMaxTimeFormatted = time.Unix(math.MaxInt64/1000-62135596801, 999999999).UTC().Format(time.RFC3339Nano)
)

const (
	// These values prevent from overflow when storing msec-precision time in int64.
	minTimeMsecs = 0 // use 0 instead of `int64(-1<<63) / 1e6` because the storage engine doesn't actually support negative time
	maxTimeMsecs = int64(1<<63-1) / 1e6
)

// GetDuration returns duration from the given argKey query arg.
func GetDuration(r *http.Request, argKey string, defaultValue int64) (int64, error) {
	argValue := r.FormValue(argKey)
	if len(argValue) == 0 {
		return defaultValue, nil
	}
	secs, err := strconv.ParseFloat(argValue, 64)
	if err != nil {
		// Try parsing string format
		d, err := logql.DurationValue(argValue, 0)
		if err != nil {
			return 0, fmt.Errorf("cannot parse %q=%q: %w", argKey, argValue, err)
		}
		secs = float64(d) / 1000
	}
	msecs := int64(secs * 1e3)
	if msecs <= 0 || msecs > maxDurationMsecs {
		return 0, fmt.Errorf("%q=%dms is out of allowed range [%d ... %d]", argKey, msecs, 0, int64(maxDurationMsecs))
	}
	return msecs, nil
}

const maxDurationMsecs = 100 * 365 * 24 * 3600 * 1000

// GetMaxQueryDuration returns the maximum duration for query from r.
func GetMaxQueryDuration(r *http.Request) time.Duration {
	dms, err := GetDuration(r, "timeout", 0)
	if err != nil {
		dms = 0
	}
	d := time.Duration(dms) * time.Millisecond
	if d <= 0 || d > *maxQueryDuration {
		d = *maxQueryDuration
	}
	if *StorageTimeout > 0 && d > *StorageTimeout {
		d = *StorageTimeout
	}
	return d
}

// GetDeadlineForQuery returns deadline for the given query r.
func GetDeadlineForQuery(r *http.Request, startTime time.Time) Deadline {
	dMax := maxQueryDuration.Milliseconds()
	return getDeadlineWithMaxDuration(r, startTime, dMax, "-search.maxQueryDuration")
}

// GetDeadlineForExport returns deadline for the given request to /api/v1/export.
func GetDeadlineForExport(r *http.Request, startTime time.Time) Deadline {
	dMax := maxExportDuration.Milliseconds()
	return getDeadlineWithMaxDuration(r, startTime, dMax, "-search.maxExportDuration")
}

func getDeadlineWithMaxDuration(r *http.Request, startTime time.Time, dMax int64, flagHint string) Deadline {
	d, err := GetDuration(r, "timeout", 0)
	if err != nil {
		d = 0
	}
	if d <= 0 || d > dMax {
		d = dMax
	}
	timeout := time.Duration(d) * time.Millisecond
	return NewDeadline(startTime, timeout, flagHint)
}

// GetBool returns boolean value from the given argKey query arg.
func GetBool(r *http.Request, argKey string) bool {
	argValue := r.FormValue(argKey)
	switch strings.ToLower(argValue) {
	case "", "0", "f", "false", "no":
		return false
	default:
		return true
	}
}

// GetInt64 returns int value from the given argKey query arg.
func GetInt64(r *http.Request, argKey string, defaultValue int64) (int64, error) {
	argValue := r.FormValue(argKey)
	if len(argValue) == 0 {
		return defaultValue, nil
	}
	return strconv.ParseInt(argValue, 10, 64)
}

// GetString returns string value from the given argKey query arg.
func GetString(r *http.Request, argKey string, defaultValue string) string {
	argValue := r.FormValue(argKey)
	if len(argValue) == 0 {
		return defaultValue
	}
	return argValue
}

// GetDenyPartialResponse returns whether partial responses are denied.
func GetDenyPartialResponse(r *http.Request) bool {
	if *denyPartialResponse {
		return true
	}
	return GetBool(r, "deny_partial_response")
}

// Deadline contains deadline with the corresponding timeout for pretty error messages.
type Deadline struct {
	deadline uint64

	timeout  time.Duration
	flagHint string
}

// NewDeadline returns deadline for the given timeout.
//
// flagHint must contain a hit for command-line flag, which could be used
// in order to increase timeout.
func NewDeadline(startTime time.Time, timeout time.Duration, flagHint string) Deadline {
	return Deadline{
		deadline: uint64(startTime.Add(timeout).Unix()),
		timeout:  timeout,
		flagHint: flagHint,
	}
}

// Exceeded returns true if deadline is exceeded.
func (d *Deadline) Exceeded() bool {
	return fasttime.UnixTimestamp() > d.deadline
}

// Deadline returns deadline in unix timestamp seconds.
func (d *Deadline) Deadline() uint64 {
	return d.deadline
}

// String returns human-readable string representation for d.
func (d *Deadline) String() string {
	return fmt.Sprintf("%.3f seconds; the timeout can be adjusted with `%s` command-line flag", d.timeout.Seconds(), d.flagHint)
}
