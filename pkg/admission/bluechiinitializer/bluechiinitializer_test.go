/* SPDX-License-Identifier: LGPL-2.1-or-later */

package bluechiinitializer_test

import (
	"context"
	"testing"
	"time"

	"github.com/engelmi/bluechi-server/pkg/admission/bluechiinitializer"
	"github.com/engelmi/bluechi-server/pkg/generated/clientset/versioned/fake"
	informers "github.com/engelmi/bluechi-server/pkg/generated/informers/externalversions"
	"k8s.io/apiserver/pkg/admission"
)

// TestWantsInternalBlueChiInformerFactory ensures that the informer factory is injected
// when the WantsInternalBlueChiInformerFactory interface is implemented by a plugin.
func TestWantsInternalBlueChiInformerFactory(t *testing.T) {
	cs := &fake.Clientset{}
	sf := informers.NewSharedInformerFactory(cs, time.Duration(1)*time.Second)
	target := bluechiinitializer.New(sf)

	wantBlueChiInformerFactory := &wantInternalBlueChiInformerFactory{}
	target.Initialize(wantBlueChiInformerFactory)
	if wantBlueChiInformerFactory.sf != sf {
		t.Errorf("expected informer factory to be initialized")
	}
}

// wantInternalBlueChiInformerFactory is a test stub that fulfills the WantsInternalBlueChiInformerFactory interface
type wantInternalBlueChiInformerFactory struct {
	sf informers.SharedInformerFactory
}

func (f *wantInternalBlueChiInformerFactory) SetInternalBlueChiInformerFactory(sf informers.SharedInformerFactory) {
	f.sf = sf
}
func (f *wantInternalBlueChiInformerFactory) Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}
func (f *wantInternalBlueChiInformerFactory) Handles(o admission.Operation) bool { return false }
func (f *wantInternalBlueChiInformerFactory) ValidateInitialization() error      { return nil }

var _ admission.Interface = &wantInternalBlueChiInformerFactory{}
var _ bluechiinitializer.WantsInternalBlueChiInformerFactory = &wantInternalBlueChiInformerFactory{}
