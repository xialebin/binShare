<?php

$arr = [7,8,9,4,5,7,4,1,2,3,5,6];

for($i=0;$i<count($arr);$i++){
    for ($j=0;$j<count($arr)-1-$i;$j++){

        if ($arr[$j]>$arr[$j+1]) {
            $temp = $arr[$j];
            $arr[$j] = $arr[$j+1];
            $arr[$j+1] = $temp;
        }
    }
}

print_r($arr);die;