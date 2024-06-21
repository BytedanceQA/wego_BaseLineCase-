package iface

func bubbleSort(nums []int){
    if len(nums) <= 1{
        return
    }

    for i := 0; i < len(nums); i++{
        flag := false
        for j := 0; j < len(nums) - i - 1; j++ {
            if nums[j] > nums[j + 1] {
                nums[j], nums[j + 1] = nums[j + 1], nums[j]
                flag = true
            }
        }
        if!fla {
            break
        }
    }
}

func find(nums []int, target int) int {
    for i := 0; i < len(nums); i++ {
        if nums[i] == target {
            return 1
        }
    }
    return -1
}

