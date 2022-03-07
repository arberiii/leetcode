package findMedianSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    length1, length2 := len(nums1), len(nums2)
    mergedLength := length1 + length2
    elementsToRemove := 0
    if mergedLength % 2 == 0 {
        elementsToRemove = mergedLength / 2 - 1
    } else {
        elementsToRemove = mergedLength / 2
    }
    counter := 0
    for counter < elementsToRemove {
        if len(nums1) == 0 {
            nums2 = nums2[1:]
        } else if len(nums2) == 0 {
            nums1 = nums1[1:]
        } else {
            if nums1[0] <= nums2[0] {
                nums1 = nums1[1:]
            } else {
                nums2 = nums2[1:]
            }
        }
        counter += 1
    }
    
    if mergedLength % 2 == 0 {
        first, nums1, nums2 := getTheSmallestOne(nums1, nums2)
        second, nums1, nums2 := getTheSmallestOne(nums1, nums2)
        return (first + second) / 2
    } else {
        result, _, _ := getTheSmallestOne(nums1, nums2)
        return result
    }
}

func getTheSmallestOne(nums1 []int, nums2 []int) (float64, []int, []int) {
    result := float64(0)
    if len(nums1) == 0 {
        result = float64(nums2[0])
        nums2 = nums2[1:]
        return result, nums1, nums2
    } else if len(nums2) == 0 {
        result = float64(nums1[0])
        nums1 = nums1[1:]
        return result, nums1, nums2
    } else {
        if nums1[0] <= nums2[0] {
            result = float64(nums1[0])
            nums1 = nums1[1:]
            return result, nums1, nums2
        } else {
            result = float64(nums2[0])
            nums2 = nums2[1:]
            return result, nums1, nums2
        }
    }
}
