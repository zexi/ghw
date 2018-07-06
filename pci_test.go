//
// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package ghw

import (
	"os"
	"testing"
)

func TestPCI(t *testing.T) {
	if _, ok := os.LookupEnv("GHW_TESTING_SKIP_PCI"); ok {
		t.Skip("Skipping PCI tests.")
	}
	info, err := PCI()
	if err != nil {
		t.Fatalf("Expected no error creating PciInfo, but got %v", err)
	}

	if len(info.Classes) == 0 {
		t.Fatalf("Expected >0 PCI classes, but found 0.")
	}

	wirelessController, exists := info.Classes["0d"]
	if !exists {
		t.Fatalf("Expected to find wireless controller class in hash for identifier '0d'")
	}
	if wirelessController.Name != "Wireless controller" {
		t.Fatalf("Expected wireless controller class name to be 'Wireless controller' but got '%v'", wirelessController.Name)
	}

	if len(wirelessController.Subclasses) == 0 {
		t.Fatalf("Expected >0 Subclasses for wirelessController, but found 0.")
	}

	foundRFController := false
	for _, sc := range wirelessController.Subclasses {
		if sc.Id == "10" {
			foundRFController = true
		}
	}
	if !foundRFController {
		t.Fatalf("Failed to find RF Controller subclass to wirelessController with ID of '10'.")
	}

	if len(info.Vendors) == 0 {
		t.Fatalf("Expected >0 PCI vendors, but found 0.")
	}

	intelInc, exists := info.Vendors["8086"]
	if !exists {
		t.Fatalf("Expected to find Intel vendor in hash for identifier '8086'")
	}
	if intelInc.Name != "Intel Corporation" {
		t.Fatalf("Expected Intel vendor name to be 'Intel Corporation' but got '%v'", intelInc.Name)
	}

	if len(info.Products) == 0 {
		t.Fatalf("Expected >0 PCI products, but found 0.")
	}

	intel10GBackplaneKey := "808610f8"
	intel10GBackplane, exists := info.Products[intel10GBackplaneKey]
	if !exists {
		t.Fatalf("Expected to find Intel 10GB Backplane Connection in products hash for key '808610f8'")
	}
	if intel10GBackplane.Name != "82599 10 Gigabit Dual Port Backplane Connection" {
		t.Fatalf("Expected Intel product '10f8' to have name '82599 10 Gigabit Dual Port Backplane Connection' but got %v", intel10GBackplane.Name)
	}

	// Make sure this product is linked in the Intel PCIVendorInfo.Products array
	foundBackplane := false
	for _, prod := range intelInc.Products {
		if prod.Id == "10f8" {
			foundBackplane = true
		}
	}
	if !foundBackplane {
		t.Fatalf("Failed to find 82599 10G backplane connection to in Intel vendor products array.")
	}
}
