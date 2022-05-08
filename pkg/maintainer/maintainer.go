package maintainer

import (
	"context"
	"net/mail"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

type Maintainer struct {
	handle framework.Handle
}

var _ framework.FilterPlugin = &Maintainer{}

const (
	Name                        = "Maintainer"
	annotationKey               = "ebiiim.com/maintainer"
	ErrReasonAnnotationNotFound = "\"" + annotationKey + "\" didn't found in Pod's Annotations"
	ErrReasonInvalidAddress     = "\"" + annotationKey + "\" in Pod's Annotations didn't match RFC 5322 (e.g. \"John Smith <john@example.com>\")"
)

func (m *Maintainer) Name() string {
	return Name
}

func (m *Maintainer) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	mnt, ok := pod.Annotations[annotationKey]
	if !ok {
		return framework.NewStatus(framework.UnschedulableAndUnresolvable, ErrReasonAnnotationNotFound)
	}
	if _, err := mail.ParseAddress(mnt); err != nil {
		return framework.NewStatus(framework.UnschedulableAndUnresolvable, ErrReasonInvalidAddress)
	}
	return nil
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &Maintainer{handle: h}, nil
}
