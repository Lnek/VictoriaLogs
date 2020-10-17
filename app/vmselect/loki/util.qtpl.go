// Code generated by qtc from "util.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/loki/util.qtpl:1
package loki

//line app/vmselect/loki/util.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaLogs/lib/storage"
)

//line app/vmselect/loki/util.qtpl:7
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/loki/util.qtpl:7
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/loki/util.qtpl:7
func streammetricNameObject(qw422016 *qt422016.Writer, mn *storage.MetricName) {
//line app/vmselect/loki/util.qtpl:7
	qw422016.N().S(`{`)
//line app/vmselect/loki/util.qtpl:9
	if len(mn.MetricGroup) > 0 {
//line app/vmselect/loki/util.qtpl:9
		qw422016.N().S(`"__name__":`)
//line app/vmselect/loki/util.qtpl:10
		qw422016.N().QZ(mn.MetricGroup)
//line app/vmselect/loki/util.qtpl:10
		if len(mn.Tags) > 0 {
//line app/vmselect/loki/util.qtpl:10
			qw422016.N().S(`,`)
//line app/vmselect/loki/util.qtpl:10
		}
//line app/vmselect/loki/util.qtpl:11
	}
//line app/vmselect/loki/util.qtpl:12
	for j := range mn.Tags {
//line app/vmselect/loki/util.qtpl:13
		tag := &mn.Tags[j]

//line app/vmselect/loki/util.qtpl:14
		qw422016.N().QZ(tag.Key)
//line app/vmselect/loki/util.qtpl:14
		qw422016.N().S(`:`)
//line app/vmselect/loki/util.qtpl:14
		qw422016.N().QZ(tag.Value)
//line app/vmselect/loki/util.qtpl:14
		if j+1 < len(mn.Tags) {
//line app/vmselect/loki/util.qtpl:14
			qw422016.N().S(`,`)
//line app/vmselect/loki/util.qtpl:14
		}
//line app/vmselect/loki/util.qtpl:15
	}
//line app/vmselect/loki/util.qtpl:15
	qw422016.N().S(`}`)
//line app/vmselect/loki/util.qtpl:17
}

//line app/vmselect/loki/util.qtpl:17
func writemetricNameObject(qq422016 qtio422016.Writer, mn *storage.MetricName) {
//line app/vmselect/loki/util.qtpl:17
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/util.qtpl:17
	streammetricNameObject(qw422016, mn)
//line app/vmselect/loki/util.qtpl:17
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/util.qtpl:17
}

//line app/vmselect/loki/util.qtpl:17
func metricNameObject(mn *storage.MetricName) string {
//line app/vmselect/loki/util.qtpl:17
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/util.qtpl:17
	writemetricNameObject(qb422016, mn)
//line app/vmselect/loki/util.qtpl:17
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/util.qtpl:17
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/util.qtpl:17
	return qs422016
//line app/vmselect/loki/util.qtpl:17
}

//line app/vmselect/loki/util.qtpl:19
func streamvaluesWithTimestamps(qw422016 *qt422016.Writer, values []float64, timestamps []int64) {
//line app/vmselect/loki/util.qtpl:20
	if len(values) == 0 {
//line app/vmselect/loki/util.qtpl:20
		qw422016.N().S(`[]`)
//line app/vmselect/loki/util.qtpl:22
		return
//line app/vmselect/loki/util.qtpl:23
	}
//line app/vmselect/loki/util.qtpl:23
	qw422016.N().S(`[`)
//line app/vmselect/loki/util.qtpl:25
	/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/loki/util.qtpl:25
	qw422016.N().S(`[`)
//line app/vmselect/loki/util.qtpl:26
	qw422016.N().F(float64(timestamps[0]) / 1e3)
//line app/vmselect/loki/util.qtpl:26
	qw422016.N().S(`,"`)
//line app/vmselect/loki/util.qtpl:26
	qw422016.N().F(values[0])
//line app/vmselect/loki/util.qtpl:26
	qw422016.N().S(`"]`)
//line app/vmselect/loki/util.qtpl:28
	timestamps = timestamps[1:]
	values = values[1:]

//line app/vmselect/loki/util.qtpl:31
	if len(values) > 0 {
//line app/vmselect/loki/util.qtpl:33
		// Remove bounds check inside the loop below
		_ = timestamps[len(values)-1]

//line app/vmselect/loki/util.qtpl:36
		for i, v := range values {
//line app/vmselect/loki/util.qtpl:37
			/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/loki/util.qtpl:37
			qw422016.N().S(`,[`)
//line app/vmselect/loki/util.qtpl:38
			qw422016.N().F(float64(timestamps[i]) / 1e3)
//line app/vmselect/loki/util.qtpl:38
			qw422016.N().S(`,"`)
//line app/vmselect/loki/util.qtpl:38
			qw422016.N().F(v)
//line app/vmselect/loki/util.qtpl:38
			qw422016.N().S(`"]`)
//line app/vmselect/loki/util.qtpl:39
		}
//line app/vmselect/loki/util.qtpl:40
	}
//line app/vmselect/loki/util.qtpl:40
	qw422016.N().S(`]`)
//line app/vmselect/loki/util.qtpl:42
}

//line app/vmselect/loki/util.qtpl:42
func writevaluesWithTimestamps(qq422016 qtio422016.Writer, values []float64, timestamps []int64) {
//line app/vmselect/loki/util.qtpl:42
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/util.qtpl:42
	streamvaluesWithTimestamps(qw422016, values, timestamps)
//line app/vmselect/loki/util.qtpl:42
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/util.qtpl:42
}

//line app/vmselect/loki/util.qtpl:42
func valuesWithTimestamps(values []float64, timestamps []int64) string {
//line app/vmselect/loki/util.qtpl:42
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/util.qtpl:42
	writevaluesWithTimestamps(qb422016, values, timestamps)
//line app/vmselect/loki/util.qtpl:42
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/util.qtpl:42
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/util.qtpl:42
	return qs422016
//line app/vmselect/loki/util.qtpl:42
}

//line app/vmselect/loki/util.qtpl:44
func streamdatasWithTimestamps(qw422016 *qt422016.Writer, values [][]byte, timestamps []int64) {
//line app/vmselect/loki/util.qtpl:45
	if len(values) == 0 {
//line app/vmselect/loki/util.qtpl:45
		qw422016.N().S(`[]`)
//line app/vmselect/loki/util.qtpl:47
		return
//line app/vmselect/loki/util.qtpl:48
	}
//line app/vmselect/loki/util.qtpl:48
	qw422016.N().S(`[`)
//line app/vmselect/loki/util.qtpl:50
	/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/loki/util.qtpl:50
	qw422016.N().S(`["`)
//line app/vmselect/loki/util.qtpl:51
	qw422016.N().DL(timestamps[0] * 1e6)
//line app/vmselect/loki/util.qtpl:51
	qw422016.N().S(`",`)
//line app/vmselect/loki/util.qtpl:51
	qw422016.N().QZ(values[0])
//line app/vmselect/loki/util.qtpl:51
	qw422016.N().S(`]`)
//line app/vmselect/loki/util.qtpl:53
	timestamps = timestamps[1:]
	values = values[1:]

//line app/vmselect/loki/util.qtpl:56
	if len(values) > 0 {
//line app/vmselect/loki/util.qtpl:58
		// Remove bounds check inside the loop below
		_ = timestamps[len(values)-1]

//line app/vmselect/loki/util.qtpl:61
		for i, v := range values {
//line app/vmselect/loki/util.qtpl:62
			/* inline metricRow call here for the sake of performance optimization */

//line app/vmselect/loki/util.qtpl:62
			qw422016.N().S(`,["`)
//line app/vmselect/loki/util.qtpl:63
			qw422016.N().DL(timestamps[i] * 1e6)
//line app/vmselect/loki/util.qtpl:63
			qw422016.N().S(`",`)
//line app/vmselect/loki/util.qtpl:63
			qw422016.N().QZ(v)
//line app/vmselect/loki/util.qtpl:63
			qw422016.N().S(`]`)
//line app/vmselect/loki/util.qtpl:64
		}
//line app/vmselect/loki/util.qtpl:65
	}
//line app/vmselect/loki/util.qtpl:65
	qw422016.N().S(`]`)
//line app/vmselect/loki/util.qtpl:67
}

//line app/vmselect/loki/util.qtpl:67
func writedatasWithTimestamps(qq422016 qtio422016.Writer, values [][]byte, timestamps []int64) {
//line app/vmselect/loki/util.qtpl:67
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/util.qtpl:67
	streamdatasWithTimestamps(qw422016, values, timestamps)
//line app/vmselect/loki/util.qtpl:67
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/util.qtpl:67
}

//line app/vmselect/loki/util.qtpl:67
func datasWithTimestamps(values [][]byte, timestamps []int64) string {
//line app/vmselect/loki/util.qtpl:67
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/util.qtpl:67
	writedatasWithTimestamps(qb422016, values, timestamps)
//line app/vmselect/loki/util.qtpl:67
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/util.qtpl:67
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/util.qtpl:67
	return qs422016
//line app/vmselect/loki/util.qtpl:67
}
