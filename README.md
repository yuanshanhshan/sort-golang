# 排序算法

## 冒泡排序
> 算法描述
比较相邻的元素。如果第一个比第二个大，就交换它们两个；
对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
针对所有的元素重复以上的步骤，除了最后一个；
重复步骤1~3，直到排序完成。

### 动画演示

[动画](./doc/images/bubbleSort.gif)

## 选择排序
> 算法描述
n个记录的直接选择排序可经过n-1趟直接选择排序得到有序结果。具体算法描述如下：
将假想墙放置在数字列表最左侧，墙的左侧为已排序子列表，右侧为未排序子列表；
找出（选择）未排序子列表中的最小（或最大）元素；
把选择的元素与未排序列表中第一个元素进行交换；
将假想墙向右移动一个位置；
反复执行 2 至 4 步操作，直至整个数字列表排序完成（需要 n - 1 轮）。

### 动画演示
[选择排序](doc/images/selectSort.gif)

## 插入排序
> 算法描述
一般来说，插入排序都采用in-place在数组上实现。具体算法描述如下：
从第一个元素开始，该元素可以认为已经被排序；
取出下一个元素，在已经排序的元素序列中从后向前扫描；
如果该元素（已排序）大于新元素，将该元素移到下一位置；
重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；将新元素插入到该位置后；
重复步骤2~5。

### 动画演示

[插入排序](./doc/images/selectSort.gif)

## 快速排序
> 算法描述
快速排序的基本思想：通过一趟排序将待排记录分隔成独立的两部分，其中一部分记录的关键字均比另一部分的关键字小，则可分别对这两部分记录继续进行排序，以达到整个序列有序。

从数列中挑出一个元素，称为 “基准”（pivot）；
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。

### 动画演示
[快速排序](doc/images/quickSort.gif)