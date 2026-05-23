package capabilityoverride

import (
	"testing"

	"github.com/flexigpt/inference-go/spec"
)

const testCompletionKey = "completion-1"

func TestCompletionKeyResolverResolveModelCapabilitiesTable(t *testing.T) {
	base := baseCapabilitiesForTest()

	tests := []struct {
		name     string
		resolver CompletionKeyResolver
		req      spec.ResolveModelCapabilitiesRequest
		wantErr  string
		check    func(t *testing.T, got *spec.ModelCapabilities)
	}{
		{
			name:     "nil capabilities errors",
			resolver: NewCompletionKeyResolver(testCompletionKey, nil),
			req: spec.ResolveModelCapabilitiesRequest{
				CompletionKey: testCompletionKey,
			},
			wantErr: "no model capabilities configured",
		},
		{
			name:     "completion key mismatch errors",
			resolver: NewCompletionKeyResolver(testCompletionKey, &base),
			req: spec.ResolveModelCapabilitiesRequest{
				CompletionKey: "completion-2",
			},
			wantErr: "capabilities not found for completionKey",
		},
		{
			name:     "completion key match resolves",
			resolver: NewCompletionKeyResolver(testCompletionKey, &base),
			req: spec.ResolveModelCapabilitiesRequest{
				CompletionKey: testCompletionKey,
			},
			check: func(t *testing.T, got *spec.ModelCapabilities) {
				t.Helper()

				if got == nil {
					t.Fatal("expected capabilities")
				}
				assertDeepEqual(t, *got, baseCapabilitiesForTest())
			},
		},
		{
			name:     "empty resolver completion key accepts any request key",
			resolver: NewCompletionKeyResolver("", &base),
			req: spec.ResolveModelCapabilitiesRequest{
				CompletionKey: "any-key",
			},
			check: func(t *testing.T, got *spec.ModelCapabilities) {
				t.Helper()

				if got == nil {
					t.Fatal("expected capabilities")
				}
				assertDeepEqual(t, *got, baseCapabilitiesForTest())
			},
		},
		{
			name:     "empty request completion key works when resolver key is also empty",
			resolver: NewCompletionKeyResolver("", &base),
			req:      spec.ResolveModelCapabilitiesRequest{},
			check: func(t *testing.T, got *spec.ModelCapabilities) {
				t.Helper()

				if got == nil {
					t.Fatal("expected capabilities")
				}
				assertDeepEqual(t, *got, baseCapabilitiesForTest())
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.resolver.ResolveModelCapabilities(t.Context(), tc.req)

			if tc.wantErr != "" {
				assertErrContains(t, err, tc.wantErr)
				if got != nil {
					t.Fatalf("expected nil capabilities on error, got %#v", got)
				}
				return
			}

			assertNoErr(t, err)
			tc.check(t, got)
		})
	}
}

func TestCompletionKeyResolverClonesInputAndOutput(t *testing.T) {
	base := baseCapabilitiesForTest()
	resolver := NewCompletionKeyResolver(testCompletionKey, &base)

	mutateModelCapabilitiesForCloneTest(&base)

	got1, err := resolver.ResolveModelCapabilities(
		t.Context(),
		spec.ResolveModelCapabilitiesRequest{
			CompletionKey: testCompletionKey,
		},
	)
	assertNoErr(t, err)

	want := baseCapabilitiesForTest()
	assertDeepEqual(t, *got1, want)

	got1.ModalitiesIn[0] = spec.ModalityAudioIn
	got1.ReasoningCapabilities.SupportedReasoningTypes[0] = spec.ReasoningTypeSingleWithLevels
	got1.CacheCapabilities.TopLevel.SupportedTTLs[0] = spec.CacheControlTTL24h

	got2, err := resolver.ResolveModelCapabilities(
		t.Context(),
		spec.ResolveModelCapabilitiesRequest{
			CompletionKey: testCompletionKey,
		},
	)
	assertNoErr(t, err)

	if got1 == got2 {
		t.Fatal("expected resolver to return a fresh capabilities pointer each time")
	}
	if got1 == resolver.Capabilities() {
		t.Fatal("expected resolved capabilities not to share resolver internal pointer")
	}
	if got2 == resolver.Capabilities() {
		t.Fatal("expected resolved capabilities not to share resolver internal pointer")
	}

	assertDeepEqual(t, *got2, want)
	assertDeepEqual(t, *resolver.Capabilities(), want)
}

func TestNewCompletionKeyResolverNilCapabilities(t *testing.T) {
	resolver := NewCompletionKeyResolver(testCompletionKey, nil)

	if resolver.CompletionKey() != testCompletionKey {
		t.Fatalf("unexpected completion key: %q", resolver.CompletionKey())
	}
	if resolver.Capabilities() != nil {
		t.Fatalf("expected nil capabilities, got %#v", resolver.Capabilities())
	}
}
