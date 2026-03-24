<!-- head  ->  end

len(list) = N
// 获取打印开始为   n-K-1  => start
foreach(list as k => v){// k =>0 -> N
	if (k <start) {
		con
	}
} -->

<?php
$list = [1,2,3,4,5];
$k = 5;

var_dump(printArrK($list, $k));
function  printArrK($list, $kNode) {
    $act = [];
    $len = length($list);
    $start = $len - $kNode;
    foreach ($list as $k => $v) {
        if ($k < $start) {
            continue;
        }
        $act[] = $v;
    }
    return $act;
}