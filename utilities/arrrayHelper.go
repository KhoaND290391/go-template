package utilities

import "google.golang.org/grpc/status"

//RemoveStringArrayItem RemoveStringArrayItem
func RemoveStringArrayItem(arr []string, item string) error {
	i := -1
	if len(arr) <= 0 {
		return status.Error(404, "Array is null or empty")
	}
	for index, val := range arr {
		if val == item {
			i = index
		}
	}
	if i == -1 {
		return status.Error(404, "Array is null or empty")
	}
	arr[i] = arr[len(arr)-1] // Copy last element to index i.
	arr[len(arr)-1] = ""     // Erase last element (write zero value).
	arr = arr[:len(arr)-1]
	return nil
}
