//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageDetails) DeepCopyInto(out *ImageDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageDetails.
func (in *ImageDetails) DeepCopy() *ImageDetails {
	if in == nil {
		return nil
	}
	out := new(ImageDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageReference) DeepCopyInto(out *ImageReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageReference.
func (in *ImageReference) DeepCopy() *ImageReference {
	if in == nil {
		return nil
	}
	out := new(ImageReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanReport) DeepCopyInto(out *ImageScanReport) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanReport.
func (in *ImageScanReport) DeepCopy() *ImageScanReport {
	if in == nil {
		return nil
	}
	out := new(ImageScanReport)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageScanReport) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanReportList) DeepCopyInto(out *ImageScanReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageScanReport, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanReportList.
func (in *ImageScanReportList) DeepCopy() *ImageScanReportList {
	if in == nil {
		return nil
	}
	out := new(ImageScanReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageScanReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanReportSpec) DeepCopyInto(out *ImageScanReportSpec) {
	*out = *in
	out.Image = in.Image
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanReportSpec.
func (in *ImageScanReportSpec) DeepCopy() *ImageScanReportSpec {
	if in == nil {
		return nil
	}
	out := new(ImageScanReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanReportStatus) DeepCopyInto(out *ImageScanReportStatus) {
	*out = *in
	in.LastChecked.DeepCopyInto(&out.LastChecked)
	in.Report.DeepCopyInto(&out.Report)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanReportStatus.
func (in *ImageScanReportStatus) DeepCopy() *ImageScanReportStatus {
	if in == nil {
		return nil
	}
	out := new(ImageScanReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanRequest) DeepCopyInto(out *ImageScanRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanRequest.
func (in *ImageScanRequest) DeepCopy() *ImageScanRequest {
	if in == nil {
		return nil
	}
	out := new(ImageScanRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageScanRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanRequestList) DeepCopyInto(out *ImageScanRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ImageScanRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanRequestList.
func (in *ImageScanRequestList) DeepCopy() *ImageScanRequestList {
	if in == nil {
		return nil
	}
	out := new(ImageScanRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ImageScanRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanRequestSpec) DeepCopyInto(out *ImageScanRequestSpec) {
	*out = *in
	if in.PullSecrets != nil {
		in, out := &in.PullSecrets, &out.PullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanRequestSpec.
func (in *ImageScanRequestSpec) DeepCopy() *ImageScanRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ImageScanRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageScanRequestStatus) DeepCopyInto(out *ImageScanRequestStatus) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageDetails)
		**out = **in
	}
	if in.ReportRef != nil {
		in, out := &in.ReportRef, &out.ReportRef
		*out = new(ScanReportRef)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageScanRequestStatus.
func (in *ImageScanRequestStatus) DeepCopy() *ImageScanRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ImageScanRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScanReportRef) DeepCopyInto(out *ScanReportRef) {
	*out = *in
	in.LastChecked.DeepCopyInto(&out.LastChecked)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScanReportRef.
func (in *ScanReportRef) DeepCopy() *ScanReportRef {
	if in == nil {
		return nil
	}
	out := new(ScanReportRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vulnerability) DeepCopyInto(out *Vulnerability) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vulnerability.
func (in *Vulnerability) DeepCopy() *Vulnerability {
	if in == nil {
		return nil
	}
	out := new(Vulnerability)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Vulnerability) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityList) DeepCopyInto(out *VulnerabilityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Vulnerability, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityList.
func (in *VulnerabilityList) DeepCopy() *VulnerabilityList {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VulnerabilityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilitySpec) DeepCopyInto(out *VulnerabilitySpec) {
	*out = *in
	in.Vulnerability.DeepCopyInto(&out.Vulnerability)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilitySpec.
func (in *VulnerabilitySpec) DeepCopy() *VulnerabilitySpec {
	if in == nil {
		return nil
	}
	out := new(VulnerabilitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VulnerabilityStatus) DeepCopyInto(out *VulnerabilityStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VulnerabilityStatus.
func (in *VulnerabilityStatus) DeepCopy() *VulnerabilityStatus {
	if in == nil {
		return nil
	}
	out := new(VulnerabilityStatus)
	in.DeepCopyInto(out)
	return out
}