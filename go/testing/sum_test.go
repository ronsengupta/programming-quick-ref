// http://golang.org/pkg/testing/
// https://golang.org/doc/code.html

package sum

import (
    "testing"
) 

func TestSum(t *testing.T) {
    actual := Sum(10, 20)
    expected := 30
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}

