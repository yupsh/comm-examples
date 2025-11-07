package comm_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/comm"
)

func ExampleComm_basic() {
	// comm file1.txt file2.txt
	// Note: This requires actual sorted files to compare
	yup.MustRun(
		Comm("testdata/file1.txt", "testdata/file2.txt"),
	)
	// Output:
	//
}

