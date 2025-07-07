/* Copyright 2017 Victor Penso, Matteo Dessalvi

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestNodesMetrics(t *testing.T) {
	// Read the input data from a file
	file, err := os.Open("test_data/sinfo.txt")
	if err != nil {
		t.Fatalf("Can not open test data: %v", err)
	}
	data, err := ioutil.ReadAll(file)
	t.Logf("%+v", ParseNodesMetrics(data))
}

func TestNodesMetricsWithIdleStates(t *testing.T) {
	// Test with the new aggregated format that includes idle# and idle~ states
	file, err := os.Open("test_data/sinfo_aggregated.txt")
	if err != nil {
		t.Fatalf("Can not open test data: %v", err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Can not read test data: %v", err)
	}
	
	nm := ParseNodesMetrics(data)
	
	// Verify basic metrics
	if nm.idle != 150 {
		t.Errorf("Expected idle to be 150, got %f", nm.idle)
	}
	if nm.alloc != 50 {
		t.Errorf("Expected alloc to be 50, got %f", nm.alloc)
	}
	
	// Verify new idle states
	if nm.idle_resume != 5 {
		t.Errorf("Expected idle_resume to be 5, got %f", nm.idle_resume)
	}
	if nm.idle_suspend != 3 {
		t.Errorf("Expected idle_suspend to be 3, got %f", nm.idle_suspend)
	}
	
	t.Logf("Result: %+v", nm)
}

func TestNodesGetMetrics(t *testing.T) {
	t.Logf("%+v", NodesGetMetrics())
}
