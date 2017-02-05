package spotcrime

import "testing"

const (
	testKey      = "privatekeyforspotcrimepublicusers-commercialuse-877.410.1607"
	testLat      = 33.39657
	testLon      = -112.03422
	httpProxyURL = "97.77.104.22"
)

func TestAPIWithProxy(t *testing.T) {
	c, err := New(testKey)

	if err != nil {
		t.Fatalf("expected no error but got %s", err.Error())
	}

	resp, err := c.GetCrimes(&Request{
		Lat:   testLat,
		Lon:   testLon,
		Proxy: httpProxyURL,
	})

	if err != nil {
		t.Fatalf("expected no error but got %s", err.Error())
	}

	if resp == nil {
		t.Fatal("expected not nil resp")
	}

	if len(resp.Results) == 0 {
		t.Fatal("expected more than zero results in response")
	}

	if resp.Results[0].CDID == 0 {
		t.Fatal("expected result CDID to be non empty")
	}

	if len(resp.Results[0].Date) == 0 {
		t.Fatal("expected result Date to be non empty")
	}

	if resp.Results[0].Lat == 0 {
		t.Fatal("expected result Lat to be non empty")
	}

	if len(resp.Results[0].Link) == 0 {
		t.Fatal("expected result Link to be non empty")
	}

	if resp.Results[0].Lon == 0 {
		t.Fatal("expected result Lon to be non empty")
	}

	if len(resp.Results[0].Type) == 0 {
		t.Fatal("expected result Type to be non empty")
	}
}
