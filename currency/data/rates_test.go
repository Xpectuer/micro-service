/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
package data

import (
	"fmt"
	"testing"

	hclog "github.com/hashicorp/go-hclog"
)

func TestNewRates(t *testing.T) {
	tr, err := NewRates(hclog.Default())
	t.Log(tr.GetRate("CNY", "USD"))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("Rates %#v", tr.rates)
}
