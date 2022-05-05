package server

import (
	"io/ioutil"
	"net/http"

	admissionv1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

type AdmissionServer struct {
	Server *http.Server
}

func (s *AdmissionServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		klog.Errorln("error to read request body, ", nil)
		return
	}
	if len(body) == 0 {
		klog.Errorln("empty request body")
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		klog.Errorf("Content-Type=%s, expect application/json\n", contentType)
		http.Error(w, "invalid Content-Type, expect `application/json`", http.StatusUnsupportedMediaType)
		return
	}

	ar := admissionv1.AdmissionReview{}
	_, groupVersionKind, err := Codec.Decode(body, nil, &ar)
	// err = json.Unmarshal(body, &ar)
	klog.Infoln(groupVersionKind)
	if err != nil {
		klog.Errorln("can't decode request")
		http.Error(w, "can't decode request", http.StatusBadRequest)
		return
	}

	var response *admissionv1.AdmissionResponse
	if r.URL.Path == "/mutate" {
		response = s.Mutate(&ar)
	} else if r.URL.Path == "/validate" {
		response = s.Validate(&ar)
	} else {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	responseAR := &admissionv1.AdmissionReview{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "admission.k8s.io/v1",
			Kind:       "AdmissionReview",
		},
	}
	response.UID = ar.Request.UID
	responseAR.Response = response

	err = Codec.Encode(responseAR, w)
	if err != nil {
		klog.Errorln("encode response error: ", err)
	}

}

func (s *AdmissionServer) Mutate(ar *admissionv1.AdmissionReview) *admissionv1.AdmissionResponse {
	// TODO 进行业务处理
	return &admissionv1.AdmissionResponse{Allowed: false, Result: &metav1.Status{Message: "forbidden everything"}}
}

func (s *AdmissionServer) Validate(ar *admissionv1.AdmissionReview) *admissionv1.AdmissionResponse {
	// TODO 进行业务处理
	return &admissionv1.AdmissionResponse{Allowed: false, Result: &metav1.Status{Message: "forbidden everything"}}
}
