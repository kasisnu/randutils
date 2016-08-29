package main

import "testing"

// Do the most basic smoke test
// TODO: Figure out why this doesn't feel right
func TestIps(t *testing.T) {
	all_ips := []string{
		"0.0.0.0",
		"1.1.1.1",
		"2.2.2.2",
		"3.3.3.3",
		"4.4.4.4",
		"5.5.5.5",
		"6.6.6.6",
		"7.7.7.7",
		"8.8.8.8",
		"9.9.9.9",
		"11.11.11.11",
		"22.22.22.22",
		"33.33.33.33",
		"44.44.44.44",
		"55.55.55.55",
		"66.66.66.66",
		"77.77.77.77",
		"88.88.88.88",
		"99.99.99.99",
		"111.111.111.111",
		"222.222.222.222",
		"255.255.255.255",
	}

	for _, ip := range all_ips {
		results := Ips(ip)
		if len(results) != 1 {
			t.Errorf("Expected: %s, Got: %q", ip, results)
		}

		result := results[0]
		if result != ip {
			t.Errorf("Expected: %s, Got: %s", ip, result)
		}
	}
}

func TestInvalidIps(t *testing.T) {
	all_ips := []string{
		"00.00.00.00",
		"000.000.000.000",
		"333.333.333.333",
		"444.444.444.444",
		"555.555.555.555",
		"666.666.666.666",
		"777.777.777.777",
		"888.888.888.888",
		"999.999.999.999",
		"256.256.256.256",
		"2222.2222.2222.2222",
	}

	for _, ip := range all_ips {
		res := Ips(ip)
		if len(res) != 0 {
			t.Errorf("Expected: NO MATCHES, Got: %s", res)
		}
	}
}
