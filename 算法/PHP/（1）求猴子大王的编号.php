<?php

//使用array_push模拟一个循环的数组·

function king($n,$m){
    //$monkeys = [1,2,3,4,5,6,7,8,9,10];
    //$m = 5;
    $monkeys = range(1,$n);
    $i = 0;
    while (count($monkeys) > 1){
        if (($i+1)%$m==0) {
            unset($monkeys[$i]);
        }else{
            array_push($monkeys,$monkeys[$i]);
            unset($monkeys[$i]);
        }
        $i++;
    }

    return current($monkeys);
}


$rel = king(2,1);
var_dump($rel);die;