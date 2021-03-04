/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */
package data

import (
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	hclog "github.com/hashicorp/go-hclog"
)

// TODO:Can we have a CACHE to fetch it quickly?
// ExchangeRates holds Current Rate from European Central Bank
type ExchangeRates struct {
	log   hclog.Logger
	rates map[string]float64
}

// NewRates Create a New Rates Data Structure
func NewRates(l hclog.Logger) (*ExchangeRates, error) {
	er := &ExchangeRates{log: l, rates: map[string]float64{}}
	err := er.getRates()
	return er, err
}

// MonitorRates checks the rates in the ECB API every interval and sends a message to the
// returned channel when there are changes
//
// Note: the ECB API only returns data once a day, this function only simulates the changes
// in rates for demonstration purposes
func (e *ExchangeRates) MonitorRates(interval time.Duration) chan struct{} {
	ret := make(chan struct{})

	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case <-ticker.C:
				// just add a random difference to the rate and return it
				// this simulates the fluctuations in currency rates
				for k, v := range e.rates {
					// change can be 10% of original value
					change := (rand.Float64() / 10)
					// is this a postive or negative change
					direction := rand.Intn(1)

					if direction == 0 {
						// new value with be min 90% of old
						change = 1 - change
					} else {
						// new value will be 110% of old
						change = 1 + change
					}

					// modify the rate
					e.rates[k] = v * change
				}

				// notify updates, this will block unless there is a listener on the other end
				ret <- struct{}{}
			}
		}
	}()

	return ret
}

// GetRate will return the rate of base and dest currency
func (e *ExchangeRates) GetRate(base, dest string) (float64, error) {
	br, ok := e.rates[base]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", base)
	}

	dr, ok := e.rates[dest]
	if !ok {
		return 0, fmt.Errorf("Rate not found for currency %s", dest)
	}
	return dr / br, nil
}

func (e *ExchangeRates) getRates() error {
	resp, err := http.DefaultClient.Get("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml")
	if err != nil {
		return nil
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Expected status code 200 got %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	md := &Cubes{}
	xml.NewDecoder(resp.Body).Decode(&md)

	for _, c := range md.CubeData {
		// str to float64
		r, err := strconv.ParseFloat(c.Rate, 64)
		if err != nil {
			return err
		}

		e.rates[c.Currency] = r
	}
	return nil
}

// Cubes parses XML
type Cubes struct {
	CubeData []Cube `xml:"Cube>Cube>Cube"`
}

// Cube fetch Info from Cubes
type Cube struct {
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}
