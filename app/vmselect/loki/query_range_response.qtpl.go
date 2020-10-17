// Code generated by qtc from "query_range_response.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line app/vmselect/loki/query_range_response.qtpl:1
package loki

//line app/vmselect/loki/query_range_response.qtpl:1
import (
	"github.com/VictoriaMetrics/VictoriaLogs/app/vmselect/netstorage"
)

// QueryRangeResponse generates response for /api/v1/query_range.See https://prometheus.io/docs/prometheus/latest/querying/api/#range-queries

//line app/vmselect/loki/query_range_response.qtpl:8
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line app/vmselect/loki/query_range_response.qtpl:8
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line app/vmselect/loki/query_range_response.qtpl:8
func StreamVectorQueryRangeResponse(qw422016 *qt422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:8
	qw422016.N().S(`{"status":"success","data":{"resultType":"matrix","result":[`)
//line app/vmselect/loki/query_range_response.qtpl:14
	if len(rs) > 0 {
//line app/vmselect/loki/query_range_response.qtpl:15
		streamvectorQueryRangeLine(qw422016, &rs[0])
//line app/vmselect/loki/query_range_response.qtpl:16
		rs = rs[1:]

//line app/vmselect/loki/query_range_response.qtpl:17
		for i := range rs {
//line app/vmselect/loki/query_range_response.qtpl:17
			qw422016.N().S(`,`)
//line app/vmselect/loki/query_range_response.qtpl:18
			streamvectorQueryRangeLine(qw422016, &rs[i])
//line app/vmselect/loki/query_range_response.qtpl:19
		}
//line app/vmselect/loki/query_range_response.qtpl:20
	}
//line app/vmselect/loki/query_range_response.qtpl:20
	qw422016.N().S(`]}}`)
//line app/vmselect/loki/query_range_response.qtpl:24
}

//line app/vmselect/loki/query_range_response.qtpl:24
func WriteVectorQueryRangeResponse(qq422016 qtio422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:24
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/query_range_response.qtpl:24
	StreamVectorQueryRangeResponse(qw422016, rs)
//line app/vmselect/loki/query_range_response.qtpl:24
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/query_range_response.qtpl:24
}

//line app/vmselect/loki/query_range_response.qtpl:24
func VectorQueryRangeResponse(rs []netstorage.Result) string {
//line app/vmselect/loki/query_range_response.qtpl:24
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/query_range_response.qtpl:24
	WriteVectorQueryRangeResponse(qb422016, rs)
//line app/vmselect/loki/query_range_response.qtpl:24
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/query_range_response.qtpl:24
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/query_range_response.qtpl:24
	return qs422016
//line app/vmselect/loki/query_range_response.qtpl:24
}

//line app/vmselect/loki/query_range_response.qtpl:26
func streamvectorQueryRangeLine(qw422016 *qt422016.Writer, r *netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:26
	qw422016.N().S(`{"metric":`)
//line app/vmselect/loki/query_range_response.qtpl:28
	streammetricNameObject(qw422016, &r.MetricName)
//line app/vmselect/loki/query_range_response.qtpl:28
	qw422016.N().S(`,"values":`)
//line app/vmselect/loki/query_range_response.qtpl:29
	streamvaluesWithTimestamps(qw422016, r.Values, r.Timestamps)
//line app/vmselect/loki/query_range_response.qtpl:29
	qw422016.N().S(`}`)
//line app/vmselect/loki/query_range_response.qtpl:31
}

//line app/vmselect/loki/query_range_response.qtpl:31
func writevectorQueryRangeLine(qq422016 qtio422016.Writer, r *netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:31
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/query_range_response.qtpl:31
	streamvectorQueryRangeLine(qw422016, r)
//line app/vmselect/loki/query_range_response.qtpl:31
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/query_range_response.qtpl:31
}

//line app/vmselect/loki/query_range_response.qtpl:31
func vectorQueryRangeLine(r *netstorage.Result) string {
//line app/vmselect/loki/query_range_response.qtpl:31
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/query_range_response.qtpl:31
	writevectorQueryRangeLine(qb422016, r)
//line app/vmselect/loki/query_range_response.qtpl:31
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/query_range_response.qtpl:31
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/query_range_response.qtpl:31
	return qs422016
//line app/vmselect/loki/query_range_response.qtpl:31
}

//line app/vmselect/loki/query_range_response.qtpl:33
func StreamStreamsQueryRangeResponse(qw422016 *qt422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:33
	qw422016.N().S(`{"status":"success","data":{"resultType":"streams","result":[`)
//line app/vmselect/loki/query_range_response.qtpl:39
	if len(rs) > 0 {
//line app/vmselect/loki/query_range_response.qtpl:40
		streamstreamsQueryRangeLine(qw422016, &rs[0])
//line app/vmselect/loki/query_range_response.qtpl:41
		rs = rs[1:]

//line app/vmselect/loki/query_range_response.qtpl:42
		for i := range rs {
//line app/vmselect/loki/query_range_response.qtpl:42
			qw422016.N().S(`,`)
//line app/vmselect/loki/query_range_response.qtpl:43
			streamstreamsQueryRangeLine(qw422016, &rs[i])
//line app/vmselect/loki/query_range_response.qtpl:44
		}
//line app/vmselect/loki/query_range_response.qtpl:45
	}
//line app/vmselect/loki/query_range_response.qtpl:45
	qw422016.N().S(`]}}`)
//line app/vmselect/loki/query_range_response.qtpl:49
}

//line app/vmselect/loki/query_range_response.qtpl:49
func WriteStreamsQueryRangeResponse(qq422016 qtio422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:49
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/query_range_response.qtpl:49
	StreamStreamsQueryRangeResponse(qw422016, rs)
//line app/vmselect/loki/query_range_response.qtpl:49
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/query_range_response.qtpl:49
}

//line app/vmselect/loki/query_range_response.qtpl:49
func StreamsQueryRangeResponse(rs []netstorage.Result) string {
//line app/vmselect/loki/query_range_response.qtpl:49
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/query_range_response.qtpl:49
	WriteStreamsQueryRangeResponse(qb422016, rs)
//line app/vmselect/loki/query_range_response.qtpl:49
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/query_range_response.qtpl:49
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/query_range_response.qtpl:49
	return qs422016
//line app/vmselect/loki/query_range_response.qtpl:49
}

//line app/vmselect/loki/query_range_response.qtpl:51
func StreamTailQueryRangeResponse(qw422016 *qt422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:51
	qw422016.N().S(`{"streams":[`)
//line app/vmselect/loki/query_range_response.qtpl:54
	if len(rs) > 0 {
//line app/vmselect/loki/query_range_response.qtpl:55
		streamstreamsQueryRangeLine(qw422016, &rs[0])
//line app/vmselect/loki/query_range_response.qtpl:56
		rs = rs[1:]

//line app/vmselect/loki/query_range_response.qtpl:57
		for i := range rs {
//line app/vmselect/loki/query_range_response.qtpl:57
			qw422016.N().S(`,`)
//line app/vmselect/loki/query_range_response.qtpl:58
			streamstreamsQueryRangeLine(qw422016, &rs[i])
//line app/vmselect/loki/query_range_response.qtpl:59
		}
//line app/vmselect/loki/query_range_response.qtpl:60
	}
//line app/vmselect/loki/query_range_response.qtpl:60
	qw422016.N().S(`]}`)
//line app/vmselect/loki/query_range_response.qtpl:63
}

//line app/vmselect/loki/query_range_response.qtpl:63
func WriteTailQueryRangeResponse(qq422016 qtio422016.Writer, rs []netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/query_range_response.qtpl:63
	StreamTailQueryRangeResponse(qw422016, rs)
//line app/vmselect/loki/query_range_response.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/query_range_response.qtpl:63
}

//line app/vmselect/loki/query_range_response.qtpl:63
func TailQueryRangeResponse(rs []netstorage.Result) string {
//line app/vmselect/loki/query_range_response.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/query_range_response.qtpl:63
	WriteTailQueryRangeResponse(qb422016, rs)
//line app/vmselect/loki/query_range_response.qtpl:63
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/query_range_response.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/query_range_response.qtpl:63
	return qs422016
//line app/vmselect/loki/query_range_response.qtpl:63
}

//line app/vmselect/loki/query_range_response.qtpl:65
func streamstreamsQueryRangeLine(qw422016 *qt422016.Writer, r *netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:65
	qw422016.N().S(`{"stream":`)
//line app/vmselect/loki/query_range_response.qtpl:67
	streammetricNameObject(qw422016, &r.MetricName)
//line app/vmselect/loki/query_range_response.qtpl:67
	qw422016.N().S(`,"values":`)
//line app/vmselect/loki/query_range_response.qtpl:68
	streamdatasWithTimestamps(qw422016, r.Datas, r.Timestamps)
//line app/vmselect/loki/query_range_response.qtpl:68
	qw422016.N().S(`}`)
//line app/vmselect/loki/query_range_response.qtpl:70
}

//line app/vmselect/loki/query_range_response.qtpl:70
func writestreamsQueryRangeLine(qq422016 qtio422016.Writer, r *netstorage.Result) {
//line app/vmselect/loki/query_range_response.qtpl:70
	qw422016 := qt422016.AcquireWriter(qq422016)
//line app/vmselect/loki/query_range_response.qtpl:70
	streamstreamsQueryRangeLine(qw422016, r)
//line app/vmselect/loki/query_range_response.qtpl:70
	qt422016.ReleaseWriter(qw422016)
//line app/vmselect/loki/query_range_response.qtpl:70
}

//line app/vmselect/loki/query_range_response.qtpl:70
func streamsQueryRangeLine(r *netstorage.Result) string {
//line app/vmselect/loki/query_range_response.qtpl:70
	qb422016 := qt422016.AcquireByteBuffer()
//line app/vmselect/loki/query_range_response.qtpl:70
	writestreamsQueryRangeLine(qb422016, r)
//line app/vmselect/loki/query_range_response.qtpl:70
	qs422016 := string(qb422016.B)
//line app/vmselect/loki/query_range_response.qtpl:70
	qt422016.ReleaseByteBuffer(qb422016)
//line app/vmselect/loki/query_range_response.qtpl:70
	return qs422016
//line app/vmselect/loki/query_range_response.qtpl:70
}
