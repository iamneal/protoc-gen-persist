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

package generator

import (
	"path/filepath"
	"strings"

	"os"

	"github.com/Sirupsen/logrus"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/tcncloud/protoc-gen-persist/persist"
	"path"
)

type FileStruct struct {
	Desc            *descriptor.FileDescriptorProto
	ImportList      *Imports
	Dependency      bool        // if is dependency
	Structures      *StructList // all structures in the file
	AllStructures   *StructList // all structures in all the files
	ServiceList     *Services
	PersistStringer *PersistStringer // print our persist file
	Opts            PersistOpts      // options passed in via parameter
}

func NewFileStruct(
	desc *descriptor.FileDescriptorProto, allStructs *StructList,
	dependency bool, opts PersistOpts) *FileStruct {
	ret := &FileStruct{
		Desc:          desc,
		ImportList:    EmptyImportList(),
		Structures:    &StructList{},
		ServiceList:   &Services{},
		AllStructures: allStructs,
		Dependency:    dependency,
		Opts:          opts,
	}
	ret.PersistStringer = &PersistStringer{}

	return ret
}

func (f *FileStruct) GetOrigName() string {
	return f.Desc.GetName()
}

func (f *FileStruct) GetPackageName() string {
	return f.GetImplPackage()
}
func (f *FileStruct) IsSameAsMyPackage(pkg string) bool {
	return f.GetImplDir() == pkg
}
func (f *FileStruct) NotSameAsMyPackage(pkg string) bool {
	return !f.IsSameAsMyPackage(pkg)
}

// extract the persist.package file option
// or return "" as default
func (f *FileStruct) GetPersistPackageOption() string {
	if f.Desc == nil || f.Desc.GetOptions() == nil {
		return ""
	}
	if proto.HasExtension(f.Desc.GetOptions(), persist.E_Package) {
		pkg, err := proto.GetExtension(f.Desc.GetOptions(), persist.E_Package)
		if err != nil {
			logrus.WithError(err).Debug("Error")
			return ""
		}
		return *pkg.(*string)
	}
	logrus.WithField("File Options", f.Desc.GetOptions()).Debug("file options")
	return ""
}

func (f *FileStruct) GetImplFileName() string {
	_, file := filepath.Split(f.Desc.GetName())
	return strings.Join([]string{
		f.GetImplDir(),
		string(os.PathSeparator),
		strings.Replace(file, ".proto", ".persist.go", -1),
	}, "")
}

func (f *FileStruct) GetImplDir() string {
	pkg := f.GetPersistPackageOption()
	if pkg == "" {
		// if the persist.package option is not present we will use go_package
		if f.Desc.GetOptions().GetGoPackage() != "" {
			pkg = f.Desc.GetOptions().GetGoPackage()
		} else {
			// last resort
			pkg = f.Desc.GetPackage()
		}
	}
	if strings.Contains(pkg, ";") {
		// we need to split by ";"
		p := strings.Split(pkg, ";")
		if len(p) > 2 {
			logrus.WithField("persist package", pkg).Panic("Invalid persist package")
		}
		return p[0]
	}
	return pkg
}

func (f *FileStruct) GetImplPackage() string {
	pkg := f.GetPersistPackageOption()
	if pkg == "" {
		// if the persist.package option is not present we will use
		// go_pacakge
		if f.Desc.GetOptions().GetGoPackage() != "" {
			pkg = f.Desc.GetOptions().GetGoPackage()
		} else {
			// last resort
			pkg = f.Desc.GetPackage()
		}
	}
	// process pkg
	if strings.Contains(pkg, ";") {
		// we need to split by ";"
		p := strings.Split(pkg, ";")
		if len(p) > 2 {
			logrus.WithField("persist package", pkg).Panic("Invalid persist package")
		}
		return p[len(p)-1]
	}

	if strings.Contains(pkg, "/") {
		// return package after last /
		p := strings.Split(pkg, "/")
		return p[len(p)-1]
	}
	return strings.Replace(pkg, ".", "_", -1)
}

// return the computed persist file taking in consideration
// the persist.package option and original file name
func (f *FileStruct) GetPersistFile() string {
	return ""
}

func (f *FileStruct) GetFileName() string {
	return strings.Replace(f.Desc.GetName(), ".proto", ".persist.go", -1)
}

func (f *FileStruct) GetServices() *Services {
	return f.ServiceList
}

func (f *FileStruct) GetFullGoPackage() string {
	if f.Desc.Options != nil && f.Desc.GetOptions().GoPackage != nil {
		switch {
		case strings.Contains(f.Desc.GetOptions().GetGoPackage(), ";"):
			idx := strings.Index(f.Desc.GetOptions().GetGoPackage(), ";")
			return f.Desc.GetOptions().GetGoPackage()[0:idx]
		default:
			return f.Desc.GetOptions().GetGoPackage()
		}
	} else {
		return strings.Replace(f.Desc.GetPackage(), ".", "_", -1)
	}
}

func (f *FileStruct) DifferentImpl() bool {
	return (f.GetFullGoPackage() != f.GetImplDir())
}

func (f *FileStruct) GetGoPackage() string {
	if f.Desc.Options != nil && f.Desc.GetOptions().GoPackage != nil {
		switch {
		case strings.Contains(f.Desc.GetOptions().GetGoPackage(), ";"):
			idx := strings.LastIndex(f.Desc.GetOptions().GetGoPackage(), ";")
			return f.Desc.GetOptions().GetGoPackage()[idx+1:]
		case strings.Contains(f.Desc.GetOptions().GetGoPackage(), "/"):
			idx := strings.LastIndex(f.Desc.GetOptions().GetGoPackage(), "/")
			return f.Desc.GetOptions().GetGoPackage()[idx+1:]
		default:
			return f.Desc.GetOptions().GetGoPackage()
		}
	} else {
		return strings.Replace(f.Desc.GetPackage(), ".", "_", -1)
	}
}

func (f *FileStruct) GetGoPath() string {
	if f.Desc.Options != nil && f.Desc.GetOptions().GoPackage != nil {
		switch {
		case strings.Contains(f.Desc.GetOptions().GetGoPackage(), ";"):
			idx := strings.LastIndex(f.Desc.GetOptions().GetGoPackage(), ";")
			return f.Desc.GetOptions().GetGoPackage()[0:idx]
		case strings.Contains(f.Desc.GetOptions().GetGoPackage(), "/"):
			return f.Desc.GetOptions().GetGoPackage()
		default:
			return f.Desc.GetOptions().GetGoPackage()
		}
	} else {
		return strings.Replace(f.Desc.GetPackage(), ".", "_", -1)
	}
}

func (f *FileStruct) ProcessImportsForType(name string) {
	typ := f.AllStructures.GetStructByProtoName(name)
	if typ != nil {
		for _, file := range *typ.GetImportedFiles() {
			if f.GetPackageName() != file.GetPackageName() {
				f.ImportList.GetOrAddImport(file.GetGoPackage(), file.GetGoPath())
			}
		}
	} else {
		logrus.WithField("all structures", f.AllStructures).Fatalf("Can't find structure %s!", name)
	}
}

func (f *FileStruct) ProcessImports() {
	importsForStructName := func(name string, m *Method) {
		// first get imports for this struct
		f.ProcessImportsForType(name)

		// make sure every field is check if it needs an import as well
		str := f.AllStructures.GetStructByProtoName(name)
		for _, mp := range str.MsgDesc.GetField() {
			if mp.GetType() == descriptor.FieldDescriptorProto_TYPE_MESSAGE ||
				mp.GetType() == descriptor.FieldDescriptorProto_TYPE_ENUM {

				// if the type is mapped, we do not need this struct imported
				// instead we need the mapped type to be imported
				if m.GetMapping(mp) == nil {
					f.ProcessImportsForType(mp.GetTypeName())
				}
			}
		}
	}
	needServiceImports := func() bool {
		for _, s := range *f.ServiceList {
			if s.IsSQL() || s.IsSpanner() {
				return true
			}
		}
		return false
	}
	if needServiceImports() {
		f.ImportList.GetOrAddImport("io", "io")
		f.ImportList.GetOrAddImport("context", "golang.org/x/net/context")
		f.ImportList.GetOrAddImport("codes", "google.golang.org/grpc/codes")
		f.ImportList.GetOrAddImport("gstatus", "google.golang.org/grpc/status")

	}
	for _, srv := range *f.ServiceList {
		if !srv.IsSQL() && !srv.IsSpanner() {
			continue
		}
		if srv.IsSpanner() {
			f.ImportList.GetOrAddImport("spanner", "cloud.google.com/go/spanner")
		}
		if srv.IsSQL() {
			f.ImportList.GetOrAddImport("sql", "database/sql")
		}
		if opt := srv.GetServiceOption(); opt != nil {
			for _, m := range opt.GetTypes() {
				pkg := m.GetGoPackage()
				f.ImportList.GetOrAddImport(GetGoPackage(pkg), GetGoPath(pkg))
			}
		}
		for _, m := range *srv.Methods {
			importsForStructName(m.Desc.GetInputType(), m)
			importsForStructName(m.Desc.GetOutputType(), m)
			opts := m.GetMethodOption()
			if opts == nil {
				continue
			}
			if beforeOpt := opts.GetBefore(); beforeOpt != nil {
				pkg := beforeOpt.GetPackage()
				f.ImportList.GetOrAddImport(GetGoPackage(pkg), GetGoPath(pkg))
			}
			if afterOpt := opts.GetAfter(); afterOpt != nil {
				pkg := afterOpt.GetPackage()
				f.ImportList.GetOrAddImport(GetGoPackage(pkg), GetGoPath(pkg))
			}
			if mappingOpt := opts.GetMapping(); mappingOpt != nil {
				for _, typMap := range mappingOpt.GetTypes() {
					pkg := typMap.GetGoPackage()
					f.ImportList.GetOrAddImport(GetGoPackage(pkg), GetGoPath(pkg))
				}
			}
		}
	}
}

type persistFile struct {
	filename  string
	path      string
	importStr string
}

func (f *FileStruct) GetPersistLibFullFilepath() persistFile {
	// the full path that generated me
	origFile := path.Base(f.GetOrigName())
	// sloppily grab the first part of the filename
	beforeDot := strings.Split(origFile, ".")[0]
	if beforeDot == "" {
		beforeDot = "_"
	}
	imp := func() string {
		if f.Opts.PersistLibRoot == "" {
			return f.GetImplDir()
		}
		return f.Opts.PersistLibRoot
	}()
	return persistFile{
		filename:  beforeDot,
		path:      path.Join(f.GetImplDir(), "/persist_lib"),
		importStr: path.Join(imp, "/persist_lib"),
	}
}
func (f *FileStruct) Process() error {
	logrus.WithFields(logrus.Fields{
		"GetFileName()": f.GetFileName(),
	}).Debug("processing file")
	// collect file defined messages
	for _, m := range f.Desc.GetMessageType() {
		s := f.AllStructures.AddMessage(m, nil, f.GetPackageName(), f)
		f.Structures.Append(s)
	}
	// collect file defined enums
	for _, e := range f.Desc.GetEnumType() {
		s := f.AllStructures.AddEnum(e, nil, f.GetPackageName(), f)
		f.Structures.Append(s)
	}

	for _, s := range f.Desc.GetService() {
		f.ServiceList.AddService(f.GetPackageName(), s, f.AllStructures, f)
	}
	if f.NeedsPersistLibDir() {
		f.ImportList.GetOrAddImport("persist_lib", f.GetPersistLibFullFilepath().importStr)
		for _, i := range *f.ImportList {
			logrus.Debugf("IMPORT LIST: %+v, %#v\n", i.GoPackageName, i.GoImportPath)
		}
	}
	return nil
}
func (f *FileStruct) NeedImport(pkg string) bool {
	if f.NotSameAsMyPackage(pkg) &&
		(f.Opts.PersistLibRoot != pkg) &&
		pkg != "" {
		return true
	}
	return false
}

func (f *FileStruct) SanatizeImports() {
	imports := Imports(make([]*Import, 0))
	for _, i := range *f.ImportList {
		if f.NeedImport(i.GoImportPath) {
			imports.GetOrAddImport(i.GoPackageName, i.GoImportPath)
		}
	}
	f.ImportList = &imports
}

// the generator may need to make an extra package in the directory beneath ours
// this library will not be imported by our persist code, so it is safe to use
// anywhere the client wants
func (f *FileStruct) NeedsPersistLibDir() bool {
	if !f.Dependency && f.ServiceList.HasPersistService() {
		return true
	}
	return false
}

func (f *FileStruct) Generate() ([]byte, error) {
	err := f.ServiceList.PreGenerate()
	if err != nil {
		return nil, err
	}
	printer := &Printer{}
	printer.P(")\n")
	for _, s := range *f.ServiceList {
		processed := make(map[string]bool)
		if s.IsSpanner() || s.IsSQL() {
			printer.P("%s\n", s.PrintBuilder())
		}
		for _, m := range *s.Methods {
			if !processed[ToParamsFuncName(m)] {
				printer.P("%s\n", m.backend.MapRequestToParams())
				processed[ToParamsFuncName(m)] = true
			}
			if !processed[FromScanableFuncName(m)] {
				printer.P("%s\n", m.backend.TranslateRowToResult())
				processed[FromScanableFuncName(m)] = true
			}
			if !processed[IterProtoName(m)] {
				printer.P("%s\n", IteratorHelper(m))
				processed[IterProtoName(m)] = true
			}
		}
		for _, m := range *s.Methods {
			printer.P("%s\n", m)
		}
	}
	// prin the imports last, because the above code might have added to our files imports
	importPrinter := &Printer{}
	importPrinter.PA([]string{
		"// This file is generated by protoc-gen-persist\n",
		"// Source File: %s\n",
		"// DO NOT EDIT !\n",
		"package %s\n",
	}, f.GetOrigName(), f.GetImplPackage())
	f.SanatizeImports()
	importPrinter.P("import(\n")
	for _, i := range *f.ImportList {
		if f.NotSameAsMyPackage(i.GoImportPath) {
			importPrinter.P("%s \"%s\"\n", i.GoPackageName, i.GoImportPath)
		}
	}
	importPrinter.P("%s", printer.String())

	return []byte(importPrinter.String()), nil
}

// FileList ----------------

type FileList []*FileStruct

func NewFileList() *FileList {
	return &FileList{}
}

func (fl *FileList) FindFile(desc *descriptor.FileDescriptorProto) *FileStruct {
	for _, f := range *fl {
		if f.Desc.GetName() == desc.GetName() {
			return f
		}
	}
	return nil
}

func (fl *FileList) GetOrCreateFile(desc *descriptor.FileDescriptorProto, allStructs *StructList, dependency bool, params PersistOpts) *FileStruct {
	if f := fl.FindFile(desc); f != nil {
		return f
	}
	f := NewFileStruct(desc, allStructs, dependency, params)
	*fl = append(*fl, f)
	return f
}

func (fl *FileList) Process() error {
	for _, file := range *fl {
		err := file.Process()
		if err != nil {
			return err
		}
	}
	return nil
}

func (fl *FileList) Append(file *FileStruct) {
	*fl = append(*fl, file)
}
