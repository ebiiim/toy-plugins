package maintainer

import (
	"context"
	"net/mail"
	"strings"

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
	annotKeyAddr                = "ebiiim.com/maintainer.address"
	annotKeyEnabled             = "ebiiim.com/maintainer.enabled"
	ErrReasonAnnotationNotFound = "\"" + annotKeyAddr + "\" didn't found in Pod's Annotations"
	ErrReasonInvalidAddress     = "\"" + annotKeyAddr + "\" in Pod's Annotations didn't match RFC 5322 (e.g. \"John Smith <john@example.com>\")"
)

func (*Maintainer) Name() string {
	return Name
}

func (m *Maintainer) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeInfo *framework.NodeInfo) *framework.Status {
	// skip this filter if not enabled
	en, ok := pod.Annotations[annotKeyEnabled]
	if !ok || strings.ToLower(en) != "true" {
		return nil
	}
	// validate the address
	addr, ok := pod.Annotations[annotKeyAddr]
	if !ok {
		return framework.NewStatus(framework.UnschedulableAndUnresolvable, ErrReasonAnnotationNotFound)
	}
	if _, err := mail.ParseAddress(addr); err != nil {
		return framework.NewStatus(framework.UnschedulableAndUnresolvable, ErrReasonInvalidAddress)
	}
	return nil
}

func New(_ runtime.Object, h framework.Handle) (framework.Plugin, error) {
	return &Maintainer{handle: h}, nil
}
