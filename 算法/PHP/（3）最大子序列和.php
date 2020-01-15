<?php


function maxSum($arr,$low,$high){

    if ($low == $high) {
        return $arr[$low];
    }

    $mid = floor(($low+$high)/2); //向下取整


    /**
     * 求包含中点的向左的序列的最大值
     */
    $left_max = null;
    $sum = 0;
    for ($i=$mid;$i>=$low;$i--) {
        $sum += $arr[$i];
        if ($sum > $left_max) {
            $left_max = $sum;
        }
    }


    /**
     * 求包含中点的向右的序列最大值
     */
    $right_max = null;
    $sum = 0;
    for ($i=$mid+1;$i<=$high;$i++) {
        $sum += $arr[$i];
        if ($sum>$right_max) {
            $right_max = $sum;
        }
    }

    $max_cross = maxSum($arr,$low,$mid) > maxSum($arr,$mid+1,$high) ? maxSum($arr,$low,$mid) : maxSum($arr,$mid+1,$high);
    return $max_cross>($left_max+$right_max) ? $max_cross : $left_max+$right_max;

}


$arr = [1,2,-3];
$rel = maxSum($arr,0,count($arr)-1);
var_dump($rel);die;