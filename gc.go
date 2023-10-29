/*
Unless explicitly stated otherwise all files in this repository are licensed
under the MIT License.
This product includes software developed at Datadog (https://www.datadoghq.com/).
Copyright 2018 Datadog, Inc.
*/

package python3

/*
#include "Python.h"
*/
import "C"

// PyGC_Collect : https://docs.python.org/3/c-api/dict.html#c.PyGC_Collect
func PyGC_Collect() int {
	return int(C.PyGC_Collect())
}
