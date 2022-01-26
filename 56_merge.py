class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        # 排序
        intervals.sort(key=lambda x: x[0])

        merged = []
        for interval in intervals:
            # 如果列表为空，或者当前区间与上一区间不重合，直接添加，将第一个区间加入 merged 数组中
            if not merged or merged[-1][1] < interval[0]:
                # 如果当前区间的左端点在数组 merged 中最后一个区间的右端点之后，那么它们不会重合，可以直接将这个区间加入数组 merged 的末尾
                merged.append(interval) # interval为list
            else:
                # 否则的话，可以与上一区间进行合并
                # 重合，需要用当前区间的右端点更新数组 merged 中最后一个区间的右端点，将其置为二者的较大值
                merged[-1][1] = max(merged[-1][1], interval[1])

        return merged
