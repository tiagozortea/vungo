package vast_test

import (
	"github.com/Vungle/vungo/vast"
	"github.com/Vungle/vungo/vast/vasttest"
	"reflect"
	"testing"
)

var CompanionAdsWrapperModelType = reflect.TypeOf(vast.CompanionAdsWrapper{})

func TestCompanionAdsWrapperMarshalUnmarshal(t *testing.T) {
	vasttest.VerifyModelAgainstFile(t, "CompanionAds", "companionadswrapper.xml", CompanionAdsWrapperModelType)
}

var companionAdsWrapperTests = []vasttest.VastTest{
	vasttest.VastTest{&vast.CompanionAdsWrapper{}, vast.ErrCompanionAdsWrapperMissCompanions, "companionadswrapper_without_companion.xml"},
	vasttest.VastTest{&vast.CompanionAdsWrapper{}, nil, "companionadswrapper_valid.xml"},
	vasttest.VastTest{&vast.CompanionAdsWrapper{}, vast.ErrCompanionWrapperResourceFormat, "companionadswrapper.xml"},
	vasttest.VastTest{&vast.CompanionAdsWrapper{}, vast.ErrUnsupportedMode, "companionadswrapper_error_required.xml"},
}

func TestCompanionAdsWrapperValidateErrors(t *testing.T) {
	for _, test := range companionAdsWrapperTests {
		vasttest.VerifyVastElementFromFile(t, "testdata/"+test.File, test.VastElement, test.Err)
	}
}
