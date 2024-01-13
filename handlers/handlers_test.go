package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	p "re-partners-take-home-assignment/packing"
	"reflect"
	"testing"
)

func TestFindOptimalPackingHandler(t *testing.T) {
	// setup test cases
	pack_sizes_1 := []int{250, 500, 1000, 2000, 5000}
	pack_sizes_2 := []int{23, 31, 53}

	test_cases := []struct {
		numItems  int
		packSizes []int
		expected  p.Solution
	}{
		{1, pack_sizes_1, p.Solution{250: 1, 500: 0, 1000: 0, 2000: 0, 5000: 0}},
		{250, pack_sizes_1, p.Solution{250: 1, 500: 0, 1000: 0, 2000: 0, 5000: 0}},
		{251, pack_sizes_1, p.Solution{250: 0, 500: 1, 1000: 0, 2000: 0, 5000: 0}},
		{501, pack_sizes_1, p.Solution{250: 1, 500: 1, 1000: 0, 2000: 0, 5000: 0}},
		{12001, pack_sizes_1, p.Solution{250: 1, 500: 0, 1000: 0, 2000: 1, 5000: 2}},
		{263, pack_sizes_2, p.Solution{23: 2, 31: 7, 53: 0}},
	}

	// construct handler
	handler := http.HandlerFunc(FindOptimalPackingHandler)

	// send request for each test case and assert that response is correct
	for _, test_case := range test_cases {
		requestBody, err := json.Marshal(findOptimalPackingRequestBody{
			PackSizes: test_case.packSizes,
			NumItems:  test_case.numItems,
		})
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest("POST", "/find-optimal-packing", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var response p.Solution
		if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
			t.Fatal(err)
		}

		if !reflect.DeepEqual(response, test_case.expected) {
			t.Errorf("POST /find-optimal-packing: (%d, %v) = %v; want %v", test_case.numItems, test_case.packSizes, response, test_case.expected)
		}
	}
}
