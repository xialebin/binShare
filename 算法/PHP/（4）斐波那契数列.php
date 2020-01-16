<?php

//递归方法
function fbi($i){
    if ($i<2) {
        return $i == 0 ? 0 : 1;
    }
    return fbi($i-1)+fbi($i-2);
}

//顺序计算方法
function fbi_1($i){

    if ($i<3) {
        return $i == 0 ? 0 : 1;
    }

    $numOne = 1;
    $numTwo = 1;
    $amount = 0;
    for($n=2;$n<$i;$n++){
        $amount = $numTwo+$numOne;
        $numTwo = $numOne;
        $numOne = $amount;
    }
    return $amount;
}

for($i=1;$i<1000;$i++){
    $a = fbi_1($i);
    echo $a.'<br/>';
}die;