package main

func searchRange(nums []int, target int) []int {
	// 存储结果 目标值的开始位置以及结束位置
	var result = []int{}

	for i := 0; i < len(nums); i++ {
		// 情况 1 只存在一个目标值
		if len(nums) == 1 { // 如果只有一个元素
			if nums[i] == target {
				result = append(result, i, i)
				return result
			}
		} else { // 如果不止一个元素
			// 判断一下数组有无越界 如果 target在最后一个元素
			if i+1 > len(nums)-1 {
				if nums[i] == target {
					result = append(result, i, i)
					return result
				}
			} else {
				// 无越界情况
				if nums[i] == target && nums[i+1] != target {
					result = append(result, i, i)
					return result
				}
			}

			// 情况 2
			if nums[i] == target && nums[i+1] == target {
				result = append(result, i)

				// 开始找到目标元素尾索引
				for j := i + 2; j < len(nums); j++ {
					// 判断一下数组越界 如果 target 在最后一个元素
					if j == len(nums)-1 {
						if nums[j] == target {
							result = append(result, j)
							return result
						} else {
							result = append(result, j-1)
							return result
						}
					}
					// 如果不在最后一个元素 则继续往后找
					if nums[j] != target {
						result = append(result, j-1)
						return result
					}
				}
			}
		}
	}
	// 如果遍历完毕没有找到目标值，则返回[-1, -1]
	result = append(result, -1, -1)
	return result
}

func main() {
	nums := []int{2, 2}
	result := searchRange(nums, 2)
	println(result)
}
