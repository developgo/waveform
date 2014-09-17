package waveform

import (
	"bytes"
	"math/rand"
	"testing"
	"time"

	"azul3d.org/audio.v1"
)

// BenchmarkNewWAV checks the performance of the New() function with a WAV file
func BenchmarkNewWAV(b *testing.B) {
	benchmarkNew(b, wavFile, nil)
}

// BenchmarkNewFLAC checks the performance of the New() function with a FLAC file
func BenchmarkNewFLAC(b *testing.B) {
	benchmarkNew(b, flacFile, nil)
}

// BenchmarkComputeValuesWAV checks the performance of the ComputeValues() function with a WAV file
func BenchmarkComputeValuesWAV(b *testing.B) {
	benchmarkComputeValues(b, wavFile, nil)
}

// BenchmarkComputeValuesFLAC checks the performance of the ComputeValues() function with a FLAC file
func BenchmarkComputeValuesFLAC(b *testing.B) {
	benchmarkComputeValues(b, flacFile, nil)
}

// BenchmarkDrawImage60 checks the performance of the DrawImage() function
// with approximately 60 seconds of computed values
func BenchmarkDrawImage60(b *testing.B) {
	benchmarkDrawImage(b, 60, nil)
}

// BenchmarkDrawImage120 checks the performance of the DrawImage() function
// with approximately 120 seconds of computed values
func BenchmarkDrawImage120(b *testing.B) {
	benchmarkDrawImage(b, 120, nil)
}

// BenchmarkDrawImage240 checks the performance of the DrawImage() function
// with approximately 240 seconds of computed values
func BenchmarkDrawImage240(b *testing.B) {
	benchmarkDrawImage(b, 240, nil)
}

// BenchmarkDrawImage480 checks the performance of the DrawImage() function
// with approximately 480 seconds of computed values
func BenchmarkDrawImage480(b *testing.B) {
	benchmarkDrawImage(b, 480, nil)
}

// BenchmarkDrawImage960 checks the performance of the DrawImage() function
// with approximately 960 seconds of computed values
func BenchmarkDrawImage960(b *testing.B) {
	benchmarkDrawImage(b, 960, nil)
}

// BenchmarkRMSF64Samples22050 checks the performance of the RMSF64Samples() function
// with 22050 samples
func BenchmarkRMSF64Samples22050(b *testing.B) {
	benchmarkRMSF64Samples(b, 22050)
}

// BenchmarkRMSF64Samples44100 checks the performance of the RMSF64Samples() function
// with 44100 samples
func BenchmarkRMSF64Samples44100(b *testing.B) {
	benchmarkRMSF64Samples(b, 44100)
}

// BenchmarkRMSF64Samples88200 checks the performance of the RMSF64Samples() function
// with 88200 samples
func BenchmarkRMSF64Samples88200(b *testing.B) {
	benchmarkRMSF64Samples(b, 88200)
}

// BenchmarkRMSF64Samples176400 checks the performance of the RMSF64Samples() function
// with 176400 samples
func BenchmarkRMSF64Samples176400(b *testing.B) {
	benchmarkRMSF64Samples(b, 176400)
}

// benchmarkNew contains common logic for benchmarking New
func benchmarkNew(b *testing.B, data []byte, options *Options) {
	for i := 0; i < b.N; i++ {
		New(bytes.NewReader(data), options)
	}
}

// benchmarkComputeValues contains common logic for benchmarking ComputeValues
func benchmarkComputeValues(b *testing.B, data []byte, options *ComputeOptions) {
	for i := 0; i < b.N; i++ {
		ComputeValues(bytes.NewReader(data), options)
	}
}

// benchmarkDrawImage contains common logic for benchmarking DrawImage
func benchmarkDrawImage(b *testing.B, count int, options *ImageOptions) {
	values := make([]float64, count)
	for i := 0; i < b.N; i++ {
		DrawImage(values, options)
	}
}

// benchmarkRMSF64Samples contains common logic for benchmarking RMSF64Samples
func benchmarkRMSF64Samples(b *testing.B, count int) {
	// Generate slice of samples
	rand.Seed(time.Now().UnixNano())
	var samples audio.F64Samples
	for i := 0; i < count; i++ {
		samples = append(samples, audio.F64(rand.Float64()))
	}

	// Reset timer and start benchmark
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RMSF64Samples(samples)
	}
}
