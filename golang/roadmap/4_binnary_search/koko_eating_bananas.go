package binarysearch

import "math"

/*
Koko Eating Bananas
You are given an integer array piles where piles[i] is the number of bananas in the ith pile. You are also given an integer h, which represents the number of hours you have to eat all the bananas.

You may decide your bananas-per-hour eating rate of k. Each hour, you may choose a pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, you may finish eating the pile but you can not eat from another pile in the same hour.

Return the minimum integer k such that you can eat all the bananas within h hours.

Example 1:

Input: piles = [1,4,3,2], h = 9

Output: 2
Explanation: With an eating rate of 2, you can eat the bananas in 6 hours. With an eating rate of 1, you would need 10 hours to eat all the bananas (which exceeds h=9), thus the minimum eating rate is 2.

Example 2:

Input: piles = [25,10,23,4], h = 4

Output: 25
*/

func MinEatingSpeed(piles []int, hour int) int {
	l, r := 1, 0

	for _, p := range piles {
		if r < p {
			r = p
		}
	}

	result := r

	for l <= r {
		k := l + (r-l)/2
		totalTime := 0

		for _, p := range piles {
			totalTime += int(math.Ceil(float64(p) / float64(k)))
		}

		if totalTime <= hour {
			result = k
			r = k - 1
		} else {
			l = k + 1
		}
	}

	return result
}
