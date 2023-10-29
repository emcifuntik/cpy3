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

// PyOS_AfterFork_Child : https://docs.python.org/3/c-api/init.html#c.PyOS_AfterFork_Child
// func PyOS_AfterFork_Child() {
// 	C.PyOS_AfterFork_Child()
// }
