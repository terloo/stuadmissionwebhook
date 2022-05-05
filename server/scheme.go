package server

import (
	admissionv1 "k8s.io/api/admission/v1"
	admissionregistryv1 "k8s.io/api/admissionregistration/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var Scheme = runtime.NewScheme()
var CodecFactory = runtimeserializer.NewCodecFactory(Scheme)
var Codec runtime.Codec

func init() {
	utilruntime.Must(corev1.AddToScheme(Scheme))
	utilruntime.Must(admissionv1.AddToScheme(Scheme))
	utilruntime.Must(admissionregistryv1.AddToScheme(Scheme))

	info, _ := runtime.SerializerInfoForMediaType(CodecFactory.SupportedMediaTypes(), runtime.ContentTypeJSON)
	jsonSerializer := info.Serializer

	Codec = CodecFactory.CodecForVersions(jsonSerializer, jsonSerializer, admissionv1.SchemeGroupVersion, admissionv1.SchemeGroupVersion)
}
