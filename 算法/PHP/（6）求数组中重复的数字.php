function getNum($arr){

    if (!$arr) {
        return 'param is empty';
    }

    $hash_arr = [];
    $rel = [];
    foreach ($arr as $k => $v){
        if (in_array($v,$hash_arr) && !in_array($v,$rel)) {
            $rel[] = $v;
        }else{
            $hash_arr[] = $v;
        }
    }
    return $rel;
}


$a = [1,5,2,2,2,3,6,6,4,4,4,5];

$rel = getNum($a);
var_dump($rel);