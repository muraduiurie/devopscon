package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type WebPage struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              WebPageSpec `json:"spec"`
}

type WebPageList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WebPage `json:"items"`
}

type WebPageSpec struct {
	Content  string `json:"content"`
	Image    string `json:"image"`
	Replicas int    `json:"replicas"`
}
