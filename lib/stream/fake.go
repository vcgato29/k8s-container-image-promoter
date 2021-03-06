/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package stream

import (
	"bytes"
	"io"
)

// Fake is a predefined stream (set with Bytes).
type Fake struct {
	stream io.Reader
	Bytes  []byte
}

// Produce a fake stream. Unlike a real stream, this does not call a subprocess
// --- instead it just reads from a predefined stream (Bytes).
func (producer *Fake) Produce() (io.Reader, error) {
	producer.stream = bytes.NewReader(producer.Bytes)
	return producer.stream, nil
}

// Close does nothing, as there is no actual subprocess to close.
func (producer *Fake) Close() error {
	return nil
}
