package fastjson

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bet365/jingo"
	jsoniter "github.com/json-iterator/go"
	jwriter "github.com/mailru/easyjson/jwriter"
)

var obj = V{S: "s", B: true, I: 233, M: map[string]string{"test": "test"}}

func BenchmarkFmtString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf(`"S": "%s", "B": %v, "I": %d`, obj.S, obj.B, obj.I)
	}
}

func BenchmarkEncodeStdlib(b *testing.B) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := enc.Encode(obj); err != nil {
			b.Fatal(err)
		}
		buf = &bytes.Buffer{}
	}
}

func BenchmarkEncodeJsoniter(b *testing.B) {
	var json = jsoniter.ConfigFastest
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := enc.Encode(obj); err != nil {
			b.Fatal(err)
		}
		buf = &bytes.Buffer{}
	}
}

func BenchmarkEncodeEasyJSON(b *testing.B) {
	buf := &jwriter.Writer{}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		obj.MarshalEasyJSON(buf)
		buf = &jwriter.Writer{}
	}
}

func BenchmarkEncodeJingo(b *testing.B) {
	buf := jingo.NewBufferFromPool()
	var enc = jingo.NewStructEncoder(V{})

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		enc.Marshal(&obj, buf)
		buf = jingo.NewBufferFromPool()
	}
}
