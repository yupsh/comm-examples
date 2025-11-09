package comm_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/comm"
)

// ExampleComm_basic shows the default behavior comparing two sorted files.
// Output has 3 columns:
//   - Column 1: Lines only in file1
//   - Column 2: Lines only in file2 (indented with tab)
//   - Column 3: Lines in both files (indented with two tabs)
func ExampleComm_basic() {
	yup.MustRun(
		Comm("testdata/fruits1.txt", "testdata/fruits2.txt"),
	)
	// Output:
	// apple
	// 		banana
	// 		cherry
	// 	kiwi
}

// ExampleComm_suppressColumn1 hides lines unique to file1.
// Shows only: lines unique to file2 + common lines
func ExampleComm_suppressColumn1() {
	yup.MustRun(
		Comm("testdata/fruits1.txt", "testdata/fruits2.txt", SuppressColumn1),
	)
	// Output:
	// 	banana
	// 	cherry
	// kiwi
}

// ExampleComm_suppressColumn2 hides lines unique to file2.
// Shows only: lines unique to file1 + common lines
func ExampleComm_suppressColumn2() {
	yup.MustRun(
		Comm("testdata/fruits1.txt", "testdata/fruits2.txt", SuppressColumn2),
	)
	// Output:
	// apple
	// 	banana
	// 	cherry
}

// ExampleComm_suppressColumn3 hides common lines.
// Shows only: lines unique to each file
func ExampleComm_suppressColumn3() {
	yup.MustRun(
		Comm("testdata/fruits1.txt", "testdata/fruits2.txt", SuppressColumn3),
	)
	// Output:
	// apple
	// 	kiwi
}

// ExampleComm_commonOnly shows only lines appearing in both files.
// This is done by suppressing columns 1 and 2.
func ExampleComm_commonOnly() {
	yup.MustRun(
		Comm("testdata/fruits1.txt", "testdata/fruits2.txt",
			SuppressColumn1,
			SuppressColumn2,
		),
	)
	// Output:
	// banana
	// cherry
}

// ExampleComm_uniqueOnly shows only lines unique to either file.
// This is done by suppressing column 3 (common lines).
func ExampleComm_uniqueOnly() {
	yup.MustRun(
		Comm("testdata/fruits1.txt", "testdata/fruits2.txt", SuppressColumn3),
	)
	// Output:
	// apple
	// 	kiwi
}
