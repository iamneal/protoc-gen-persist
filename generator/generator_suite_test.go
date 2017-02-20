// Copyright 2017, TCN Inc.
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:

//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//     * Neither the name of TCN Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package generator_test

import (
	"compress/gzip"

	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/tcncloud/protoc-gen-persist/examples"

	"bytes"
	"io/ioutil"
	"testing"
)

func TestGenerator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Generator Suite")
}

var files []*descriptor.FileDescriptorProto

func LoadDescriptor(file string) *descriptor.FileDescriptorProto {

	bReader, err := gzip.NewReader(bytes.NewReader(proto.FileDescriptor(file)))
	defer bReader.Close()
	Expect(err).To(BeNil())

	buf, err := ioutil.ReadAll(bReader)
	Expect(err).To(BeNil())

	var descr descriptor.FileDescriptorProto
	err = proto.Unmarshal(buf, &descr)
	Expect(err).To(BeNil())
	return &descr

}

var _ = BeforeSuite(func() {
	files = append(files, LoadDescriptor("examples/example1.proto"))
	files = append(files, LoadDescriptor("github.com/golang/protobuf/ptypes/timestamp/timestamp.proto"))
})
