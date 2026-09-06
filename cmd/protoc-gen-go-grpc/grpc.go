package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	contextPackage = protogen.GoImportPath("context")
	grpcPackage    = protogen.GoImportPath("google.golang.org/grpc")
	codesPackage   = protogen.GoImportPath("google.golang.org/grpc/codes")
	statusPackage  = protogen.GoImportPath("google.golang.org/grpc/status")
)

type serviceGenerateHelperInterface interface {
	formatFullMethodSymbol(service *protogen.Service, method *protogen.Method) string
	genFullMethods(g *protogen.GeneratedFile, service *protogen.Service)
	generateClientStruct(g *protogen.GeneratedFile, clientName string)
	generateNewClientDefinitions(g *protogen.GeneratedFile, service *protogen.Service, clientName string)
	generateUnimplementedServerType(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service)
	generateServerFunctions(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service, serverType string, serviceDescVar string)
	formatHandlerFuncName(service *protogen.Service, hname string) string
}

type serviceGenerateHelper struct{}

func (serviceGenerateHelper) formatFullMethodSymbol(service *protogen.Service, method *protogen.Method) string {
	_ = "STUB: not implemented"
	return ""
}

func (serviceGenerateHelper) genFullMethods(g *protogen.GeneratedFile, service *protogen.Service) {
	_ = "STUB: not implemented"
	return
}

func (serviceGenerateHelper) generateClientStruct(g *protogen.GeneratedFile, clientName string) {
	_ = "STUB: not implemented"
	return
}

func (serviceGenerateHelper) generateNewClientDefinitions(g *protogen.GeneratedFile, _ *protogen.Service, clientName string) {
	_ = "STUB: not implemented"
	return
}

func (serviceGenerateHelper) generateUnimplementedServerType(_ *protogen.Plugin, _ *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	_ = "STUB: not implemented"
	return
}

func (serviceGenerateHelper) generateServerFunctions(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service, serverType string, serviceDescVar string) {
	_ = "STUB: not implemented"
	return
}

func (serviceGenerateHelper) formatHandlerFuncName(_ *protogen.Service, hname string) string {
	_ = "STUB: not implemented"
	return ""
}

var helper serviceGenerateHelperInterface = serviceGenerateHelper{}

const fileDescriptorProtoPackageFieldNumber = 2

const fileDescriptorProtoSyntaxFieldNumber = 12

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	_ = "STUB: not implemented"
	return nil
}

func protocVersion(gen *protogen.Plugin) string { _ = "STUB: not implemented"; return "" }

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	_ = "STUB: not implemented"
	return
}

func genServiceComments(g *protogen.GeneratedFile, service *protogen.Service) {
	_ = "STUB: not implemented"
	return
}

func genService(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, service *protogen.Service) {
	_ = "STUB: not implemented"
	return
}

func clientSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	_ = "STUB: not implemented"
	return ""
}

func clientStreamInterface(g *protogen.GeneratedFile, method *protogen.Method) string {
	_ = "STUB: not implemented"
	return ""
}

func genClientMethod(_ *protogen.Plugin, _ *protogen.File, g *protogen.GeneratedFile, method *protogen.Method, index int) {
	_ = "STUB: not implemented"
	return
}

func serverSignature(g *protogen.GeneratedFile, method *protogen.Method) string {
	_ = "STUB: not implemented"
	return ""
}

func genServiceDesc(file *protogen.File, g *protogen.GeneratedFile, serviceDescVar string, serverType string, service *protogen.Service, handlerNames []string) {
	_ = "STUB: not implemented"
	return
}

func serverStreamInterface(g *protogen.GeneratedFile, method *protogen.Method) string {
	_ = "STUB: not implemented"
	return ""
}

func genServerMethod(_ *protogen.Plugin, _ *protogen.File, g *protogen.GeneratedFile, method *protogen.Method, hnameFuncNameFormatter func(string) string) string {
	_ = "STUB: not implemented"
	return ""
}

func genLeadingComments(g *protogen.GeneratedFile, loc protoreflect.SourceLocation) {
	_ = "STUB: not implemented"
	return
}

const deprecationComment = "// Deprecated: Do not use."

func unexport(s string) string { _ = "STUB: not implemented"; return "" }
