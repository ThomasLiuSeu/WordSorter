package main

func SortSlice(strings []string, out chan<- []string) {
    sorted := &List{}

    for _, value := range strings {
        sorted.Add(value)
    }

    out <- sorted.ToSlice()
}

func JoinSortedSlices(sliceA, sliceB []string, out chan<- []string) {
    sorted := make([]string, len(sliceA) + len(sliceB))
    sliceAPos := 0
    sliceBPos := 0

    for i := 0; i < len(sorted); i++ {
        if sliceAPos >= len(sliceA) {
            sorted[i] = sliceB[sliceBPos]
            sliceBPos++
        } else if sliceBPos >= len(sliceB) {
            sorted[i] = sliceA[sliceAPos]
            sliceAPos++
        } else if sliceA[sliceAPos] < sliceB[sliceBPos] {
            sorted[i] = sliceA[sliceAPos]
            sliceAPos++
        } else {
            sorted[i] = sliceB[sliceBPos]
            sliceBPos++
        }
    }

    out <- sorted
}

