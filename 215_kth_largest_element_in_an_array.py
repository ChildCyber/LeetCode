import heapq

# 小根堆
class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        #   构造大小为 k 的小根堆
        heap = [x for x in nums[:k]]
        heapq.heapify(heap)
        n = len(nums)

        for i in range(k, n): # 前k个元素已经入堆，不需要再处理
            if nums[i] > heap[0]: # 当前元素大于堆顶元素
                heapq.heappop(heap) # 弹出堆顶，即当前堆中最小元素
                heapq.heappush(heap, nums[i]) # 当前元素入堆
        # 堆顶元素即为结果
        return heap[0]
