package bus

import (
	"encoding/json"
	"testing"

	"github.com/keepdevops/cofiswarm-observer-sdk/pkg/servicecomponent"
)

func TestInfoRouteReturnsAdapter(t *testing.T) {
	subj := servicecomponent.Prefix + ".adapter.openai-compat.info"
	out, err := Routes("openai-compat")[subj](nil)
	if err != nil {
		t.Fatal(err)
	}
	if r := out.(infoReply); !r.OK || r.Adapter != "openai-compat" {
		t.Fatalf("got %+v", r)
	}
}

func TestHealthRouteOK(t *testing.T) {
	subj := servicecomponent.Prefix + ".adapter.openai-compat.health"
	out, _ := Routes("openai-compat")[subj](nil)
	if r := out.(healthReply); !r.OK || r.Status != "ok" {
		t.Fatalf("got %+v", r)
	}
}

func TestReplyCarriesSchemaVersion(t *testing.T) {
	subj := servicecomponent.Prefix + ".adapter.openai-compat.info"
	out, _ := Routes("openai-compat")[subj](nil)
	b, _ := json.Marshal(out)
	var m map[string]any
	_ = json.Unmarshal(b, &m)
	if m["schema_version"] != servicecomponent.SchemaVersion {
		t.Fatalf("schema_version = %v, want %s", m["schema_version"], servicecomponent.SchemaVersion)
	}
}
