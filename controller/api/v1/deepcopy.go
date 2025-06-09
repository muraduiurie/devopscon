package v1

import "k8s.io/apimachinery/pkg/runtime"

func (in *WebPage) DeepCopyInto(out *WebPage) {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	out.Spec = WebPageSpec{
		Image:    in.Spec.Image,
		Content:  in.Spec.Content,
		Replicas: in.Spec.Replicas,
	}
}

func (in *WebPage) DeepCopyObject() runtime.Object {
	out := WebPage{}
	in.DeepCopyInto(&out)

	return &out
}

func (in *WebPageList) DeepCopyObject() runtime.Object {
	out := WebPageList{}
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta

	if in.Items != nil {
		out.Items = make([]WebPage, len(in.Items))
		for i := range in.Items {
			in.Items[i].DeepCopyInto(&out.Items[i])
		}
	}

	return &out
}
