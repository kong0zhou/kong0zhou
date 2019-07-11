<?php 
//设置最大执行时间
set_time_limit(0);
function getHtml($url){
  // 1. 初始化
   $ch = curl_init();
   // 2. 设置选项，包括URL

   //设置需要获取的url地址
   curl_setopt($ch,CURLOPT_URL,$url);
   //设置是否以字符串形式返回
   curl_setopt($ch,CURLOPT_RETURNTRANSFER,1);
   //
   curl_setopt($ch,CURLOPT_HEADER,0);
   // 3. 执行并获取HTML文档内容
   $output = curl_exec($ch);
   if($output === FALSE ){
    $output = '';
   }
   // 4. 释放curl句柄
   curl_close($ch);
   return $output;
}
function getPageData($url){
  // 获取整个网页内容
  $html = getHtml($url);
  preg_match_all("/<p.*>.*<\/p>/",$html,$result_array);
  return $result_array;
}

$arr = getPageData("https://blog.csdn.net/YDTG1993/article/details/83861629");
foreach($arr as $value){
  echo json_encode($value);
}

?>